// Package config manages configurations.
package config

// Config contains app config.
type Config struct {
	Address string `env:"ADDRESS" json:"address"`
}

// NewConfig creates config instance.
func NewConfig() *Config {
	return &Config {
		Address: Address,
	}
}