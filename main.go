package main

import (
	"otp-verification-module/config"
	"otp-verification-module/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func init() {
	config.EnvVariableInitilizer()
}

func main() {

	r := gin.Default()
	store := cookie.NewStore([]byte("iamsuperkey"))
	r.Use(sessions.Sessions("mysession", store))

	r.POST("/sendOTP", controller.SendOTP)
	r.POST("/verifyOTP", controller.VerifyOTP)
	r.Run()

}
