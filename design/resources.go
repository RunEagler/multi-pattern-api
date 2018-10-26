package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// /actionsの定義をする
var _ = Resource("samples", func() {
	// actionsリソースのベースパス

	Action("ping", func() {

		Description("疎通テスト用")
		Routing(
			GET("/ping"),
		)
		Response(OK, UserMedia)
		Response(InternalServerError)
	})

	/*
	Action("crud", func() {

		Description("疎通テスト用")
		Routing(
			GET("/crud/:id"),
			GET("/crud"),
			POST("/crud"),
			PUT("/crud/:id"),
			DELETE("/crud/:id"),
		)
		Response(OK, UserMedia)
		Response(NotFound)
		Response(BadRequest)
		Response(InternalServerError)
	})*/
})
