package common

import (
	"github.com/goadesign/goa"
	"net/http"
	"fmt"
)

//Service :service
var Service = goa.New("test service")

// LogInfo logs the message with the route.
func LogInfo(service *goa.Service, req *http.Request, msg string, keyvals ...interface{}) {
	service.LogInfo(fmt.Sprintf("%s=%s %s", req.Method, req.URL, msg), keyvals...)
}

// LogError logs the error with the route.
func LogError(service *goa.Service, req *http.Request, msg string, keyvals ...interface{}) {
	service.LogError(fmt.Sprintf("%s=%s %s", req.Method, req.URL, msg), keyvals...)
}
