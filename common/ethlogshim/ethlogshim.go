package ethlogshim

import (
	"context"
	"log/slog"
	"os"

	ethlog "github.com/zircuit-labs/l2-geth-public/log"
)

type EthLoggerShim struct {
	logger *slog.Logger
}

func (l *EthLoggerShim) With(ctx ...any) ethlog.Logger {
	return l.New(ctx...)
}

func (l *EthLoggerShim) Log(level slog.Level, msg string, ctx ...any) {
	l.logger.Log(context.Background(), level, msg, ctx...)
}

func (l *EthLoggerShim) Write(level slog.Level, msg string, ctx ...any) {
	l.Log(level, msg, ctx...)
}

func (l *EthLoggerShim) Enabled(ctx context.Context, level slog.Level) bool {
	return l.logger.Enabled(ctx, level)
}

func NewEthLoggerShim(logger *slog.Logger) *EthLoggerShim {
	return &EthLoggerShim{logger: logger}
}

func (l *EthLoggerShim) New(ctx ...interface{}) ethlog.Logger {
	return &EthLoggerShim{logger: l.logger.With(ctx...)}
}

func (l *EthLoggerShim) Handler() slog.Handler {
	// not implemented
	return nil
}

func (l *EthLoggerShim) SetHandler(h slog.Handler) {
	// not implemented
}

func (l *EthLoggerShim) Trace(msg string, ctx ...interface{}) {
	// slog does not support Trace level
	l.logger.Debug(msg, ctx...)
}

func (l *EthLoggerShim) Debug(msg string, ctx ...interface{}) {
	l.logger.Debug(msg, ctx...)
}

func (l *EthLoggerShim) Info(msg string, ctx ...interface{}) {
	l.logger.Info(msg, ctx...)
}

func (l *EthLoggerShim) Warn(msg string, ctx ...interface{}) {
	l.logger.Warn(msg, ctx...)
}

func (l *EthLoggerShim) Error(msg string, ctx ...interface{}) {
	l.logger.Error(msg, ctx...)
}

func (l *EthLoggerShim) Crit(msg string, ctx ...interface{}) {
	// slog does not support Crit level
	l.logger.Error(msg, ctx...)
	os.Exit(1)
}
