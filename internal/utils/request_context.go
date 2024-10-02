package utils

import (
	"context"
	"github.com/chamodshehanka/go-notification-api/internal/utils/constants"
	"github.com/gin-gonic/gin"
)

func GetRequestIdFromGinCtx(rCtx *gin.Context) string {
	return rCtx.GetHeader("request-id")
}

func GetRequestContext(correlationId string) context.Context {
	ctx := context.Background()
	return context.WithValue(ctx, constants.CorrelationIdKeyName, correlationId)
}

func GetCorrelationIdFromContext(ctx context.Context) string {
	val := ctx.Value(constants.CorrelationIdKeyName).(string)
	return val
}
