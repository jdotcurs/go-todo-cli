package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Description string
	DueDate     time.Time
}
