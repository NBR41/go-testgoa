//+build wireinject

package main

import (
	"context"
	"github.com/NBR41/go-testgoa/controllers"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/NBR41/go-testgoa/internal/model/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

func setupDocker(ctx context.Context, flags *cliFlags) (*config, func(), error) {
	wire.Build(
		applicationSet,
		provideDockerModeler,
		provideDevMailSender,
	)
	return nil, nil, nil
}

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
