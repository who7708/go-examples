//go:build windows
// +build windows

package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	fmt.Println("Starting...")
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	fmt.Println("Stopping...")
	return nil
}

func (p *program) run() {
	fmt.Println("Run...")
	r := gin.Default()
	// 设置gin运行模式
	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8888")
}

func main() {
	svcConfig := &service.Config{
		Name:        "GinService",
		DisplayName: "Gin Service",
		Description: "This is a service written in Golang.",
	}

	prg := &program{}

	s, err := service.New(prg, svcConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 输出所有 os.Args
	fmt.Println(os.Args)

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			err = s.Install()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Installed!")
			return
		}
		if os.Args[1] == "uninstall" {
			err = s.Uninstall()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Uninstalled!")
			return
		}

		if os.Args[1] == "start" {
			err = s.Start()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Started!")
			return
		}

		if os.Args[1] == "stop" {
			err = s.Stop()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Stopped!")
			return
		}
		if os.Args[1] == "restart" {
			err = s.Restart()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Restarted!")
			return
		}
	}

	err = s.Run()
	if err != nil {
		fmt.Println(err)
		// return
	}
}

// go build -o go-service main.go
// mac /Library/LaunchDaemons/GinService.plist
// linux /etc/systemd/system/GinService.service
// windows sc create GinService binPath= "C:\Users\Administrator\go\src\golang1115\my-workspace\api\main.exe" start= auto


