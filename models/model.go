package models

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type OTPResponse struct {
	OTP   string `json:"otp"`
	Email string `json:"email"`
}
