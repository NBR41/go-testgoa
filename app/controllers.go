// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "my-inventory": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/NBR41/go-testgoa/design
// --out=$(GOPATH)/src/github.com/NBR41/go-testgoa
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// AuthenticateController is the controller interface for the Authenticate actions.
type AuthenticateController interface {
	goa.Muxer
	Auth(*AuthAuthenticateContext) error
}

// MountAuthenticateController "mounts" a Authenticate resource controller on the given service.
func MountAuthenticateController(service *goa.Service, ctrl AuthenticateController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/authenticate", ctrl.MuxHandler("preflight", handleAuthenticateOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAuthAuthenticateContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*AuthenticatePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Auth(rctx)
	}
	h = handleAuthenticateOrigin(h)
	service.Mux.Handle("POST", "/authenticate", ctrl.MuxHandler("auth", h, unmarshalAuthAuthenticatePayload))
	service.LogInfo("mount", "ctrl", "Authenticate", "action", "Auth", "route", "POST /authenticate")
}

// handleAuthenticateOrigin applies the CORS response headers corresponding to the origin.
func handleAuthenticateOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:4200") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, Content-Type, Accept")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalAuthAuthenticatePayload unmarshals the request body into the context request data Payload field.
func unmarshalAuthAuthenticatePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &authenticatePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// BooksController is the controller interface for the Books actions.
type BooksController interface {
	goa.Muxer
	Create(*CreateBooksContext) error
	Delete(*DeleteBooksContext) error
	List(*ListBooksContext) error
	Show(*ShowBooksContext) error
	Update(*UpdateBooksContext) error
}

// MountBooksController "mounts" a Books resource controller on the given service.
func MountBooksController(service *goa.Service, ctrl BooksController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/books", ctrl.MuxHandler("preflight", handleBooksOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/books/:book_id", ctrl.MuxHandler("preflight", handleBooksOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateBooksContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateBooksPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleSecurity("JWTSec", h)
	h = handleBooksOrigin(h)
	service.Mux.Handle("POST", "/books", ctrl.MuxHandler("create", h, unmarshalCreateBooksPayload))
	service.LogInfo("mount", "ctrl", "Books", "action", "Create", "route", "POST /books", "security", "JWTSec")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteBooksContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleSecurity("JWTSec", h)
	h = handleBooksOrigin(h)
	service.Mux.Handle("DELETE", "/books/:book_id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Books", "action", "Delete", "route", "DELETE /books/:book_id", "security", "JWTSec")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListBooksContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleBooksOrigin(h)
	service.Mux.Handle("GET", "/books", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Books", "action", "List", "route", "GET /books")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowBooksContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleBooksOrigin(h)
	service.Mux.Handle("GET", "/books/:book_id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Books", "action", "Show", "route", "GET /books/:book_id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateBooksContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateBooksPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleSecurity("JWTSec", h)
	h = handleBooksOrigin(h)
	service.Mux.Handle("PUT", "/books/:book_id", ctrl.MuxHandler("update", h, unmarshalUpdateBooksPayload))
	service.LogInfo("mount", "ctrl", "Books", "action", "Update", "route", "PUT /books/:book_id", "security", "JWTSec")
}

// handleBooksOrigin applies the CORS response headers corresponding to the origin.
func handleBooksOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:4200") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, Content-Type, Accept")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateBooksPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateBooksPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createBooksPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateBooksPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateBooksPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateBooksPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// HealthController is the controller interface for the Health actions.
type HealthController interface {
	goa.Muxer
	Health(*HealthHealthContext) error
}

// MountHealthController "mounts" a Health resource controller on the given service.
func MountHealthController(service *goa.Service, ctrl HealthController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/_ah/health", ctrl.MuxHandler("preflight", handleHealthOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewHealthHealthContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Health(rctx)
	}
	h = handleHealthOrigin(h)
	service.Mux.Handle("GET", "/_ah/health", ctrl.MuxHandler("health", h, nil))
	service.LogInfo("mount", "ctrl", "Health", "action", "Health", "route", "GET /_ah/health")
}

// handleHealthOrigin applies the CORS response headers corresponding to the origin.
func handleHealthOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:4200") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, Content-Type, Accept")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// PasswordController is the controller interface for the Password actions.
type PasswordController interface {
	goa.Muxer
	Get(*GetPasswordContext) error
	Update(*UpdatePasswordContext) error
}

