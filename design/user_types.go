package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var emailAttribute = func() {
	Attribute("email", String, "user email", func() {
		Format("email")
		Example("me@foo.bar")
	})
}
var nicknameAttribute = func() {
	Attribute("nickname", String, "user nickname", func() {
		MinLength(1)
		MaxLength(32)
	})
}
var passwordAttribute = func() {
	Attribute("password", String, "user password", func() {
		MinLength(1)
		MaxLength(32)
	})
}
