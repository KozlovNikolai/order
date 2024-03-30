package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config is base structure of package
type Config struct {
	ScyllaHosts         []string `envconfig:"SCYLLA_HOSTS"`
	ScyllaKeyspace      string   `envconfig:"SCYLLA_KEYSPACE"`
	ScyllaMigrationsDir string   `envconfig:"SCYLLA_MIGRATIONS_DIR"`
}

// Load loads values into fields of structure Config
func Load() (Config, error) {
	var c Config
	err := envconfig.Process("", &c)
	return c, err
}
