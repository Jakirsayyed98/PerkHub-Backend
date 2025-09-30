package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type LogConfig struct {
	Level            string  // debug, info, warn, error
	Format           string  // json, console
	Output           string  // stdout, file
	FilePath         string  // base path without date
	RequestBody      bool    // log request body
	ResponseBody     bool    // log response body
	LatencyThreshold float64 // slow request threshold in seconds
	MaxSize          int     // MB per file
	MaxBackups       int     // number of rotated files to keep
	MaxAge           int     // days to keep logs
	Compress         bool    // compress old files
}

func LoadLogConfig() *LogConfig {
	return &LogConfig{
		Level:            getEnv("LOG_LEVEL", "info"),
		Format:           getEnv("LOG_FORMAT", "json"),
		Output:           getEnv("LOG_OUTPUT", "stdout"),
		FilePath:         getEnv("LOG_FILE_PATH", "./logs/perkhub"), // base path, date added dynamically
		RequestBody:      getBool("LOG_REQUEST_BODY", true),
		ResponseBody:     getBool("LOG_RESPONSE_BODY", true),
		LatencyThreshold: getFloat("LOG_LATENCY_THRESHOLD", 2.0),
		MaxSize:          getInt("LOG_MAX_SIZE", 50),
		MaxBackups:       getInt("LOG_MAX_BACKUPS", 30),
		MaxAge:           getInt("LOG_MAX_AGE", 30),
		Compress:         getBool("LOG_COMPRESS", true),
	}
}

// helpers
func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func getBool(key string, defaultVal bool) bool {
	if val := os.Getenv(key); val != "" {
		return strings.ToLower(val) == "true"
	}
	return defaultVal
}

func getInt(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		i, err := strconv.Atoi(val)
		if err != nil {
			log.Printf("invalid int for %s: %v", key, err)
			return defaultVal
		}
		return i
	}
	return defaultVal
}

func getFloat(key string, defaultVal float64) float64 {
	if val := os.Getenv(key); val != "" {
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			log.Printf("invalid float for %s: %v", key, err)
			return defaultVal
		}
		return f
	}
	return defaultVal
}

// GetDailyLogFile returns the log filename for today
func (cfg *LogConfig) GetDailyLogFile() string {
	date := time.Now().Format("2006-01-02")
	return cfg.FilePath + "-" + date + ".log"
}
