package scale

import (
	"github.com/gin-gonic/gin"
        "io/ioutil"
        "net/http"
	"os"
        "fmt"
        "log"
        
)
var apiHost string

func getenv(env string) string {
	env_value := os.Getenv(env)
	if env_value == "" {
		fmt.Println("FATAL: NEED ENV", env)
		fmt.Println("Exit...........")
		os.Exit(2)
	}
	fmt.Println("ENV:", env, env_value)
	return env_value
}

func GenRequest(method, url, token string, body []byte) (*http.Response, error) {
 var req *http.Request
 var err error
 apiHost = getenv("APIHOST") 
 url = "https://" + apiHost + url
 if len(body) == 0 {
  req, err = http.NewRequest(method, url, nil)
 } else {
  req, err = http.NewRequest(method, url, bytes.NewReader(body))
 }
 if err != nil {
  return nil, err
 }

 req.Header.Set("Content-Type", "application/json")
 req.Header.Set("Authorization", token)
 return httpClientG.Do(req)
}

func GetScaleDepFromNS(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
        token := getenv("APITOKEN")
	req, err := GenRequest("GET", "/apis/apps/v1beta1/namespaces/"+namespace+"/deployments/"+name+"/scale", token, []byte{})
	if err != nil {
		log.Error("GetScaleDepFromNS error ", err)
	}
	log.Info("Get Scale Dep From NameSpace ", "result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("GetScaleDepFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}
