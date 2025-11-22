package git

import (
	"github.com/charmbracelet/log"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// CloneResult represents the result of a clone operation
type CloneResult struct {
	Skipped bool
	Error   error
}

// CloneProject clones a GitLab project to the target directory (legacy function for backward compatibility)
func CloneProject(project *gitlab.Project, targetDir string, token string, logger *log.Logger) (bool, error) {
	cloner := NewCloner(WithLogger(logger))
	return cloner.CloneProject(project, targetDir, token)
}
