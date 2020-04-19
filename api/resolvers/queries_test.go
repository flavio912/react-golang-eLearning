package resolvers_test

import (
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers/gqltest"
)

func TestInfo(t *testing.T) {
	gqltest.RunTests(t, []*gqltest.Test{
		{
			Schema: schema,
			Query: `
				{
					info
				}
			`,
			ExpectedResult: `
				{
					"info": "This is the TTC server api"
				}
			`,
		},
	})
}

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

	accessTest(
		t, schema,
		`{
			 admin(uuid: "00000000-0000-0000-0000-000000000001") { uuid }
		 }`,
		[]gqltest.TestQueryError{
			{
				ResolverError: &errors.ErrUnauthorized,
				Path:          []interface{}{"admin"},
			},
		}, true, true, false, false,
	)
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
							{"uuid":"00000000-0000-0000-0000-000000000004"},
							{"uuid":"00000000-0000-0000-0000-000000000003"},
							{"uuid":"00000000-0000-0000-0000-000000000002"},
							{"uuid":"00000000-0000-0000-0000-000000000001"}
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
		{
			Name:    "Should page",
			Context: adminContext,
			Schema:  schema,
			Query: `
				{
					admins (page: {
						offset: 1
						limit: 2
					}) {
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
							{"uuid":"00000000-0000-0000-0000-000000000003"},
							{"uuid":"00000000-0000-0000-0000-000000000002"}
						],
						"pageInfo": {
							"given": 2,
							"limit": 2,
							"offset": 1,
							"total": 4
						}
					}
				}
			`,
		},
	})

	accessTest(
		t, schema,
		`{
			 admins { edges { uuid } }
		 }`,
		[]gqltest.TestQueryError{
			{
				ResolverError: &errors.ErrUnauthorized,
				Path:          []interface{}{"admins"},
			},
		}, true, true, false, false,
	)
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

	accessTest(
		t, schema,
		`{
			 manager(uuid: "00000000-0000-0000-0000-000000000001") { uuid }
		 }`,
		[]gqltest.TestQueryError{
			{
				ResolverError: &errors.ErrUnauthorized,
				Path:          []interface{}{"manager"},
			},
		}, true, true, true, false, // the deafult manager can see itself
	)
	accessTest(
		t, schema,
		`{
			 manager(uuid: "00000000-0000-0000-0000-000000000002") { uuid }
		 }`,
		[]gqltest.TestQueryError{
			{
				ResolverError: &errors.ErrUnauthorized,
				Path:          []interface{}{"manager"},
			},
		}, true, true, false, false, // the deafult manager cannot see others
	)
}

func TestManagers(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Name:    "Should return all managers",
			Context: adminContext,
			Schema:  schema,
			Query: `
				{
					managers {
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
					"managers":{
						"edges":[
							{"uuid":"00000000-0000-0000-0000-000000000004"},
							{"uuid":"00000000-0000-0000-0000-000000000003"},
							{"uuid":"00000000-0000-0000-0000-000000000002"},
							{"uuid":"00000000-0000-0000-0000-000000000001"}
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
		{
			Name:    "Should order",
			Context: adminContext,
			Schema:  schema,
			Query: `
				{
					managers (orderBy: {
						ascending: true
						field: "first_name"
					}) {
						edges {
							firstName
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
					"managers":{
						"edges":[
							{"firstName":"Jimothy"},
							{"firstName":"Steve"},
							{"firstName":"Slim"},
							{"firstName":"Manager"}
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
		{
			Name:    "Should filter",
			Context: adminContext,
			Schema:  schema,
			Query: `
				{
					managers (filter: {
						name: "S"
					}) {
						edges {
							firstName
							lastName
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
					"managers":{
						"edges":[
							{"firstName": "Jimothy", "lastName": "Bobnes"},
							{"firstName": "Steve", "lastName": "Person"},
							{"firstName": "Slim", "lastName":"Shady"}
						],
						"pageInfo": {
							"given": 3,
							"limit": 100,
							"offset": 0,
							"total": 3
						}
					}
				}
			`,
		},
		{
			Name:    "Should page",
			Context: adminContext,
			Schema:  schema,
			Query: `
				{
					managers (page: {
						offset: 1
						limit: 2
					}) {
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
					"managers":{
						"edges":[
							{"uuid":"00000000-0000-0000-0000-000000000003"},
							{"uuid":"00000000-0000-0000-0000-000000000002"}
						],
						"pageInfo": {
							"given": 2,
							"limit": 2,
							"offset": 1,
							"total": 4
						}
					}
				}
			`,
		},
		{
			Name:    "filter must validate",
			Context: adminContext,
			Schema:  schema,
			Query: `
				{
					managers (filter: {
						telephone: "sa#q345654sdf"
					}) {
						edges {
							uuid
						}
					}
				}
			`,
			ExpectedResult: `{"managers":null}`,
			ExpectedErrors: []gqltest.TestQueryError{
				{
					Message: helpers.StringPointer("Telephone: sa#q345654sdf does not validate as numeric"),
					Path:    []interface{}{"managers"},
				},
			},
		},
	})

	accessTest(
		t, schema,
		`{
			 managers { edges { uuid } }
		 }`,
		[]gqltest.TestQueryError{
			{
				ResolverError: &errors.ErrUnauthorized,
				Path:          []interface{}{"managers"},
			},
		}, true, true, false, false,
	)
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

	accessTest(
		t, schema,
		`{
			company(uuid:"00000000-0000-0000-0000-000000000001") { uuid }
		 }`,
		[]gqltest.TestQueryError{
			{
				ResolverError: &errors.ErrUnauthorized,
				Path:          []interface{}{"company"},
			},
		}, true, true, false, false,
	)
}

func TestCompanies(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Name:    "Should return all companies",
			Context: adminContext,
			Schema:  schema,
			Query: `
					{
						companies {
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
						"companies":{
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
		{
			Name:    "Should order",
			Context: adminContext,
			Schema:  schema,
			Query: `
					{
						companies (orderBy: {
							field: "name"
						}) {
							edges {
								name
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
						"companies":{
							"edges":[
								{"name": "TestCompany"},
								{"name": "Microsoft"},
								{"name": "Fake Work Place"},
								{"name": "ACME"}
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
		{
			Name:    "Should filter",
			Context: adminContext,
			Schema:  schema,
			Query: `
					{
						companies (filter: {
							name: "m"
						}) {
							edges {
								name
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
						"companies":{
							"edges":[
								{"name":"TestCompany"},
								{"name":"ACME"},
								{"name":"Microsoft"}
							],
							"pageInfo": {
								"given": 3,
								"limit": 100,
								"offset": 0,
								"total": 3
							}
						}
					}
				`,
		},
		{
			Name:    "Should page",
			Context: adminContext,
			Schema:  schema,
			Query: `
					{
						companies (page: {
							offset: 1
							limit: 2
						}) {
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
						"companies":{
							"edges":[
								{"uuid":"00000000-0000-0000-0000-000000000002"},
								{"uuid":"00000000-0000-0000-0000-000000000003"}
							],
							"pageInfo": {
								"given": 2,
								"limit": 2,
								"offset": 1,
								"total": 4
							}
						}
					}
				`,
		},
	})

	accessTest(
		t, schema,
		`{
			 companies { edges { uuid } }
		 }`,
		[]gqltest.TestQueryError{
			{
				ResolverError: &errors.ErrUnauthorized,
				Path:          []interface{}{"companies"},
			},
		}, true, true, false, false,
	)
}
