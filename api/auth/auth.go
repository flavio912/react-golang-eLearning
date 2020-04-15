/*
Package auth deals with generating and checking JWTs
*/
package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
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

// Role - A role type
type Role string

const (
	// AdminRole - The admin role jwt mapping
	AdminRole Role = "admin"
	// ManagerRole - The manager role jwt mapping
	ManagerRole Role = "manager"
	// DelegateRole - The delegate role jwt mapping
	DelegateRole Role = "delegate"
)

// RoleToString gets the string representation of a role
func RoleToString(role Role) string {
	switch role {
	case AdminRole:
		return "admin"
	case ManagerRole:
		return "manager"
	case DelegateRole:
		return "delegate"
	default:
		return ""
	}
}

// UserClaims - Claims other than the default added to the JWT
type UserClaims struct {
	UUID    string
	Company string
	Role    Role
}

type trueClaims struct {
	jwt.StandardClaims
	Claims UserClaims `json:"claims"`
}

// ValidateToken - Checks the signature on a token and returns claims
func ValidateToken(token string) (UserClaims, error) {

	claims := &trueClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(helpers.Config.Jwt.Secret), nil
	})

	if err != nil {
		return UserClaims{}, err
	}
	if !tkn.Valid {
		return UserClaims{}, errors.New("Token invalid")
	}

	return claims.Claims, nil
}

/*GenerateToken - TAKE CARE BEFORE USING DIRECTLY @temmerson
Generates a jwt token given a secret and some claims
*/
func GenerateToken(claims UserClaims, expiresInHours float64) (string, error) {
	finalClaims := &trueClaims{
		Claims: claims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expiresInHours) * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, finalClaims)
	tokenString, errToken := token.SignedString([]byte(helpers.Config.Jwt.Secret))
	if errToken != nil {
		return "", errors.New("jwt error: " + errToken.Error())
	}

	return tokenString, nil
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
