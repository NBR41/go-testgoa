//+build wireinject

package main

import (
	"github.com/google/go-cloud/wire"
)

func setupDev() (*config, func(), error) {
	wire.Build(
		newConfig,
		provideTokenHelper,
		providePassworder,
		provideDevModeler,
		provideDevMailSender,
		provideAPI,
	)
	return nil, nil, nil
}

func setupProd() (*config, func(), error) {
	wire.Build(
		newConfig,
		provideTokenHelper,
		providePassworder,
		provideProdModeler,
		provideProdMailSender,
		provideAPI,
	)
	return nil, nil, nil
}
