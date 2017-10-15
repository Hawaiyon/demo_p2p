// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"demo_p2p_bak/restapi/operations/transaction"
	"demo_p2p_bak/restapi/operations/user"
)

// NewP2PAPI creates a new P2P instance
func NewP2PAPI(spec *loads.Document) *P2PAPI {
	return &P2PAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		TransactionAddLoanHandler: transaction.AddLoanHandlerFunc(func(params transaction.AddLoanParams) middleware.Responder {
			return middleware.NotImplemented("operation TransactionAddLoan has not yet been implemented")
		}),
		TransactionAddRepayHandler: transaction.AddRepayHandlerFunc(func(params transaction.AddRepayParams) middleware.Responder {
			return middleware.NotImplemented("operation TransactionAddRepay has not yet been implemented")
		}),
		UserCreateUserHandler: user.CreateUserHandlerFunc(func(params user.CreateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation UserCreateUser has not yet been implemented")
		}),
		TransactionGetDebtInfoHandler: transaction.GetDebtInfoHandlerFunc(func(params transaction.GetDebtInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation TransactionGetDebtInfo has not yet been implemented")
		}),
		TransactionGetLoanByIDHandler: transaction.GetLoanByIDHandlerFunc(func(params transaction.GetLoanByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation TransactionGetLoanByID has not yet been implemented")
		}),
		TransactionGetRepayByIDHandler: transaction.GetRepayByIDHandlerFunc(func(params transaction.GetRepayByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation TransactionGetRepayByID has not yet been implemented")
		}),
		UserGetUserByIDHandler: user.GetUserByIDHandlerFunc(func(params user.GetUserByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation UserGetUserByID has not yet been implemented")
		}),
	}
}

/*P2PAPI This is a P2P trade system. */
type P2PAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// TransactionAddLoanHandler sets the operation handler for the add loan operation
	TransactionAddLoanHandler transaction.AddLoanHandler
	// TransactionAddRepayHandler sets the operation handler for the add repay operation
	TransactionAddRepayHandler transaction.AddRepayHandler
	// UserCreateUserHandler sets the operation handler for the create user operation
	UserCreateUserHandler user.CreateUserHandler
	// TransactionGetDebtInfoHandler sets the operation handler for the get debt info operation
	TransactionGetDebtInfoHandler transaction.GetDebtInfoHandler
	// TransactionGetLoanByIDHandler sets the operation handler for the get loan by Id operation
	TransactionGetLoanByIDHandler transaction.GetLoanByIDHandler
	// TransactionGetRepayByIDHandler sets the operation handler for the get repay by Id operation
	TransactionGetRepayByIDHandler transaction.GetRepayByIDHandler
	// UserGetUserByIDHandler sets the operation handler for the get user by ID operation
	UserGetUserByIDHandler user.GetUserByIDHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *P2PAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *P2PAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *P2PAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *P2PAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *P2PAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *P2PAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *P2PAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the P2PAPI
func (o *P2PAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.TransactionAddLoanHandler == nil {
		unregistered = append(unregistered, "transaction.AddLoanHandler")
	}

	if o.TransactionAddRepayHandler == nil {
		unregistered = append(unregistered, "transaction.AddRepayHandler")
	}

	if o.UserCreateUserHandler == nil {
		unregistered = append(unregistered, "user.CreateUserHandler")
	}

	if o.TransactionGetDebtInfoHandler == nil {
		unregistered = append(unregistered, "transaction.GetDebtInfoHandler")
	}

	if o.TransactionGetLoanByIDHandler == nil {
		unregistered = append(unregistered, "transaction.GetLoanByIDHandler")
	}

	if o.TransactionGetRepayByIDHandler == nil {
		unregistered = append(unregistered, "transaction.GetRepayByIDHandler")
	}

	if o.UserGetUserByIDHandler == nil {
		unregistered = append(unregistered, "user.GetUserByIDHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *P2PAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *P2PAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *P2PAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *P2PAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *P2PAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *P2PAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the p2 p API
func (o *P2PAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *P2PAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/loan"] = transaction.NewAddLoan(o.context, o.TransactionAddLoanHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/repayment"] = transaction.NewAddRepay(o.context, o.TransactionAddRepayHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user"] = user.NewCreateUser(o.context, o.UserCreateUserHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/debt"] = transaction.NewGetDebtInfo(o.context, o.TransactionGetDebtInfoHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/loan/{transactionId}"] = transaction.NewGetLoanByID(o.context, o.TransactionGetLoanByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repayment/{transactionId}"] = transaction.NewGetRepayByID(o.context, o.TransactionGetRepayByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/{userID}"] = user.NewGetUserByID(o.context, o.UserGetUserByIDHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *P2PAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *P2PAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}