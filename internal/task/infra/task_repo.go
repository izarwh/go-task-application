package infra

import (
	"task_planner_application/internal/pkg/apperror"
	"task_planner_application/internal/task/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ITaskRepo interface {
	GetTask(id uuid.UUID) (*domain.TaskDao, error)
	GetTasks(page, limit int, filter domain.TaskFilter) ([]domain.TaskDao, error)
	CreateTask(task *domain.TaskDao) error
	UpdateTask(id uuid.UUID, task *domain.TaskDao) error
	SoftDeleteTask(id uuid.UUID) error
}

type taskRepo struct {
	DB *gorm.DB
}

func NewTaskRepo(db *gorm.DB) ITaskRepo {
	return &taskRepo{DB: db}
}

func (r *taskRepo) GetTask(id uuid.UUID) (*domain.TaskDao, error) {
	var task domain.TaskDao
	if err := r.DB.First(&task).Where("id = ?", id).Error; err != nil {
		return nil, apperror.FromDB(err, "task")
	}
	return &task, nil
}

func (r *taskRepo) GetTasks(page, limit int, filter domain.TaskFilter) ([]domain.TaskDao, error) {
	var tasks []domain.TaskDao
	offsets := (page - 1) * limit

	if err := r.DB.Offset(offsets).Limit(limit).Find(&tasks).Scopes(func(d *gorm.DB) *gorm.DB {
		return filterTasks(d, filter)
	}).Error; err != nil {
		return nil, apperror.FromDB(err, "task")
	}

	return tasks, nil
}

func (r *taskRepo) CreateTask(task *domain.TaskDao) error {
	err := r.DB.Create(task).Error
	if err != nil {
		return apperror.FromDB(err, "task")
	}
	return nil
}

func (r *taskRepo) UpdateTask(id uuid.UUID, task *domain.TaskDao) error {
	err := r.DB.Model(&domain.TaskDao{}).Where("id = ?", id).Updates(task).Error
	if err != nil {
		return apperror.FromDB(err, "task")
	}
	return nil
}

func (r *taskRepo) SoftDeleteTask(id uuid.UUID) error {
	err := r.DB.Model(&domain.TaskDao{}).Where("id = ?", id).Update("is_deleted", true).Error
	if err != nil {
		return apperror.FromDB(err, "task")
	}
	return nil
}

func filterTasks(DB *gorm.DB, filter domain.TaskFilter) *gorm.DB {
	query := DB.Model(&domain.TaskDao{})

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.DueDate != nil {
		query = query.Where("due_date = ?", filter.DueDate)
	}
	if filter.Title != "" {
		query = query.Where("title LIKE ?", "%"+filter.Title+"%")
	}

	return query
}
