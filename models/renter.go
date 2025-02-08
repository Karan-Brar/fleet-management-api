package models

type Renter struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number" dynamodbav:"PhoneNumber"`
	Address string `json:"address" dynamodbav:"Address"`
	DrivingLicenseNumber string `json:"driving_license_number" dynamodbav:"DrivingLicenseNumber"`
	RentalHistory []string `json:"rental_history" dynamodbav:"RentalHistory"`
}