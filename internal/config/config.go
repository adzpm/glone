package config

// Config holds the application configuration
type Config struct {
	GitLabHost  string
	GitLabUser  string
	GitLabToken string
	Group       string
	TargetDir   string
}

// Validate checks that all required fields are set
func (c *Config) Validate() error {
	if c.GitLabHost == "" {
		return ErrMissingGitLabHost
	}
	if c.GitLabUser == "" {
		return ErrMissingGitLabUser
	}
	if c.GitLabToken == "" {
		return ErrMissingGitLabToken
	}
	return nil
}

// Merge merges non-empty values from other into c
func (c *Config) Merge(other *Config) {
	if other == nil {
		return
	}
	if c.GitLabHost == "" && other.GitLabHost != "" {
		c.GitLabHost = other.GitLabHost
	}
	if c.GitLabUser == "" && other.GitLabUser != "" {
		c.GitLabUser = other.GitLabUser
	}
	if c.GitLabToken == "" && other.GitLabToken != "" {
		c.GitLabToken = other.GitLabToken
	}
}
