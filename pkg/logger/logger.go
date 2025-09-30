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

	encCfg := zap.NewProductionEncoderConfig()
	encCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	var encoder zapcore.Encoder
	if logCfg.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encCfg)
	} else {
		encoder = zapcore.NewConsoleEncoder(encCfg)
	}

	var ws zapcore.WriteSyncer
	if logCfg.Output == "file" {
		// ensure folder exists
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

		ws = zapcore.AddSync(lumberjackLogger)
	} else {
		ws = zapcore.AddSync(os.Stdout)
	}

	core := zapcore.NewCore(encoder, ws, lvl)
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
