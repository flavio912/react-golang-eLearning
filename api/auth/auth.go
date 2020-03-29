/*
Package auth deals with generating and checking JWTs
*/
package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

// ValidateToken - Checks the signature on a token and returns claims
func ValidateToken(claims jwt.Claims, token string, secret string) (jwt.Claims, error) {
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return claims, err
	}
	if !tkn.Valid {
		return claims, errors.New("Token invalid")
	}
	return claims, nil
}

/*GenerateToken - TAKE CARE BEFORE USING DIRECTLY @temmerson
Generates a jwt token given a secret and some claims
*/
func GenerateToken(claims jwt.Claims, secret string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, errToken := token.SignedString([]byte(secret))
	if errToken != nil {
		return "", errors.New("jwt error: " + errToken.Error())
	}
	return tokenString, nil
}
