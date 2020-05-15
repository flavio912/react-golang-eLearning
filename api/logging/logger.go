package logging

import (
	"context"
	"fmt"

	"github.com/getsentry/sentry-go"
)

type Logger struct {
	Hub *sentry.Hub
}

func GetLoggerFromCtx(ctx context.Context) Logger {
	if hub := sentry.GetHubFromContext(ctx); hub != nil {
		return Logger{Hub: hub}
	}

	return Logger{}
}

// HubScope runs a function with the hub scope
func (l *Logger) HubScope(scopedFunction func(*sentry.Hub, *sentry.Scope)) {
	if l.Hub != nil {
		l.Hub.WithScope(func(scope *sentry.Scope) {
			scopedFunction(l.Hub, scope)
		})
	}
}

// CaptureException is a wrapper around sentry.CaptureException
func (l *Logger) CaptureException(err error, level sentry.Level) {
	l.HubScope(func(hub *sentry.Hub, scope *sentry.Scope) {
		scope.SetLevel(level)
		hub.CaptureException(err)
	})
}

// CaptureMessage is a wrapper around sentry.CaptureMessage
func (l *Logger) CaptureMessage(message string, level sentry.Level) {
	l.HubScope(func(hub *sentry.Hub, scope *sentry.Scope) {
		scope.SetLevel(level)
		hub.CaptureMessage(message)
	})
}

// CaptureEvent is a wrapper around sentry.CaptureEvent
func (l *Logger) CaptureEvent(event *sentry.Event, level sentry.Level) {
	l.HubScope(func(hub *sentry.Hub, scope *sentry.Scope) {
		scope.SetLevel(level)
		hub.CaptureEvent(event)
	})
}

// LogMessage a wrapper around glog and sentry to log in both places
func (l *Logger) LogMessage(level sentry.Level, message string) {
	glogAtLevel(level, message)
	l.CaptureMessage(message, level)
}

// LogException logs to glog and sentry
func (l *Logger) LogException(level sentry.Level, err error) {
	glogAtLevel(level, err.Error())
	l.CaptureException(err, level)
}

// Log logs both a sentry exception as well as a glog message
func (l *Logger) Log(level sentry.Level, err error, message string) {
	glogAtLevel(level, fmt.Sprintf("%s : %s", message, err.Error()))
	l.CaptureException(err, level)
}

// Log logs both a sentry exception as well as a glog message
func (l *Logger) Logf(level sentry.Level, err error, message string, args ...interface{}) {
	m := fmt.Sprintf(message, args...)
	glogAtLevel(level, fmt.Sprintf("%s : %s", m, err.Error()))
	l.Hub.AddBreadcrumb(&sentry.Breadcrumb{Type: "Message", Message: m}, nil)
	l.CaptureException(err, level)
}
