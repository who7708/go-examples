package main

import (
    "github.com/gin-gonic/gin"
    "os/exec"
    "os"
)

const (
    serviceName = "SampleApp"
    jarPath = "D:\\jar-service\\sample-app.jar"
)

func main() {
    router := gin.Default()

    router.GET("/api/status", func(c *gin.Context) {
        out, _ := exec.Command("sc", "query", serviceName).Output()
        running := false
        if string(out) != "" && (string(out), "RUNNING") {
            running = true
        }
        c.JSON(200, gin.H{"running": running})
    })

    router.POST("/api/start", func(c *gin.Context) {
        err := exec.Command("sc", "start", serviceName).Run()
        c.JSON(200, gin.H{"status": "started", "error": err})
    })
    router.POST("/api/stop", func(c *gin.Context) {
        err := exec.Command("sc", "stop", serviceName).Run()
        c.JSON(200, gin.H{"status": "stopped", "error": err})
    })
    router.POST("/api/upgrade", func(c *gin.Context) {
        file, _ := c.FormFile("jar")
        savePath := "./uploaded_new.jar"
        c.SaveUploadedFile(file, savePath)
        backupPath := jarPath + ".bak"
        // Stop service
        exec.Command("sc", "stop", serviceName).Run()
        // Backup old jar
        os.Rename(jarPath, backupPath)
        // Copy new jar
        os.Rename(savePath, jarPath)
        // Start service
        exec.Command("sc", "start", serviceName).Run()
        c.JSON(200, gin.H{"status": "upgraded"})
    })

    router.Run(":8080")
}