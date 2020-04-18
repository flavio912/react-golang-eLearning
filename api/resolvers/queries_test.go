package resolvers_test

import (
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers/gqltest"
)

func TestAdmin(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
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

func TestAdmins(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Name:    "Should return all admins",
			Context: adminContext,
			Schema:  schema,
			Query: `
				{
					admins {
						edges {
							uuid
						}
						pageInfo {
							total
							offset
							limit
							given
						}
					}
				}
			`,
			ExpectedResult: `
				{
					"admins":{
						"edges":[
							{"uuid":"00000000-0000-0000-0000-000000000001"},
							{"uuid":"00000000-0000-0000-0000-000000000002"},
							{"uuid":"00000000-0000-0000-0000-000000000003"},
							{"uuid":"00000000-0000-0000-0000-000000000004"}
						],
						"pageInfo": {
							"given": 4,
							"limit": 100,
							"offset": 0,
							"total": 4
						}
					}
				}
			`,
		},
	})
}

func TestManager(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
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
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
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
