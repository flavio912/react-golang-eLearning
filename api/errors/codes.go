package errors

import (
	"github.com/lib/pq"
)

// Postgres Error codes

// CodeUniqueViolation is the postgres error code for a uniqueness violation
var CodeUniqueViolation pq.ErrorCode = "23505"
