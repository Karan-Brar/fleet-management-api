package repository

import (
	"errors"
	"fleet-management/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type RentalRepository interface {
	CreateRental(rental models.Rental) error
	GetRentalsByCarID(id string) ([]models.Rental, error)
}

type rentalsRepository struct {
    db *dynamodb.DynamoDB
}

func NewRentalRepository(db *dynamodb.DynamoDB) RentalRepository {
	return &rentalsRepository{db : db}
}

func (r *rentalsRepository) CreateRental(rental models.Rental) error{
	item, err := dynamodbattribute.MarshalMap(rental)

	if err != nil {
		return err
	}

	if _,ok := item["RentalID"]; !ok {
		return errors.New("missing required field: ID")
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("Rentals"),
		Item: item,
	}

	_, err = r.db.PutItem(input)
	return err
}

func (r *rentalsRepository) GetRentalsByCarID(carID string) ([]models.Rental, error) {
	    input := &dynamodb.QueryInput{
        TableName: aws.String("Rentals"),
        KeyConditionExpression: aws.String("CarID = :carID"),
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":carID": {S: aws.String(carID)},
        },
    }

	result, err := r.db.Query(input)
	if err != nil {
		return nil, err
	}

	var rentals []models.Rental
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &rentals)
	return rentals, err
}

