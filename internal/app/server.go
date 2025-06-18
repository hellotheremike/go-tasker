package app

import (
	"fmt"
	"log"

	"github.com/hellotheremike/go-tasker/internal/config"
)

func StartServer() {

	router := SetupRouter()

	router.Run(fmt.Sprintf(":%s", config.Load().DATABASE_URL))

	log.Println("Server is running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
