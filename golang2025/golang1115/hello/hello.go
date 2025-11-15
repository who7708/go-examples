package main

import "fmt"
import "example.com/greetings"

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}

// 声明一个main包。在 Go 中，作为应用程序执行的代码必须在main包中.

// 在 hello 目录中
// go mod edit -replace example.com/greetings=../greetings
// 该命令指定 example.com/greetings 应替换为 ../greetings，用于查找依赖项。运行该命令后，hello 目录中的 go.mod 文件应包含一个 replace 指令:
