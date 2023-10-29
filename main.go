package main

import (
	"example/back/config"
	"example/back/routes"
	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "os"
    "github.com/joho/godotenv"
)

func main() {
    godotenv.Load()
    router := gin.Default()
    corsConfig := cors.DefaultConfig()
    corsConfig.AllowOrigins = []string{os.Getenv("URL_REACT")}
    corsConfig.AllowCredentials = true
    corsConfig.AddAllowMethods("OPTIONS")
    router.Use(cors.New(corsConfig))
    config.DatabaseConnection()
    routes.ContactRouter(router)
    router.Run("localhost:8080")
}
