package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggerInstance *zap.SugaredLogger

type messageType string

const (
	Info  messageType = "Info"
	Error messageType = "Error"
	Warn  messageType = "Warn"
)

func InitializedLogger() {
	consoleEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:     "time",
		LevelKey:    "level",
		MessageKey:  "msg",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime:  zapcore.TimeEncoderOfLayout("[2006-01-02 15:04]"),
	})

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	)

	loggerSetup := zap.New(core, zap.AddCaller())
	loggerInstance = loggerSetup.Sugar()
}

func LoggerMesssage(message string, msgType messageType) {
	switch msgType {
	case Info:
		loggerInstance.Infof("Message: %s", message)
	case Error:
		loggerInstance.Errorf("Error: %s", message)
	case Warn:
		loggerInstance.Warnf("Error: %s", message)
	default:
		loggerInstance.Debugf("Debug: %s", message)
	}
}
