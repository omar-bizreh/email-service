package services

import (
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/omar-bizreh/email-service/aws/contracts"
)

// AwsService SDK Wraper
type AwsService struct {
	Session             contracts.SessionContainer
	client              *ssm.SSM
	isClientInitialized bool
}

// Init SSM Client
func (awsService *AwsService) initClient() {
	svc := ssm.New(&awsService.Session.AwsSession)
	awsService.client = svc
	awsService.isClientInitialized = true
}

// GetParam gets parameter from SSM Param Store
func (awsService *AwsService) GetParam(name string) (value *string, err error) {
	if !awsService.isClientInitialized {
		awsService.initClient()
	}
	param, err := awsService.client.GetParameter(&ssm.GetParameterInput{Name: &name})

	if err != nil {
		return nil, err
	}

	return param.Parameter.Value, nil
}
