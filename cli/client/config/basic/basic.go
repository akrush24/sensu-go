package basic

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/Sirupsen/logrus"
	"github.com/sensu/sensu-go/types"
	"github.com/spf13/pflag"
)

const (
	clusterFilename = "cluster"
	profileFilename = "profile"
)

var logger = logrus.WithFields(logrus.Fields{
	"component": "cli-config",
})

// Config contains the CLI configuration
type Config struct {
	Cluster
	Profile
	path string
}

// Cluster contains the Sensu cluster access information
type Cluster struct {
	APIUrl string `json:"api-url"`
	*types.Tokens
}

// Profile contains the active configuration
type Profile struct {
	Format string `json:"format"`
}

// Load imports the CLI configuration and returns an initialized Config struct
func Load(flags *pflag.FlagSet, path string) *Config {
	config := &Config{}

	// Save the path for later use
	config.path = path

	// Load the profile config file
	if err := config.open(filepath.Join(path, profileFilename)); err != nil {
		logger.Debug(err)
	}

	// Load the cluster config file
	if err := config.open(filepath.Join(path, clusterFilename)); err != nil {
		logger.Debug(err)
	}

	// Load the flags config
	config.flags(flags)

	return config
}

func (c *Config) flags(flags *pflag.FlagSet) {
	if flags == nil {
		return
	}

	// Set the API URL
	if value, err := flags.GetString("api-url"); err == nil && value != "" {
		c.Cluster.APIUrl = value
	}
}

func (c *Config) open(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(content, c)
}
