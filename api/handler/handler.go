package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/loader"
)

type GraphQL struct {
	Schema  *graphql.Schema
	Loaders loader.Map
}

func (g *GraphQL) Serve() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		ctx := r.Context()

		ctx = g.Loaders.Attach(ctx)
		resp := g.Schema.Exec(ctx, payload.Query, payload.OperationName, payload.Variables)

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

		json, err := json.Marshal(resp)
		if err != nil {
			log.Printf("json.Marshal: %s", err)
			return
		}

		w.Write(json)
	})
}
