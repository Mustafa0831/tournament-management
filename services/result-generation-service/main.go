package main

import (
	"log"

	"github.com/Mustafa0831/tournament-management/services/result-generation-service/handlers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// Инициализация конфигурации
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка чтения конфигурационного файла, %s", err)
	}

	port := viper.GetString("server.port")

	router := gin.Default()

	// Определите маршруты
	router.POST("/generate-results/divisionA", handlers.GenerateDivisionAResults)
	router.POST("/generate-results/divisionB", handlers.GenerateDivisionBResults)
	router.POST("/generate-results/playoff", handlers.GeneratePlayoffResults)

	log.Printf("Result Generation Service is running on port %s", port)
	router.Run(":" + port)
}
