package domain

import "task_planner_application/internal/common"

type TaskResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Status      string `json:"status"`
}

type TaskListResponse struct {
	common.PaginationMetadata `json:"pagination"`
	Items                     []TaskResponse `json:"items"`
}

func TasksDaoMapper(tasks []TaskDao) []TaskResponse {
	var res []TaskResponse
	for _, task := range tasks {
		res = append(res, taskDaoToResponse(task))
	}
	return res
}

func TaskDaoMapper(dao TaskDao) TaskResponse {
	return taskDaoToResponse(dao)
}

func taskDaoToResponse(dao TaskDao) TaskResponse {
	return TaskResponse{
		ID:          dao.ID.String(),
		Title:       dao.Title,
		Description: string(dao.Description),
		DueDate:     dao.DueDate.Format("2006-01-02"),
		Status:      string(dao.Status),
	}
}
