package helpers

// StringNotNilOrEmpty returns true of a string pointer is not nil or ""
func StringNotNilOrEmpty(item *string) bool {
	if item != nil && *item != "" {
		return true
	}
	return false
}

// NilStringToEmpty returns a string or the empty string
func NilStringToEmpty(item *string) string {
	if item == nil {
		return ""
	}
	return *item
}

// NilFloatToZero returns a float or 0
func NilFloatToZero(item *float64) float64 {
	if item == nil {
		return 0
	}
	return *item
}
