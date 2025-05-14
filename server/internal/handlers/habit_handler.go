package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/habits/server/internal/models"
)

type HabitHandler struct {
	// Service layer
}

func NewHabitHandler() *HabitHandler {
	return &HabitHandler{}
}

func (h *HabitHandler) GetHabits(c *gin.Context) {
	habits := []models.Habit{}
	c.JSON(http.StatusOK, habits)
}

func (h *HabitHandler) CreateHabit(c *gin.Context) {
	var habit models.Habit
	if err := c.ShouldBindJSON(&habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, habit)
}
