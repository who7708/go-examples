package main

import (
	"log"
	"mysql-app/config"
	"mysql-app/handlers"
	"net/http"
)

func main() {
	// 初始化数据库连接
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	defer db.Close()
	log.Println("数据库连接成功")

	// 设置路由
	http.HandleFunc("/users", handlers.GetUsers)
	http.HandleFunc("/users/create", handlers.CreateUser)

	// 启动HTTP服务器
	log.Println("服务器启动在 :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
