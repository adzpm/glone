package git

import (
	"github.com/adzpm/glup/internal/logger"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// CloneResult represents the result of a clone operation
type CloneResult struct {
	Skipped bool
	Error   error
}

// CloneProject clones a GitLab project to the target directory (legacy function for backward compatibility)
func CloneProject(project *gitlab.Project, targetDir string, token string, lgr logger.Logger) (bool, error) {
	cloner := NewCloner(WithLogger(lgr))
	return cloner.CloneProject(project, targetDir, token)
}
