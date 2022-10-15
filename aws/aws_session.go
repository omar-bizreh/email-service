package awswrapper

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type SessionHandler struct {
	region string
}

// InitSession Initialize new AWS sesion
func (receiver *SessionHandler) InitSession() session.Session {
	configs := aws.NewConfig().WithRegion(receiver.region)
	session := session.Must(session.NewSession(configs))

	return *session
}
