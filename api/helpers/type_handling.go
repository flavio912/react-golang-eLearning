package helpers

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

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

func StringPointer(str string) *string {
	_string := str
	return &_string
}

func FloatPointer(flo float64) *float64 {
	_float := flo
	return &_float
}

func BoolPointer(boolean bool) *bool {
	_boolean := boolean
	return &_boolean
}

func IntPointer(integer int) *int {
	_integer := integer
	return &_integer
}

func Int32Pointer(integer int32) *int32 {
	_integer := integer
	return &_integer
}

func UintPointer(unsignedInt uint) *uint {
	_unsignedInt := unsignedInt
	return &_unsignedInt
}

func Int32ToUint(integer32 int32) uint {
	if integer32 < 0 {
		return 0
	}
	return uint(integer32)
}

func Int32sToUints(integer32s []int32) []uint {
	var res = make([]uint, len(integer32s))
	for index, val := range integer32s {
		res[index] = Int32ToUint(val)
	}
	return res
}

func UUIDPointer(uuid gentypes.UUID) *gentypes.UUID {
	_uuid := uuid
	return &_uuid
}
