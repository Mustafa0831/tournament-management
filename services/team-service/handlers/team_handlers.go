package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTeams(c *gin.Context) {
	// Здесь вы будете получать список команд из базы данных
	c.JSON(http.StatusOK, gin.H{
		"teams": []string{"Team A", "Team B"},
	})
}

func CreateTeam(c *gin.Context) {
	// Здесь вы будете обрабатывать создание новой команды
	c.JSON(http.StatusCreated, gin.H{
		"message": "Команда создана",
	})
}
