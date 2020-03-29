/*
Package auth deals with generating and checking JWTs
*/
package auth

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword Hashes and salts a given string
func HashPassword(pwd string) (string, error) {

	passwordInBytes := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(passwordInBytes, 12)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(hash), nil
}

// ValidatePassword - Compare password and a hash
func ValidatePassword(hash string, password string) error {
	// Compare given password to the hashed value
	byteHash := []byte(hash)
	bytePassword := []byte(password)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		return errors.New("Password and hash do not match")
	}
	return nil
}

type Claims struct {
	jwt.StandardClaims
	Claims interface{} `json:"claims"`
}

// ValidateToken - Checks the signature on a token and returns claims
func ValidateToken(claims interface{}, token string, secret string) error {
	newClaims := &Claims{
		Claims: claims,
	}
	tkn, err := jwt.ParseWithClaims(token, newClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return err
	}
	if !tkn.Valid {
		return errors.New("Token invalid")
	}

	claims = newClaims
	return nil
}

/*GenerateToken - TAKE CARE BEFORE USING DIRECTLY @temmerson
Generates a jwt token given a secret and some claims
*/
func GenerateToken(claims interface{}, expiresInHours float64, secret string) (string, error) {
	finalClaims := &Claims{
		Claims: claims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expiresInHours) * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, finalClaims)
	tokenString, errToken := token.SignedString([]byte(secret))
	if errToken != nil {
		return "", errors.New("jwt error: " + errToken.Error())
	}

	return tokenString, nil
}
