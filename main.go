package main

import (
	"otp-verification-module/config"
	"otp-verification-module/controller"

	"github.com/gin-gonic/gin"
)

func init() {
	config.EnvVariableInitilizer()
}

func main() {

	r := gin.Default()

	r.POST("/sendOTP", controller.SendOTP)
	r.Run()

}
