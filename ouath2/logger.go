package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func registerLogger() (*zap.Logger, error) {
	level := zap.NewAtomicLevel()
	err := level.UnmarshalText([]byte("info"))
	if err != nil {
		return nil, err
	}
	var samplingConfig = &zap.SamplingConfig{
		Initial:    10,
		Thereafter: 10,
	}
	var encoderConfig = zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	zapConfig := &zap.Config{
		Level:             level,
		Sampling:          samplingConfig,
		Encoding:          "json",
		EncoderConfig:     encoderConfig,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stdout"},
		DisableStacktrace: true,
	}

	return zapConfig.Build()
}
