package errors

var (
	// ErrUserNotFound given when database could not find a user
	ErrUserNotFound = FullError{
		Type:     "ErrUserNotFound",
		Message:  "Could not find user",
		Title:    "User not found",
		HelpText: "Couldn't find a user with that email address",
	}
	// ErrUserExists usually for when someone tries to create a user that already exists
	ErrUserExists = FullError{
		Type:     "ErrUserExists",
		Message:  "User with that email/identifier already exists",
		Title:    "User already exists",
		HelpText: "Oh no, a user with that email/identifier already exists",
	}
	//ErrAdminNotFound given when something like a uuid meant to be an admin is invalid
	ErrAdminNotFound = FullError{
		Type:     "ErrAdminNotFound",
		Message:  "There is no user matching the information given",
		Title:    "Could not find the specified admin user",
		HelpText: "Oops, no admin with those details exists. Please check the details and try again",
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
	// ErrUUIDInvalid used when a given uuid is not valid (not for when its not found)
	ErrUUIDInvalid = SimpleError{
		Type:    "ErrUUIDInvalid",
		Message: "Given UUID was invalid",
	}
	// ErrCompanyNotFound when a company cannot be found from uuid or otherwise
	ErrCompanyNotFound = FullError{
		Type:     "ErrCompanyNotFound",
		Message:  "Cound not find the specified company",
		Title:    "Company not found",
		HelpText: "Sorry, we couldn't find the company you were looking for",
	}
	//ErrDeleteFailed is used for failed deletion where we know the resource was not deleted
	ErrDeleteFailed = FullError{
		Type:     "ErrDeleteFailed",
		Message:  "Could not delete the requested item",
		Title:    "Unable to delete",
		HelpText: "Sorry, we we're able to delete the item. Please try again.",
	}
)
