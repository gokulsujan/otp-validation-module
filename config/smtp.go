package config

import "os"

var SMTPVar = struct {
	Server   string
	Port     string
	Username string
	Password string
}{
	Server:   "smtp.gmail.com",
	Port:     "587",
	Username: os.Getenv("smtpusername"),
	Password: os.Getenv("smtppassword"),
}
