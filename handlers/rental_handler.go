package handlers

import (
	"fleet-management/models"
    "fleet-management/repository"
    "github.com/gin-gonic/gin"
    "net/http"
)

type RentalHandler struct {
	repo repository.RentalRepository
}

func NewRentalHandler(repo repository.RentalRepository) *RentalHandler {
	return &RentalHandler{repo: repo}
}

func (h * RentalHandler) CreateRental (c *gin.Context) {
	var rental models.Rental

	if err := c.ShouldBindJSON(&rental); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

		if err := h.repo.CreateRental(rental); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err. Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Rental Created successfully"})
}

func (h *RentalHandler) GetAllRentals (c *gin.Context) {
	id := c.Param("carID")

	rentals, err := h.repo.GetRentalsByCarID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
	}

	if rentals == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "No Rentals found for this specific car"})
        return
    }

	c.JSON(http.StatusOK, rentals)
}

