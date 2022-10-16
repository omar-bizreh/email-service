package contracts

import "github.com/aws/aws-sdk-go/aws/session"

type SessionContainer struct {
	AwsSession session.Session
}
