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

})

var _ = Resource("books", func() {

	BasePath("/books")
	Action("list", func() {

		Description("疎通テスト用")
		Routing(
			GET(""),
		)
		Response(OK, BooksMedia)
		Response(InternalServerError)
	})

	Action("create", func() {
		Description("本登録")
		Routing(
			POST("/"),
		)
		Params(func() {
			Param("title", String, func() {
				Example("Clean Code")
			})
			Param("type", Integer, func() {
				Enum(1, 2, 3)
				Example(2)
			})
			Param("is_read", Boolean, func() {
				Example(true)
			})
			Param("rating", Number, func() {
				Minimum(1)
				Maximum(5)
				Example(4.5)
			})
			Param("date_of_publication", String, func() {
				Pattern("^(19[0-9]{2}|20[0-9]{2})-([1-9]|1[0-2])-([1-9]|[12][0-9]|3[01])$")
			})
			Required("title", "type", "is_read", "date_of_publication")
		})
		Response(Created)
		Response(BadRequest)
		Response(InternalServerError)
	})

	Action("update", func() {
		Description("本登録")
		Routing(
			PUT("/:id"),
		)
		Params(func() {
			Param("id", Integer)
			Required("id")
		})
		Response(NoContent)
		Response(BadRequest)
		Response(InternalServerError)
	})

	Action("retrieve", func() {
		Description("本情報取得")
		Routing(
			GET("/:id"),
		)
		Params(func() {
			Param("id", Integer)
			Required("id")
		})
		Response(OK, BookMedia)
		Response(BadRequest)
		Response(InternalServerError)

	})
	Action("delete", func() {
		Description("本情報削除")
		Routing(
			DELETE("/books/:id"),
		)
		Params(func() {
			Param("id", Integer)
			Required("id")
		})
		Response(NoContent)
		Response(BadRequest)
		Response(InternalServerError)
	})
})
