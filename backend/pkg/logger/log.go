package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

type RequestID string

var RequestIDKey RequestID = "request-id"

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
		ForceQuote:    true,
	})
}

func Info(ctx context.Context, caller, format string, values ...interface{}) {
	requestID := "-"
	ctxVal := ctx.Value(RequestIDKey)
	if ctxVal != nil {
		requestID = ctxVal.(string)
	}

	logrus.WithFields(logrus.Fields{
		"caller":     caller,
		"request-id": requestID,
	}).Infof(format, values...)
}

func Error(ctx context.Context, caller, format string, values ...interface{}) {
	requestID := "-"
	ctxVal := ctx.Value(RequestIDKey)
	if ctxVal != nil {
		requestID = ctxVal.(string)
	}

	logrus.WithFields(logrus.Fields{
		"caller":     caller,
		"request-id": requestID,
	}).Errorf(format, values...)
}

func Warn(ctx context.Context, caller, format string, values ...interface{}) {
	requestID := "-"
	ctxVal := ctx.Value(RequestIDKey)
	if ctxVal != nil {
		requestID = ctxVal.(string)
	}

	logrus.WithFields(logrus.Fields{
		"caller":     caller,
		"request-id": requestID,
	}).Warnf(format, values...)
}

func Fatal(ctx context.Context, caller, format string, values ...interface{}) {
	requestID := "-"
	ctxVal := ctx.Value(RequestIDKey)
	if ctxVal != nil {
		requestID = ctxVal.(string)
	}

	logrus.WithFields(logrus.Fields{
		"caller":     caller,
		"request-id": requestID,
	}).Fatalf(format, values...)
}
