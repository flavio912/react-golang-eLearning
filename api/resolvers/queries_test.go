package resolvers_test

import (
	"testing"

	"github.com/graph-gophers/graphql-go/gqltesting"
)

func TestAdmin(t *testing.T) {
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: adminContext,
			Schema:  schema,
			Query: `
				{
					admin(uuid: "00000000-0000-0000-0000-000000000001") {
						uuid
						firstName
						email
					}
				}
			`,
			ExpectedResult: `
				{
					"admin": {
					  "uuid": "00000000-0000-0000-0000-000000000001",
					  "email": "test123@test.com",
					  "firstName": "Jim"
					}
				}
			`,
		},
	})
}

func TestManager(t *testing.T) {
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: adminContext,
			Schema:  schema,
			Query: `
				{
					manager(uuid: "00000000-0000-0000-0000-000000000001") {
						uuid
						firstName
						email
					}
				}
			`,
			ExpectedResult: `
				{
					"manager": {
					  "uuid": "00000000-0000-0000-0000-000000000001",
					  "email": "man@managers.com",
					  "firstName": "Manager"
					}
				}
			`,
		},
	})
}

func TestCompany(t *testing.T) {
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: adminContext,
			Schema:  schema,
			Query: `
				{
					company(uuid:"00000000-0000-0000-0000-000000000001") {
							uuid
							name
							address {
							  postCode
							}
					}
				}
			`,
			ExpectedResult: `
				{
					"company": {
					  "uuid": "00000000-0000-0000-0000-000000000001",
						"name": "TestCompany",
						"address": {
							"postCode": "IP24RF"
						}
					}
				}
			`,
		},
	})
}
