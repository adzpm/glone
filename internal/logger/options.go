package logger

import (
	"io"
	"os"

	"github.com/charmbracelet/log"
)

// Options holds logger configuration options
type Options struct {
	Level           log.Level
	Output          io.Writer
	ReportTimestamp bool
	ReportCaller    bool
}

// Option is a function that modifies Options
type Option func(*Options)

// WithLevel sets the log level
func WithLevel(level log.Level) Option {
	return func(o *Options) {
		o.Level = level
	}
}

// WithOutput sets the output writer
func WithOutput(w io.Writer) Option {
	return func(o *Options) {
		o.Output = w
	}
}

// WithReportTimestamp sets whether to report timestamp
func WithReportTimestamp(report bool) Option {
	return func(o *Options) {
		o.ReportTimestamp = report
	}
}

// WithReportCaller sets whether to report caller location
func WithReportCaller(report bool) Option {
	return func(o *Options) {
		o.ReportCaller = report
	}
}

// defaultOptions returns default options
func defaultOptions() *Options {
	return &Options{
		Level:           log.InfoLevel,
		Output:          os.Stderr,
		ReportTimestamp: true,
		ReportCaller:    false,
	}
}
