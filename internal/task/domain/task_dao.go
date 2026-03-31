package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusCompleted TaskStatus = "completed"
)

type TaskDao struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Title       string     `gorm:"type:varchar(255);not null"`
	Description string     `gorm:"type:varchar(255)"`
	Status      TaskStatus `gorm:"type:varchar(50);default:'pending'"`
	DueDate     *time.Time `gorm:"type:date"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	IsDeleted   bool           `gorm:"default:false"`
}

func (t *TaskDao) TableName() string {
	return "tasks"
}
