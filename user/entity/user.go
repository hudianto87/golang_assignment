package entity

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" binding:"required,min=3"`
	Email     string    `json:"email" binding:"required,email"`
	Address   string    `json:"address" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
