package errors

var (
	// ErrUserNotFound given when database could not find a user
	ErrUserNotFound = FullError{
		Type:     "ErrUserNotFound",
		Message:  "Could not find user",
		Title:    "User not found",
		HelpText: "Couldn't find a user with that email address",
	}
	// ErrAuthFailed given when password or email is not correct
	ErrAuthFailed = FullError{
		Type:     "ErrAuthFailed",
		Message:  "Email or password incorrect",
		Title:    "Email or password incorrect",
		HelpText: "Please try again",
	}
	// ErrUnauthorized is given when user attempts an action above their auth
	ErrUnauthorized = SimpleError{
		Type:    "ErrUnauthorized",
		Message: "User does not have permissions for request",
	}
	// ErrTokenInvalid is given when the JWT is not valid or expired
	ErrTokenInvalid = SimpleError{
		Type:    "ErrTokenInvalid",
		Message: "Given token is invalid or expired",
	}
	// ErrNotFound given when some item cannot be found
	ErrNotFound = SimpleError{
		Type:    "ErrNotFound",
		Message: "Unable to find the requested item[s]",
	}
	// ErrWhileHandling is used when an unknown error occured
	ErrWhileHandling = SimpleError{
		Type:    "ErrWhileHandling",
		Message: "There was an error while handling the request",
	}
)
