package config

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/joho/godotenv"
    "os"
)

func InitializeDynamoDB() (*dynamodb.DynamoDB, error) {
	godotenv.Load();

	creds := credentials.NewStaticCredentials(
        os.Getenv("AWS_ACCESS_KEY_ID"),
        os.Getenv("AWS_SECRET_ACCESS_KEY"),
        "", // No session token needed for basic credentials
    )

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Credentials: creds,
	})

	if err != nil {
		return nil, err
	}

	return dynamodb.New(sess), nil
}