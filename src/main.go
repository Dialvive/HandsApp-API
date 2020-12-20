package main

import (
	"API/controllers"
	"API/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	// Routes for countries
	r.GET("/countries", controllers.GetCountries)
	r.POST("/countires", controllers.CreateCountry)

	r.Run(":8080")
}
