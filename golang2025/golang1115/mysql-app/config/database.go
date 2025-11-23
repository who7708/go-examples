package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DB 全局数据库连接实例
var DB *sql.DB

// InitDB 初始化数据库连接
func InitDB() (*sql.DB, error) {
	// 数据库连接参数 - 请根据实际情况修改
	username := "root"
	password := "root"
	host := "localhost"
	port := "3309"
	database := "test"

	// 构建DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Asia%%2FShanghai&timeout=10s&readTimeout=30s&writeTimeout=30s",
		username, password, host, port, database)

	// 打开数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("打开数据库连接失败: %v", err)
	}

	// 设置连接池配置
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	// 测试连接
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("数据库连接测试失败: %v", err)
	}

	// 设置全局DB实例
	DB = db

	log.Println("数据库连接池初始化完成")
	return db, nil
}

// GetDB 获取数据库连接实例
func GetDB() *sql.DB {
	return DB
}
