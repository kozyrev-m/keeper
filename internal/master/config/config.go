// Package config manages configurations.
package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/kozyrev-m/keeper/internal/master/usage"
)

var (
	configFile  string
	address     string
	databaseDSN string
)

// Config contains app config.
type Config struct {
	Address     string `env:"ADDRESS" json:"address"`
	DatabaseDSN string `env:"DATABASE_DSN" json:"database_dsn"`
	SessionKey  string `env:"SESSION_KEY" json:"session_key"`
}

// NewConfig creates config instance.
func NewConfig() *Config {
	c := &Config{
		Address: Address,
	}

	c.configure()

	return c
}

// configure sets actual values to 'Config'.
func (c *Config) configure() *Config {
	c.parseFlags()

	// 1. get config from file
	c.configFromFile()

	// DEBUG: log.Printf("config from file:%+v", c)

	// 2. get config from flags
	c.configFromFlags()

	// DEBUG: log.Printf("config from flags:%+v", c)

	// 3. get config from env
	c.configFromEnv()

	// DEBUG: log.Printf("config from env:%+v", c)

	return c
}

// configFromFile gets configs from file.
func (c *Config) configFromFile() {
	// define path to config file by env
	if !strings.EqualFold(os.Getenv("CONFIG_KEEPER"), "") {
		configFile = os.Getenv("CONFIG_KEEPER")
	}

	if len(configFile) == 0 {
		return
	}

	configJSON, err := usage.ReadFromFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(configJSON), c); err != nil {
		log.Fatal(err)
	}
}

// configFromFlags gets config from flags.
func (c *Config) configFromFlags() {
	// skip if don't set flag -a
	if address != Address {
		c.Address = address
	}

	// skip if don't set flag -d
	if databaseDSN != "" {
		c.DatabaseDSN = databaseDSN
	}
}

// configFromEnv gets configs from environments.
func (c *Config) configFromEnv() {
	if err := env.Parse(c); err != nil {
		log.Fatal(err)
	}
}

// parseFlags parses flags.
func (c *Config) parseFlags() {
	// define path to config file by flag
	flag.StringVar(&configFile, "c", "", "path to config file")
	flag.StringVar(&configFile, "config", "", "path to config file")

	flag.StringVar(&address, "a", "", "http-server address")
	flag.StringVar(&databaseDSN, "d", "", "database dsn")

	flag.Parse()
}
