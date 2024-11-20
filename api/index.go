package handler

import (
	"log"
	"net/http"

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

func Handler() {
	r := gin.Default()
	routes.Routes(r)
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "The site is running smoothly!")
	})
}
