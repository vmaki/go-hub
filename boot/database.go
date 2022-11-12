package boot

import (
	"go-hub/config"
	"go-hub/pkg/database"
)

func SetupDB() {
	database.InitDatabase(config.Cfg.DataBases)
}
