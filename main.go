package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ahmadraza-khan/jobwebsite/config"
	"github.com/ahmadraza-khan/jobwebsite/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	config.ConnectDB()
}

func main() {
	r := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	routes.Routes(r)
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "The website is running successfully!")
	})
	log.Fatal(r.Run(port))
}
