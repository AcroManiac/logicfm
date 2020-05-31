// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/ahamtat/logicfm/internal/adapters/swagger/builder/models"
	"github.com/ahamtat/logicfm/pkg/version"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ahamtat/logicfm/internal/adapters/swagger/builder/restapi/operations"
	"github.com/ahamtat/logicfm/internal/adapters/swagger/builder/restapi/operations/info"
	"github.com/ahamtat/logicfm/internal/adapters/swagger/builder/restapi/operations/rule"
)

//go:generate swagger generate server --target ../../internal --name Builderapi --spec ../../api/builder-api-swagger.yaml --exclude-main

func configureFlags(api *operations.BuilderapiAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BuilderapiAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.RuleAddNewHandler == nil {
		api.RuleAddNewHandler = rule.AddNewHandlerFunc(func(params rule.AddNewParams) middleware.Responder {
			return middleware.NotImplemented("operation rule.AddNew has not yet been implemented")
		})
	}
	if api.RuleDeleteHandler == nil {
		api.RuleDeleteHandler = rule.DeleteHandlerFunc(func(params rule.DeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation rule.Delete has not yet been implemented")
		})
	}
	api.InfoGetHandler = info.GetHandlerFunc(func(params info.GetParams) middleware.Responder {
		return info.NewGetInfoOK().WithPayload(&models.Info{
			Branch:        version.Branch,
			Commit:        version.Commit,
			Date:          version.Date,
			Name:          "LogicFM Builder API",
			Release:       version.Release,
			RepositoryURL: version.Repo,
		})
	})
	if api.RuleUpdateHandler == nil {
		api.RuleUpdateHandler = rule.UpdateHandlerFunc(func(params rule.UpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation rule.Update has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
