package database

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var (
	DB    *gorm.DB
	SQLDB *sql.DB
)

func InitDatabase(dsn string, maxOpenConnections, maxIdleConnections, maxLifeSeconds int) {
	dbConfig := mysql.New(mysql.Config{
		DSN: dsn,
	})

	// 连接数据库，并设置 GORM 的日志模式
	connect(dbConfig, logger.Default.LogMode(logger.Info))

	SQLDB.SetMaxOpenConns(maxOpenConnections)                             // 设置最大连接数
	SQLDB.SetMaxIdleConns(maxIdleConnections)                             // 设置最大空闲连接数
	SQLDB.SetConnMaxLifetime(time.Duration(maxLifeSeconds) * time.Second) // 设置每个链接的过期时间
}

func connect(dbConfig gorm.Dialector, _logger logger.Interface) {
	var err error

	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})

	if err != nil {
		panic("连接数据库失败, err: " + err.Error())
	}

	SQLDB, err = DB.DB()
	if err != nil {
		panic("获取数据库失败, err: " + err.Error())
	}
}
