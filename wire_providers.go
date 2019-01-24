package main

import (
	"errors"
	"github.com/NBR41/go-testgoa/controllers"
	"github.com/NBR41/go-testgoa/internal/api/google"
	"github.com/NBR41/go-testgoa/internal/mail"
	"github.com/NBR41/go-testgoa/internal/mail/content"
	"github.com/NBR41/go-testgoa/internal/mail/display"
	"github.com/NBR41/go-testgoa/internal/mail/smtp"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/NBR41/go-testgoa/internal/model/local"
	"github.com/NBR41/go-testgoa/internal/model/sql"
	"github.com/NBR41/go-testgoa/internal/security"
	"os"
)

var conf *config

type config struct {
	fmod  controllers.Fmodeler
	token controllers.TokenHelper
	mail  controllers.MailSender
	api   controllers.APIHelper
	li    controllers.Lister
}

func newConfig(
	fmod controllers.Fmodeler, token controllers.TokenHelper,
	mail controllers.MailSender, api controllers.APIHelper,
	li controllers.Lister,
) *config {
	return &config{fmod: fmod, token: token, mail: mail, api: api, li: li}
}

func initws() (*config, func(), error) {
	if os.Getenv("ISPROD") == "1" {
		return setupProd()
	}
	return setupDev()
}

func provideLister() controllers.Lister {
	return &controllers.ListBuilder{}
}

func provideTokenHelper() controllers.TokenHelper {
	return &security.JWTHelper{}
}

func providePassworder() model.Passworder {
	return security.NewPasswordHelper([]byte("0UArPJLVC3h667sQ"))
}

func provideDevModeler(pass model.Passworder) controllers.Fmodeler {
	return controllers.Fmodeler(func() (controllers.Modeler, error) {
		return local.New(pass), nil
	})
}

func provideProdModeler(pass model.Passworder) (controllers.Fmodeler, error) {
	connString := os.Getenv("DB_CONN_STR")
	if connString == "" {
		return nil, errors.New("invalid connection string")
	}
	return controllers.Fmodeler(func() (controllers.Modeler, error) {
		return sql.New(sql.GetConnGetter(connString), pass)
	}), nil
}

func provideDevMailSender() controllers.MailSender {
	return mail.New(content.NewFactory("http://localhost:4200"), &display.Actionner{})
}

func provideProdMailSender() controllers.MailSender {
	return mail.New(content.NewFactory("http://localhost:4200"), &smtp.Actionner{})
}

func provideAPI() controllers.APIHelper {
	return google.New(&google.HttpCaller{})
}
