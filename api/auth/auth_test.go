package auth

import (
	"encoding/json"
	"testing"
)

func TestHashAndValidate(t *testing.T) {
	tables := []struct {
		password     string
		testPassword string
		isValid      bool
	}{
		{"password", "passwordd", false},
		{"test", "test", true},
		{"dasdgjfuf73j93ed38SF", "dasdgjfuf73j93ed38SF", true},
	}

	for _, table := range tables {
		hashedPassword, err := HashPassword(table.password)
		if err != nil {
			t.Error("Unable to generate hash")
		}

		validationErr := ValidatePassword(hashedPassword, table.testPassword)

		// If password valid but not accepted
		if table.isValid && validationErr != nil {
			t.Errorf("Password valid but not accepted: %s", table.password)
		}

		// If password is invalid but is accepted
		if !table.isValid && validationErr == nil {
			t.Errorf("Password incorrectly accepted: %s", table.password)
		}
	}

}

func TestGenAndValidateToken(t *testing.T) {
	secret := "thisisasupersecretpassword"

	type Claims struct {
		UUID    string `json:"uuid"`
		HasCake string `json:"hasCake"`
	}

	claims := &Claims{
		UUID:    "asdasd-asdadsad-asdad-asd",
		HasCake: "true",
	}

	token, err := GenerateToken(claims, 3, secret)
	if err != nil {
		t.Error("Error generating token")
	}

	var returnedClaims Claims
	validErr := ValidateToken(&returnedClaims, token, secret)
	if validErr != nil {
		t.Error(validErr.Error())
	}

	jsonRetClaims, _ := json.Marshal(returnedClaims)
	jsonGivenClaims, _ := json.Marshal(claims)
	if returnedClaims.UUID != claims.UUID {
		t.Errorf("UUID claims do not match: returned > %s , given > %s", string(jsonRetClaims), string(jsonGivenClaims))
	}

	if returnedClaims.HasCake != claims.HasCake {
		t.Errorf("hasCake claims do not match: returned > %s , given > %s", string(jsonRetClaims), string(jsonGivenClaims))
	}
}
