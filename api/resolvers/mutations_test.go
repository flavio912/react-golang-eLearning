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
			defaultContext(),
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
				Context: defaultContext(),
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
				Context: defaultContext(),
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
			defaultContext(),
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

	t.Run("noResp param is respected", func(t *testing.T) {
		gqltest.RunTest(t, &gqltest.Test{

			Name:    "Blank response expected",
			Context: defaultContext(),
			Schema:  schema,
			Query: `
					mutation {
						managerLogin(input:{email: "man@managers.com", password: "iamamanager", noResp: true}) {
							token
						}
					}
				`,
			ExpectedResult: `{"managerLogin":{"token":""}}`,
			ExpectedErrors: nil,
		})
	})

	t.Run("must fail properly", func(t *testing.T) {
		gqltest.RunTests(t, []*gqltest.Test{
			{
				Name:    "bad email",
				Context: defaultContext(),
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
				Context: defaultContext(),
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
			Context: adminContext(),
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
				Context: adminContext(),
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
				Context: defaultContext(),
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
				Context: adminContext(),
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
			Context: adminContext(),
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
			Context: adminContext(),
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
			Context: adminContext(),
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
			Context: adminContext(),
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
			Context: adminContext(),
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
			Context: adminContext(),
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

	prepareTestDatabase()
	accessTest(t, schema, accessTestOpts{
		Query: `
				mutation {
					deleteManager(input: {
						uuid: "00000000-0000-0000-0000-000000000002"
					})
				}
			`,
		Path:            []interface{}{"deleteManager"},
		MustAuth:        true,
		AdminAllowed:    true,
		ManagerAllowed:  false,
		DelegateAllowed: false,
		CleanDB:         true,
	})
}

func TestCreateAdmin(t *testing.T) {
	prepareTestDatabase()

	t.Run("should successfully create a manager", func(t *testing.T) {
		gqltest.RunTests(t, []*gqltest.Test{{
			Name:    "create manager",
			Context: adminContext(),
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
				Context: adminContext(),
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
				Context: adminContext(),
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
			Context: adminContext(),
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
			Context: adminContext(),
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
			Context: adminContext(),
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
			Context: adminContext(),
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

	t.Run("Test loaders reset", func(t *testing.T) {
		prepareTestDatabase()

		gqltest.RunTests(t, []*gqltest.Test{
			{
				Name:    "get admin into loader ctx",
				Context: adminContext(),
				Schema:  schema,
				Query: `
					{
						admin(uuid: "00000000-0000-0000-0000-000000000002") {
							uuid
							firstName
							lastName
							email
						}
					}
				`,
				ExpectedResult: `
					{
						"admin": {
							"uuid": "00000000-0000-0000-0000-000000000002",
							"email": "steve@wombat.com",
							"firstName": "Steve",
							"lastName": "Wombat"
						}
					}
				`,
			},
			{
				Name:    "update some fields",
				Context: adminContext(),
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
				Name:    "check loader has been flushed",
				Context: adminContext(),
				Schema:  schema,
				Query: `
					{
						admin(uuid: "00000000-0000-0000-0000-000000000002") {
							uuid
							firstName
							lastName
							email
						}
					}
				`,
				ExpectedResult: `
					{
						"admin": {
							"uuid": "00000000-0000-0000-0000-000000000002",
							"email": "steve@wombat.com",
							"firstName": "edfadd",
							"lastName": "dsa"
						}
					}
				`,
			},
		})
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
			AdminAllowed:    true,
			ManagerAllowed:  false,
			DelegateAllowed: false,
		},
	)
}

func TestDeleteAdmin(t *testing.T) {
	gqltest.RunTests(t, []*gqltest.Test{
		{
			Context: adminContext(),
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

func TestCreateCompany(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Name:    "create company",
			Context: adminContext(),
			Schema:  schema,
			Query: `
				mutation {
					createCompany(input: {
						companyName: "Cool Co"
						addressLine1: "100 Cool Lane"
						addressLine2: ""
						county: "Coolington"
						postCode: "CO0L3ST"
						country: "UK"
					}) {
						approved
						name
						address {
							addressLine1
							addressLine2
							county
							postCode
							country
						}
					}
				}
			`,
			ExpectedResult: `
				{
					"createCompany":{
						"address":{
							"addressLine1":"100 Cool Lane",
							"addressLine2":"",
							"country":"UK",
							"county":"Coolington",
							"postCode":"CO0L3ST"
						},
						"approved":true,
						"name":"Cool Co"
					}
				}
			`,
		},
		{
			Name:    "should validate",
			Context: adminContext(),
			Schema:  schema,
			Query: `
				mutation {
					createCompany(input: {
						companyName: ""
						addressLine1: ""
						addressLine2: ""
						county: ""
						postCode: "reallylong"
						country: ""
					}) {
						name
					}
				}
			`,
			ExpectedResult: `{"createCompany":null}`,
			ExpectedErrors: []gqltest.TestQueryError{{
				Message: helpers.StringPointer("PostCode: reallylong does not validate as stringlength(6|7)"),
				Path:    []interface{}{"createCompany"},
			}},
		},
	})

	accessTest(t, schema, accessTestOpts{
		Query: `
			mutation {
				createCompany(input: {
					companyName: ""
					addressLine1: ""
					addressLine2: ""
					county: ""
					postCode: "1234567"
					country: ""
				}) {
					name
				}
			}
		`,
		Path:            []interface{}{"createCompany"},
		MustAuth:        true,
		AdminAllowed:    true,
		ManagerAllowed:  false,
		DelegateAllowed: false,
		CleanDB:         false,
	})
}

func TestUpdateCompany(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Name:    "Update Some Fields",
			Context: adminContext(),
			Schema:  schema,
			Query: `
				mutation {
					updateCompany(input: {
						uuid: "00000000-0000-0000-0000-000000000002"
						companyName: "C132"
						addressLine1: "ajfd"
						postCode: "1234567"
					}) {
						uuid
						name
						approved
						address {
							addressLine1
							addressLine2
							postCode
							county
							country
						}
					}
				}
			`,
			ExpectedResult: `
				{
					"updateCompany":{
						"address":{
							"addressLine1":"ajfd",
							"addressLine2":"Address line two 2",
							"country":"UK2",
							"county":"York2",
							"postCode":"1234567"
						},
						"approved":true,
						"name":"C132",
						"uuid":"00000000-0000-0000-0000-000000000002"
					}
				}
			`,
		},
		{
			Name:    "Update All Fields",
			Context: adminContext(),
			Schema:  schema,
			Query: `
				mutation {
					updateCompany(input: {
						uuid: "00000000-0000-0000-0000-000000000002"
						companyName: "C132"
						addressLine1: "afdsa"
						addressLine2: "asdfa"
						postCode: "asdf12"
						county: "aadfjk"
						country: "USA"
						approved: false
					}) {
						uuid
						name
						approved
						address {
							addressLine1
							addressLine2
							postCode
							county
							country
						}
					}
				}
			`,
			ExpectedResult: `
				{
					"updateCompany":{
						"address":{
							"addressLine1":"afdsa",
							"addressLine2":"asdfa",
							"country":"USA",
							"county":"aadfjk",
							"postCode":"asdf12"
						},
						"approved":false,
						"name":"C132",
						"uuid":"00000000-0000-0000-0000-000000000002"
					}
				}
			`,
		},
		{
			Name:    "UUID does not exist",
			Context: adminContext(),
			Schema:  schema,
			Query: `
				mutation {
					updateCompany(input: {
						uuid: "00000000-0000-0000-0000-000000000000"
					}) {
						uuid
					}
				}
			`,
			ExpectedResult: `
				{
					"updateCompany": null
				}
			`,
			ExpectedErrors: []gqltest.TestQueryError{
				{
					ResolverError: &errors.ErrCompanyNotFound,
					Path:          []interface{}{"updateCompany"},
				},
			},
		},
		{
			Name:    "Fail validation",
			Context: adminContext(),
			Schema:  schema,
			Query: `
				mutation {
					updateCompany(input: {
						uuid: "00000000-0000-0000-0000-000000000001"
						country: "123!"
						county: "not^%!£$*"
					}) {
						uuid
					}
				}
			`,
			ExpectedResult: `
				{
					"updateCompany": null
				}
			`,
			ExpectedErrors: []gqltest.TestQueryError{
				{
					Message: helpers.StringPointer("County: not^%!£$* does not validate as alpha;Country: 123! does not validate as alpha"),
					Path:    []interface{}{"updateCompany"},
				},
			},
		},
	})

	accessTest(
		t, schema, accessTestOpts{
			Query: `
				mutation {
					updateCompany(input: {
						uuid: "00000000-0000-0000-0000-000000000003"
					}) {
						uuid
					}
				}
			`,
			Path:            []interface{}{"updateCompany"},
			MustAuth:        true,
			AdminAllowed:    true,
			ManagerAllowed:  false,
			DelegateAllowed: false,
		},
	)
}

func TestApproveCompany(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Name:    "create company",
			Context: adminContext(),
			Schema:  schema,
			Query: `
				mutation {
					approveCompany(uuid: "00000000-0000-0000-0000-000000000004") {
						approved
						name
					}
				}
			`,
			ExpectedResult: `
				{
					"approveCompany":{
						"approved":true,
						"name":"Microsoft"
					}
				}
			`,
		},
	})

	accessTest(t, schema, accessTestOpts{
		Query: `
			mutation {
				approveCompany(uuid: "00000000-0000-0000-0000-000000000004") {
					name
				}
			}
		`,
		Path:            []interface{}{"approveCompany"},
		MustAuth:        true,
		AdminAllowed:    true,
		ManagerAllowed:  false,
		DelegateAllowed: false,
		CleanDB:         false,
	})
}

func TestCreateCategory(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Name:    "create category",
			Context: adminContext(),
			Schema:  schema,
			Query: `
				mutation {
					createCategory(
						input:{
							name: "best category ever made",
							color: "#fffffa"
						}
					) {
						color
						name
					}
				}
			`,
			ExpectedResult: `
				{
					"createCategory":{
							"color": "#fffffa",
							"name": "best category ever made"
					}
				}
			`,
		},
	})

	// this needs to be done ...
	// accessTest(t, schema, accessTestOpts{
	// 	Query: `
	// 		mutation {
	// 			approveCompany(uuid: "00000000-0000-0000-0000-000000000004") {
	// 				name
	// 			}
	// 		}
	// 	`,
	// 	Path:            []interface{}{"approveCompany"},
	// 	MustAuth:        true,
	// 	AdminAllowed:    true,
	// 	ManagerAllowed:  false,
	// 	DelegateAllowed: false,
	// 	CleanDB:         false,
	// })
}

func TestSaveOnlineCourse(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Name:    "Create online course",
			Context: adminContext(),
			Schema:  schema,
			Query: `
				mutation {
					saveOnlineCourse(
						input:{
							name: "Test online course",
							excerpt: "{}",
							introduction:"{}",
							backgroundCheck: true,
							accessType: open,
							price: 34.3,
							color: "#fff",
							howToComplete: "{}",
        			whatYouLearn: ["What 1", "What 2"],
							requirements: ["req 1", "req 2"]
						}
					) {
						name
						excerpt
						introduction
						backgroundCheck
						price
						color
						howToComplete
						whatYouLearn
						requirements
					}
				}
			`,
			ExpectedResult: `
				{
					"saveOnlineCourse":{
							"name": "Test online course",
							"excerpt": "{}",
							"introduction":"{}",
							"backgroundCheck": true,
							"price": 34.3,
							"color": "#fff",
							"howToComplete": "{}",
							"whatYouLearn": ["What 1", "What 2"],
							"requirements": ["req 1", "req 2"]
					}
				}
			`,
		},
	})
}

