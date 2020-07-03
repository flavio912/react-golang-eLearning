package errors

import (
	"fmt"
)

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
	//ErrManagerNotFound given when a manager with the given uuid cannot be found
	ErrManagerNotFound = FullError{
		Type:     "ErrManagerNotFound",
		Message:  "There is no user matching the information given",
		Title:    "Could not find the specified manager user",
		HelpText: "Oops, no manager with those details exists. Please check the details and try again",
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
	ErrOrderUnauthorized = SimpleError{
		Type:    "ErrOrderUnauthorized",
		Message: "Model isn't allowed to be ordered by this field",
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
	ErrFileTooLarge = FullError{
		Type:     "ErrFileTooLarge",
		Message:  "The specified file is too large, must be < 20MB",
		Title:    "File is too large",
		HelpText: "The specified file is too large, must be < 20MB",
	}
	ErrUploadTokenInvalid = SimpleError{
		Type:    "ErrUploadTokenInvalid",
		Message: "The given upload token was not valid, please check it and try again",
	}
	ErrGeneratingToken = SimpleError{
		Type:    "ErrGeneratingToken",
		Message: "Unable to generate the requested token",
	}
	ErrNotUploaded = SimpleError{
		Type:    "ErrNotUploaded",
		Message: "The item meant to be uploaded is not present",
	}
	ErrDuplicationFailure = SimpleError{
		Type:    "ErrDuplicationFailure",
		Message: "Unable to duplicate the requested item",
	}
	ErrTagsNotFound = SimpleError{
		Type:    "ErrTagsNotFound",
		Message: "Could not find valid tags for all given tag UUIDs",
	}
	ErrTagAlreadyExists = SimpleError{
		Type:    "ErrTagAlreadyExists",
		Message: "Could not create tag as it already exists",
	}
	ErrCategoryAlreadyExists = SimpleError{
		Type:    "ErrCategoryAlreadyExists",
		Message: "Could not create category as it already exists",
	}
	ErrNoEmailProvided = FullError{
		Type:     "ErrNoEmailProvided",
		Message:  "If not email is provided, you must have a password generated for the user",
		Title:    "No email or generated password",
		HelpText: "You must either have a password generated for the user or provide an email address",
	}
	ErrDelegateFinalised = SimpleError{
		Type:    "ErrDelegateFinalised",
		Message: "Delegate already has a password set, please use reset password to reset it instead",
	}
	ErrCSRFTokenInvalid = SimpleError{
		Type:    "ErrCSRFTokenInvalid",
		Message: "The given CSRF token was blank or invalid",
	}
	ErrLessonNotFound = FullError{
		Type:     "ErrLessonNotFound",
		Message:  "There is no lesson matching the information given",
		Title:    "Could no find the specified lesson",
		HelpText: "Oi, this lesson does not exist. Check the details and try again",
	}
	ErrUnauthorizedToBook = SimpleError{
		Type:    "ErrUnauthorized",
		Message: "You cannot book these courses as you are not authorized",
	}
	ErrNotAllFound = SimpleError{
		Type:    "ErrNotAllFound",
		Message: "Not all items you were looking for were found",
	}
	ErrDelegateDoesNotExist = func(uuid string) *SimpleError {
		return &SimpleError{
			Type:    "ErrDelegateDoesNotExist",
			Message: fmt.Sprintf("The given delegate does not exist: %s", uuid),
		}
	}
	ErrRequiredField = func(name string) *SimpleError {
		return &SimpleError{
			Type:    "ErrRequiredField",
			Message: fmt.Sprintf("The field '%s', is required", name),
		}
	}
	ErrInputValidation = func(paramName string, issue string) *SimpleError {
		return &SimpleError{
			Type:    "ErrInputValidation",
			Message: fmt.Sprintf("The field '%s' is invalid: %s", paramName, issue),
		}
	}
	ErrAlreadyTakenTest = SimpleError{
		Type:    "ErrAlreadyTakenTest",
		Message: "You cannot take this test again",
	}
	ErrNotEnoughAnswersGiven = SimpleError{
		Type:    "ErrNotEnoughAnswersGiven",
		Message: "Not enough answers were given to complete the test, please try again",
	}
	ErrBlogNotFound = func(uuid string) *SimpleError {
		return &SimpleError{
			Type:    "ErrBlogNotFound",
			Message: fmt.Sprintf("The given blog does not exist: %s", uuid),
		}
	}
)
