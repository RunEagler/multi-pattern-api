//go:generate goagen bootstrap -d GolandProject/multi-pattern-api/design

package main

import (
	"GolandProject/multi-pattern-api/app"
	"GolandProject/multi-pattern-api/controller"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("goa simple sample")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "samples" controller
	c := controller.NewSamplesController(service)
	app.MountSamplesController(service, c)
	// Mount "swagger" controller
	c2 := controller.NewSwaggerController(service)
	app.MountSwaggerController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
