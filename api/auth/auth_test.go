package auth

import (
	"encoding/json"
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
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
	helpers.LoadConfig("../config.yml")

	claims := UserClaims{
		UUID: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000000"),
		Role: AdminRole,
	}

	token, err := GenerateToken(claims, 3)
	if err != nil {
		t.Error("Error generating token")
	}

	returnedClaims, validErr := ValidateToken(token)
	if validErr != nil {
		t.Error(validErr.Error())
	}

	jsonRetClaims, _ := json.Marshal(returnedClaims)
	jsonGivenClaims, _ := json.Marshal(claims)
	if returnedClaims.UUID != claims.UUID {
		t.Errorf("UUID claims do not match: returned > %s , given > %s", string(jsonRetClaims), string(jsonGivenClaims))
	}

	if returnedClaims.Role != claims.Role {
		t.Errorf("Role claims do not match: returned > %s , given > %s", string(jsonRetClaims), string(jsonGivenClaims))
	}
}
