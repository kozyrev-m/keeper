// Package config manages configurations.
package config

// Config contains app config.
type Config struct {
	Address string `env:"ADDRESS" json:"address"`
	DatabaseDSN string `env:"DATABASE_DSN" json:"database_dsn"`
	SessionKey string `env:"SESSION_KEY" json:"session_key"`
}

// NewConfig creates config instance.
func NewConfig() *Config {
	return &Config {
		Address: Address,
		DatabaseDSN: DatabaseDSN,
		SessionKey: SessionKey,
	}
}