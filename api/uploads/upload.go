package uploads

import (
	"fmt"
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
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
	return
}

/* GenerateUploadURL creates an aws url that allows users to upload files and images
Returns (presignedURL, successToken, error). TAKE CARE, a specific uploadIdent should
be given for each purpose so that users cannot use successTokens from one upload
to validate an upload at another endpoint. There must be a different uploadIdent for
every different acceptedTypes, urlBase and maxFileSize
*/
func GenerateUploadURL(
	imageType string,
	contentLength int32,
	acceptedTypes []string,
	maxFileSize int32,
	urlBase string,
	uploadIdent string,
) (string, string, error) {
	// Check contentLength is within limits
	if contentLength > maxFileSize {
		return "", "", &errors.ErrFileTooLarge
	}

	// Check image type is valid
	var allowed bool
	for _, tpe := range acceptedTypes {
		if tpe == imageType {
			allowed = true
		}
	}
	if !allowed {
		return "", "", &errors.ErrUnauthorized
	}

	// Generate a secure url
	str, err := auth.GenerateRandomString(40)
	if err != nil {
		glog.Errorf("Unable to generate random string: %s", err.Error())
		return "", "", &errors.ErrWhileHandling
	}

	imageKey := fmt.Sprintf("%s/%s.%s", urlBase, str, imageType)
	svc := s3.New(Session)
	// TODO: Setup CORS on S3 to only allow requests from our URL (not sure what our url is atm)
	// Create S3 service client
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket:        aws.String(helpers.Config.AWS.UploadsBucket),
		Key:           aws.String(imageKey),
		ContentLength: aws.Int64(int64(contentLength)),
	})
	url, err := req.Presign(2 * time.Minute)
	if err != nil {
		glog.Errorf("Unable to generate presigned url: %s", err.Error())
		return "", "", &errors.ErrWhileHandling
	}

	// Generate a success token
	successToken, tokenErr := auth.GenerateUploadToken(imageKey, uploadIdent)
	if tokenErr != nil {
		return "", "", tokenErr
	}

	return url, successToken, err
}

// VerifyUploadSuccess checks if the given profile success token is valid
func VerifyUploadSuccess(token string, uploadIdent string) (string, error) {
	// Verify token is valid
	claims, err := auth.ValidateUploadToken(token, uploadIdent)
	if err != nil {
		return "", err
	}

	// Check that the image actually exists
	svc := s3.New(Session)
	_, err = svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(helpers.Config.AWS.UploadsBucket),
		Key:    aws.String(claims.Key),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case "NotFound":
				return "", &errors.ErrNotUploaded
			}
		}
		glog.Errorf("Image has not been uploaded: %s", err.Error())
		return "", &errors.ErrWhileHandling
	}

	return claims.Key, nil
}

// GetImgixURL takes the s3 key and adds the imgix url to get a full URL
func GetImgixURL(key string) string {
	//TODO: Add checks on the key so we
	// always get a valid URL (or at least so we can log errors)
	return helpers.Config.Imgix.BaseURL + key
}
