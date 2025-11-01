package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID			uint 			`gorm:"primaryKey" json:"id"`
	Title		string			`gorm:"not null" json:"title"`
	Description	string			`json:"description"`
	UserID		uint			`gorm:"not null" json:"user_id"`
	User		User			`gorm:"foreignKey:UserID" json:"user,omitempty"`
	Tasks		[]Task			`gorm:"foreignKey:ProjectID" json:"tasks,omitempty"`
	CreatedAt	time.Time		`gorm:"created_at"`
	UpdatedAt	time.Time		`gorm:"updated_at"`
	DeletedAt	gorm.DeletedAt	`gorm:"index" json:"-"`
}

