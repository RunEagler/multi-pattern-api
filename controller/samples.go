package controller

import (
	"GolandProject/multi-pattern-api/app"
	"github.com/goadesign/goa"
)

// SamplesController implements the samples resource.
type SamplesController struct {
	*goa.Controller
}

// NewSamplesController creates a samples controller.
func NewSamplesController(service *goa.Service) *SamplesController {
	return &SamplesController{Controller: service.NewController("SamplesController")}
}

// Ping runs the ping action.
func (c *SamplesController) Ping(ctx *app.PingSamplesContext) error {
	// SamplesController_Ping: start_implement

	// Put your logic here
	res := &app.User{
		Name: "test_user",
		Age:  34,
	}
	return ctx.OK(res)
	// SamplesController_Ping: end_implement
}
