// go:build windows
package main

import (
	"fmt"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
	"log"
	"os"
	"time"
)

func controlService(serviceName string, cmd svc.Cmd) error {
	// 1. 连接到服务控制管理器
	m, err := mgr.Connect()
	if err != nil {
		log.Printf("连接服务管理器失败: %v", err)
		return fmt.Errorf("连接服务管理器失败: %v", err)
	}
	defer m.Disconnect()

	// 2. 打开服务
	s, err := m.OpenService(serviceName)
	if err != nil {
		log.Printf("打开服务 %s 失败: %v", serviceName, err)
		return fmt.Errorf("打开服务 %s 失败: %v", serviceName, err)
	}
	defer s.Close()

	// 3. 获取服务当前状态（可选，用于检查）
	status, err := s.Query()
	if err != nil {
		log.Printf("查询服务状态失败: %v", err)
		return fmt.Errorf("查询服务状态失败: %v", err)
	}
	log.Printf("当前状态: %v\n", status.State)

	// 4. 发送控制命令
	_, err = s.Control(cmd)
	if err != nil {
		log.Printf("发送控制命令 %v 失败: %v", cmd, err)
		return fmt.Errorf("发送控制命令 %v 失败: %v", cmd, err)
	}

	// 5. 等待状态变更（例如等待停止）
	if cmd == svc.Stop {
		timeout := time.Now().Add(10 * time.Second)
		for time.Now().Before(timeout) {
			status, err = s.Query()
			if err != nil {
				log.Printf("等待时查询状态失败: %v", err)
				return fmt.Errorf("等待时查询状态失败: %v", err)
			}
			if status.State == svc.Stopped {
				log.Println("服务已确认停止")
				return nil
			}
			time.Sleep(300 * time.Millisecond)
		}
		log.Printf("等待服务停止超时")
		return fmt.Errorf("等待服务停止超时")
	}

	log.Printf("命令 %v 已发送成功\n", cmd)
	return nil
}

func main() {
	// 检查命令行参数
	if len(os.Args) < 2 {
		log.Fatalf("用法: %s <服务名称>", os.Args[0])
	}

	// 从命令行参数获取服务名称
	serviceName := os.Args[1]
	log.Printf("目标服务: %s", serviceName)

	// 停止服务
	log.Println("正在停止服务...")
	if err := controlService(serviceName, svc.Stop); err != nil {
		log.Printf("停止失败: %v\n", err)
	}

	// ... 执行更新操作 ...

	// 启动服务
	log.Println("正在启动服务...")
	// 注意：启动使用的是 s.Start() 方法，而非 Control 调用
	if err := startServiceNative(serviceName); err != nil {
		log.Printf("启动失败: %v\n", err)
	}
}

// 单独处理服务启动
func startServiceNative(serviceName string) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()

	s, err := m.OpenService(serviceName)
	if err != nil {
		return err
	}
	defer s.Close()

	// 启动服务可以传入启动参数（第二个参数）
	err = s.Start()
	if err != nil {
		return fmt.Errorf("启动服务失败: %v", err)
	}
	log.Println("服务启动指令已成功发送")
	return nil
}

// 功能特性
// 控制 Windows 服务的启动和停止
// 通过命令行参数指定目标服务名称
// 提供服务状态查询和命令执行反馈

// # 在 Windows 系统上直接构建
// go build -o service-controller.exe main.go
//
// # 或者指定目标平台进行交叉编译
// GOOS=windows GOARCH=amd64 go build -o service-controller.exe main.go
