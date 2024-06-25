package ctxutil

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/xid"
	"google.golang.org/grpc/metadata"
)

// GetTraceIDFromContext get trace-id from context
func GetTraceIDFromContext(ctx context.Context) string {
	v, ok := ctx.Value(XTraceID).(string)
	if !ok {
		return uuid.New().String()
	}
	return v
}

func TraceIDWithContext(ctx context.Context) context.Context {
	traceID := GetTraceIDFromContext(ctx)
	if traceID == "" {
		traceID = NewTraceID()
	}

	return ContextWithXTraceID(ctx, traceID)
}

// ContextWithXTraceID returns a context.Context with given trace-id value.
func ContextWithXTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, XTraceID, traceID)
}

// MetadataXTraceID returns a context.Context with given trace-id value.
func MetadataXTraceID(ctx context.Context, requestID string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, metadataXTraceIDKey, requestID)
}

// GetTraceIDFromContext get trace-id from context
func MetadataFromContext(ctx context.Context) string {
	var requestID string
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return NewTraceID()
	}
	requestIDMeta, ok := md[metadataXTraceIDKey]
	if !ok || len(requestIDMeta) == 0 {
		requestID = NewTraceID()
	} else {
		requestID = requestIDMeta[0]
	}

	return requestID
}

func NewTraceID() string {
	return xid.New().String()
}
