package services

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"

	awswrapper "github.com/omar-bizreh/email-service/aws"
)

// MailService Send emails
type MailService struct {
}

func getSenderAppKey() (*string, error) {
	awsSession := awswrapper.SessionHandler{Region: os.Getenv("AWS_REGION")}
	sessionContainer := awsSession.InitSession()
	awsClient := AwsService{Session: sessionContainer}

	param, err := awsClient.GetParam("SENDER_CREDETNAILS")

	return param, err
}

// SendEmail sends email DUH!
func (service *MailService) SendEmail(to string, subject string, body string) {
	from := os.Getenv("supportEmail")
	strBuilder := new(strings.Builder)

	strBuilder.WriteString("To: " + to + "\n")
	strBuilder.WriteString("From: Support" + "<" + from + ">\n")
	strBuilder.WriteString("Subject:" + subject + " \n")
	strBuilder.WriteString(body)

	key, err := getSenderAppKey()

	if err != nil {
		log.Fatal(err.Error())
	}

	smtpHost := os.Getenv("SMTP")
	smtpPort := os.Getenv("PORT")

	err = smtp.SendMail(smtpHost+":"+smtpPort, smtp.PlainAuth("", from, *key, smtpHost), from, []string{to}, []byte(strBuilder.String()))
	if err != nil {
		fmt.Println(err.Error())
	}
}
