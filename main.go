//go:generate goagen bootstrap -d github.com/NBR41/go-testgoa/design

package main

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/controllers"
	"github.com/NBR41/go-testgoa/internal/security"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	goajwt "github.com/goadesign/goa/middleware/security/jwt"
)

func main() {
	conf, f, err := initws()
	if err != nil {
		panic(err)
	}
	defer f()

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

	app.UseJWTSecMiddleware(service, goajwt.New(security.JWTAuthKey, customMiddleware, app.NewJWTSecSecurity()))

	// Mount "authenticate" controller
	c := controllers.NewAuthenticateController(service, conf.fmod, conf.token)
	app.MountAuthenticateController(service, c)
	// Mount "books" controller
	c2 := controllers.NewBooksController(service, conf.fmod)
	app.MountBooksController(service, c2)
	// Mount "health" controller
	c3 := controllers.NewHealthController(service, conf.fmod)
	app.MountHealthController(service, c3)
	// Mount "ownerships" controller
	c4 := controllers.NewOwnershipsController(service, conf.fmod, conf.api)
	app.MountOwnershipsController(service, c4)
	// Mount "password" controller
	c5 := controllers.NewPasswordController(service, conf.fmod, conf.token, conf.mail)
	app.MountPasswordController(service, c5)
	// Mount "users" controller
	c6 := controllers.NewUsersController(service, conf.fmod, conf.token, conf.mail)
	app.MountUsersController(service, c6)
	// Mount "validation" controller
	c7 := controllers.NewValidationController(service, conf.fmod, conf.token, conf.mail)
	app.MountValidationController(service, c7)
	// Mount swagger controller onto service
	c8 := controllers.NewSwagger(service)
	app.MountSwaggerController(service, c8)
	// Mount "token" controller
	c9 := controllers.NewTokenController(service, conf.fmod, conf.token)
	app.MountTokenController(service, c9)

	// Start service
	if err := service.ListenAndServe(":8089"); err != nil {
		service.LogError("startup", "err", err)
	}

}
