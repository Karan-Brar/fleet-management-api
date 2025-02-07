package models

type Car struct {
	ID string `json:"id" dynamodbav:"ID"`
	Make string `json:"make" dynamodbav:"Make"`
	Model string `json:"model" dynamodbav:"Model"`
	Year int `json:"year" dynamodbav:"Year"`
	VIN string `json:"vin" dynamodbav:"VIN"`
	Status string `json:"status" dynamodbav:"Status"`
	DailyRate float64 `json:"status" dynamodbav:"DailyRate`
	Features []string `json:"features" dynamodbav:"Features"`
	Mileage int `json:"mileage" dynamodbav:"Mileage"`
	Transmission string   `json:"transmission" dynamodbav:"Transmission"`
	ImageURL string   `json:"image_url" dynamodbav:"ImageURL"`
}