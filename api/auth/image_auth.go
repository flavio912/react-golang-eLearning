package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
)

// ImageClaims are verification claims for presigned URLs
type ImageClaims struct {
	jwt.StandardClaims
	Key        string
	UploadType string
}

// GenerateUploadToken creates a jwt to guarantee the key given is the one
// that gets returned as the uploaded image. The uploadType is key to security
// make sure that you specify a different uploadType for different user level actions
func GenerateUploadToken(key string, uploadType string) (string, error) {
	claims := ImageClaims{
		Key:        key,
		UploadType: uploadType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(2) * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, errToken := token.SignedString([]byte(helpers.Config.Jwt.UploadsSecret))
	if errToken != nil {
		glog.Errorf("Unable to generate token: %s", errToken.Error())
		return "", &errors.ErrGeneratingToken
	}

	return tokenString, nil
}

// ValidateUploadToken - Checks the signature and uploadType
// on a token and returns claims if it is valid
func ValidateUploadToken(token string, uploadType string) (ImageClaims, error) {
	var claims ImageClaims
	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(helpers.Config.Jwt.UploadsSecret), nil
	})

	if err != nil {
		glog.Infof("Token validation failed: %s - TOKEN: %s", err.Error(), token)
		return ImageClaims{}, &errors.ErrUploadTokenInvalid
	}

	if !tkn.Valid {
		glog.Infof("Token is invalid: %s", token)
		return ImageClaims{}, &errors.ErrUploadTokenInvalid
	}

	if claims.UploadType != uploadType {
		glog.Infof("Upload token type '%s' doesn't match expected type '%s'", claims.UploadType, uploadType)
		return ImageClaims{}, &errors.ErrUploadTokenInvalid
	}

	return claims, nil
}
