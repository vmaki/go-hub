package boot

import "go-hub/pkg/config"

func SetupConfig(env string) {
	config.LoadEnv(env)
}
