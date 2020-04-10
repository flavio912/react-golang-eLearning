package middleware_test

import (
	"reflect"
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

// func _checkStrucTypeEquals(a interface{}, b interface{}, fields []string) (bool, error) {
// 	aRef := reflect.ValueOf(a).Elem()
// 	bRef := reflect.ValueOf(b).Elem()

// 	for _, f := range fields {
// 		aVal := aRef.FieldByName(f)
// 		if !aVal.IsValid() {
// 			return false, fmt.Errorf("No such feild in a: %s", f)
// 		}

// 		bVal := bRef.FieldByName(f)
// 		if !bVal.IsValid() {
// 			return false, fmt.Errorf("No such feild in a: %s", f)
// 		}

// 		if reflect.DeepEqual(aVal, bVal) {
// 			fmt.Printf("[%s] %s != %s", f, aVal, bVal)
// 			return false, fmt.Errorf("[%s] %s != %s", f, aVal, bVal)
// 		}
// 	}

// 	return true, nil
// }

func TestCreateAdminUser(t *testing.T) {
	newAdmin := &models.Admin{
		Email:     "test@test.com",
		Password:  "test",
		FirstName: "FName",
		LastName:  "LName",
	}
	err := database.GormDB.Create(newAdmin)
	if err.Error != nil {
		t.Fatal(err.GetErrors())
	}

	admin := &models.Admin{}
	q := database.GormDB.Where("email = ?", "test@test.com").First(admin)
	if q.Error != nil {
		if q.RecordNotFound() {
			t.Error("GORM didn't create a user")
		}

		t.Errorf("GORM Errored:\n%#v", q.GetErrors())
	}
	// t.Logf("created admin: %v", admin)
}

func TestAddAdmin(t *testing.T) {
	// fake grant
	grant := &middleware.Grant{auth.UserClaims{}, true, true, true}

	newAdmin := gentypes.AddAdminInput{
		Email:     "admi1n@admin.com",
		Password:  "aderrmin123",
		FirstName: "Admin",
		LastName:  "Man",
	}
	createdAdmin, _ := grant.AddAdmin(newAdmin)

	// check, err := _checkStrucTypeEquals(&newAdmin, &createdAdmin, []string{"Email", "FirstName", "LastName"})
	// if err != nil {
	// 	t.Errorf("Check failed %sv", err)
	// }
	// if !check {
	// 	t.Error("Feilds not equal")
	// }

	if !reflect.DeepEqual(gentypes.Admin{
		UUID:      createdAdmin.UUID,
		Email:     newAdmin.Email,
		FirstName: newAdmin.FirstName,
		LastName:  newAdmin.LastName,
	}, createdAdmin) {
		t.Error("didn't create an equal admin")
	}
}
