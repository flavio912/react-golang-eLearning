package middleware

// SimpleError - A less complex error
type SimpleError struct {
	Type    string // Error type
	Message string // Information for developer
}

func (e *SimpleError) Error() string {
	return e.Message
}

// FullError - Detailed error message for use in resolvers
type FullError struct {
	Type     string // Error type
	Message  string // Information for developer
	Title    string // Title of error for user
	HelpText string // Information for user
}

func (e *FullError) Error() string {
	return e.Type
}

// Extensions - Add resolver extensions to show error details on responses
func (e *FullError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"type":     e.Type,
		"message":  e.Message,
		"title":    e.Title,
		"helpText": e.HelpText,
	}
}
