package database

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var (
	DB    *gorm.DB
	SQLDB *sql.DB
)

func InitDatabase(driver, username, password, host, port, database, charset string, maxOpenConnections, maxIdleConnections, maxLifeSeconds int) {
	var dbConfig gorm.Dialector

	switch driver {
	case "mysql":
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			username,
			password,
			host,
			port,
			database,
			charset,
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	case "sqlite":
		database := database
		dbConfig = sqlite.Open(database)
	default:
		panic(errors.New("该数据库类型暂不支持"))
	}

	// 连接数据库，并设置 GORM 的日志模式
	connect(dbConfig, logger.Default.LogMode(logger.Info))

	// 设置最大连接数
	SQLDB.SetMaxOpenConns(maxOpenConnections)
	// 设置最大空闲连接数
	SQLDB.SetMaxIdleConns(maxIdleConnections)
	// 设置每个链接的过期时间
	SQLDB.SetConnMaxLifetime(time.Duration(maxLifeSeconds) * time.Second)
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
