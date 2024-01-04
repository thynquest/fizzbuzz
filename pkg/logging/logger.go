package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger : used to log message in the system
var logger *zap.SugaredLogger

type logEvent struct {
	message string
}

var (
	infoMessage    = logEvent{"%s: INFO %s"}
	warningMessage = logEvent{"%s: WARNING %s"}
	errorMessage   = logEvent{"%s: ERROR %s"}
	fatalMessage   = logEvent{"%s: FATAL %s"}
)

func init() {
	if logger == nil {
		lp := zap.NewProductionConfig()
		lp.DisableStacktrace = true
		lp.EncoderConfig.CallerKey = zapcore.OmitKey
		build, _ := lp.Build()
		logger = build.Sugar()
	}
}

// Info : log info message
func Info(name string, messages ...interface{}) {
	logger.Infof(infoMessage.message, name, messages)
}

// Warning :
func Warning(name string, messages ...interface{}) {
	logger.Warnf(warningMessage.message, name, messages)
}

// Error :
func Error(name string, messages ...interface{}) {
	logger.Errorf(errorMessage.message, name, messages)
}

// Fatal :
func Fatal(name string, messages ...interface{}) {
	logger.Fatalf(fatalMessage.message, name, messages)
}
