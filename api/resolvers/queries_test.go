package resolvers_test

import (
	"testing"

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
			Context: adminContext(),
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
		t, schema, accessTestOpts{
			Query:           `{admin(uuid: "00000000-0000-0000-0000-000000000001") { uuid }}`,
			Path:            []interface{}{"admin"},
			MustAuth:        true,
			AdminAllowed:    true,
			ManagerAllowed:  false,
			DelegateAllowed: false,
		},
	)
}

func TestAdmins(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Name:    "Should return all admins",
			Context: adminContext(),
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
			Context: adminContext(),
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
		t, schema, accessTestOpts{
			Query:           `{admins { edges { uuid } }}`,
			Path:            []interface{}{"admins"},
			MustAuth:        true,
			AdminAllowed:    true,
			ManagerAllowed:  false,
			DelegateAllowed: false,
		},
	)
}

func TestManager(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Context: adminContext(),
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
		t, schema, accessTestOpts{
			Query:           `{manager(uuid: "00000000-0000-0000-0000-000000000001") { uuid }}`,
			Path:            []interface{}{"manager"},
			MustAuth:        true,
			AdminAllowed:    true,
			ManagerAllowed:  true, // the deafult manager can see itself
			DelegateAllowed: false,
		},
	)
	accessTest(
		t, schema, accessTestOpts{
			Query:           `{manager(uuid: "00000000-0000-0000-0000-000000000002") { uuid }}`,
			Path:            []interface{}{"manager"},
			MustAuth:        true,
			AdminAllowed:    true,
			ManagerAllowed:  false,
			DelegateAllowed: false,
		},
	)
}

