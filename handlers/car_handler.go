package handlers

import (
	"fleet-management/models"
    "fleet-management/repository"
    "github.com/gin-gonic/gin"
    "net/http"
)

type CarHandler struct {
	repo repository.CarRepository
}

func NewCarHandler(repo repository.CarRepository) *CarHandler {
	return &CarHandler{repo: repo}
}

func (h *CarHandler) CreateCar (c *gin.Context) {
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.CreateCar(car); err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err. Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Car Created successfully"})
}

func (h *CarHandler) GetCar (c *gin.Context) {
	id := c.Param("id")

	car, err := h.repo.GetCar(id)
	    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if car == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
        return
    }

    c.JSON(http.StatusOK, car)
}