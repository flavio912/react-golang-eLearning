package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	graphql "github.com/graph-gophers/graphql-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/resolvers"
)

var (
	// We can pass an option to the schema so we don’t need to
	// write a method to access each type’s field:
	opts = []graphql.SchemaOpt{graphql.UseFieldResolvers()}
)

// Reads and parses the schema from file. Associates resolver. Panics if can't read.
func parseSchema(path string, resolver interface{}) *graphql.Schema {
	bstr, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	schemaString := string(bstr)
	parsedSchema, err := graphql.ParseSchema(schemaString, resolver, opts...)
	if err != nil {
		panic(err)
	}
	return parsedSchema
}

func main() {
	errDb := database.SetupDatabase()
	if errDb != nil {
		panic(errDb)
	}

	schema := parseSchema("./schema.graphql", &resolvers.RootResolver{})
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		type Payload struct {
			Query         string
			OperationName string
			Variables     map[string]interface{}
		}
		var payload Payload

		errParse := json.NewDecoder(r.Body).Decode(&payload)
		if errParse != nil {
			http.Error(w, errParse.Error(), http.StatusBadRequest)
			return
		}

		// todo: Pass JWT to resolvers for now; this should be moved to middleware
		token := strings.ReplaceAll(r.Header.Get("Authorization"), "Bearer ", "")
		ctx := context.WithValue(context.Background(), "token", token)

		resp := schema.Exec(ctx, payload.Query, payload.OperationName, payload.Variables)

		if len(resp.Errors) > 0 {
			type ErrorResponse struct {
				Data   interface{} `json:"data"`
				Errors interface{} `json:"errors"`
			}

			errResp := ErrorResponse{
				Data:   resp.Data,
				Errors: resp.Errors,
			}

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errResp)
			return
		}
		json, err := json.MarshalIndent(resp, "", "\t")
		if err != nil {
			// todo: does this need to be handled
			log.Printf("json.MarshalIndent: %s", err)
			return
		}

		fmt.Fprint(w, string(json))
	})

	log.Println("serving on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
