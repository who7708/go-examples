package web

import (
	"github.com/gin-gonic/gin"
	"go-spring-boot-manager/internal/springboot"
)

func SetupRouter(manager *springboot.ServiceManager) *gin.Engine {
	r := gin.Default()
	h := NewHandler(manager)
	r.Static("/static", "./internal/web/static")
	r.GET("/", func(c *gin.Context) {
		c.File("./internal/web/static/index.html")
	})
	r.POST("/api/start", h.Start)
	r.POST("/api/stop", h.Stop)
	r.POST("/api/upgrade", h.Upgrade)
	r.GET("/api/status", h.Status)
	return r
}
