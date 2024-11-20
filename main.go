package main

import (
	"log"
	"net/http"

	"github.com/ahmadraza-khan/jobwebsite/config"
	"github.com/ahmadraza-khan/jobwebsite/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func init() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	config.ConnectDB()
	router = gin.Default()
	routes.Routes(router)
	router.GET("/", func(ctx *gin.Context) {
		ctx.File("index.html")
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
