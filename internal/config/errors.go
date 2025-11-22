package config

import "errors"

var (
	ErrMissingGitLabHost  = errors.New("gitlab-host is required")
	ErrMissingGitLabUser  = errors.New("gitlab-user is required")
	ErrMissingGitLabToken = errors.New("gitlab-token is required")
)
