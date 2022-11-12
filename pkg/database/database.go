package database

import (
	"database/sql"
	"go-hub/config"
	sysLogger "go-hub/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type DBInfo struct {
	DB    *gorm.DB
	SQLDB *sql.DB
}

var DBCollections map[string]*DBInfo

func InitDatabase(dbs map[string]config.DatabaseConfig) {
	for name, v := range dbs {
		if DBCollections == nil {
			DBCollections = make(map[string]*DBInfo, len(config.Cfg.DataBases))
		}

		db, sqlDb, err := Connect(v.Dsn, sysLogger.NewGormLogger())
		if err != nil {
			panic(err)
		}

		dbStruct := &DBInfo{
			DB:    db,
			SQLDB: sqlDb,
		}

		DBCollections[name] = dbStruct
		DBCollections[name].SQLDB.SetMaxOpenConns(v.MaxOpenConnections)
		DBCollections[name].SQLDB.SetMaxIdleConns(v.MaxIdleConnections)
		DBCollections[name].SQLDB.SetConnMaxLifetime(time.Duration(v.MaxLifeSeconds) * time.Second)
	}
}

func Connect(dsn string, _logger logger.Interface) (*gorm.DB, *sql.DB, error) {
	dbConfig := mysql.New(mysql.Config{
		DSN: dsn,
	})

	db, err := gorm.Open(
		dbConfig,
		&gorm.Config{
			Logger: _logger,
		},
	)
	if err != nil {
		return nil, nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	return db, sqlDb, nil
}

func DB(name ...string) *gorm.DB {
	if len(name) > 0 {
		if collect, ok := DBCollections[name[0]]; ok {
			return collect.DB
		}

		return nil
	}

	return DBCollections["default"].DB
}
