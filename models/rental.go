package models

type Rental struct {
    ID string `json:"id" dynamodbav:"RentalID"`
	CarID    string `json:"car_id" dynamodbav:"CarID"`
    RenterID string `json:"renter_id" dynamodbav:"RenterID"`
    StartDate string `json:"start_date" dynamodbav:"StartDate"`
    EndDate   string `json:"end_date" dynamodbav:"EndDate"`
    Status    string `json:"status" dynamodbav:"Status"` // e.g., "active", "completed", "cancelled"
    TotalCost float64  `json:"total_cost" dynamodbav:"TotalCost"`
    AdditionalDrivers []string `json:"additional_drivers" dynamodbav:"AdditionalDrivers"`
    DamageReport string   `json:"damage_report" dynamodbav:"DamageReport"`
}