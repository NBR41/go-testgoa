//go:generate goagen bootstrap -d github.com/NBR41/go-testgoa/design

package main

import (
	"context"
	"flag"
	"log"

	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/controllers"
	"github.com/NBR41/go-testgoa/internal/security"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	goajwt "github.com/goadesign/goa/middleware/security/jwt"
)

type cliFlags struct {
	dbHost     string
	dbName     string
	dbUser     string
	dbPassword string

	cloudSQLRegion string
}

var envFlag string

func main() {
	cf := new(cliFlags)
	flag.StringVar(&envFlag, "env", "local", "environment to run under")
	addr := flag.String("listen", ":8089", "port to listen for HTTP on")
	flag.StringVar(&cf.dbHost, "db_host", "", "database host or Cloud SQL instance name")
	flag.StringVar(&cf.dbName, "db_name", "myinventory", "database name")
	flag.StringVar(&cf.dbUser, "db_user", "myinventory", "database user")
	flag.StringVar(&cf.dbPassword, "db_password", "", "database user password")
	flag.StringVar(&cf.cloudSQLRegion, "cloud_sql_region", "", "region of the Cloud SQL instance (GCP only)")
	flag.Parse()

	ctx := context.Background()
	var conf *config
	var cleanup func()
	var err error
	switch envFlag {
	case "gcp":
		conf, cleanup, err = setupGCP(ctx, cf)
	case "docker":
		conf, cleanup, err = setupDocker(ctx, cf)
	case "local":
		conf, cleanup, err = setupLocal(ctx, cf)
	default:
		log.Fatalf("unknown -env=%s", envFlag)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

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
	// Mount "authors" controller
	app.MountAuthorsController(service, controllers.NewAuthorsController(service, conf.fmod))
	// Mount "authorships" controller
	app.MountAuthorshipsController(service, controllers.NewAuthorshipsController(service, conf.fmod))
	// Mount "books" controller
	app.MountBooksController(service, controllers.NewBooksController(service, conf.fmod, conf.api))
	// Mount "categories" controller
	app.MountCategoriesController(service, controllers.NewCategoriesController(service, conf.fmod))
	// Mount "classes" controller
	app.MountClassesController(service, controllers.NewClassesController(service, conf.fmod))
	// Mount "classifications" controller
	app.MountClassificationsController(service, controllers.NewClassificationsController(service, conf.fmod))
	// Mount "collections" controller
	app.MountCollectionsController(service, controllers.NewCollectionsController(service, conf.fmod))
	// Mount "editions" controller
	app.MountEditionsController(service, controllers.NewEditionsController(service, conf.fmod))
	// Mount "editors" controller
	app.MountEditorsController(service, controllers.NewEditorsController(service, conf.fmod))
	// Mount "health" controller
	app.MountHealthController(service, controllers.NewHealthController(service, conf.fmod))
	// Mount "ownerships" controller
	app.MountOwnershipsController(service, controllers.NewOwnershipsController(service, conf.fmod))
	// Mount "password" controller
	app.MountPasswordController(service, controllers.NewPasswordController(service, conf.fmod, conf.token, conf.mail))
	// Mount "prints" controller
	app.MountPrintsController(service, controllers.NewPrintsController(service, conf.fmod))
	// Mount "relationAuthor" controller
	app.MountRelationAuthorController(service, controllers.NewRelationAuthorController(service, conf.fmod, conf.li))
	// Mount "relationCategory" controller
	app.MountRelationCategoryController(service, controllers.NewRelationCategoryController(service, conf.fmod, conf.li))
	// Mount "relationClass" controller
	app.MountRelationClassController(service, controllers.NewRelationClassController(service, conf.fmod, conf.li))
	// Mount "relationCollection" controller
	app.MountRelationCollectionController(service, controllers.NewRelationCollectionController(service, conf.fmod, conf.li))
	// Mount "relationEditors" controller
	app.MountRelationEditorsController(service, controllers.NewRelationEditorsController(service, conf.fmod, conf.li))
	// Mount "relationEditorsCollections" controller
	app.MountRelationEditorsCollectionsController(service, controllers.NewRelationEditorsCollectionsController(service, conf.fmod, conf.li))
	// Mount "relationEditorsPrints" controller
	app.MountRelationEditorsPrintsController(service, controllers.NewRelationEditorsPrintsController(service, conf.fmod, conf.li))
	// Mount "relationEditorsSeries" controller
	app.MountRelationEditorsSeriesController(service, controllers.NewRelationEditorsSeriesController(service, conf.fmod, conf.li))
	// Mount "relationPrints" controller
	app.MountRelationPrintsController(service, controllers.NewRelationPrintsController(service, conf.fmod, conf.li))
	// Mount "relationPrintsCollections" controller
	app.MountRelationPrintsCollectionsController(service, controllers.NewRelationPrintsCollectionsController(service, conf.fmod, conf.li))
	// Mount "relationPrintsEditors" controller
	app.MountRelationPrintsEditorsController(service, controllers.NewRelationPrintsEditorsController(service, conf.fmod, conf.li))
	// Mount "relationPrintsSeries" controller
	app.MountRelationPrintsSeriesController(service, controllers.NewRelationPrintsSeriesController(service, conf.fmod, conf.li))
	// Mount "relationRole" controller
	app.MountRelationRoleController(service, controllers.NewRelationRoleController(service, conf.fmod, conf.li))
	// Mount "relationSeries" controller
	app.MountRelationSeriesController(service, controllers.NewRelationSeriesController(service, conf.fmod, conf.li))
	// Mount "relationSeriesCollections" controller
	app.MountRelationSeriesCollectionsController(service, controllers.NewRelationSeriesCollectionsController(service, conf.fmod, conf.li))
	// Mount "relationSeriesEditors" controller
	app.MountRelationSeriesEditorsController(service, controllers.NewRelationSeriesEditorsController(service, conf.fmod, conf.li))
	// Mount "relationSeriesPrints" controller
	app.MountRelationSeriesPrintsController(service, controllers.NewRelationSeriesPrintsController(service, conf.fmod, conf.li))
	// Mount "roles" controller
	app.MountRolesController(service, controllers.NewRolesController(service, conf.fmod))
	// Mount "series" controller
	app.MountSeriesController(service, controllers.NewSeriesController(service, conf.fmod))
	// Mount "swagger" controller
	app.MountSwaggerController(service, controllers.NewSwaggerController(service))
	// Mount "token" controller
	app.MountTokenController(service, controllers.NewTokenController(service, conf.fmod, conf.token))
	// Mount "users" controller
	app.MountUsersController(service, controllers.NewUsersController(service, conf.fmod, conf.token, conf.mail))
	// Mount "validation" controller
	app.MountValidationController(service, controllers.NewValidationController(service, conf.fmod, conf.token, conf.mail))

	// Start service
	if err := service.ListenAndServe(*addr); err != nil {
		service.LogError("startup", "err", err)
	}
}
