package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// init 为函数中使用的变量设置初始值
// init 函数会在包被引用时自动调用
func init() {
	fmt.Println("greetings package initialized")
	// rand.Seed(time.Now().UnixNano())
	rand.NewSource(time.Now().UnixNano())
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// 如果 name 为空字符串，返回错误
	if name == "" {
		return "", errors.New("姓名不能为空")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v well met!",
	}
	return formats[rand.Intn(len(formats))]
}

// 在 Go 中，名称以大写字母开头的函数可以由不在同一包中的函数调用。这在 Go 中称为导出的名称
// 在 Go 中， :=运算符是在一行中声明和初始化变量的快捷方式（Go 使用右侧的值来确定变量的类型）
