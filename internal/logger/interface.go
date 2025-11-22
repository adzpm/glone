package logger

// Logger defines the interface for logging operations
type Logger interface {
	Info(msg any, args ...any)
	Infof(fmt string, args ...any)

	Warn(msg any, args ...any)
	Warnf(fmt string, args ...any)

	Error(msg any, args ...any)
	Errorf(fmt string, args ...any)

	Debug(msg any, args ...any)
	Debugf(fmt string, args ...any)

	Fatal(msg any, args ...any)
	Fatalf(fmt string, args ...any)
}
