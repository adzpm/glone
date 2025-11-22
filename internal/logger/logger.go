package logger

import (
	"github.com/charmbracelet/log"
)

// New creates and initializes a new logger instance with options
func New(opts ...Option) *log.Logger {
	options := defaultOptions()
	for _, opt := range opts {
		opt(options)
	}

	l := log.New(options.Output)
	l.SetLevel(options.Level)
	l.SetReportTimestamp(options.ReportTimestamp)
	l.SetReportCaller(options.ReportCaller)
	return l
}
