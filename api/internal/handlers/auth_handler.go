package handlers

import (
	"bookmark-api/internal/repositories"
	"bookmark-api/internal/services"
	"bookmark-api/internal/utils"
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var userRepo = repositories.UserRepository{}

func Signup(c *gin.Context) {
	var req SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to hash passwords",
		})
		return
	}

	err = userRepo.Create(
		req.Name,
		req.Email,
		string(hashedPassword),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email already exists",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
	})
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := userRepo.FindByEmail(
		req.Email,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	token, err := utils.GenerateToken(
		user.ID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate a token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func ForgotPassword(c *gin.Context) {
	var req ForgotPasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := userRepo.FindByEmail(req.Email)
	if err != nil {
		// Don't reveal if email exists (security best practice)
		c.JSON(http.StatusOK, gin.H{
			"message": "if email exists, reset link will be sent",
		})
		return
	}

	resetToken := generateResetToken()
	expiresAt := time.Now().Add(24 * time.Hour)

	err = userRepo.CreatePasswordResetToken(user.ID, resetToken, expiresAt)
	if err != nil {
		log.Printf("CreatePasswordResetToken error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create reset token",
		})
		return
	}

	emailService := services.NewEmailService()
	err = emailService.SendPasswordResetEmail(user.Email, resetToken)
	if err != nil {
		log.Printf("Failed to send reset email: %v", err)
		// Still return success to user (security: don't leak if email failed)
	}

	// Always return success message
	c.JSON(http.StatusOK, gin.H{
		"message": "if email exists, password reset link has been sent",
	})
}

func ResetPassword(c *gin.Context) {
	var req ResetPasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := userRepo.FindUserByResetToken(req.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid or expired reset token",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.NewPassword),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	err = userRepo.UpdatePassword(user.ID, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update password",
		})
		return
	}

	err = userRepo.InvalidateResetToken(req.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to invalidate reset token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "password reset successfully",
	})
}

func generateResetToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}
