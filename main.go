package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"my-edx-go/database"
	authtransport "my-edx-go/modules/auth/transport"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	store, err := database.NewStore()

	if err != nil {
		log.Fatal("Database error")
	}

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/login", authtransport.LoginTransport(store.DB))
	}

	err = router.Run(":5000")
	if err != nil {
		log.Fatal("Error")
	}
}
