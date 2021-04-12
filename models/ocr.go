package models

import (
	"gorm.io/gorm"
)

// Create a custom model which wraps postgres connection pool
type OcrModel struct {
	DB *gorm.DB
}

// Create custom model type for async request response
type OcrAsync struct {
	gorm.Model
	TaskId string `json:"task_id"`
}

// Create custom model type for sync request response
type OcrSync struct {
	gorm.Model
	Text string `json:"text"`
}