func TestManagers(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Name:    "Should return all managers",
			Context: adminContext(),
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
							{"uuid":"00000000-0000-0000-0000-000000000001"},
							{"uuid":"00000000-0000-0000-0000-000000000005"}
						],
						"pageInfo": {
							"given": 5,
							"limit": 100,
							"offset": 0,
							"total": 5
						}
					}
				}
			`,
		},
		{
			Name:    "Should order",
			Context: adminContext(),
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
							{"firstName":"Manager"},
							{"firstName":"Slim"},
							{"firstName":"Steve"},
							{"firstName":"Tony"}
						],
						"pageInfo": {
							"given": 5,
							"limit": 100,
							"offset": 0,
							"total": 5
						}
					}
				}
			`,
		},
		{
			Name:    "Should filter",
			Context: adminContext(),
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
							{"firstName": "Slim", "lastName":"Shady"},
							{"firstName": "Tony", "lastName":"NicePerson"}
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
			Context: adminContext(),
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
							"total": 5
						}
					}
				}
			`,
		},
		{
			Name:    "filter must validate",
			Context: adminContext(),
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
					Message: helpers.StringPointer("UserFilter.Telephone: sa#q345654sdf does not validate as numeric"),
					Path:    []interface{}{"managers"},
				},
			},
		},
	})

	accessTest(
		t, schema, accessTestOpts{
			Query:           `{managers { edges { uuid } }}`,
			Path:            []interface{}{"managers"},
			MustAuth:        true,
			AdminAllowed:    true,
			ManagerAllowed:  false,
			DelegateAllowed: false,
		},
	)
}

func TestDelegate(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Context: adminContext(),
			Schema:  schema,
			Query: `
				{
					delegate(uuid: "00000000-0000-0000-0000-000000000001") {
						uuid
						TTC_ID
						email
						firstName
						lastName
						telephone
						jobTitle
						company {
							uuid
							name
						}
					}
				}
			`,
			ExpectedResult: `
				{
					"delegate":{
						"TTC_ID":"delegate-test-1",
						"company":{
							"name":"TestCompany",
							"uuid":"00000000-0000-0000-0000-000000000001"
						},
						"email":"del@delegates.com",
						"firstName":"Delegate",
						"jobTitle":"Doer",
						"lastName":"Man",
						"telephone":"7912935287",
						"uuid":"00000000-0000-0000-0000-000000000001"
					}
				}
			`,
		},
	})

	accessTest(
		t, schema, accessTestOpts{
			Query:           `{delegate(uuid: "00000000-0000-0000-0000-000000000001") { uuid }}`,
			Path:            []interface{}{"delegate"},
			MustAuth:        true,
			AdminAllowed:    true,
			ManagerAllowed:  true, // the deafult manager can see itself
			DelegateAllowed: false,
		},
	)
}

func TestDelegates(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Name:    "Should return all delegates",
			Context: adminContext(),
			Schema:  schema,
			Query: `
				{
					delegates {
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
					"delegates":{
						"edges":[
							{"uuid":"00000000-0000-0000-0000-000000000004"},
							{"uuid":"00000000-0000-0000-0000-000000000003"},
							{"uuid":"00000000-0000-0000-0000-000000000002"},
							{"uuid":"00000000-0000-0000-0000-000000000001"}
						],
						"pageInfo":{
							"given":4,
							"limit":100,
							"offset":0,
							"total":4
						}
					}
				}
			`,
		},
		{
			Name:    "Should order",
			Context: adminContext(),
			Schema:  schema,
			Query: `
				{
					delegates (orderBy: {
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
					"delegates": {
						"edges": [
							{"firstName": "David"},
							{"firstName": "Delegate"},
							{"firstName": "James"},
							{"firstName": "Person"}
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
			Context: adminContext(),
			Schema:  schema,
			Query: `
				{
					delegates (filter: {
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
					"delegates": {
						"edges": [
							{"firstName": "James","lastName": "Bay"},
							{"firstName": "Person","lastName": "Pearson"}
						],
						"pageInfo": {
							"given": 2,
							"limit": 100,
							"offset": 0,
							"total": 2
						}
					}
				}
			`,
		},
		{
			Name:    "Should page",
			Context: adminContext(),
			Schema:  schema,
			Query: `
				{
					delegates (page: {
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
					"delegates":{
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
			Context: adminContext(),
			Schema:  schema,
			Query: `
				{
					delegates (filter: {
						telephone: "sa#q345654sdf"
					}) {
						edges {
							uuid
						}
					}
				}
			`,
			ExpectedResult: `{"delegates":null}`,
			ExpectedErrors: []gqltest.TestQueryError{
				{
					Message: helpers.StringPointer("UserFilter.Telephone: sa#q345654sdf does not validate as numeric"),
					Path:    []interface{}{"delegates"},
				},
			},
		},
	})

	accessTest(
		t, schema, accessTestOpts{
			Query:           `{delegates { edges { uuid } }}`,
			Path:            []interface{}{"delegates"},
			MustAuth:        true,
			AdminAllowed:    true,
			ManagerAllowed:  true,
			DelegateAllowed: false,
		},
	)
}
func TestCompany(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Context: adminContext(),
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
		t, schema, accessTestOpts{
			Query:           `{company(uuid:"00000000-0000-0000-0000-000000000001") { uuid }}`,
			Path:            []interface{}{"company"},
			MustAuth:        true,
			AdminAllowed:    true,
			ManagerAllowed:  true,
			DelegateAllowed: false,
		},
	)
	accessTest(
		t, schema, accessTestOpts{
			Query:           `{company(uuid:"00000000-0000-0000-0000-000000000002") { uuid }}`,
			Path:            []interface{}{"company"},
			MustAuth:        true,
			AdminAllowed:    true,
			ManagerAllowed:  false, // manager can only see it's own company
			DelegateAllowed: false,
		},
	)
}

func TestCompanies(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Name:    "Should return all companies",
			Context: adminContext(),
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
			Context: adminContext(),
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
			Context: adminContext(),
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
			Context: adminContext(),
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
		t, schema, accessTestOpts{
			Query:           `{companies { edges { uuid } }}`,
			Path:            []interface{}{"companies"},
			MustAuth:        true,
			AdminAllowed:    true,
			ManagerAllowed:  false,
			DelegateAllowed: false,
		},
	)
}

func TestOnlineCourses(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Name:    "Should return all online courses",
			Context: adminContext(),
			Schema:  schema,
			Query: `
					{
						onlineCourses {
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
						"onlineCourses":{
							"edges":[
								{"uuid":"00000000-0000-0000-0000-000000000001"},
								{"uuid":"00000000-0000-0000-0000-000000000002"},
								{"uuid":"00000000-0000-0000-0000-000000000003"}
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
			Name:    "Should return unrestricted courses",
			Context: adminContext(),
			Schema:  schema,
			Query: `
					{
						onlineCourses {
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
						"onlineCourses":{
							"edges":[
								{"uuid":"00000000-0000-0000-0000-000000000001"},
								{"uuid":"00000000-0000-0000-0000-000000000002"},
								{"uuid":"00000000-0000-0000-0000-000000000003"}
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
	})
}

func TestClassroomCourses(t *testing.T) {
	prepareTestDatabase()

	gqltest.RunTests(t, []*gqltest.Test{
		{
			Name:    "Should return all classroom courses",
			Context: adminContext(),
			Schema:  schema,
			Query: `
					{
						classroomCourses {
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
						"classroomCourses":{
							"edges":[
								{"uuid":"00000000-0000-0000-0000-000000000013"},
								{"uuid":"00000000-0000-0000-0000-000000000014"},
								{"uuid":"00000000-0000-0000-0000-000000000012"}
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
	})
}
