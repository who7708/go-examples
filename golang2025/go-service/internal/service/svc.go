package service

import (
	"github.com/kardianos/service"
	"go-spring-boot-manager/internal/springboot"
	"go-spring-boot-manager/internal/web"
	"log"
)

type program struct {
	srv *springboot.ServiceManager
}

func (p *program) Start(s service.Service) error {
	go func() {
		router := web.SetupRouter(p.srv)
		router.Run(":8080")
	}()
	return nil
}
func (p *program) Stop(s service.Service) error {
	return p.srv.Stop()
}

func RunService(jarPath string) {
	srvManager := springboot.NewServiceManager(jarPath)
	svcConfig := &service.Config{
		Name:        "GoSpringBootManager",
		DisplayName: "Go Spring Boot 管理服务",
		Description: "使用 Go 管理 Spring Boot 服务，带 web 界面",
	}
	prg := &program{srv: srvManager}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
