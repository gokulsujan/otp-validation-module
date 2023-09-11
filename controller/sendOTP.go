package controller

import (
	"math/rand"
	"net/http"
	"net/smtp"
	"otp-verification-module/config"
	"otp-verification-module/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func generateOTP() int {
	return rand.Intn(9999 - 1000)
}

func SendEmailOTP(name, email, otp string) error {
	auth := smtp.PlainAuth("", config.SMTPVar.Username, config.SMTPVar.Password, config.SMTPVar.Server)

	msg := "Subject: WebPortal OTP\nHey " + name + "Your OTP is " + otp

	err := smtp.SendMail(config.SMTPVar.Server+":"+config.SMTPVar.Port, auth, config.SMTPVar.Username, []string{email}, []byte(msg))
	return err
}

func SendOTP(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse the user data from json"})
		return
	}
	otp := generateOTP()

	if err := SendEmailOTP(user.Name, user.Email, strconv.Itoa(otp)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	} else {
		c.JSON(http.StatusAccepted, gin.H{"message": "Email deliverd to " + user.Email})
	}

}
