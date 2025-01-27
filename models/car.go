package models

type Car struct {
	ID string `json:"id" dynamodbav:"ID"`
	Make string `json:"make" dynamodbav:"Make"`
	Model string `json:"model" dynamodbav:"Model"`
	Year int `json:"year" dynamodbav:"Year"`
	VIN string `json:"vin" dynamodbav:"VIN"`
	Status string `json:"status" dynamodbav:"Status"`
}