// MountPasswordController "mounts" a Password resource controller on the given service.
func MountPasswordController(service *goa.Service, ctrl PasswordController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/password", ctrl.MuxHandler("preflight", handlePasswordOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetPasswordContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handlePasswordOrigin(h)
	service.Mux.Handle("GET", "/password", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "Password", "action", "Get", "route", "GET /password")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdatePasswordContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*PasswordChangePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handlePasswordOrigin(h)
	service.Mux.Handle("POST", "/password", ctrl.MuxHandler("update", h, unmarshalUpdatePasswordPayload))
	service.LogInfo("mount", "ctrl", "Password", "action", "Update", "route", "POST /password")
}

// handlePasswordOrigin applies the CORS response headers corresponding to the origin.
func handlePasswordOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:4200") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, Content-Type, Accept")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalUpdatePasswordPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdatePasswordPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &passwordChangePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// UsersController is the controller interface for the Users actions.
type UsersController interface {
	goa.Muxer
	Create(*CreateUsersContext) error
	Delete(*DeleteUsersContext) error
	List(*ListUsersContext) error
	Show(*ShowUsersContext) error
	Update(*UpdateUsersContext) error
}

// MountUsersController "mounts" a Users resource controller on the given service.
func MountUsersController(service *goa.Service, ctrl UsersController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/users", ctrl.MuxHandler("preflight", handleUsersOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/users/:user_id", ctrl.MuxHandler("preflight", handleUsersOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateUsersContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UserCreatePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleUsersOrigin(h)
	service.Mux.Handle("POST", "/users", ctrl.MuxHandler("create", h, unmarshalCreateUsersPayload))
	service.LogInfo("mount", "ctrl", "Users", "action", "Create", "route", "POST /users")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteUsersContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleSecurity("JWTSec", h)
	h = handleUsersOrigin(h)
	service.Mux.Handle("DELETE", "/users/:user_id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Users", "action", "Delete", "route", "DELETE /users/:user_id", "security", "JWTSec")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListUsersContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleUsersOrigin(h)
	service.Mux.Handle("GET", "/users", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Users", "action", "List", "route", "GET /users")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowUsersContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleUsersOrigin(h)
	service.Mux.Handle("GET", "/users/:user_id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Users", "action", "Show", "route", "GET /users/:user_id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateUsersContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateUsersPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleSecurity("JWTSec", h)
	h = handleUsersOrigin(h)
	service.Mux.Handle("PUT", "/users/:user_id", ctrl.MuxHandler("update", h, unmarshalUpdateUsersPayload))
	service.LogInfo("mount", "ctrl", "Users", "action", "Update", "route", "PUT /users/:user_id", "security", "JWTSec")
}

// handleUsersOrigin applies the CORS response headers corresponding to the origin.
func handleUsersOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:4200") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, Content-Type, Accept")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateUsersPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateUsersPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &userCreatePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateUsersPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateUsersPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateUsersPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// OwnershipsController is the controller interface for the Ownerships actions.
type OwnershipsController interface {
	goa.Muxer
	Add(*AddOwnershipsContext) error
	Create(*CreateOwnershipsContext) error
	Delete(*DeleteOwnershipsContext) error
	List(*ListOwnershipsContext) error
	Show(*ShowOwnershipsContext) error
}

// MountOwnershipsController "mounts" a Ownerships resource controller on the given service.
func MountOwnershipsController(service *goa.Service, ctrl OwnershipsController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/users/:user_id/ownerships/isbn", ctrl.MuxHandler("preflight", handleOwnershipsOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/users/:user_id/ownerships", ctrl.MuxHandler("preflight", handleOwnershipsOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/users/:user_id/ownerships/:book_id", ctrl.MuxHandler("preflight", handleOwnershipsOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAddOwnershipsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*AddOwnershipsPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Add(rctx)
	}
	h = handleSecurity("JWTSec", h)
	h = handleOwnershipsOrigin(h)
	service.Mux.Handle("POST", "/users/:user_id/ownerships/isbn", ctrl.MuxHandler("add", h, unmarshalAddOwnershipsPayload))
	service.LogInfo("mount", "ctrl", "Ownerships", "action", "Add", "route", "POST /users/:user_id/ownerships/isbn", "security", "JWTSec")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateOwnershipsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateOwnershipsPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleSecurity("JWTSec", h)
	h = handleOwnershipsOrigin(h)
	service.Mux.Handle("POST", "/users/:user_id/ownerships", ctrl.MuxHandler("create", h, unmarshalCreateOwnershipsPayload))
	service.LogInfo("mount", "ctrl", "Ownerships", "action", "Create", "route", "POST /users/:user_id/ownerships", "security", "JWTSec")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteOwnershipsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleSecurity("JWTSec", h)
	h = handleOwnershipsOrigin(h)
	service.Mux.Handle("DELETE", "/users/:user_id/ownerships/:book_id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Ownerships", "action", "Delete", "route", "DELETE /users/:user_id/ownerships/:book_id", "security", "JWTSec")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListOwnershipsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSecurity("JWTSec", h)
	h = handleOwnershipsOrigin(h)
	service.Mux.Handle("GET", "/users/:user_id/ownerships", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Ownerships", "action", "List", "route", "GET /users/:user_id/ownerships", "security", "JWTSec")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowOwnershipsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleSecurity("JWTSec", h)
	h = handleOwnershipsOrigin(h)
	service.Mux.Handle("GET", "/users/:user_id/ownerships/:book_id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Ownerships", "action", "Show", "route", "GET /users/:user_id/ownerships/:book_id", "security", "JWTSec")
}

// handleOwnershipsOrigin applies the CORS response headers corresponding to the origin.
func handleOwnershipsOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:4200") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, Content-Type, Accept")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalAddOwnershipsPayload unmarshals the request body into the context request data Payload field.
func unmarshalAddOwnershipsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &addOwnershipsPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalCreateOwnershipsPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateOwnershipsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createOwnershipsPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger.json", "public/swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "public/swagger/swagger.json", "route", "GET /swagger.json")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			}
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:4200") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, Content-Type, Accept")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// TokenController is the controller interface for the Token actions.
type TokenController interface {
	goa.Muxer
	Access(*AccessTokenContext) error
	Auth(*AuthTokenContext) error
}

