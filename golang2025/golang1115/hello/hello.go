package main

import (
	"log"

	"example.com/greetings"
)

func main() {
	// 设置预定义Logger的属性，包括
	// 日志条目前缀和禁用打印的标志
	//  时间、源文件和行号.
	log.SetPrefix("greetings: ")
	// log.SetFlags(0)

	message, err := greetings.Hello("Gladys")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(message)

	// 一个名字切片.
	names := []string{"Gladys", "Samantha", "Darrin"}
	message1, err1 := greetings.Hellos(names)
	if err1 != nil {
		log.Fatal(err1)
	}
	log.Println(message1)
}

// 声明一个main包。在 Go 中，作为应用程序执行的代码必须在main包中.

// 在 hello 目录中
// 1. 先导入 import "example.com/greetings"
// 2. 运行 go mod edit -replace example.com/greetings=../greetings
// 该命令指定 example.com/greetings 应替换为 ../greetings，用于查找依赖项。运行该命令后，hello 目录中的 go.mod 文件应包含一个 replace 指令:
// 3. 运行 go mod tidy
// 4. 运行 go run .
