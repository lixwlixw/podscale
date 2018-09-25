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
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 0,
	}
	s.ListenAndServe()
}

func handle() (router *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
        router.GET("/namespaces/:namespace/deployments/:name/scale", scale.ListReplicas)
	return
}
