package boot

import (
	cfg "go-hub/config"
	"go-hub/pkg/config"
)

func SetupConfig(env string) {
	config.LoadEnv(env, &cfg.Cfg)
}
