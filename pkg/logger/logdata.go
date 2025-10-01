package logger

import (
	"time"

	"go.uber.org/zap"
)

// LogData is the structured logging payload
type LogData struct {
	Message   string      // main message
	Request   interface{} // optional request body
	Response  interface{} // optional response body
	Error     error       // optional error
	UserID    string      // optional user ID
	ReqID     string      // optional request ID
	StartTime time.Time   // optional start time
	EndTime   time.Time   // optional end time
	Latency   float64     // optional latency in seconds
}

// LogInfo logs informational events (success)
func LogInfo(data LogData) {
	fields := buildFields(data)
	Get().Info(data.Message, fields...)
}

// LogError logs errors
func LogError(data LogData) {
	fields := buildFields(data)
	Get().Error(data.Message, fields...)
}

// LogWarn logs warnings (e.g., slow requests)
func LogWarn(data LogData) {
	fields := buildFields(data)
	Get().Warn(data.Message, fields...)
}

// LogDebug logs debug events
func LogDebug(data LogData) {
	fields := buildFields(data)
	Get().Debug(data.Message, fields...)
}

// buildFields converts LogData to zap fields
func buildFields(data LogData) []zap.Field {
	fields := []zap.Field{}

	if data.ReqID != "" {
		fields = append(fields, zap.String("req_id", data.ReqID))
	}
	if data.UserID != "" {
		fields = append(fields, zap.String("user_id", data.UserID))
	}
	if !data.StartTime.IsZero() {
		fields = append(fields, zap.Time("start_time", data.StartTime))
	}

	endTime := time.Now()
	fields = append(fields, zap.Time("end_time", endTime))

	if data.Latency <= 0 {
		data.Latency = endTime.Sub(data.StartTime).Seconds()
	}
	fields = append(fields, zap.Float64("latency", data.Latency))

	if data.Request != nil {
		fields = append(fields, zap.Any("request", data.Request))
	}
	if data.Response != nil {
		fields = append(fields, zap.Any("response", data.Response))
	}
	if data.Error != nil {
		fields = append(fields, zap.Error(data.Error))
	}

	return fields
}
