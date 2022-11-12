package boot

import (
	"go-hub/config"
	"go-hub/pkg/database"
)

func SetupDB() {
	db := config.Cfg.DataBase

	database.InitDatabase(
		db.Dsn,
		db.MaxOpenConnections,
		db.MaxIdleConnections,
		db.MaxLifeSeconds,
	)

	// database.DB.AutoMigrate(&model.UserModel{})
}
