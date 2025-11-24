// 在 api/main.go 中
package main

import (
	"example.com/service" // 直接引用工作区内的模块
	"example.com/utils"
)

func main() {
	// 使用 service 和 utils 中的功能
	svc := service.NewService()
	messago work syncge := utils.FormatMessage("Hello")
	// ...
}
