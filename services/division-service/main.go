package main

import (
	"log"

	"github.com/Mustafa0831/tournament-management/services/division-service/handlers"
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

	// Определите ваши маршруты
	router.POST("/divisions", handlers.CreateDivisions)
	router.GET("/divisions", handlers.GetDivisions)

	log.Printf("Division Service is running on port %s", port)
	router.Run(":" + port)
}
