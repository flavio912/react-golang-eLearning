package logging

import (
	"context"
	"fmt"

	"github.com/golang/glog"

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
func CaptureException(ctx context.Context, err error, level sentry.Level) {
	HubScope(ctx, func(hub *sentry.Hub, scope *sentry.Scope) {
		scope.SetLevel(level)
		hub.CaptureException(err)
	})
}

// CaptureMessage is a wrapper around sentry.CaptureMessage
func CaptureMessage(ctx context.Context, message string, level sentry.Level) {
	HubScope(ctx, func(hub *sentry.Hub, scope *sentry.Scope) {
		scope.SetLevel(level)
		hub.CaptureMessage(message)
	})
}

// CaptureEvent is a wrapper around sentry.CaptureEvent
func CaptureEvent(ctx context.Context, event *sentry.Event, level sentry.Level) {
	HubScope(ctx, func(hub *sentry.Hub, scope *sentry.Scope) {
		scope.SetLevel(level)
		hub.CaptureEvent(event)
	})
}

func glogAtLevel(level sentry.Level, message string) {
	switch level {
	case sentry.LevelFatal:
		glog.Errorf("**Fatal** - %s", message)
	case sentry.LevelError:
		glog.Error(message)
	case sentry.LevelWarning:
		glog.Warning(message)
	case sentry.LevelInfo:
		glog.Info(message)
	case sentry.LevelDebug:
		glog.Info(message)
	}
}

// LogMessage a wrapper around glog and sentry to log in both places
func LogMessage(ctx context.Context, level sentry.Level, message string) {
	glogAtLevel(level, message)
	CaptureMessage(ctx, message, level)
}

// LogException logs to glog and sentry
func LogException(ctx context.Context, level sentry.Level, err error) {
	glogAtLevel(level, err.Error())
	CaptureException(ctx, err, level)
}

// Log logs both a sentry exception as well as a glog message
func Log(ctx context.Context, level sentry.Level, message string, err error) {
	glogAtLevel(level, fmt.Sprintf("%s : %s", message, err.Error()))
	CaptureException(ctx, err, level)
}
