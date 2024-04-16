package models

import (
	"time"

	"gorm.io/gorm"
)

type Url struct {
	*gorm.Model
	OriginalUrl string `json:"original_url"`
	Hash        string `json:"hash"`
	UserId      int    `json:"user_id"`
}

type ListAllUrls struct {
	ID           uint      `json:"id"`
	Original_url string    `json:"original_url"`
	Hash         string    `json:"hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateUrl struct {
	Original_url string `validate:"required,http_url"`
}
