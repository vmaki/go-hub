package boot

import (
	"go-hub/config"
	"go-hub/pkg/logger"
)

func SetupLogger() {
	log := config.Cfg.Logger
	logger.InitLogger(log.Level, log.Type, log.Path, log.MaxSize, log.MaxBackup, log.MaxAge, log.Compress)
}
