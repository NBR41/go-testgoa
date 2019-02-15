package main

import (
	"github.com/NBR41/go-testgoa/controllers"
	"github.com/NBR41/go-testgoa/internal/api/google"
	"github.com/NBR41/go-testgoa/internal/mail"
	"github.com/NBR41/go-testgoa/internal/mail/content"
	"github.com/NBR41/go-testgoa/internal/mail/display"
	"github.com/NBR41/go-testgoa/internal/mail/smtp"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/NBR41/go-testgoa/internal/security"
	"github.com/google/wire"
)

var conf *config

type config struct {
	fmod  controllers.Fmodeler
	token controllers.TokenHelper
	mail  controllers.MailSender
	api   controllers.APIHelper
	li    controllers.Lister
}

// applicationSet is the Wire provider set for the Guestbook application that
// does not depend on the underlying platform.
var applicationSet = wire.NewSet(
	newConfig,
	provideTokenHelper,
	providePassworder,
	provideLister,
	provideAPI,
)

func newConfig(
	fmod controllers.Fmodeler, token controllers.TokenHelper,
	mail controllers.MailSender, api controllers.APIHelper,
	li controllers.Lister,
) *config {
	return &config{fmod: fmod, token: token, mail: mail, api: api, li: li}
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

func provideAPI() controllers.APIHelper {
	return google.New(&google.HttpCaller{})
}

func provideDevMailSender() controllers.MailSender {
	return mail.New(content.NewFactory("http://localhost:4200"), &display.Actionner{})
}

func provideProdMailSender() controllers.MailSender {
	return mail.New(content.NewFactory("http://localhost:4200"), &smtp.Actionner{})
}
