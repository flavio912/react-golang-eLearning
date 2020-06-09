package email

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	//go get -u github.com/aws/aws-sdk-go
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/golang/glog"
)

// Session - AWS session
var Session *session.Session

// Initialize sets up an AWS session
func Initialize() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1"),
	})
	if err != nil {
		glog.Errorf("AWS error: %s", err.Error())
		panic("Could not setup aws connection")
	}
	Session = sess
}

// Sends a html + text email to one or more recipients
func SendRawMail(recipients []*string, subject string, htmlBody string, textBody string) error {
	var (
		charSet = "UTF-8"
		sender  = helpers.Config.AWS.SESSendAddress
	)

	if Session == nil {
		glog.Error("Unable to send email, aws must be initialized")
		return &errors.ErrWhileHandling
	}

	// Create an SES session.
	svc := ses.New(Session)

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: recipients,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(charSet),
					Data:    aws.String(htmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(charSet),
					Data:    aws.String(textBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(charSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sender),
	}

	// If testing don't actually send emails
	if helpers.Config.IsTesting {
		return nil
	}

	// Attempt to send the email.
	_, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				glog.Infof("Email rejected: %s", aerr.Error())
				return aerr
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				glog.Errorf("Mail from domain not allowed: %s", aerr.Error())
				return aerr
			default:
				glog.Errorf("Email send error: %s", aerr.Error())
				return aerr
			}
		} else {
			glog.Error(err.Error())
			return err
		}
	}

	return nil
}
