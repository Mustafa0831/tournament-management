package main

import (
	"github.com/spf13/viper"
	"log"

	"github.com/Mustafa0831/tournament-management/services/team-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка чтения конфигурационного файла, %s", err)
	}

	port := viper.GetString("server.port")
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/teams", handlers.GetTeams)
	router.POST("/teams", handlers.CreateTeam)

	log.Printf("Team Service is running on port %s", port)
	err := router.Run(":" + port)
	if err != nil {
		return
	}
}
