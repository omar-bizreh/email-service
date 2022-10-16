package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	eventcontracts "github.com/omar-bizreh/email-service/eventContracts"
)

func HandleRequest(ctx context.Context, name eventcontracts.SendEmailEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", name.From), nil
}

func main() {
	goEnvErr := godotenv.Load()

	if goEnvErr != nil {
		log.Fatal(goEnvErr.Error())
	}

	lambda.Start(HandleRequest)
}
