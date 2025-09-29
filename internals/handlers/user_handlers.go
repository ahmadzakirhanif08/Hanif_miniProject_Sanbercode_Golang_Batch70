package handlers

import (
	"log"
	"net/http"

	"miniProject/internals/models"
	"miniProject/internals/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(s services.UserService) *UserHandler {
	return &UserHandler{userService: s}
}

// Login menangani POST request ke /api/users/login
func (h *UserHandler) Login(c *gin.Context) {
	var req models.UserLoginRequest

	// 1. Binding Input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid: " + err.Error()})
		return
	}

	// 2. Panggil Service Login
	token, err := h.userService.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		// Untuk alasan keamanan, hindari detail error yang terlalu spesifik ke klien
		log.Printf("Login failed for user %s: %v", req.Username, err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login gagal: Username atau password salah."})
		return
	}

	// 3. Response Berhasil
	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"token": token,
	})
}