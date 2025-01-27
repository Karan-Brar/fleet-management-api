package repository

import (
	"errors"
	"fleet-management/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type CarRepository interface {
	CreateCar(car models.Car) error
	GetCar(id string) (*models.Car, error)
}

type carRepository struct {
	db *dynamodb.DynamoDB
}

func NewCarRepository(db *dynamodb.DynamoDB) CarRepository {
	return &carRepository{db : db}
}

func (r *carRepository) CreateCar(car models.Car) error {
	item, err := dynamodbattribute.MarshalMap(car)

	if err != nil {
		return err
	}

	if _, ok := item["ID"]; !ok {
        return errors.New("missing required field: ID")
    }

	input := &dynamodb.PutItemInput{
		TableName: aws.String("Cars"),
        Item:      item,
	}

	_, err = r.db.PutItem(input)
	return err
}

func (r *carRepository) GetCar(id string) (*models.Car, error){
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Cars"),
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

	var car models.Car

	err = dynamodbattribute.UnmarshalMap(result.Item, &car)
	return &car, err
}



