package handlers

import (
	"fleet-management/models"
    "fleet-management/repository"
    "github.com/gin-gonic/gin"
    "net/http"
)

type RenterHandler struct {
	repo repository.RenterRepository
}


func NewRenterHandler(repo repository.RenterRepository) *RenterHandler {
	return &RenterHandler{repo: repo}
}

func (h *RenterHandler) CreateRenter (c *gin.Context) {
	var renter models.Renter

	if err := c.ShouldBindJSON(&renter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.CreateRenter(renter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err. Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Renter Created successfully"})
}

func (h *RenterHandler) GetRenter (c *gin.Context) {
	id := c.Param("id")

	renter, err := h.repo.GetRenter(id)
	    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if renter == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Renter not found"})
        return
    }

    c.JSON(http.StatusOK, renter)
}

