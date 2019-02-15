//+build wireinject

package main

import (
	"context"
	"github.com/NBR41/go-testgoa/controllers"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/NBR41/go-testgoa/internal/model/local"
	"github.com/google/wire"
)

func setupLocal(ctx context.Context, flags *cliFlags) (*config, func(), error) {
	wire.Build(
		applicationSet,
		provideLocalModeler,
		provideDevMailSender,
	)
	return nil, nil, nil
}

func provideLocalModeler(pass model.Passworder) controllers.Fmodeler {
	return controllers.Fmodeler(func() (controllers.Modeler, error) {
		return local.New(pass), nil
	})
}
