package email

import (
	"bytes"
	"html/template"
	"path"
	"runtime"

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

func SendEmail(templateName string, subject string, email string, params interface{}) error {
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "./templates/"+templateName)
	t, err := template.ParseFiles(filepath)
	if err != nil {
		glog.Errorf("Unable to create email template: %s", err.Error())
		return &errors.ErrWhileHandling
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, params)
	if err != nil {
		glog.Errorf("Unable to execute email template: %s", err.Error())
		return &errors.ErrWhileHandling
	}

	// Run in goroutine so as not to disrupt return of request
	go func() {
		err := SendRawMail([]*string{aws.String(email)}, subject, buf.String(), "") // TODO: add text version
		if err != nil {
			// TODO: Tell someone the email failed
			glog.Errorf("Unable to send finalise account email to: %s - error: %s", email, err.Error())
			return
		}
	}()
	return nil
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
