package controller

import (
	"net/http"
	"otp-verification-module/models"
	"time"

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
		expiry := session.Get("expirationtime" + response.Email)
		if expiryTime, ok := expiry.(time.Time); ok && time.Now().Before(expiryTime) {
			c.JSON(http.StatusAccepted, gin.H{"message": "OTP Verified"})
		} else {
			c.JSON(http.StatusOK, "OTP expired")
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid OTP"})
	}

}
