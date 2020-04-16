package resolvers_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/graph-gophers/graphql-go/gqltesting"
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
			UUID:    "00000000-0000-0000-0000-000000000001",
			Company: "",
			Role:    auth.AdminRole,
		},
		IsAdmin:    true,
		IsManager:  false,
		IsDelegate: false,
	}, *grant)
}

func TestUpdateAdmin(t *testing.T) {
	gqltesting.RunTests(t, []*gqltesting.Test{
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

func TestDeleteAdmin(t *testing.T) {
	gqltesting.RunTests(t, []*gqltesting.Test{
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
