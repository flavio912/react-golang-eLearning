package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAddressesByIDs(t *testing.T) {
	prepareTestDatabase()

	tests := []struct {
		name   string
		ids    []uint
		retLen int
	}{
		{"Empty IDs", []uint{}, 0},
		{"One ID", []uint{0}, 1},
		{"Two IDs", []uint{0, 1}, 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			addresses, err := usersRepo.GetAddressesByIDs(test.ids)
			assert.Nil(t, err)
			assert.Len(t, addresses, test.retLen)
		})
	}
}
