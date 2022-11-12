package config

type DatabaseConfig struct {
	Dsn                string
	MaxOpenConnections int
	MaxIdleConnections int
	MaxLifeSeconds     int
}
