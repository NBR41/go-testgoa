//+build wireinject

package main

import (
	"context"
	"database/sql"
	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/certs"
	"github.com/NBR41/go-testgoa/controllers"
	"github.com/NBR41/go-testgoa/internal/model"
	modelsql "github.com/NBR41/go-testgoa/internal/model/sql"
	"github.com/google/wire"
	"gocloud.dev/gcp"
	"gocloud.dev/gcp/gcpcloud"
	"gocloud.dev/mysql/cloudmysql"
)

func setupGCP(ctx context.Context, flags *cliFlags) (*config, func(), error) {
	wire.Build(
		applicationSet,
		gcpcloud.GCP,
		gcpSQLParams,
		provideGCPModeler,
		provideDevMailSender,
	)
	return nil, nil, nil
}

func provideGCPModeler(ctx context.Context, remoteCertSource *certs.RemoteCertSource, params *cloudmysql.Params, pass model.Passworder) controllers.Fmodeler {
	return controllers.Fmodeler(func() (controllers.Modeler, error) {
		return modelsql.New(
			modelsql.ConnGetter(func() (*sql.DB, error) {
				return cloudmysql.Open(ctx, remoteCertSource, params)
			}),
			pass,
		)
	})
}

// gcpSQLParams is a Wire provider function that returns the Cloud SQL
// connection parameters based on the command-line flags. Other providers inside
// gcpcloud.GCP use the parameters to construct a *sql.DB.
func gcpSQLParams(id gcp.ProjectID, flags *cliFlags) *cloudmysql.Params {
	return &cloudmysql.Params{
		ProjectID: string(id),
		Region:    flags.cloudSQLRegion,
		Instance:  flags.dbHost,
		Database:  flags.dbName,
		User:      flags.dbUser,
		Password:  flags.dbPassword,
	}
}
