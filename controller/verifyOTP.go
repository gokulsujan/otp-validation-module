package controller

import (
	"net/http"
	"otp-verification-module/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func VerifyOTP(c *gin.Context) {
	var response models.OTPResponse
	if err := c.ShouldBindJSON(&response); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err})
		return
	}

	session := sessions.Default(c)
	if session.Get("iamsuperkey"+response.Email) == response.OTP {
		c.JSON(http.StatusAccepted, gin.H{"message": "OTP Verified"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid OTP"})
	}

}
