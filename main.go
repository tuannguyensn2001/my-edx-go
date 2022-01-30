package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"my-edx-go/database"
	"my-edx-go/routes"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.SetFormatter(&log.JSONFormatter{})

	dbstring := os.Getenv("DB_STRING")

	store, err := database.NewStore(dbstring)

	if err != nil {
		log.Fatal("Database error")
	}

	router := gin.Default()

	router.Use(cors.Default())

	v1 := router.Group("/api/v1")
	{
		init := routes.InitRoute(v1, store)
		init("auth")
	}

	err = router.Run(":5000")
	if err != nil {
		log.Fatal("Error")
	}
}
