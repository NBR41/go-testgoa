//go:generate goagen bootstrap -d github.com/NBR41/go-testgoa/design

package main

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/appsec"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"

	goajwt "github.com/goadesign/goa/middleware/security/jwt"
)

func main() {
	// Create service
	service := goa.New("my-inventory")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	customMiddleware, err := goa.NewMiddleware(jwtUserValiadtion)
	if err != nil {
		service.LogError("security middleware", "err", err)
	}

	app.UseJWTSecMiddleware(
		service,
		goajwt.New(goajwt.NewSimpleResolver([]goajwt.Key{appsec.JWTAuthKey}), customMiddleware, app.NewJWTSecSecurity()),
	)

	// Mount "authenticate" controller
	c := NewAuthenticateController(service)
	app.MountAuthenticateController(service, c)
	// Mount "books" controller
	c2 := NewBooksController(service)
	app.MountBooksController(service, c2)
	// Mount "health" controller
	c3 := NewHealthController(service)
	app.MountHealthController(service, c3)
	// Mount "ownerships" controller
	c4 := NewOwnershipsController(service)
	app.MountOwnershipsController(service, c4)
	// Mount "password" controller
	c5 := NewPasswordController(service)
	app.MountPasswordController(service, c5)
	// Mount "users" controller
	c6 := NewUsersController(service)
	app.MountUsersController(service, c6)
	// Mount "validation" controller
	c7 := NewValidationController(service)
	app.MountValidationController(service, c7)
	// Mount swagger controller onto service
	c8 := NewSwagger(service)
	app.MountSwaggerController(service, c8)
	// Mount "token" controller
	c9 := NewTokenController(service)
	app.MountTokenController(service, c9)

	// Start service
	if err := service.ListenAndServe(":8089"); err != nil {
		service.LogError("startup", "err", err)
	}

}
