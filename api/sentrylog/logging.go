package sentrylog

import (
	"context"

	"github.com/getsentry/sentry-go"
)

/*
	Some helper functions for logging with sentry so we
	don't get sentry from context every time
*/

// HubScope runs a function with the hub scope
func HubScope(ctx context.Context, scopedFunction func(*sentry.Hub, *sentry.Scope)) {
	if hub := sentry.GetHubFromContext(ctx); hub != nil {
		hub.WithScope(func(scope *sentry.Scope) {
			scopedFunction(hub, scope)
		})
	}
}

// CaptureException is a wrapper around sentry.CaptureException
func CaptureException(ctx context.Context, err error) {
	HubScope(ctx, func(hub *sentry.Hub, scope *sentry.Scope) {
		hub.CaptureException(err)
	})
}

// CaptureMessage is a wrapper around sentry.CaptureMessage
func CaptureMessage(ctx context.Context, message string) {
	HubScope(ctx, func(hub *sentry.Hub, scope *sentry.Scope) {
		hub.CaptureMessage(message)
	})
}

// CaptureEvent is a wrapper around sentry.CaptureEvent
func CaptureEvent(ctx context.Context, event *sentry.Event) {
	HubScope(ctx, func(hub *sentry.Hub, scope *sentry.Scope) {
		hub.CaptureEvent(event)
	})
}
