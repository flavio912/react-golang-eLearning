// forked from https://github.com/graph-gophers/graphql-go/blob/master/gqltesting/testing.go

package gqltest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	graphql "github.com/graph-gophers/graphql-go"
	gqlerrors "github.com/graph-gophers/graphql-go/errors"
	"github.com/stretchr/testify/assert"
)

type TestQueryError struct {
	Message       *string
	ResolverError error
	Path          []interface{}
}

// Test is a GraphQL test case to be used with RunTest(s).
type Test struct {
	Name           string
	Context        context.Context
	Schema         *graphql.Schema
	Query          string
	OperationName  string
	Variables      map[string]interface{}
	ExpectedResult string
	ExpectedErrors []TestQueryError
}

// RunTests runs the given GraphQL test cases as subtests.
func RunTests(t *testing.T, tests []*Test) {
	if len(tests) == 1 {
		RunTest(t, tests[0])
		return
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d %s", i+1, test.Name), func(t *testing.T) {
			RunTest(t, test)
		})
	}
}

// RunTest runs a single GraphQL test case.
func RunTest(t *testing.T, test *Test) {
	if test.Context == nil {
		test.Context = context.Background()
	}
	result := test.Schema.Exec(test.Context, test.Query, test.OperationName, test.Variables)

	CheckErrors(t, test.ExpectedErrors, result.Errors)

	if test.ExpectedResult == "" {
		if result.Data != nil {
			t.Errorf("\ngot: %s\nwant: null", result.Data)
		}
		return
	}

	// Verify JSON to avoid red herring errors.
	got, err := formatJSON(result.Data)
	if err != nil {
		t.Fatalf("got: invalid JSON: %s", err)
	}
	want, err := formatJSON([]byte(test.ExpectedResult))
	if err != nil {
		t.Fatalf("want: invalid JSON: %s", err)
	}

	if !bytes.Equal(got, want) {
		t.Logf("got:  %s", got)
		t.Logf("want: %s", want)
		t.Fail()
	}
}

func formatJSON(data []byte) ([]byte, error) {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	formatted, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return formatted, nil
}

func CheckErrors(t *testing.T, want []TestQueryError, got []*gqlerrors.QueryError) {
	var gotTestErrors []TestQueryError
	for _, e := range got {
		gotTestErrors = append(gotTestErrors, TestQueryError{
			Message:       helpers.StringPointer(e.Message),
			Path:          e.Path,
			ResolverError: e.ResolverError,
		})
	}
	sortErrors(gotTestErrors)
	sortErrors(want)

	if len(want) != len(gotTestErrors) {
		t.Error("Unequal number of errors:")
		t.Errorf("want: %#v", want)
		t.Error("got:")
	}

	for i, g := range gotTestErrors {
		if len(want) < len(gotTestErrors) {
			if g.Message != nil {
				t.Errorf("%s", *g.Message)
			}
			t.Errorf("%#v", g.ResolverError)
			continue
		}

		assert.Equal(t, want[i].Path, g.Path)
		if want[i].Message != nil {
			assert.Equal(t, *want[i].Message, *g.Message)
		}
		if want[i].ResolverError != nil {
			assert.Equal(t, want[i].ResolverError, g.ResolverError)
		}
	}
}

func sortErrors(errors []TestQueryError) {
	if len(errors) <= 1 {
		return
	}
	sort.Slice(errors, func(i, j int) bool {
		return fmt.Sprintf("%s", errors[i].Path) < fmt.Sprintf("%s", errors[j].Path)
	})
}
