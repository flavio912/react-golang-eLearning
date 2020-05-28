package resolvers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers/gqltest"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

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
