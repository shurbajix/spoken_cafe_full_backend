package auth

import (
	"context"
	"net/http"
	"spoken_cafe_backend/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

type registerRequest struct {
	Username string `json:"username" binding:"required`
	Password string `json:"password" binding:"required,min=6"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req regisyerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	err := h.service.Register(ctx, req.Username, req.Password)
	if err != nil {
		c.Json(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password binding:"required`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background
	token, err := h.service.Login(ctx, req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
