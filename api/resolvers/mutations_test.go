package resolvers_test

import (
	"encoding/json"
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers/gqltest"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

func TestAdminLogin(t *testing.T) {
	t.Run("Must auth and return correct grant", func(t *testing.T) {
		res := schema.Exec(
			defaultContext,
			`mutation {
				adminLogin(input:{email: "test123@test.com", password: "iamasuperadmin"}) {
					token
				}
			}`,
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
	})

	t.Run("must fail properly", func(t *testing.T) {
		gqltest.RunTests(t, []*gqltest.Test{
			{
				Name:    "bad email",
				Context: defaultContext,
				Schema:  schema,
				Query: `
					mutation {
						adminLogin(input:{email: "doesnot@exist.com", password: "iamasuperadmin"}) {
							token
						}
					}
				`,
				ExpectedResult: `{"adminLogin":null}`,
				ExpectedErrors: []gqltest.TestQueryError{
					{
						Path:          []interface{}{"adminLogin"},
						ResolverError: &errors.ErrAdminNotFound,
					},
				},
			},
			{
				Name:    "bad password",
				Context: defaultContext,
				Schema:  schema,
				Query: `
					mutation {
						adminLogin(input:{email: "test123@test.com", password: "notmypass"}) {
							token
						}
					}
				`,
				ExpectedResult: `{"adminLogin":null}`,
				ExpectedErrors: []gqltest.TestQueryError{
					{
						Path:          []interface{}{"adminLogin"},
						ResolverError: &errors.ErrAuthFailed,
					},
				},
			},
		})
	})
}

func TestManagerLogin(t *testing.T) {
	prepareTestDatabase()

	t.Run("Must auth and return correct grant", func(t *testing.T) {
		res := schema.Exec(
			defaultContext,
			`mutation {
				managerLogin(input:{email: "man@managers.com", password: "iamamanager"}) {
					token
				}
			}`,
			"",
			map[string]interface{}{},
		)

		assert.Nil(t, res.Errors)

		// test that the token works
		var data map[string]interface{}
		err := json.Unmarshal(res.Data, &data)
		assert.Nil(t, err)

		token := data["managerLogin"].(map[string]interface{})["token"].(string)
		assert.Nil(t, err, "Error converting the token: \n%#v", data)

		// use the token to auth
		grant, err := middleware.Authenticate(token)
		assert.Nil(t, err)
		assert.Equal(t, middleware.Grant{
			Claims: auth.UserClaims{
				UUID:    gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
				Company: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
				Role:    auth.ManagerRole,
			},
			IsAdmin:    false,
			IsManager:  true,
			IsDelegate: false,
		}, *grant)
	})

	t.Run("must fail properly", func(t *testing.T) {
		gqltest.RunTests(t, []*gqltest.Test{
			{
				Name:    "bad email",
				Context: defaultContext,
				Schema:  schema,
				Query: `
					mutation {
						managerLogin(input:{email: "doesnot@exist.com", password: "iamamanager"}) {
							token
						}
					}
				`,
				ExpectedResult: `{"managerLogin":null}`,
				ExpectedErrors: []gqltest.TestQueryError{
					{
						Path:          []interface{}{"managerLogin"},
						ResolverError: &errors.ErrUserNotFound,
					},
				},
			},
			{
				Name:    "bad password",
				Context: defaultContext,
				Schema:  schema,
				Query: `
					mutation {
						managerLogin(input:{email: "man@managers.com", password: "notmypass"}) {
							token
						}
					}
				`,
				ExpectedResult: `{"managerLogin":null}`,
				ExpectedErrors: []gqltest.TestQueryError{
					{
						Path:          []interface{}{"managerLogin"},
						ResolverError: &errors.ErrAuthFailed,
					},
				},
			},
		})
	})
}

func TestCreateManager(t *testing.T) {
	prepareTestDatabase()

	t.Run("should successfully create a manager", func(t *testing.T) {
		gqltest.RunTests(t, []*gqltest.Test{{
			Name:    "create manager",
			Context: adminContext,
			Schema:  schema,
			Query: `
				mutation {
					createManager(input: {
						companyUUID: "00000000-0000-0000-0000-000000000001"
						email:     "asdf@fdsa.com"
						firstName: "James"
						lastName:  "Bay"
						telephone: "07894561230"
						jobTitle:  "Overlord"
						password: "realpassword"
					}) {
						firstName
						lastName
						email
						jobTitle
						telephone
					}
				}
			`,
			ExpectedResult: `
				{
					"createManager":{
						"firstName":"James",
						"lastName":"Bay",
						"email":"asdf@fdsa.com",
						"jobTitle":"Overlord",
						"telephone":"07894561230"
					}
				}
			`,
		}})

		// check you can auth with the new creds
		_, err := middleware.GetManagerAccessToken("asdf@fdsa.com", "realpassword")
		assert.Nil(t, err)
	})

	t.Run("it should fail correctly", func(t *testing.T) {
		gqltest.RunTests(t, []*gqltest.Test{
			{
				Name:    "doesn't validate",
				Context: adminContext,
				Schema:  schema,
				Query: `
				mutation {
					createManager(input: {
						companyUUID: "00000000-0000-0000-0000-000000000001"
						email:     "not an email"
						firstName: "James"
						lastName:  "Bay"
						telephone: "asdf"
						jobTitle:  "Overlord"
						password: "1"
					}) {
						uuid
					}
				}
			`,
				ExpectedResult: `{"createManager":null}`,
				ExpectedErrors: []gqltest.TestQueryError{
					{
						Message: helpers.StringPointer("Email: not an email does not validate as email;" +
							"Telephone: asdf does not validate as numeric;" +
							"Password: 1 does not validate as stringlength(5|30)",
						),
						Path: []interface{}{"createManager"},
					},
				},
			},
			{
				Name:    "must be authed",
				Context: defaultContext,
				Schema:  schema,
				Query: `
				mutation {
					createManager(input: {
						companyUUID: "00000000-0000-0000-0000-000000000001"
						email:     "e@mail.com"
						firstName: "James"
						lastName:  "Bay"
						telephone: "07932446835"
						jobTitle:  "Overlord"
						password: "123456789"
					}) {
						uuid
					}
				}
			`,
				ExpectedResult: `{"createManager":null}`,
				ExpectedErrors: []gqltest.TestQueryError{
					{
						ResolverError: &errors.ErrUnauthorized,
						Path:          []interface{}{"createManager"},
					},
				},
			},
			{
				Name:    "must be unique email",
				Context: adminContext,
				Schema:  schema,
				Query: `
				mutation {
					createManager(input: {
						companyUUID: "00000000-0000-0000-0000-000000000001"
						email:     "man@managers.com"
						firstName: "James"
						lastName:  "Bay"
						telephone: "07932446835"
						jobTitle:  "Overlord"
						password: "123456789"
					}) {
						uuid
					}
				}
			`,
				ExpectedResult: `{"createManager":null}`,
				ExpectedErrors: []gqltest.TestQueryError{
					{
						ResolverError: &errors.ErrUserExists,
						Path:          []interface{}{"createManager"},
					},
				},
			},
		})
	})
}

func TestUpdateManager(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
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
			ExpectedErrors: []gqltest.TestQueryError{
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
			ExpectedErrors: []gqltest.TestQueryError{
				{
					Message: helpers.StringPointer("Email: not^%!£$* does not validate as email;FirstName: 123! does not validate as alpha"),
					Path:    []interface{}{"updateManager"},
				},
			},
		},
	})

	accessTest(
		t, schema, accessTestOpts{
			Query: `
				mutation {
					updateManager(input: {
						uuid: "00000000-0000-0000-0000-000000000003"
					}) {
						uuid
					}
				}
			`,
			Path:            []interface{}{"updateManager"},
			MustAuth:        true,
			AdminAllowed:    true,
			ManagerAllowed:  false,
			DelegateAllowed: false,
		},
	)
}

func TestDeleteManager(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Name:    "Delete manager",
			Context: adminContext,
			Schema:  schema,
			Query: `
				mutation {
					deleteManager(input: {
						uuid: "00000000-0000-0000-0000-000000000002"
					})
				}
			`,
			ExpectedResult: `{"deleteManager": true}`,
		},
		{
			Name:    "Check deleted manager was deleted",
			Context: adminContext,
			Schema:  schema,
			Query: `
				mutation {
					deleteManager(input: {
						uuid: "00000000-0000-0000-0000-000000000002"
					})
				}
			`,
			ExpectedResult: "null",
			ExpectedErrors: []gqltest.TestQueryError{{
				ResolverError: &errors.ErrUserNotFound,
				Path:          []interface{}{"deleteManager"},
			}},
		},
	})

	accessTest(t, schema, accessTestOpts{
		Query: `
				mutation {
					deleteManager(input: {
						uuid: "00000000-0000-0000-0000-000000000001"
					})
				}
			`,
		Path:            []interface{}{"deleteManager"},
		MustAuth:        true,
		AdminAllowed:    true,
		ManagerAllowed:  true,
		DelegateAllowed: false,
		CleanDB:         true,
	})
}

func TestCreateAdmin(t *testing.T) {
	prepareTestDatabase()

	t.Run("should successfully create a manager", func(t *testing.T) {
		gqltest.RunTests(t, []*gqltest.Test{{
			Name:    "create manager",
			Context: adminContext,
			Schema:  schema,
			Query: `
				mutation {
					createAdmin(input: {
						email:     "adminman@fdsa.com"
						firstName: "James"
						lastName:  "May"
						password: "realpassword"
					}) {
						firstName
						lastName
						email
					}
				}
			`,
			ExpectedResult: `
				{
					"createAdmin": {
						"firstName":"James",
						"lastName":"May",
						"email":"adminman@fdsa.com"
					}
				}
			`,
		}})

		// check you can auth with the new creds
		_, err := middleware.GetAdminAccessToken("adminman@fdsa.com", "realpassword")
		assert.Nil(t, err)
	})

	t.Run("it should fail correctly", func(t *testing.T) {
		gqltest.RunTests(t, []*gqltest.Test{
			{
				Name:    "doesn't validate",
				Context: adminContext,
				Schema:  schema,
				Query: `
				mutation {
					createAdmin(input: {
						email:     "not an email"
						firstName: "James"
						lastName:  "Bay"
						password: "1"
					}) {
						uuid
					}
				}
			`,
				ExpectedResult: `{"createAdmin":null}`,
				ExpectedErrors: []gqltest.TestQueryError{
					{
						Message: helpers.StringPointer("Email: not an email does not validate as email;" +
							"Password: 1 does not validate as stringlength(8|30)",
						),
						Path: []interface{}{"createAdmin"},
					},
				},
			},
			{
				Name:    "must be unique email",
				Context: adminContext,
				Schema:  schema,
				Query: `
				mutation {
					createAdmin(input: {
						email:     "test123@test.com"
						firstName: "James"
						lastName:  "Bay"
						password: "123456789"
					}) {
						uuid
					}
				}
			`,
				ExpectedResult: `{"createAdmin":null}`,
				ExpectedErrors: []gqltest.TestQueryError{
					{
						ResolverError: &errors.ErrUserExists,
						Path:          []interface{}{"createAdmin"},
					},
				},
			},
		})
	})

	accessTest(t, schema, accessTestOpts{
		Query: `
				mutation {
					createAdmin(input: {
						email:     "adminman@fdsa.com"
						firstName: "James"
						lastName:  "May"
						password: "realpassword"
					}) 
					{
						uuid
					}
				}
			`,
		Path:            []interface{}{"createAdmin"},
		MustAuth:        true,
		AdminAllowed:    true,
		ManagerAllowed:  false,
		DelegateAllowed: false,
		CleanDB:         true,
	})
}

