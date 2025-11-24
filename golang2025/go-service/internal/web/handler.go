package web

import (
	"github.com/gin-gonic/gin"
	"go-spring-boot-manager/internal/springboot"
)

type Handler struct {
	Manager *springboot.ServiceManager
}

func NewHandler(mgr *springboot.ServiceManager) *Handler {
	return &Handler{Manager: mgr}
}

func (h *Handler) Start(c *gin.Context) {
	err := h.Manager.Start()
	c.JSON(200, gin.H{"status": "started", "error": err})
}

func (h *Handler) Stop(c *gin.Context) {
	err := h.Manager.Stop()
	c.JSON(200, gin.H{"status": "stopped", "error": err})
}

func (h *Handler) Status(c *gin.Context) {
	status := h.Manager.Status()
	c.JSON(200, gin.H{"running": status})
}

func (h *Handler) Upgrade(c *gin.Context) {
	file, _ := c.FormFile("jar")
	jarPath := "./springboot.jar"
	savePath := "./upgrade_tmp.jar"
	if file == nil {
		c.JSON(400, gin.H{"error": "No file uploaded"})
		return
	}
	c.SaveUploadedFile(file, savePath)
	err := h.Manager.Upgrade(savePath)
	c.JSON(200, gin.H{"status": "upgraded", "error": err})
}
