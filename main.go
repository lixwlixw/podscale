package main

import (
	"github.com/lixwlixw/podscale/scale"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := handle()
	s := &http.Server{
		Addr:    ":10012",
		Handler: router,
		//ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 0,
	}
	//监听端口
	s.ListenAndServe()
}

func handle() (router *gin.Engine) {
	//设置全局环境：1.开发环境（datafoundry_docker.DebugMode） 2.线上环境（datafoundry_docker.ReleaseMode）
	gin.SetMode(gin.ReleaseMode)
	//获取路由实例
	router = gin.Default()
	router.GET("/", scale.Hello)
	return
}
