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
		Member("password", String, "user password", func() {
			MinLength(1)
			MaxLength(32)
		})
		Required("password")
		Member("token", String, "token", func() {
			MinLength(1)
		})
	})

	AuthenticatePayload = Type("AuthenticatePayload", func() {
		Member("login", String, "email or nickname", func() {
			MinLength(5)
		})
		Required("login")
		Reference(PasswordPayload)
	})
)
