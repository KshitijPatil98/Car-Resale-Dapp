package main

import (
	"fmt"
	"net/http"
	"restapitest/models"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	server.GET("/cars", getCars)
	server.POST("/car", registerCar)
	server.Run(":8081")

}

func getCars(context *gin.Context) {

	cars := models.GetAll()
	if len(cars) == 0 {
		context.JSON(http.StatusOK, gin.H{"message": "No cars present"})
		return

	}
	context.JSON(http.StatusOK, cars)

}

func registerCar(context *gin.Context) {

	var car models.Car
	err := context.ShouldBindJSON(&car)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please ensure your payload is correct"})
		return
	}

	fmt.Print("The payload is as follows:", car)
	car.Register()
	context.JSON(http.StatusOK, gin.H{"message": "Successfully registered the car"})

}
