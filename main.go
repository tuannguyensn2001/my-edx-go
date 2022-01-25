package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	authtransport "my-edx-go/modules/auth/transport"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/login", authtransport.LoginTransport())
	}

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error")
	}
}
