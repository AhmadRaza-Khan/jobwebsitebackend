package routes

import (
	"github.com/ahmadraza-khan/jobwebsite/controllers"
	"github.com/ahmadraza-khan/jobwebsite/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Set-Cookie"},
		AllowCredentials: true,
	}))

	v1 := router.Group("/api/v1")
	v1.POST("/signup", controllers.SignUp)
	v1.POST("/login", controllers.Login)
	// protected routes
	v1.POST("/apply", middleware.Authentication, controllers.Apply)
	v1.POST("/application", middleware.Authentication, controllers.GetApplicationStatus)
	v1.GET("/logout", middleware.Authentication, controllers.Logout)
}
