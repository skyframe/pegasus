package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level int32

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

func (l Level) ToZap() zap.AtomicLevel {
	switch l {
	case PanicLevel:
		return zap.NewAtomicLevelAt(zapcore.PanicLevel)
	case FatalLevel:
		return zap.NewAtomicLevelAt(zapcore.FatalLevel)
	case ErrorLevel:
		return zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case WarnLevel:
		return zap.NewAtomicLevelAt(zapcore.WarnLevel)
	default:
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case DebugLevel:
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case TraceLevel:
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}
}
