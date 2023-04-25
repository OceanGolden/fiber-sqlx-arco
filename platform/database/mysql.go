package database

import (
	"fiber-sqlx-arco/pkg/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

func MySQLConnection() (*sqlx.DB, error) {
	mysqlConnURL, err := utils.ConnectionURLBuilder("mysql")
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Connect("mysql", mysqlConnURL)
	if err != nil {
		return nil, fmt.Errorf("错误,无法连接数据库, %w", err)
	}

	db.SetMaxOpenConns(100)          // 最大打开的连接数 不超过数据库服务自身支持的并发连接数
	db.SetMaxIdleConns(50)           // 最大闲置的连接数 一般建议maxIdleConns的值为MaxOpenConns的1/2
	db.SetConnMaxLifetime(time.Hour) // 连接的最大可复用时间 不超过数据库的超时参数值。

	return db, nil
}
