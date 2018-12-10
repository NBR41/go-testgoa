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
	app.MountAuthenticateController(service, controllers.NewAuthenticateController(service, conf.fmod, conf.token))
	// Mount "books" controller
	app.MountBooksController(service, controllers.NewBooksController(service, conf.fmod))
	// Mount "health" controller
	app.MountHealthController(service, controllers.NewHealthController(service, conf.fmod))
	// Mount "ownerships" controller
	app.MountOwnershipsController(service, controllers.NewOwnershipsController(service, conf.fmod, conf.api))
	// Mount "password" controller
	app.MountPasswordController(service, controllers.NewPasswordController(service, conf.fmod, conf.token, conf.mail))
	// Mount "users" controller
	app.MountUsersController(service, controllers.NewUsersController(service, conf.fmod, conf.token, conf.mail))
	// Mount "validation" controller
	app.MountValidationController(service, controllers.NewValidationController(service, conf.fmod, conf.token, conf.mail))
	// Mount swagger controller onto service
	app.MountSwaggerController(service, controllers.NewSwagger(service))
	// Mount "token" controller
	app.MountTokenController(service, controllers.NewTokenController(service, conf.fmod, conf.token))
	// Mount "authors" controller
	app.MountAuthorsController(service, controllers.NewAuthorsController(service, conf.fmod))
	// Mount "categories" controller
	app.MountCategoriesController(service, controllers.NewCategoriesController(service, conf.fmod))
	// Mount "collections" controller
	app.MountCollectionsController(service, controllers.NewCollectionsController(service))
	// Mount "edition_types" controller
	app.MountEditionTypesController(service, controllers.NewEditionTypesController(service, conf.fmod))
	// Mount "editors" controller
	app.MountEditorsController(service, controllers.NewEditorsController(service, conf.fmod))
	// Mount "genres" controller
	app.MountGenresController(service, controllers.NewGenresController(service, conf.fmod))
	// Mount "roles" controller
	app.MountRolesController(service, controllers.NewRolesController(service, conf.fmod))
	// Mount "series" controller
	app.MountSeriesController(service, controllers.NewSeriesController(service))

	// Start service
	if err := service.ListenAndServe(":8089"); err != nil {
		service.LogError("startup", "err", err)
	}

}
