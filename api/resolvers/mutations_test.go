package resolvers_test

import (
	"context"
	"encoding/json"
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/testhelpers"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

func TestAdminLogin(t *testing.T) {
	ctx := context.Background()
	res := schema.Exec(
		ctx,
		`mutation {
			adminLogin(input:{email: "test123@test.com", password: "iamasuperadmin"}) {
				token
			}
		} `,
		"",
		map[string]interface{}{},
	)

	assert.Nil(t, res.Errors)

	// test that the token works
	var data map[string]interface{}
	err := json.Unmarshal(res.Data, &data)
	assert.Nil(t, err)

	token := data["adminLogin"].(map[string]interface{})["token"].(string)
	assert.Nil(t, err, "Error converting the token: \n%#v", data)

	// use the token to auth
	grant, err := middleware.Authenticate(token)
	assert.Nil(t, err)
	assert.Equal(t, middleware.Grant{
		Claims: auth.UserClaims{
			UUID:    gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
			Company: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000000"),
			Role:    auth.AdminRole,
		},
		IsAdmin:    true,
		IsManager:  false,
		IsDelegate: false,
	}, *grant)
}

func TestUpdateAdmin(t *testing.T) {
	testhelpers.RunTests(t, []*testhelpers.Test{
		{
			Context: adminContext,
			Schema:  schema,
			Query: `
				mutation {
					updateAdmin(input: {
						uuid: "00000000-0000-0000-0000-000000000002"
						firstName: "edfadd",
						lastName: "dsa"
					}) {
						uuid
						email
						firstName
						lastName
					}
				}
			`,
			ExpectedResult: `
				{
					"updateAdmin": {
					  "uuid": "00000000-0000-0000-0000-000000000002",
					  "email": "steve@wombat.com",
						"firstName": "edfadd",
						"lastName": "dsa"
					}
				}
			`,
		},
	})
}

func TestUpdateManager(t *testing.T) {
	testhelpers.RunTests(t, []*testhelpers.Test{
		{
			Name:    "Update Some Fields",
			Context: adminContext,
			Schema:  schema,
			Query: `
				mutation {
					updateManager(input: {
						uuid: "00000000-0000-0000-0000-000000000002"
						firstName: "edfadd",
						lastName: "dsa"
					}) {
						uuid
						email
						firstName
						lastName
					}
				}
			`,
			ExpectedResult: `
				{
					"updateManager": {
					  "uuid": "00000000-0000-0000-0000-000000000002",
					  "email": "ver@diff.com",
						"firstName": "edfadd",
						"lastName": "dsa"
					}
				}
			`,
		},
		{
			Name:    "Update All Fields",
			Context: adminContext,
			Schema:  schema,
			Query: `
				mutation {
					updateManager(input: {
						uuid: "00000000-0000-0000-0000-000000000002"
						email: "dsa@das.dfa",
						firstName: "asdf",
						lastName: "fdsa",
						telephone: "07886515216",
						jobTitle: "overlord",
					}) {
						uuid
						email
						firstName
						lastName
						telephone
						jobTitle
					}
				}
			`,
			ExpectedResult: `
				{
					"updateManager": {
					  "uuid": "00000000-0000-0000-0000-000000000002",
					  "email": "dsa@das.dfa",
						"firstName": "asdf",
						"lastName": "fdsa",
						"telephone": "07886515216",
						"jobTitle": "overlord"
					}
				}
			`,
		},
		{
			Name:    "UUID does not exist",
			Context: adminContext,
			Schema:  schema,
			Query: `
				mutation {
					updateManager(input: {
						uuid: "00000000-0000-0000-0000-000000000000"
					}) {
						uuid
					}
				}
			`,
			ExpectedResult: `
				{
					"updateManager": null
				}
			`,
			ExpectedErrors: []testhelpers.TestQueryError{
				{
					ResolverError: &errors.ErrManagerNotFound,
					Path:          []interface{}{"updateManager"},
				},
			},
		},
		{
			Name:    "Fail validation",
			Context: adminContext,
			Schema:  schema,
			Query: `
				mutation {
					updateManager(input: {
						uuid: "00000000-0000-0000-0000-000000000000"
						firstName: "123!"
						email: "not^%!£$*"
					}) {
						uuid
					}
				}
			`,
			ExpectedResult: `
				{
					"updateManager": null
				}
			`,
			ExpectedErrors: []testhelpers.TestQueryError{
				{
					Message: helpers.StringPointer("Email: not^%!£$* does not validate as email;FirstName: 123! does not validate as alpha"),
					Path:    []interface{}{"updateManager"},
				},
			},
		},
	})
}

func TestDeleteAdmin(t *testing.T) {
	testhelpers.RunTests(t, []*testhelpers.Test{
		{
			Context: adminContext,
			Schema:  schema,
			Query: `
				mutation {
					deleteAdmin(input: {
						uuid: "00000000-0000-0000-0000-000000000002"
					})
				}
			`,
			ExpectedResult: `
				{
					"deleteAdmin": true
				}
			`,
		},
	})
}
