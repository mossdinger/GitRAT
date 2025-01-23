package main

// Configuration Struct
type Config struct {
	CommandRepoUrl  string // The input repository URL
	OutputRepoUrl   string // The output repository URL
	RefreshInterval int    // Delay in seconds
	IsIgnoringError bool   // Determines error handling behavior
	Timeout         int    // Timeout in seconds
}

// LoadConfig initializes and returns a default configuration
func LoadConfig() *Config {
	return &Config{
		// RepoUrlFormat: https://{Token}:x-oauth-basic@github.com/{github_username}/{repo_name}
		CommandRepoUrl:  "",
		OutputRepoUrl:   "",
		RefreshInterval: 10,   // seconds
		Timeout:         0,    // seconds
		IsIgnoringError: true, // continue even if error
	}
}
