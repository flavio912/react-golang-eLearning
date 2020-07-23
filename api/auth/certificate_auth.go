package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
)

// CertificateClaims
type CertificateClaims struct {
	jwt.StandardClaims
	HistoricalCourseUUID gentypes.UUID
}

// GenerateCertificateToken creates a certificate token allowed to get the information required to build a certificate
func GenerateCertificateToken(historicalCourseUUID gentypes.UUID) (string, error) {
	claims := CertificateClaims{
		HistoricalCourseUUID: historicalCourseUUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(10) * time.Minute).Unix(),
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

// ValidateCertificateToken - returns the historicalCourseUUID if valid
func ValidateCertificateToken(token string) (gentypes.UUID, error) {
	var claims CertificateClaims
	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(helpers.Config.Jwt.UploadsSecret), nil
	})

	if err != nil {
		glog.Infof("Token validation failed: %s - TOKEN: %s", err.Error(), token)
		return gentypes.UUID{}, &errors.ErrUploadTokenInvalid
	}

	if !tkn.Valid {
		glog.Infof("Token is invalid: %s", token)
		return gentypes.UUID{}, &errors.ErrUploadTokenInvalid
	}

	return claims.HistoricalCourseUUID, nil
}
