package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"task_planner_application/internal/common/helper"
	"task_planner_application/internal/pkg/apperror"
	"time"

	"github.com/go-playground/validator/v10"
)

type TaskRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=255"`
	Description string `json:"description" validate:"required"`
	DueDate     string `json:"due_date" validate:"required,datetime=2006-01-02"`
	Status      string `json:"status" validate:"required,oneof=pending completed"`
}

type TaskQuery struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	TaskFilter
}

type TaskFilter struct {
	Status  string
	DueDate *time.Time
	Title   string
}

func TaskRequestMapper(req TaskRequest) TaskDao {
	return TaskDao{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     helper.StringTimeToTime(req.DueDate),
		Status:      TaskStatus(req.Status),
	}
}

func (req *TaskRequest) BindAndValidate(r *http.Request, validators *validator.Validate) error {
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	return helper.ValidateStruct(validators, req)
}

func (f *TaskQuery) BindQuery(r *http.Request) error {
	query := r.URL.Query()

	f.Page = 1
	if p := query.Get("page"); p != "" {
		val, err := strconv.Atoi(p)
		if err != nil || val <= 0 {
			return apperror.BadRequest("Invalid page number", err)
		}
		f.Page = val
	}

	f.Limit = 10
	if l := query.Get("limit"); l != "" {
		val, err := strconv.Atoi(l)
		if err != nil || val <= 0 {
			return apperror.BadRequest("Invalid limit number", err)
		}

		if val > 100 {
			val = 100
		}
		f.Limit = val
	}

	f.Status = query.Get("status")
	f.Title = query.Get("title")

	if dueDateStr := query.Get("due_date"); dueDateStr != "" {
		parsedDate, err := time.Parse(time.DateOnly, dueDateStr)
		if err == nil {
			f.DueDate = &parsedDate
		} else {
			parsedDateRFC, err := time.Parse(time.RFC3339, dueDateStr)
			if err == nil {
				f.DueDate = &parsedDateRFC
			} else {
				return apperror.BadRequest(
					fmt.Sprintf("Invalid date format for '%s'. Use YYYY-MM-DD or RFC3339", dueDateStr),
					err,
				)
			}
		}
	}

	return nil
}
