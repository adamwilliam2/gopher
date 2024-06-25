package ctxutil

type ctxKey string

const (
	// XTraceID request id
	XTraceID            ctxKey = "x-trace-id"
	metadataXTraceIDKey string = "x-trace-id"

	XRealIP            ctxKey = "x-real-ip"
	metadataXRealIPKey string = "x-real-ip"

	XDeviceID            ctxKey = "x-device-id"
	metadataXDeviceIDKey string = "x-device-id"
)
