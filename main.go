package main

import (
	"example/back/config"
	"example/back/routes"
	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    router := gin.Default()
    corsConfig := cors.DefaultConfig()
    corsConfig.AllowOrigins = []string{"http://localhost:3000"}
    corsConfig.AllowCredentials = true
    corsConfig.AddAllowMethods("OPTIONS")
    router.Use(cors.New(corsConfig))
    config.DatabaseConnection()
    routes.ContactRouter(router)
    router.Run("localhost:8080")
}
