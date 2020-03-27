/*
Package auth deals with generating and checking JWTs
*/
package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func GenerateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"role":   "user",
	})
	tokenString, errToken := token.SignedString([]byte("verysecret"))
	if errToken != nil {
		return "", errors.New("jwt error: " + errToken.Error())
	}
	return tokenString, nil
}
