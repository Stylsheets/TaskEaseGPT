package log

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is a custom logger struct that embeds logr.Logger.
type Logger struct {
	logr.Logger
}

// Options is a struct for configuring the Logger.
type Options struct {
	Development bool
	Level       int
	Debug       bool
}

// New returns a new instance of Logger with the given options.
func New(opts Options) *Logger {
	var zapLogger *zap.Logger
	var err error

	if opts.Development {
		zapLogger, err = zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
	} else {
		zapLogger, err = zap.NewProduction()
		if err != nil {
			panic(err)
		}
	}

	if opts.Debug {
		zapLogger = zapLogger.WithOptions(zap.AddStacktrace(zap.ErrorLevel))
	}

	zapLogger = zapLogger.WithOptions(zap.IncreaseLevel(zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.Level(opts.Level)
	})))

	logger := zapr.NewLogger(zapLogger)

	return &Logger{
		Logger: logger,
	}
}
