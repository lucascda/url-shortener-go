package models

import "gorm.io/gorm"

type Url struct {
	*gorm.Model
	OriginalUrl string `json:"original_url"`
	Hash        string `json:"hash"`
	UserId      int    `json:"user_id"`
}

type CreateUrl struct {
	Original_url string
}
