package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

func TestFulfilPendingOrder(t *testing.T) {
	prepareTestDatabase()
	t.Run("Fulfil order successfully", func(t *testing.T) {
		err := middleware.FulfilPendingOrder("supersecretclientsecret")
		assert.Nil(t, err)
	})
}
