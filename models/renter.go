package models

type Renter struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	CarID string `json:"car_id"`
}