package gitlab

import (
	"github.com/charmbracelet/log"
)

// ClientOptions holds GitLab client configuration options
type ClientOptions struct {
	Logger   *log.Logger
	BaseURL  string
	SkipAuth bool
}

// ClientOption is a function that modifies ClientOptions
type ClientOption func(*ClientOptions)

// WithLogger sets the logger
func WithLogger(logger *log.Logger) ClientOption {
	return func(o *ClientOptions) {
		o.Logger = logger
	}
}

// WithBaseURL sets a custom base URL
func WithBaseURL(url string) ClientOption {
	return func(o *ClientOptions) {
		o.BaseURL = url
	}
}

// WithSkipAuth skips authentication check
func WithSkipAuth(skip bool) ClientOption {
	return func(o *ClientOptions) {
		o.SkipAuth = skip
	}
}

// defaultClientOptions returns default client options
func defaultClientOptions() *ClientOptions {
	return &ClientOptions{
		Logger:   nil,
		BaseURL:  "",
		SkipAuth: false,
	}
}
