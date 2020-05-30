package email

import (
	"bytes"
	"html/template"
	"path"
	"runtime"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/golang/glog"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
)

func SendFinaliseAccountEmail(token string, fName string, email string) error {
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "./templates/finaliseAccount.html")
	t, err := template.ParseFiles(filepath)
	if err != nil {
		glog.Errorf("Unable to create email template: %s", err.Error())
		return &errors.ErrWhileHandling
	}
	var buf bytes.Buffer

	data := struct {
		Token string
		FName string
	}{
		Token: token,
		FName: fName,
	}

	err = t.Execute(&buf, data)
	if err != nil {
		glog.Errorf("Unable to execute email template: %s", err.Error())
		return &errors.ErrWhileHandling
	}

	// Run in goroutine so as not to disrupt return of request
	go func() {
		err := SendRawMail([]*string{aws.String(email)}, "Finalise your TTC account", buf.String(), "") // TODO: add text version
		if err != nil {
			// TODO: Tell someone the email failed
			glog.Errorf("Unable to send finalise account email to: %s - error: %s", email, err.Error())
			return
		}
	}()
	return nil
}
