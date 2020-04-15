package gentypes

import (
	"fmt"

	"github.com/google/uuid"
)

// UUID is a custom graphql schema type for representing UUIDs
type UUID struct {
	uuid.UUID
}

func (_ UUID) ImplementsGraphQLType(name string) bool {
	return name == "UUID"
}

func (u *UUID) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case uuid.UUID:
		u.UUID = input
		return nil
	case string:
		var err error
		u.UUID, err = uuid.Parse(input)
		return err
	default:
		return fmt.Errorf("wrong type")
	}
}

// Parse is a wrapper around uuid.Parse for our own type
func (*UUID) Parse(input string) (UUID, error) {
	id, err := uuid.Parse(input)
	if err != nil {
		return UUID{}, err
	}
	return UUID{
		UUID: id,
	}, nil
}