// MountTokenController "mounts" a Token resource controller on the given service.
func MountTokenController(service *goa.Service, ctrl TokenController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/token/access_token", ctrl.MuxHandler("preflight", handleTokenOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/token/auth", ctrl.MuxHandler("preflight", handleTokenOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAccessTokenContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Access(rctx)
	}
	h = handleSecurity("JWTSec", h)
	h = handleTokenOrigin(h)
	service.Mux.Handle("GET", "/token/access_token", ctrl.MuxHandler("access", h, nil))
	service.LogInfo("mount", "ctrl", "Token", "action", "Access", "route", "GET /token/access_token", "security", "JWTSec")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAuthTokenContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Auth(rctx)
	}
	h = handleSecurity("JWTSec", h)
	h = handleTokenOrigin(h)
	service.Mux.Handle("GET", "/token/auth", ctrl.MuxHandler("auth", h, nil))
	service.LogInfo("mount", "ctrl", "Token", "action", "Auth", "route", "GET /token/auth", "security", "JWTSec")
}

// handleTokenOrigin applies the CORS response headers corresponding to the origin.
func handleTokenOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:4200") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, Content-Type, Accept")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// ValidationController is the controller interface for the Validation actions.
type ValidationController interface {
	goa.Muxer
	Get(*GetValidationContext) error
	Validate(*ValidateValidationContext) error
}

// MountValidationController "mounts" a Validation resource controller on the given service.
func MountValidationController(service *goa.Service, ctrl ValidationController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/validation/:user_id", ctrl.MuxHandler("preflight", handleValidationOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/validation", ctrl.MuxHandler("preflight", handleValidationOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetValidationContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handleSecurity("JWTSec", h)
	h = handleValidationOrigin(h)
	service.Mux.Handle("GET", "/validation/:user_id", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "Validation", "action", "Get", "route", "GET /validation/:user_id", "security", "JWTSec")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewValidateValidationContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*ValidateValidationPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Validate(rctx)
	}
	h = handleValidationOrigin(h)
	service.Mux.Handle("POST", "/validation", ctrl.MuxHandler("validate", h, unmarshalValidateValidationPayload))
	service.LogInfo("mount", "ctrl", "Validation", "action", "Validate", "route", "POST /validation")
}

// handleValidationOrigin applies the CORS response headers corresponding to the origin.
func handleValidationOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:4200") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, Content-Type, Accept")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalValidateValidationPayload unmarshals the request body into the context request data Payload field.
func unmarshalValidateValidationPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &validateValidationPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
