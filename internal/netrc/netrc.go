package netrc

import (
	"github.com/adzpm/glup/internal/config"
	"github.com/charmbracelet/log"
)

// LoadCredentials loads GitLab credentials from .netrc file (legacy function for backward compatibility)
func LoadCredentials(logger *log.Logger) (*config.Config, error) {
	loader, err := NewLoader(WithLogger(logger))
	if err != nil {
		return nil, err
	}
	return loader.LoadCredentials()
}
