package scale

import (
	"github.com/gin-gonic/gin"
	"github.com/pivotal-golang/lager"
        "io/ioutil"
        "net/http"
	"os"
        "bytes"
	"time"
	"crypto/tls"
)

const (
      JSON = "application/json"
)
var apiHost string
var token string
var log lager.Logger
var httpClientG = &http.Client{
                Transport: httpClientB.Transport,
		Timeout:   time.Duration(10) * time.Second,
		}
var httpClientB = &http.Client{
		Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
	        Timeout: 0,
			}

func init() {
	log = lager.NewLogger("Deployment")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG)) 
        }

func init() {
	apiHost = os.Getenv("APIHOST")
   	}
func GenRequest(method, url, token string, body []byte) (*http.Response, error) {
 var req *http.Request
 var err error
 apiHost = os.Getenv("APIHOST")
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

func ListReplicas(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token = os.Getenv("APITOKEN")
	req, err := GenRequest("GET", "/apis/apps/v1beta1/namespaces/"+namespace+"/deployments/"+name+"/scale", token , []byte{})
	if err != nil {
		log.Error("Get Deployment Scale error ", err)
	}
	log.Info("Get Deployment Scale From NameSpace " , map[string]interface{}{"result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("Get Deployment Scale Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

func ScaleReplicas(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	token = os.Getenv("APITOKEN")
	req, err := GenRequest("PATCH", "/apis/apps/v1beta1/namespaces/"+namespace+"/deployments/"+name+"/scale", token , body []byte{})
	if err != nil {
		log.Error("Set Deployment Scale error ", err)
	}
	log.Info("Set Deployment Scale From NameSpace " , map[string]interface{}{"result": req.StatusCode})
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("Set Deployment Scale Read req.Body error", err)
	}
	defer req.Body.Close()
	c.Data(req.StatusCode, JSON, result)
}

