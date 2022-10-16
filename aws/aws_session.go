package awswrapper

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/omar-bizreh/email-service/aws/contracts"
)

type SessionHandler struct {
	Region            string
	sessionInitalized bool
}

// InitSession Initialize new AWS sesion
func (receiver *SessionHandler) InitSession() (container contracts.SessionContainer) {
	configs := aws.NewConfig().WithRegion(receiver.Region)
	session := session.Must(session.NewSession(configs))

	awsContainer := &contracts.SessionContainer{AwsSession: *session}

	receiver.sessionInitalized = true

	return *awsContainer
}
