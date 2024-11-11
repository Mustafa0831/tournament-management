package main

import (
	"log"

	"github.com/Mustafa0831/tournament-management/services/match-service/handlers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка чтения конфигурационного файла, %s", err)
	}

	port := viper.GetString("server.port")

	router := gin.Default()

	router.POST("/matches", handlers.CreateMatches)
	router.GET("/matches", handlers.GetMatches)

	log.Printf("Match Service is running on port %s", port)
	router.Run(":" + port)
}
