package resolvers_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers/gqltest"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

func TestDelegateLogin(t *testing.T) {
	prepareTestDatabase()

	t.Run("Must auth and return correct grant", func(t *testing.T) {
		res := schema.Exec(
			defaultContext(),
			`mutation {
				delegateLogin(input:{TTC_ID: "delegate-test-1", password: "iamamanager"}) {
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

		token := data["delegateLogin"].(map[string]interface{})["token"].(string)
		assert.Nil(t, err, "Error converting the token: \n%#v", data)

		// use the token to auth
		grant, err := middleware.Authenticate(token)
		assert.Nil(t, err)
		assert.Equal(t, middleware.Grant{
			Claims: auth.UserClaims{
				UUID:    gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
				Company: gentypes.MustParseToUUID("00000000-0000-0000-0000-000000000001"),
				Role:    auth.DelegateRole,
			},
			IsAdmin:      false,
			IsManager:    false,
			IsDelegate:   true,
			IsIndividual: false,
		}, *grant)
	})

	t.Run("must fail properly", func(t *testing.T) {
		gqltest.RunTests(t, []*gqltest.Test{
			{
				Name:    "bad ttc_id",
				Context: defaultContext(),
				Schema:  schema,
				Query: `
					mutation {
						delegateLogin(input:{TTC_ID: "this-is-not-a-ttc-id", password: "iamamanager"}) {
							token
						}
					}
				`,
				ExpectedResult: `{"delegateLogin":null}`,
				ExpectedErrors: []gqltest.TestQueryError{
					{
						Path:          []interface{}{"delegateLogin"},
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
						delegateLogin(input:{TTC_ID: "delegate-test-1", password: "notmypass"}) {
							token
						}
					}
				`,
				ExpectedResult: `{"delegateLogin":null}`,
				ExpectedErrors: []gqltest.TestQueryError{
					{
						Path:          []interface{}{"delegateLogin"},
						ResolverError: &errors.ErrAuthFailed,
					},
				},
			},
		})
	})
}

func TestCreateDelegate(t *testing.T) {
	prepareTestDatabase()

	t.Run("should successfully create a delegate", func(t *testing.T) {
		gqltest.RunTests(t, []*gqltest.Test{{
			Name:    "create delegate",
			Context: adminContext(),
			Schema:  schema,
			Query: `
				mutation {
					createDelegate(input: {
						companyUUID: "00000000-0000-0000-0000-000000000001"
						email:     "ttc@test.com"
						firstName: "Harry"
						lastName:  "Styles"
						telephone: "07894561230"
						jobTitle:  "Dev"
						password: "realpassword"
					}) {
						TTC_ID
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
					"createDelegate":{
						"TTC_ID":"testcompany-harrystyles",
						"email":"ttc@test.com",
						"firstName":"Harry",
						"jobTitle":"Dev",
						"lastName":"Styles",
						"telephone":"07894561230"
					}
				}
			`,
		}})

		// check you can auth with the new creds
		_, err := middleware.GetDelegateAccessToken("testcompany-harrystyles", "realpassword")
		assert.Nil(t, err)
	})
}
