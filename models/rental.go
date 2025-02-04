package models

type Rental struct {
	CarID    string `json:"car_id" dynamodbav:"CarID"`
    RentalID string `json:"rental_id" dynamodbav:"RentalID"`
    RenterID string `json:"renter_id" dynamodbav:"RenterID"`
    StartDate string `json:"start_date" dynamodbav:"StartDate"`
    EndDate   string `json:"end_date" dynamodbav:"EndDate"`
    Status    string `json:"status" dynamodbav:"Status"`
}