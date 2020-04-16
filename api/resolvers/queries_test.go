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
