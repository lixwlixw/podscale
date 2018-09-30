package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lixwlixw/podscale/scale"
)

var secrets = gin.H{
	"admin": gin.H{"email": "foo@gaojihealth.com", "phone": "123433"},
}

func main() {
	router := handle()
	s := &http.Server{
		Addr:           ":10012",
		Handler:        router,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 0,
	}
	s.ListenAndServe()
}
func getSecrets(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	if _, ok := secrets[user]; ok {
		c.JSON(200, "200")
	} else {
		c.JSON(200, "401")
	}
}

func handle() (router *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	authorized := router.Group("/namespaces", gin.BasicAuth(gin.Accounts{
		"admin": "gaojihealthadmin",
	}))
	authorized.GET("/:namespace/deployments/:name/scale", scale.ListReplicas, getSecrets)
	authorized.POST("/:namespace/deployments/:name/scale", scale.ScaleReplicas, getSecrets)
	return
}
