package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	UserIDAttribute = Type("UserIDAttribute", func() {
		Description("UserID Attribute Type")
		Attribute("user_id", Integer, "Unique User ID", func() {
			Minimum(1)
		})
	})

	BookIDAttribute = Type("BookIDAttribute", func() {
		Description("BookID Attribute Type")
		Attribute("book_id", Integer, "Unique Book ID", func() {
			Minimum(1)
		})
	})

	BookNameAttribute = Type("BookNameAttribute", func() {
		Description("Book Name Attribute type")
		Attribute("book_name", String, "Book Name", func() {
			MinLength(1)
			MaxLength(128)
		})
	})

	EmailAttribute = Type("EmailAttribute", func() {
		Description("Email Attribute Type")
		Attribute("email", String, "user email", func() {
			Format("email")
			Example("me@foo.bar")
		})
	})

	NicknameAttribute = Type("NicknameAttribute", func() {
		Description("Nickname Attribute Type")
		Attribute("nickname", String, "user nickname", func() {
			MinLength(1)
			MaxLength(32)
		})
	})

	PasswordAttribute = Type("PasswordAttribute", func() {
		Description("Password Attribute Type")
		Attribute("password", String, "user password", func() {
			MinLength(1)
			MaxLength(32)
		})
	})

	TokenAttribute = Type("TokenAttribute", func() {
		Description("Token Attribute Type")
		Attribute("token", String, "token", func() {
			MinLength(1)
		})
	})

	PasswordPayload = Type("PasswordPayload", func() {
		Member("password", PasswordAttribute)
		Required("password")
	})

	PasswordChangePayload = Type("PasswordChangePayload", func() {
		Member("password", PasswordAttribute)
		Required("password")
		Member("confirm", PasswordAttribute)
		Required("confirm")
		Member("token", TokenAttribute)
	})

	AuthenticatePayload = Type("AuthenticatePayload", func() {
		Member("login", String, "email or nickname", func() {
			MinLength(5)
		})
		Required("login")
		Reference(PasswordPayload)
	})
)
