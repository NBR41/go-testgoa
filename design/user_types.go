package design

import (
	// All symbols
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	PasswordPayload = Type("PasswordPayload", func() {
		attrUserPassword()
		Required("password")
	})

	PasswordChangePayload = Type("PasswordChangePayload", func() {
		Reference(PasswordPayload)
		Member("password")
		attrToken()
		Required("password", "token")
	})

	AuthenticatePayload = Type("AuthenticatePayload", func() {
		Reference(PasswordPayload)
		attrUserLogin()
		Member("password")
		Required("login", "password")
	})

	UserCreatePayload = Type("UserCreatePayload", func() {
		Reference(PasswordPayload)
		attrUserEmail()
		attrUserNickname()
		Member("password")
		Required("email", "nickname", "password")

	})
)
