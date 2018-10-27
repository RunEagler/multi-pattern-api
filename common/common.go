package common

import (
	"github.com/goadesign/goa"
	"net/http"
	"fmt"
	"time"
)

// LogInfo logs the message with the route.
func LogInfo(service *goa.Service, req *http.Request, msg string, keyvals ...interface{}) {
	service.LogInfo(fmt.Sprintf("%s=%s %s", req.Method, req.URL, msg), keyvals...)
}

// LogError logs the error with the route.
func LogError(service *goa.Service, req *http.Request, msg string, keyvals ...interface{}) {
	service.LogError(fmt.Sprintf("%s=%s %s", req.Method, req.URL, msg), keyvals...)
}

func TestMessage(title string, paramName string) string {

	return fmt.Sprintf("test case[%v] :%s", title, paramName)
}

func PtrInt(x int) *int {
	return &x
}

func PtrFloat(x float64) *float64 {
	return &x
}

func PtrString(x string) *string {
	return &x
}

func PtrTime(x time.Time) *time.Time {
	return &x
}

func PtrBool(x bool) *bool {
	return &x
}
