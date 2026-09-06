package handler

import (
	"net/http"
	"user-service/internal/adapters/http/dto"
	"user-service/internal/domain/ports"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service ports.UserService
}

func NewAuthHandler(service ports.UserService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/login", h.Login)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{Token: token})
}
