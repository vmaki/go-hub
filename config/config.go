package config

var Cfg Configuration

type Configuration struct {
	Application ApplicationConfig
	Logger      LoggerConfig
	DataBase    DatabaseConfig
}
