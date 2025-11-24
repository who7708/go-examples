package main

import (
	"go-spring-boot-manager/internal/service"
	"os"
)

func main() {
	jarPath := "./springboot.jar"
	if len(os.Args) > 1 {
		jarPath = os.Args[1]
	}
	service.RunService(jarPath)
}
