package config

type LoggerConfig struct {
	Path      string
	Level     string
	Type      string
	MaxSize   int
	MaxBackup int
	MaxAge    int
	Compress  bool
}
