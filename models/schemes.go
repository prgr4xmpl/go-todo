package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Name        string    `gorm:"type:string;required json:"name"`
	Description string    `gorm:"type:string; json:"description"`
	Done        bool      `json:"done"`
	Created     time.Time `json:"created"`
	Modified    time.Time `json:"modified"`
}
