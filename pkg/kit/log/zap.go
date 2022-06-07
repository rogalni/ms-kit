package log

import (
	"context"

	"github.com/rogalni/ms-kit/pkg/kit/config"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func Setup() {
	if config.Kit.IsDevMode {
		log, _ = zap.NewDevelopment(zap.AddCallerSkip(1))
	} else {
		log, _ = zap.NewProduction(zap.AddCallerSkip(1), zap.Fields(defaultFields()...))
	}
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
			zap.String("trace_id", tl.sc.TraceID().String()),
			zap.String("span_id", tl.sc.SpanID().String()),
		}
	}
	return []zapcore.Field{}
}

func defaultFields() []zapcore.Field {
	return []zapcore.Field{
		zap.String("service", config.Kit.ServiceName),
	}
}
