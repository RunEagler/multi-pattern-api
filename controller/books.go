package controller

import (
	"GolandProject/multi-pattern-api/app"
	"github.com/goadesign/goa"
	"GolandProject/multi-pattern-api/helper"
	"net/http"
	"github.com/jinzhu/gorm"
)

// BooksController implements the books resource.
type BooksController struct {
	Service    *goa.Service
	Controller *goa.Controller
	Postgresql *gorm.DB
}

// NewBooksController creates a books controller.
func NewBooksController(service *goa.Service, postgresql *gorm.DB) *BooksController {
	return &BooksController{
		Service:    service,
		Controller: service.NewController("BooksController"),
		Postgresql: postgresql,
	}
}

// Create runs the create action.
func (c *BooksController) Create(ctx *app.CreateBooksContext) error {
	// BooksController_Create: start_implement

	// Put your logic here
	bookResource := helper.NewBookResource(c.Service, ctx.Request, c.Postgresql)

	statusCode := bookResource.Create(ctx.Title, ctx.Type, ctx.IsRead, ctx.DateOfPublication, ctx.Rating)

	switch statusCode {
	case http.StatusOK:
		return ctx.Created()
	case http.StatusBadRequest:
		return ctx.BadRequest()
	case http.StatusInternalServerError:
		return ctx.InternalServerError()
	default:
		return ctx.InternalServerError()
	}
	// BooksController_Create: end_implement
}

// Delete runs the delete action.
func (c *BooksController) Delete(ctx *app.DeleteBooksContext) error {
	// BooksController_Delete: start_implement

	// Put your logic here

	return nil
	// BooksController_Delete: end_implement
}

// List runs the list action.
func (c *BooksController) List(ctx *app.ListBooksContext) error {
	// BooksController_List: start_implement

	// Put your logic here

	res := app.BookCollection{}
	return ctx.OK(res)
	// BooksController_List: end_implement
}

// Retrieve runs the retrieve action.
func (c *BooksController) Retrieve(ctx *app.RetrieveBooksContext) error {
	// BooksController_Retrieve: start_implement

	// Put your logic here

	res := &app.Book{}
	return ctx.OK(res)
	// BooksController_Retrieve: end_implement
}

// Update runs the update action.
func (c *BooksController) Update(ctx *app.UpdateBooksContext) error {
	// BooksController_Update: start_implement

	// Put your logic here

	return nil
	// BooksController_Update: end_implement
}
