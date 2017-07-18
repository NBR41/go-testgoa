package design

import (
	// All symbols
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	PasswordPayload = Type("PasswordPayload", func() {
		Member("password", String, "user password", func() {
			MinLength(1)
			MaxLength(32)
		})
		Required("password")
	})

	PasswordChangePayload = Type("PasswordChangePayload", func() {
		Reference(PasswordPayload)
		Member("password")
		Member("token", String, "token", func() {
			MinLength(1)
		})
		Required("password", "token")
	})

	AuthenticatePayload = Type("AuthenticatePayload", func() {
		Reference(PasswordPayload)
		Member("login", String, "email or nickname", func() {
			MinLength(5)
		})
		Member("password")
		Required("login", "password")

	})

	UserCreatePayload = Type("UserCreatePayload", func() {
		Reference(PasswordPayload)
		Member("email", String, "user email", func() {
			Format("email")
			Example("me@foo.bar")
		})
		Required()
		Member("nickname", String, "user nickname", func() {
			MinLength(1)
			MaxLength(32)
		})
		Member("password")
		Required("email", "nickname", "password")

	})
)
