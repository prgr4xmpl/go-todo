package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
