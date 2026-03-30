package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"task_planner_application/internal/common/helper"
	"task_planner_application/internal/pkg/apperror"
	"task_planner_application/internal/pkg/logger"
	"task_planner_application/internal/task/domain"
	"task_planner_application/internal/task/services"

	"github.com/go-playground/validator/v10"
)

type ITaskHandler interface {
	CreateTask(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
	GetTask(w http.ResponseWriter, r *http.Request)
	GetTasks(w http.ResponseWriter, r *http.Request)
}

type TaskHandler struct {
	service  services.ITaskService
	validate *validator.Validate
}

func NewTaskHandler(service services.ITaskService) ITaskHandler {
	return &TaskHandler{
		service:  service,
		validate: validator.New(),
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req domain.TaskRequest

	if err := req.BindAndValidate(r, h.validate); err != nil {
		logger.Error(ctx, "invalid request", err)
		h.sendError(ctx, w, apperror.BadRequest("Invalid request", err))
		return
	}

	task, err := h.service.CreateTask(ctx, &req)
	if err != nil {
		logger.Error(ctx, "failed to create task", err)
		h.sendError(ctx, w, err)
		return
	}

	logger.Info(ctx, "task created", "task_id", task.ID)
	h.sendJSON(ctx, w, http.StatusCreated, map[string]any{
		"message": "Task created successfully",
		"task":    task,
	})
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req domain.TaskRequest

	taskID := r.PathValue("task_id")
	uuid := helper.StringUUIDToUUID(taskID)

	if err := req.BindAndValidate(r, h.validate); err != nil {
		logger.Error(ctx, "invalid request", err)
		h.sendError(ctx, w, apperror.BadRequest("Invalid request", err))
		return
	}

	task, err := h.service.UpdateTask(ctx, *uuid, &req)
	if err != nil {
		logger.Error(ctx, "failed to update task", err)
		h.sendError(ctx, w, err)
		return
	}

	logger.Info(ctx, "task updated", "task_id", task.ID)
	h.sendJSON(ctx, w, http.StatusOK, map[string]any{
		"message": "Task updated successfully",
		"task":    task,
	})
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	taskID := r.PathValue("task_id")
	uuid := helper.StringUUIDToUUID(taskID)

	if err := h.service.DeleteTask(ctx, *uuid); err != nil {
		logger.Error(ctx, "failed to delete task", err)
		h.sendError(ctx, w, err)
		return
	}

	logger.Info(ctx, "task deleted", "task_id", taskID)
	h.sendJSON(ctx, w, http.StatusOK, map[string]any{
		"message": "Task deleted successfully",
	})
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req domain.TaskQuery

	if err := req.BindQuery(r); err != nil {
		logger.Error(ctx, "failed to bind query", err)
		h.sendError(ctx, w, err)
		return
	}

	tasks, err := h.service.GetTasks(ctx, req.Page, req.Limit, req.TaskFilter)
	if err != nil {
		logger.Error(ctx, "failed to get tasks", err)
		h.sendError(ctx, w, err)
		return
	}

	logger.Info(ctx, "tasks retrieved", "count", len(tasks.Items))
	h.sendJSON(ctx, w, http.StatusOK, map[string]any{
		"tasks": tasks,
	})
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	taskID := r.PathValue("task_id")
	uuid := helper.StringUUIDToUUID(taskID)

	task, err := h.service.GetTask(ctx, *uuid)
	if err != nil {
		logger.Error(ctx, "failed to get task", err)
		h.sendError(ctx, w, err)
		return
	}

	logger.Info(ctx, "task retrieved", "task_id", taskID)
	h.sendJSON(ctx, w, http.StatusOK, map[string]any{
		"task": task,
	})
}

func (h *TaskHandler) sendError(ctx context.Context, w http.ResponseWriter, err error) {
	customErr := apperror.Extract(err)
	h.sendJSON(ctx, w, customErr.StatusCode, map[string]any{
		"code":    customErr.Code,
		"message": customErr.Message,
	})
}

func (h *TaskHandler) sendJSON(ctx context.Context, w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Error(ctx, "failed to encode response", err)
	}
}
