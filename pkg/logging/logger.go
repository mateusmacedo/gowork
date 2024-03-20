package logging

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	SetLevel(level string)
	Console() *zap.Logger
}

var levelType string

func SetLevel(level string) {
	levelType = level
}

func Console(ctx context.Context, monitoring bool) *zap.Logger {
	level, err := zap.ParseAtomicLevel(levelType)
	if err != nil {
		level, _ = zap.ParseAtomicLevel("info")
	}

	config := zap.NewProductionConfig()
	config.Level = level
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	config.DisableStacktrace = true
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	if ctx != nil {
		for ctxKey, fieldEntry := range contextLogFields {
			if val := ctx.Value(ctxKey); val != nil {
				logger = marshalLogField(logger, val, fieldEntry)
			}
		}
	}

	if monitoring {
		hostname, _ := os.Hostname()
		logger = logger.With(zap.String("host", hostname))
	}

	defer func() {
		if err := logger.Sync(); err != nil {
			return
		}
	}()
	return logger
}
