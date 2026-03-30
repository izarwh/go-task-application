package services

import (
	"context"
	"encoding/json"
	"fmt"
	"task_planner_application/internal/common"
	"task_planner_application/internal/pkg/logger"
	"task_planner_application/internal/task/domain"
	"task_planner_application/internal/task/infra"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type ITaskService interface {
	GetTask(ctx context.Context, id uuid.UUID) (*domain.TaskResponse, error)
	GetTasks(ctx context.Context, page, limit int, filter domain.TaskFilter) (*domain.TaskListResponse, error)
	CreateTask(ctx context.Context, req *domain.TaskRequest) (*domain.TaskResponse, error)
	UpdateTask(ctx context.Context, id uuid.UUID, req *domain.TaskRequest) (*domain.TaskResponse, error)
	DeleteTask(ctx context.Context, id uuid.UUID) error
}

type taskService struct {
	taskRepo    infra.ITaskRepo
	redisClient *redis.Client
}

func NewTaskService(ITaskRepo infra.ITaskRepo, redisClient *redis.Client) ITaskService {
	return &taskService{
		taskRepo:    ITaskRepo,
		redisClient: redisClient,
	}
}

func (s *taskService) GetTask(ctx context.Context, id uuid.UUID) (*domain.TaskResponse, error) {
	cacheKey := fmt.Sprintf("task:%s", id.String())

	// 1. Try fetching from Redis
	cachedTask, err := s.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var taskResponse domain.TaskResponse
		if err := json.Unmarshal([]byte(cachedTask), &taskResponse); err == nil {
			return &taskResponse, nil
		}
	}

	// 2. Fetch from DB if Redis misses or fails
	task, err := s.taskRepo.GetTask(id)
	if err != nil {
		logger.Error(ctx, "failed to get task", err)
		return nil, err
	}

	taskResponse := domain.TaskDaoMapper(*task)

	// 3. Set to Redis
	if taskBytes, err := json.Marshal(taskResponse); err == nil {
		s.redisClient.Set(ctx, cacheKey, taskBytes, 15*time.Minute)
	}

	return &taskResponse, nil
}

func (s *taskService) CreateTask(ctx context.Context, req *domain.TaskRequest) (*domain.TaskResponse, error) {
	task := domain.TaskRequestMapper(*req)
	if err := s.taskRepo.CreateTask(&task); err != nil {
		return nil, err
	}

	taskResponse := domain.TaskDaoMapper(task)

	return &taskResponse, nil
}

func (s *taskService) UpdateTask(ctx context.Context, id uuid.UUID, req *domain.TaskRequest) (*domain.TaskResponse, error) {
	// check if the data exist
	task, err := s.taskRepo.GetTask(id)
	if err != nil {
		return nil, err
	}
	acuiredTask := domain.TaskRequestMapper(*req)

	if err := s.taskRepo.UpdateTask(acuiredTask.ID, task); err != nil {
		return nil, err
	}

	taskResponse := domain.TaskDaoMapper(*task)

	// invalidate cache
	cacheKey := fmt.Sprintf("task:%s", id.String())
	s.redisClient.Del(ctx, cacheKey)

	return &taskResponse, nil
}

func (s *taskService) DeleteTask(ctx context.Context, id uuid.UUID) error {
	task, err := s.taskRepo.GetTask(id)
	if err != nil {
		return err
	}

	if err := s.taskRepo.SoftDeleteTask(task.ID); err != nil {
		return err
	}

	// invalidate cache
	cacheKey := fmt.Sprintf("task:%s", id.String())
	s.redisClient.Del(ctx, cacheKey)

	return nil
}

func (s *taskService) GetTasks(ctx context.Context, page, limit int, filter domain.TaskFilter) (*domain.TaskListResponse, error) {
	tasks, err := s.taskRepo.GetTasks(page, limit, filter)
	if err != nil {
		return nil, err
	}

	totalPages := (len(tasks) + limit - 1) / limit
	return &domain.TaskListResponse{
		PaginationMetadata: common.PaginationMetadata{
			Page:       page,
			PageSize:   limit,
			TotalItems: len(tasks),
			TotalPages: totalPages,
		},
		Items: domain.TasksDaoMapper(tasks),
	}, nil
}
