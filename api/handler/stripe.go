package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"

	"github.com/stripe/stripe-go"
)

var ordersRepository = middleware.NewOrdersRepository(&logging.Logger{Hub: sentry.CurrentHub()})

func ServeStripeWebook() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		const MaxBodyBytes = int64(65536)
		req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
		payload, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		// TODO: Verify webhook - https://stripe.com/docs/payments/handling-payment-events#signature-checking

		event := stripe.Event{}

		if err := json.Unmarshal(payload, &event); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse webhook body json: %v\n", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Unmarshal the event data into an appropriate struct depending on its Type
		switch event.Type {
		case "payment_intent.succeeded":
			var paymentIntent stripe.PaymentIntent
			err := json.Unmarshal(event.Data.Raw, &paymentIntent)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if err := ordersRepository.FulfilPendingOrder(paymentIntent.ClientSecret); err != nil {
				sentry.CaptureException(err)
				sentry.CaptureMessage(fmt.Sprintf("Unable to fulfil order from stripe: %s", paymentIntent.ClientSecret))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

		case "payment_intent.canceled":
			var paymentIntent stripe.PaymentIntent
			err := json.Unmarshal(event.Data.Raw, &paymentIntent)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if err := ordersRepository.CancelPendingOrder(paymentIntent.ClientSecret); err != nil {
				sentry.CaptureException(err)
				sentry.CaptureMessage(fmt.Sprintf("Unable to cancel order from stripe: %s", paymentIntent.ClientSecret))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

		default:
			// fmt.Fprintf(os.Stderr, "Unexpected event type: %s\n", event.Type)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
