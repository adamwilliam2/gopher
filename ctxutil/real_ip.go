package ctxutil

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// GetRealIPFromContext get trace-id from context
func GetRealIPFromContext(ctx context.Context) string {
	v, ok := ctx.Value(XRealIP).(string)
	if !ok {
		return ""
	}
	return v
}

// ContextWithXRealIP returns a context.Context with given trace-id value.
func ContextWithXRealIP(ctx context.Context, readIP string) context.Context {
	return context.WithValue(ctx, XRealIP, readIP)
}

// MetadataXRealIP returns a context.Context with given real-ip value.
func MetadataXRealIP(ctx context.Context, readIP string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, metadataXRealIPKey, readIP)
}
