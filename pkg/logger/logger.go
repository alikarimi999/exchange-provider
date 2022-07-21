package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(agent, msg string)
	Info(agent, msg string)
	Warn(agent, msg string)
	Error(agent, msg string)
	Fatal(agent, msg string)
}

type logger struct {
	logger *zap.Logger
}

func NewLogger(logfile string, verbose bool) Logger {

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	logFile, _ := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)

	var core zapcore.Core

	switch verbose {
	case true:
		core = zapcore.NewTee(
			zapcore.NewCore(fileEncoder, writer, zapcore.InfoLevel),
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		)
	case false:
		core = zapcore.NewCore(fileEncoder, writer, zapcore.WarnLevel)
	}

	l := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))

	return &logger{
		logger: l,
	}
}

func (l *logger) Debug(agent, msg string) {
	l.logger.Debug(wrap(agent, msg))
}

func (l *logger) Info(agent, msg string) {
	l.logger.Info(wrap(agent, msg))
}

func (l *logger) Warn(agent, msg string) {
	l.logger.Warn(wrap(agent, msg))
}

func (l *logger) Error(agent, msg string) {
	l.logger.Error(wrap(agent, msg))
}

func (l *logger) Fatal(agent, msg string) {
	l.logger.Fatal(wrap(agent, msg))
}

func wrap(a, m string) string {
	return fmt.Sprintf("[%s] %s", a, m)
}
