package main

import jwt_provider "my-edx-go/providers/jwt"

func main() {

	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMzMTA5MTUsImlhdCI6MTY0MzMwNzMxNX0.SA0zD6PvET5Y9MYDv9eqrI7rStGhNxHZuvBJRET1H2s"

	jwt_provider.ValidateToken(tokenString)
	//log.SetFormatter(&log.JSONFormatter{})
	//
	//store, err := database.NewStore()
	//
	//if err != nil {
	//	log.Fatal("Database error")
	//}
	//
	//router := gin.Default()
	//
	//router.Use(cors.Default())
	//
	//v1 := router.Group("/api/v1")
	//{
	//	//v1.POST("/login", authtransport.LoginTransport(store.DB))
	//	init := routes.InitRoute(v1, store)
	//	init("auth")
	//}
	//
	//err = router.Run(":5000")
	//if err != nil {
	//	log.Fatal("Error")
	//}
}
