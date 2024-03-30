package config

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	require := require.New(t)

	err := os.Setenv("SCYLLA_HOSTS", "host1.com,host2.com")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("SCYLLA_MIGRATIONS_DIR", "./cql")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("SCYLLA_KEYSPACE", "reporting")
	if err != nil {
		log.Fatal(err)
	}
	actual, err := Load()
	require.NoError(err)

	expected := Config{
		ScyllaHosts:         []string{"host1.com", "host2.com"},
		ScyllaKeyspace:      "reporting",
		ScyllaMigrationsDir: "./cql",
	}
	require.Equal(expected, actual)
}
