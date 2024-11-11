package main

import (
	"github.com/Mustafa0831/tournament-management/services/playoff-service/handlers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
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
	router.POST("/playoffs", handlers.CreatePlayoffs)
	router.GET("/playoffs", handlers.GetPlayoffs)

	log.Printf("Playoff Service is running on port %s", port)
	router.Run(":" + port)
}
