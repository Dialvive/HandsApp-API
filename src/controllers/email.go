package controllers

import (
	"API/models"
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/gin-gonic/gin"
)

//MeiliSearchWords returns search results
func SendEmail(c *gin.Context) {
	var input models.Email
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	emailHost := "smtp.gmail.com"
	emailFrom := "handsapp.bugs@gmail.com"
	emailPassword := "yourEmailPassword"
	emailPort := 587

	emailAuth := smtp.PlainAuth("", emailFrom, emailPassword, emailHost)

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := input.Subject + "!\n"
	msg := []byte(subject + mime + "\n" + input.Body + "\n ReplyTo:" + input.ReplyTo)
	addr := fmt.Sprintf("%s:%s", emailHost, emailPort)
	to := []string{"haikode@protonmail.com"}

	if err := smtp.SendMail(addr, emailAuth, emailFrom, to, msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		panic(err)
	} else {
		c.JSON(http.StatusOK, "Delivered")
	}
}
