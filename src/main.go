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
	r.POST("/country", controllers.CreateCountry)
	r.GET("/country/:id", controllers.FindCountry)
	r.PUT("/country/:id", controllers.UpdateCountry)
	r.DELETE("/country/:id", controllers.DeleteCountry)

	r.Run(":8080")
}
