package models

import "time"

type Habit struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id,omitempty"`
	Name        string    `json:"name" binding:"required,min=3,max=100"`
	Description string    `json:"description" binding:"max=500"`
	TargetCount uint16    `json:"target_count" binding:"required,min=1"`
	CreatedAt   time.Time `json:"created_at"`
}
