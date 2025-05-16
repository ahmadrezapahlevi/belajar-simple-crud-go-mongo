package main

import (
	"go-mongo-crud/config"
	"go-mongo-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	router := gin.Default()
	routes.UserRoute(router)

	router.Run(":8080") // jalan di localhost:8080
}