package config

type DatabaseConfig struct {
	Driver             string
	Username           string
	Password           string
	Host               string
	Port               string
	DataBase           string
	Charset            string
	MaxOpenConnections int
	MaxIdleConnections int
	MaxLifeSeconds     int
}
