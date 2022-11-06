package services

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/omar-bizreh/email-service/aws/contracts"
	evenetContracts "github.com/omar-bizreh/email-service/eventContracts"
)

type AwsDynamo struct {
	Session             contracts.SessionContainer
	client              *dynamodb.DynamoDB
	isClientInitialized bool
}

// initClient initializes DynamoDB Client
func (handler *AwsDynamo) initClient() *dynamodb.DynamoDB {
	client := dynamodb.New(&handler.Session.AwsSession)
	handler.isClientInitialized = true
	return client
}

func (handler *AwsDynamo) PutItem(input *evenetContracts.SendEmailEvent) error {
	if !handler.isClientInitialized {
		handler.initClient()
	}

	av, err := dynamodbattribute.MarshalMap(input)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, errr := handler.client.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("inquiries"),
		Item:      av,
	})

	if errr != nil {
		log.Fatal(err.Error())
	}
	return nil
}
