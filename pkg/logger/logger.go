package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"PerkHub/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *zap.Logger
var logCfg *config.LogConfig

func Init() {
	logCfg = config.LoadLogConfig()

	var lvl zapcore.Level
	_ = lvl.Set(logCfg.Level)

	// Encoder config (common)
	encCfg := zap.NewProductionEncoderConfig()
	encCfg.TimeKey = "ts"
	encCfg.LevelKey = "level"
	encCfg.CallerKey = "caller"
	encCfg.MessageKey = "msg"
	encCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	// ---------------------------
	// File core (JSON, no colors)
	// ---------------------------
	logDir := filepath.Dir(logCfg.FilePath)
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		_ = os.MkdirAll(logDir, os.ModePerm)
	}

	// generate date-based filename
	date := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("%s-%s.log", logDir+"/perkhub", date)

	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    logCfg.MaxSize,
		MaxBackups: logCfg.MaxBackups,
		MaxAge:     logCfg.MaxAge,
		Compress:   logCfg.Compress,
	}

	fileEncoder := zapcore.NewJSONEncoder(encCfg)
	fileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(lumberjackLogger), lvl)

	// ---------------------------
	// Console core (colored text)
	// ---------------------------
	consoleEncCfg := encCfg
	consoleEncCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncCfg)
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), lvl)

	// ---------------------------
	// Combine file + console
	// ---------------------------
	core := zapcore.NewTee(fileCore, consoleCore)
	log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func Get() *zap.Logger {
	if log == nil {
		Init()
	}
	return log
}

func Cfg() *config.LogConfig {
	return logCfg
}
