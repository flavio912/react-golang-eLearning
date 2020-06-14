package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

func TestFulfilPendingOrder(t *testing.T) {
	prepareTestDatabase()
	t.Run("Fulfil order successfully", func(t *testing.T) {
		err := middleware.FulfilPendingOrder("supersecretclientsecret")
		assert.Nil(t, err)
	})
}

func TestCancelPendingOrder(t *testing.T) {
	prepareTestDatabase()

	t.Run("Cancels and can't be fulfiled", func(t *testing.T) {
		err := middleware.CancelPendingOrder("supersecretclientsecret")
		assert.Nil(t, err)

		fulfilErr := middleware.FulfilPendingOrder("supersecretclientsecret")
		assert.Equal(t, &errors.ErrNotFound, fulfilErr)
	})
}
