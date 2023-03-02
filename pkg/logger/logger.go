package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New returns a new logger that writes to file and
// console. Pass in an empty string to only log to console.
// Returns an error if file interaction fails.
func New(file string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(cfg.EncoderConfig)

	if file != "" {
		fileEncoder := zapcore.NewJSONEncoder(cfg.EncoderConfig)

		logFile, err := os.OpenFile(
			file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}

		writer := zapcore.AddSync(logFile)
		defaultLogLvl := zap.NewAtomicLevelAt(zapcore.InfoLevel)

		core := zapcore.NewTee(
			zapcore.NewCore(fileEncoder, writer, defaultLogLvl),
			zapcore.NewCore(
				consoleEncoder, zapcore.Lock(os.Stdout), defaultLogLvl),
		)

		logger := zap.New(core, zap.AddCaller(),
			zap.AddStacktrace(zap.ErrorLevel))
		return logger, nil
	}

	core := zapcore.NewCore(
		consoleEncoder, zapcore.Lock(os.Stdout), zapcore.InfoLevel)

	logger := zap.New(core, zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel))

	return logger, nil
}
