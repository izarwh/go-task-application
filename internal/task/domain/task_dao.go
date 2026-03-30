package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskDao struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Title       string     `gorm:"type:varchar(255);not null"`
	Description string     `gorm:"type:jsonb"`
	Status      string     `gorm:"type:varchar(50);default:'pending'"`
	DueDate     *time.Time `gorm:"type:date"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	IsDeleted   bool           `gorm:"default:false"`
}

type TaskFilter struct {
	Status  string `gorm:"type:varchar(50)"`
	DueDate *time.Time
	Title   string `gorm:"type:varchar(255)"`
}
