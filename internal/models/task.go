package models

import (
	"time"

	"gorm.io/gorm"
)

type TaskStatus string

const (
	StatusTodo			TaskStatus = "todo"
	StatusInProgress	TaskStatus = "in_progress"
	StatusDone			TaskStatus = "done"
)

type Task struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `json:"description"`
	Status      TaskStatus     `gorm:"type:varchar(20);default:'todo'" json:"status"`
	ProjectID   uint           `gorm:"not null" json:"project_id"`
	Project     Project        `gorm:"foreignKey:ProjectID" json:"project,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}