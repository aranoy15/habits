package models

import "time"

type Progress struct {
	ID          string    `json:"id"`
	HabitID     string    `json:"habit_id"`
	Date        time.Time `json:"date"`
	IsCompleted bool      `json:"is_completed"`
	Notes       string    `json:"notes"`
}
