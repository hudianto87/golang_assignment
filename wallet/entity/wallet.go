package entity

import "time"

type Wallet struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required,min=3"`
	UserID      string    `json:"user_id" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Balance     float32   `json:"balance"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
