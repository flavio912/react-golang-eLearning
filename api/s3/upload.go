package s3

import (
	"fmt"
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
)

var validImageTypes = []string{"png", "jpg"}
var maxFileSize = int64(20000000) // 20MB in bytes

// generateProfileUrl creates an aws url that allows users to upload their profile image
// Note: the following header must be added to the request: `x-amz-acl: public-read`
func generateProfileURL(imageType string, contentLength int64) (string, error) {
	// Check contentLength is within limits
	if contentLength > maxFileSize {
		return "", &errors.ErrFileTooLarge
	}

	// Check image type is valid
	var allowed bool
	for _, tpe := range validImageTypes {
		if tpe == imageType {
			allowed = true
		}
	}
	if !allowed {
		return "", &errors.ErrUnauthorized
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)

	// Generate a secure url
	str, err := auth.GenerateRandomString(40)
	if err != nil {
		glog.Errorf("Unable to generate random string: %s", err.Error())
		return "", &errors.ErrWhileHandling
	}

	// TODO: Setup CORS on S3 to only allow requests from our URL (not sure what our url is atm)
	// Create S3 service client
	svc := s3.New(sess)
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket:        aws.String(helpers.Config.AWS.UploadsBucket),
		Key:           aws.String(fmt.Sprintf("profile/%s.%s", str, imageType)),
		ACL:           aws.String("public-read"),
		ContentLength: aws.Int64(contentLength),
	})
	url, err := req.Presign(1 * time.Minute)
	if err != nil {
		glog.Errorf("Unable to generate presigned url: %s", err.Error())
		return "", &errors.ErrWhileHandling
	}

	return url, err
}
