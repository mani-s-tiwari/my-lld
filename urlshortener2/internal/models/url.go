package models

import "time"

type URL struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ShortCode   string    `gorm:"uniqueIndex;not null" json:"short_code"`
	OriginalURL string    `gorm:"not null" json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiresAt   time.Time `json:"expired_at"`
	ClickCount  int64     `json:"click_count"`
	IsActive    bool      `json:"is_active"`
}

type CreateURLRequest struct {
	OriginalURL string    `json:"original_url" binding:"required,url"`
	CustomCode  string    `json:"custom_code,omitempty"`
	ExpiresAt   time.Time `json:"expires_at,omitempty"`
}

type URLResponse struct {
	ShortCode   string    `json:"short_code"`
	ShortURL    string    `json:"short_url"`
	OriginalURL string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiresAt   time.Time `json:"expires_at"`
	ClickCount  int64     `json:"click_count"`
}
