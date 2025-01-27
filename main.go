package main

import (
	"fleet-management/config"
	"fleet-management/handlers"
	"fleet-management/repository"
	"github.com/gin-gonic/gin"
)

func main(){

	db, err := config.InitializeDynamoDB()
	if err != nil{
		panic(err)
	}

	carRepo := repository.NewCarRepository(db)

	carHandler := handlers.NewCarHandler(carRepo)

	r := gin.Default()

	    // Define routes
    r.POST("/cars", carHandler.CreateCar)
    r.GET("/cars/:id", carHandler.GetCar)

	r.Run()
}