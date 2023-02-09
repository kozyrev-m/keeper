package config

// Default configs.
const (
	// http-server address .
	Address = "127.0.0.1:8080"

	// Local postgres db.
	DatabaseDSN = "host=localhost port=5433 dbname=keeper_dev user=admin password=12345 sslmode=disable"

	// Session key.
	SessionKey = "1a2b3c4d5e6fffffffffffffffffffffffffffff"
)
