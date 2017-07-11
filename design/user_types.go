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

	UserIDParam = func() {
		Param("user_id", UserIDAttribute)
	}

	BookIDAttribute = Type("BookIDAttribute", func() {
		Description("BookID Attribute Type")
		Attribute("book_id", Integer, "Unique Book ID", func() {
			Minimum(1)
		})
	})

	BookIDParam = func() {
		Param("book_id", BookIDAttribute)
	}

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

	PasswordPayload = Type("PasswordPayload", func() {
		Reference(PasswordAttribute)
		Required("password")
	})

	AuthenticatePayload = Type("AuthenticatePayload", func() {
		Member("login", String, "email or nickname", func() {
			MinLength(5)
		})
		Required("login")
		Reference(PasswordPayload)
	})
)
