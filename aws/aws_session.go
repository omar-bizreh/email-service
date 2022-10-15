package awswrapper

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type SessionHandler struct {
	region string
}


func (receiver *SessionHandler) InitSession() {
	configs := aws.NewConfig().WithRegion("eu-west-1")
	session.Must(session.NewSession(configs))
}