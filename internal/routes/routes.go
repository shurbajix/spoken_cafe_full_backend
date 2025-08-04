package routes

import (
	"spoken_cafe_backend/internal/handlers/auth"
	"spoken_cafe_backend/internal/handlers/lessons"
	"spoken_cafe_backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	//Auth routes (public)
	api.POST("/login", auth.LoginHandler)
	api.POST("/register", auth.RegisterHandler)

	//Lessons routes (public)
	api.GET("/lessons", lessons.GetLessons)

	//Protected routes
	authorized := api.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/profile", auth.ProfileHandler)
	}
}
