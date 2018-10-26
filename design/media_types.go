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
