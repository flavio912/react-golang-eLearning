package middleware_test

import (
	"testing"

	"github.com/getsentry/sentry-go"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

var ordersRepository = middleware.NewOrdersRepository(&logging.Logger{Hub: &sentry.Hub{}})

func TestFulfilPendingOrder(t *testing.T) {
	prepareTestDatabase()
	t.Run("Fulfil order successfully", func(t *testing.T) {
		_, err := ordersRepository.FulfilPendingOrder("supersecretclientsecret")
		assert.Nil(t, err)
	})
}

func TestCancelPendingOrder(t *testing.T) {
	prepareTestDatabase()

	t.Run("Cancels and can't be fulfiled", func(t *testing.T) {
		err := ordersRepository.CancelPendingOrder("supersecretclientsecret")
		assert.Nil(t, err)

		_, fulfilErr := ordersRepository.FulfilPendingOrder("supersecretclientsecret")
		assert.Equal(t, &errors.ErrNotFound, fulfilErr)
	})
}
