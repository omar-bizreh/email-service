package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	awswrapper "github.com/omar-bizreh/email-service/aws"
	eventcontracts "github.com/omar-bizreh/email-service/eventContracts"
	"github.com/omar-bizreh/email-service/services"
)

func HandleRequest(ctx context.Context, name eventcontracts.SendEmailEvent) (string, error) {

	awsSession := awswrapper.SessionHandler{Region: os.Getenv("region")}
	dynamo := services.AwsDynamo{Session: awsSession.InitSession()}

	dynamo.PutItem(&name)

	return "Success", nil
}

func main() {
	goEnvErr := godotenv.Load()

	if goEnvErr != nil {
		log.Fatal(goEnvErr.Error())
	}

	lambda.Start(HandleRequest)
}
