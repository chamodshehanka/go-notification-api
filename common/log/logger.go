package log

import (
	"context"
	"github.com/chamodshehanka/go-notification-api/internal/utils"
	"github.com/sirupsen/logrus"
)

const (
	SeverityLevel = "SeverityLevel"
	ErrorCode     = "ErrorCode"
	CorrelationId = "CorrelationId"
)

func Fatal(severityLevel string, errorCode string, errorMessage string) {
	logrus.WithFields(logrus.Fields{
		SeverityLevel: severityLevel,
		ErrorCode:     errorCode,
	}).Fatal(errorMessage)
}

func Error(ctx context.Context, severityLevel string, errorCode string, errorMessage string) {
	requestId := utils.GetCorrelationIdFromContext(ctx)
	logrus.WithFields(logrus.Fields{
		SeverityLevel: severityLevel,
		ErrorCode:     errorCode,
		CorrelationId: requestId,
	}).Error(errorMessage)
}

func Info(ctx context.Context, message string) {
	fields := make(map[string]interface{})
	requestId := utils.GetCorrelationIdFromContext(ctx)
	fields[CorrelationId] = requestId
	logrus.WithFields(fields).Info(message)
}

func Debug(ctx context.Context, fields map[string]interface{}, message string) {
	fds := make(map[string]interface{})
	requestId := utils.GetCorrelationIdFromContext(ctx)
	fds[CorrelationId] = requestId
	if fields != nil {
		for k, v := range fields {
			fds[k] = v
		}
	}
	logrus.WithFields(fds).Debug(message)
}
