// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"context"
	sql2 "database/sql"
	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/certs"
	"github.com/NBR41/go-testgoa/controllers"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/NBR41/go-testgoa/internal/model/local"
	"github.com/NBR41/go-testgoa/internal/model/sql"
	"github.com/go-sql-driver/mysql"
	"gocloud.dev/gcp"
	"gocloud.dev/gcp/cloudsql"
	"gocloud.dev/mysql/cloudmysql"
)

// Injectors from inject_docker.go:

func setupDocker(ctx context.Context, flags *cliFlags) (*config, func(), error) {
	passworder := providePassworder()
	fmodeler := provideDockerModeler(flags, passworder)
	tokenHelper := provideTokenHelper()
	mailSender := provideDevMailSender()
	apiHelper := provideAPI()
	lister := provideLister()
	mainConfig := newConfig(fmodeler, tokenHelper, mailSender, apiHelper, lister)
	return mainConfig, func() {
	}, nil
}

// Injectors from inject_gcp.go:

func setupGCP(ctx context.Context, flags *cliFlags) (*config, func(), error) {
	roundTripper := gcp.DefaultTransport()
	credentials, err := gcp.DefaultCredentials(ctx)
	if err != nil {
		return nil, nil, err
	}
	tokenSource := gcp.CredentialsTokenSource(credentials)
	httpClient, err := gcp.NewHTTPClient(roundTripper, tokenSource)
	if err != nil {
		return nil, nil, err
	}
	remoteCertSource := cloudsql.NewCertSource(httpClient)
	projectID, err := gcp.DefaultProjectID(credentials)
	if err != nil {
		return nil, nil, err
	}
	params := gcpSQLParams(projectID, flags)
	passworder := providePassworder()
	fmodeler := provideGCPModeler(ctx, remoteCertSource, params, passworder)
	tokenHelper := provideTokenHelper()
	mailSender := provideDevMailSender()
	apiHelper := provideAPI()
	lister := provideLister()
	mainConfig := newConfig(fmodeler, tokenHelper, mailSender, apiHelper, lister)
	return mainConfig, func() {
	}, nil
}

// Injectors from inject_local.go:

func setupLocal(ctx context.Context, flags *cliFlags) (*config, func(), error) {
	passworder := providePassworder()
	fmodeler := provideLocalModeler(passworder)
	tokenHelper := provideTokenHelper()
	mailSender := provideDevMailSender()
	apiHelper := provideAPI()
	lister := provideLister()
	mainConfig := newConfig(fmodeler, tokenHelper, mailSender, apiHelper, lister)
	return mainConfig, func() {
	}, nil
}

// inject_docker.go:

func provideDockerModeler(flags *cliFlags, pass model.Passworder) controllers.Fmodeler {
	cfg := &mysql.Config{
		Net:                  "tcp",
		Addr:                 flags.dbHost,
		DBName:               flags.dbName,
		User:                 flags.dbUser,
		Passwd:               flags.dbPassword,
		Params:               map[string]string{"charset": "utf8mb4,utf8"},
		AllowNativePasswords: true,
	}
	return controllers.Fmodeler(func() (controllers.Modeler, error) {
		return sql.New(sql.GetConnGetter(cfg.FormatDSN()), pass)
	})
}

// inject_gcp.go:

func provideGCPModeler(ctx context.Context, remoteCertSource *certs.RemoteCertSource, params *cloudmysql.Params, pass model.Passworder) controllers.Fmodeler {
	return controllers.Fmodeler(func() (controllers.Modeler, error) {
		return sql.New(sql.ConnGetter(func() (*sql2.DB, error) {
			return cloudmysql.Open(ctx, remoteCertSource, params)
		}), pass,
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

// inject_local.go:

func provideLocalModeler(pass model.Passworder) controllers.Fmodeler {
	model2 := local.New(pass)
	return controllers.Fmodeler(func() (controllers.Modeler, error) {
		return model2, nil
	})
}
