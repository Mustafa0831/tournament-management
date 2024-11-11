package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMatches(c *gin.Context) {
	// Логика для создания матчей
	c.JSON(http.StatusCreated, gin.H{
		"message": "Matches created successfully",
	})
}

func GetMatches(c *gin.Context) {
	// Логика для получения списка матчей
	c.JSON(http.StatusOK, gin.H{
		"matches": []string{"Match 1", "Match 2"},
	})
}
