package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

// 在 Go 中，名称以大写字母开头的函数可以由不在同一包中的函数调用。这在 Go 中称为导出的名称
// 在 Go 中， :=运算符是在一行中声明和初始化变量的快捷方式（Go 使用右侧的值来确定变量的类型）