func TestCreateTest(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Name:    "Create test course",
			Context: adminContext(),
			Schema:  schema,
			Query: `
				mutation {
					createTest(input: {
							name: "Cake",
							attemptsAllowed: 3,
							passPercentage: 30,
							questionsToAnswer: 7,
							randomiseAnswers: false,
							questions: []
					}) {
							test {
								name
								complete
								attemptsAllowed
								passPercentage
								questionsToAnswer
								randomiseAnswers
								questions {
									text
								}
							}
						}
				}
			`,
			ExpectedResult: `
				{
					"createTest": {
            "test": {
                "name": "Cake",
                "complete": false,
                "attemptsAllowed": 3,
                "passPercentage": 30,
                "questionsToAnswer": 7,
                "randomiseAnswers": false,
                "questions": null
            }
        	}
				}
			`,
		},
	})
}

func TestCreateTutor(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Name:    "Creates tutor",
			Context: adminContext(),
			Schema:  schema,
			Query: `
				mutation {
					createTutor(input: {
						name: "Walter White"
						cin: 420
					}) {
						name
						cin
					}
				}
			`,
			ExpectedResult: `
				{
					"createTutor": {
						"name": "Walter White",
						"cin": 420
					}
				}
			`,
		},
	})

	accessTest(t, schema, accessTestOpts{
		Query: `
			mutation {
				createTutor(input: {
					name: "Savage"
					cin: 21
				}) {
					uuid
				}
			}
		`,
		Path:            []interface{}{"createTutor"},
		MustAuth:        true,
		AdminAllowed:    true,
		ManagerAllowed:  false,
		DelegateAllowed: false,
	})
}

func TestUpdateTutor(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Name:    "Update some fields",
			Context: adminContext(),
			Schema:  schema,
			Query: `
				mutation {
					updateTutor(input: {
						uuid: "386bd256-82e0-4d8a-91af-b4a117e0eda8"
						name: "Richard Feynman"
						cin: 69
					}) {
						name
						cin
					}
				}
			`,
			ExpectedResult: `
				{
					"updateTutor": {
						"name": "Richard Feynman",
						"cin": 69
					}
				}
			`,
		},
	})

	accessTest(t, schema, accessTestOpts{
		Query: `
			mutation {
				updateTutor(input: {
					uuid: "386bd256-82e0-4d8a-91af-b4a117e0eda8"
				}) {
					uuid
				}
			}
		`,
		Path:            []interface{}{"updateTutor"},
		MustAuth:        true,
		AdminAllowed:    true,
		ManagerAllowed:  false,
		DelegateAllowed: false,
	})
}
