package repository

import (
	"errors"
	"fleet-management/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type RenterRepository interface {
	CreateRenter(renter models.Renter) error
	GetRenter(id string) (*models.Renter, error)
}

type renterRepository struct {
	db *dynamodb.DynamoDB
}

func NewRenterRepository(db *dynamodb.DynamoDB) RenterRepository {
	return &renterRepository{db : db}
}

func (r *renterRepository) CreateRenter(renter models.Renter) error {
	item, err := dynamodbattribute.MarshalMap(renter)

		if err != nil {
		return err
	}

	
	if _, ok := item["ID"]; !ok {
        return errors.New("missing required field: ID")
    }

		input := &dynamodb.PutItemInput{
		TableName: aws.String("Renters"),
        Item:      item,
	}

	_, err = r.db.PutItem(input)
	return err
}

func (r *renterRepository) GetRenter(id string) (*models.Renter, error){
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Renters"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {S: aws.String(id)},
		},
	}

	result, err := r.db.GetItem(input)

	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	var renter models.Renter

	err = dynamodbattribute.UnmarshalMap(result.Item, &renter)
	return &renter, err
}