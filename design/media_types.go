package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var UserMedia = MediaType("application/vnd.user+json", func() {
	Description("ユーザー情報")
	Attributes(func() {
		Attribute("name", String, func() {
			Example("John")
		})
		Attribute("age", Integer, func() {
			Example(23)
		})
		Required("name", "age")
	})
	View("default", func() {
		Attribute("name")
		Attribute("age")
	})
})

var BookMedia = MediaType("application/vnd.book+json", func() {

	Attributes(func() {
		Attribute("title", String, func() {
			Example("Clean Code")
		})
		Attribute("is_read", Boolean, func() {
			Example(true)
		})
		Attribute("type", Integer, func() {
			Enum(1, 2, 3)
			Example(2)
		})
		Attribute("rating", Number, func() {
			Minimum(1)
			Maximum(5)
			Example(4.5)
		})
		Attribute("date_of_publication", String, func() {
			Example("2018-10-10")
		})
		Required("title", "type", "is_read", "date_of_publication")

	})
	View("default", func() {
		Attribute("title")
		Attribute("type")
		Attribute("is_read")
		Attribute("rating")
		Attribute("date_of_publication")
	})
})

var BooksMedia = CollectionOf(BookMedia)
