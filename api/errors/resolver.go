package errors

var (
	ErrUnableToResolve = SimpleError{
		Type:    "ErrUnableToResolve",
		Message: "Could not resolve this request",
	}
)