func TestUpdateAdmin(t *testing.T) {
	gqltest.RunTests(t, []*gqltest.Test{
		{
			Name:    "update some fields",
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
		{
			Name:    "update some all",
			Context: adminContext,
			Schema:  schema,
			Query: `
				mutation {
					updateAdmin(input: {
						uuid: "00000000-0000-0000-0000-000000000002"
						email: "new@email.com"
						firstName: "asdf",
						lastName: "fdsa"
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
					  "email": "new@email.com",
						"firstName": "asdf",
						"lastName": "fdsa"
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
					updateAdmin(input: {
						uuid: "00000000-0000-0000-0000-000000000000"
					}) {
						uuid
					}
				}
			`,
			ExpectedResult: `
				{
					"updateAdmin": null
				}
			`,
			ExpectedErrors: []gqltest.TestQueryError{
				{
					ResolverError: &errors.ErrAdminNotFound,
					Path:          []interface{}{"updateAdmin"},
				},
			},
		},
		{
			Name:    "Fail validation",
			Context: adminContext,
			Schema:  schema,
			Query: `
					mutation {
						updateAdmin(input: {
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
						"updateAdmin": null
					}
				`,
			ExpectedErrors: []gqltest.TestQueryError{
				{
					Message: helpers.StringPointer("FirstName: 123! does not validate as alpha;" +
						"Email: not^%!£$* does not validate as email"),
					Path: []interface{}{"updateAdmin"},
				},
			},
		},
	})

	accessTest(
		t, schema, accessTestOpts{
			Query: `
				mutation {
					updateAdmin(input: {
						uuid: "00000000-0000-0000-0000-000000000003"
					}) {
						uuid
					}
				}
			`,
			Path:            []interface{}{"updateAdmin"},
			MustAuth:        true,
			AdminAllowed:    false,
			ManagerAllowed:  false,
			DelegateAllowed: false,
		},
	)
}

func TestDeleteAdmin(t *testing.T) {
	gqltest.RunTests(t, []*gqltest.Test{
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

	accessTest(
		t, schema, accessTestOpts{
			Query: `
				mutation {
					deleteAdmin(input: {
						uuid: "00000000-0000-0000-0000-000000000002"
					})
				}
			`,
			Path:            []interface{}{"deleteAdmin"},
			MustAuth:        true,
			AdminAllowed:    true,
			ManagerAllowed:  false,
			DelegateAllowed: false,
			CleanDB:         true,
		},
	)
}
