package scale

import (
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
        "io/ioutil"
        "net/http"
	"os"
        "fmt"
        "bytes"
	"time"
	"crypto/tls"
        
)
const (
      JSON = "application/json"
)
var apiHost string
var token string

var httpClientB = &http.Client{
		Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 0,
		}

var httpClientG = &http.Client{
		Transport: httpClientB.Transport,
		Timeout:   time.Duration(10) * time.Second,
		}

var log lager.Logger
func init() {
	log = lager.NewLogger("DeploymentConfig")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG)) 
        }

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
 apiHost = os.Getenv("APIHOST") 
 url = "https:/" + apiHost + url
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

func ListReplicas(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token = os.Getenv("APITOKEN")
	req, err := GenRequest("GET", apiHost+"/apis/apps/v1beta1/namespaces/"+namespace+"/deployments/"+name+"/scale", token , []byte{})
	if err != nil {
		log.Error("GetScaleDepFromNS error ", err)
	}
	log.Info("Get Scale Dep From NameSpace " , map[string]interface{}{"result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("GetScaleDepFromNS Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

