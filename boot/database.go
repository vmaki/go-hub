package boot

import (
	"go-hub/config"
	"go-hub/pkg/database"
)

func SetupDB() {
	db := config.Cfg.DataBase

	database.InitDatabase(
		db.Driver,
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.DataBase,
		db.Charset,
		db.MaxOpenConnections,
		db.MaxIdleConnections,
		db.MaxLifeSeconds,
	)

	// database.DB.AutoMigrate(&model.UserModel{})
}
