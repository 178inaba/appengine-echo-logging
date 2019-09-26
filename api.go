package log

import (
	"golang.org/x/net/context"
)

// Debugf formats its arguments according to the format, analogous to fmt.Printf,
// and records the text as a log message at Debug level. The message will be associated
// with the request linked with the provided context.
func Debugf(ctx context.Context, format string, args ...interface{}) {
}

// Infof is like Debugf, but at Info level.
func Infof(ctx context.Context, format string, args ...interface{}) {
}

// Warningf is like Debugf, but at Warning level.
func Warningf(ctx context.Context, format string, args ...interface{}) {
}

// Errorf is like Debugf, but at Error level.
func Errorf(ctx context.Context, format string, args ...interface{}) {
}

// Criticalf is like Debugf, but at Critical level.
func Criticalf(ctx context.Context, format string, args ...interface{}) {
}
