package utils

import (
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

type TelemetryHook struct{}

func (h *TelemetryHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *TelemetryHook) Fire(e *logrus.Entry) error {
	span := trace.SpanFromContext(e.Context)
	spanContext := span.SpanContext()
	traceId := spanContext.TraceID()
	spanId := spanContext.SpanID()

	if _, ok := e.Data["traceId"]; !ok && traceId.IsValid() {
		e.Data["traceId"] = traceId.String()
	}
	if _, ok := e.Data["spanId"]; !ok && spanId.IsValid() {
		e.Data["spanId"] = spanId.String()
	}
	return nil
}

func NewLogger(cfg *Config) *logrus.Logger {
	logger := logrus.New()

	if level, err := logrus.ParseLevel(cfg.LogLevel); err == nil {
		logger.SetLevel(level)
	}

	logger.Formatter = &logrus.JSONFormatter{}
	logger.AddHook(&TelemetryHook{})
	return logger
}
