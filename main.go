package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed frontend/dist/*
var webFiles embed.FS

func main() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	// 静态文件
	fsys, _ := fs.Sub(webFiles, "frontend/dist")
	fileServer := http.FileServer(http.FS(fsys))
	r.NoRoute(func(c *gin.Context) {
		fileServer.ServeHTTP(c.Writer, c.Request)
	})
	// jwt鉴权登录
	r.POST("/api/login", Login)
	// 需鉴权的路由组
	r.GET("/api/form", authMiddleware(), GetForm)
	r.POST("/api/form", authMiddleware(), PostForm)
	r.GET("/s", GetSub)
	r.Run(":8080")
}
