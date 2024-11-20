package main

import (
	"log"

	"github.com/ahmadraza-khan/jobwebsite/config"
	"github.com/ahmadraza-khan/jobwebsite/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	config.ConnectDB()

	router = gin.Default()

	routes.Routes(router)

	router.GET("/", func(ctx *gin.Context) {
		ctx.File("index.html")
	})
}

func Handler(c *gin.Context) {
	router.HandleContext(c)
}
