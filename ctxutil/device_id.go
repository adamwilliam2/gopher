package ctxutil

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// GetDeviceIDFromContext get device-id from context
func GetDeviceIDFromContext(ctx context.Context) string {
	v, ok := ctx.Value(XDeviceID).(string)
	if !ok {
		return ""
	}
	return v
}

// ContextWithXDeviceID returns a context.Context with given device-id value.
func ContextWithXDeviceID(ctx context.Context, deviceID string) context.Context {
	return context.WithValue(ctx, XDeviceID, deviceID)
}

// MetadataXDeviceID returns a context.Context with given device-id value.
func MetadataXDeviceID(ctx context.Context, readIP string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, metadataXDeviceIDKey, readIP)
}
