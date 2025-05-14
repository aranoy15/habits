package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"time":   time.Now(),
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Route not found",
		})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "Method not allowed",
		})
	})

	api := router.Group("/api")
	v1 := api.Group("/v1")
	{
		habits := v1.Group("/habits")
		{
			habitHandler := new(HabitHandler)
			habits.GET("", habitHandler.GetHabits)
			habits.POST("", habitHandler.CreateHabit)
		}
	}
}
