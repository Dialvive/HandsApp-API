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

	emailFrom := "handsapp.bugs@gmail.com"
	emailAuth := smtp.PlainAuth("", "handsapp.bugs@gmail.com", "PASSWORD", "smtp.gmail.com")
	addr := fmt.Sprintf("%s:%s", "smtp.gmail.com", "587")

	deliverTo := []string{"dialvive@protonmail.com"}
	msg := []byte("To:" + deliverTo[0] + "\r\n" +
		"Subject: " + input.Subject + "\r\n" +
		"\r\n" +
		input.Body +
		"\r\n" +
		"ReplyTo: " + input.ReplyTo + "\r\n")

	if err := smtp.SendMail(addr, emailAuth, emailFrom, deliverTo, msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		panic(err)
	} else {
		c.JSON(http.StatusOK, "Delivered")
	}
}
