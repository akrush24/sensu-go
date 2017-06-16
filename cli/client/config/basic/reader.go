package basic

import "github.com/sensu/sensu-go/types"

// APIUrl returns the active cluster API URL
func (c *Config) APIUrl() string {
	return c.Cluster.APIUrl
}

// Format returns the user's preferred format
func (c *Config) Format() string {
	return c.Profile.Format
}

// Tokens returns the active cluster JWT
func (c *Config) Tokens() *types.Tokens {
	return c.Cluster.Tokens
}
