package log

import (
	"context"

	"github.com/rogalni/ms-kit/internal/config"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func Setup() {
	f := zap.Fields(defaultFields()...)
	log, _ = zap.NewProduction(f)
}

type TracedLog struct {
	sc trace.SpanContext
}

func Ctx(c context.Context) *TracedLog {
	sc := trace.SpanFromContext(c).SpanContext()
	return &TracedLog{sc: sc}
}

func (tl *TracedLog) Debug(msg string) {
	log.Debug(msg, tracingFields(tl)...)
}

func (tl *TracedLog) Info(msg string) {
	log.Info(msg, tracingFields(tl)...)
}

func (tl *TracedLog) Warn(msg string) {
	log.Warn(msg, tracingFields(tl)...)
}

func (tl *TracedLog) Error(msg string) {
	log.Error(msg, tracingFields(tl)...)
}

func tracingFields(tl *TracedLog) []zapcore.Field {
	if tl.sc.IsValid() {
		return []zapcore.Field{
			zap.String("trace", tl.sc.TraceID().String()),
			zap.String("span", tl.sc.SpanID().String()),
		}
	}
	return []zapcore.Field{}
}

func defaultFields() []zapcore.Field {
	return []zapcore.Field{
		zap.String("service", config.EnvOr(config.EnvServiceName, "ms-kit-service")),
	}
}
