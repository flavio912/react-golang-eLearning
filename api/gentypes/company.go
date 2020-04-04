package gentypes

import (
	"github.com/google/uuid"
)

type Company struct {
	CreatedAt *string `valid:"rfc3339"`
	UUID      uuid.UUID
	Name      string
}

type CompanyFilter struct {
	UUID *string
	Name *string
}

type OrderBy struct {
	Ascending *bool //defaults to false, thus decending
	Field     string
}
