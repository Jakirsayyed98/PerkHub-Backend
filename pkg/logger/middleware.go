package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestLogger logs all API requests and responses
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		reqID := uuid.New().String()
		c.Set("req_id", reqID)

		var reqBody map[string]interface{}
		// Read request body if enabled
		if Cfg().RequestBody && c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			_ = json.Unmarshal(bodyBytes, &reqBody)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // reset body for handlers
		}

		// Capture response using Gin writer
		writer := &bodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = writer

		c.Next() // process request

		latency := time.Since(start).Seconds()
		status := c.Writer.Status()

		var respBody map[string]interface{}
		if Cfg().ResponseBody && writer.body.Len() > 0 {
			_ = json.Unmarshal(writer.body.Bytes(), &respBody)
		}

		logData := LogData{
			Message:   fmt.Sprintf("%s %s", c.Request.Method, c.FullPath()),
			Request:   reqBody,
			Response:  respBody,
			StartTime: start,
			EndTime:   time.Now(),
			Latency:   latency,
			UserID:    c.GetString("user_id"),
			ReqID:     reqID,
		}

		// -------------------------
		// Log by HTTP status level
		// -------------------------
		switch {
		case status >= 400:
			LogError(logData) // anything >=400 is treated as error
		default:
			LogInfo(logData)
		}

		// -------------------------
		// Log slow requests
		// -------------------------
		if latency > Cfg().LatencyThreshold {
			LogWarn(LogData{
				Message:   "slow request detected",
				Latency:   latency,
				ReqID:     reqID,
				UserID:    c.GetString("user_id"),
				Request:   reqBody,
				Response:  respBody,
				StartTime: start,
				EndTime:   time.Now(),
			})
		}
	}
}

// bodyWriter intercepts response body
type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
