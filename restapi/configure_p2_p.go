// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"demo_p2p_bak/models"
	"demo_p2p_bak/models/dbmodels"
	"demo_p2p_bak/restapi/operations"
	"demo_p2p_bak/restapi/operations/transaction"
	"demo_p2p_bak/restapi/operations/user"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"strings"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name P2P --spec ../swagger.yml

func configureFlags(api *operations.P2PAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.P2PAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// 创建用户
	api.UserCreateUserHandler = user.CreateUserHandlerFunc(func(params user.CreateUserParams) middleware.Responder {
		if err := dbmodels.AddUser(params.Body); err != nil {
			return user.NewCreateUserDefault(500).WithPayload(
				&models.APIError{Code: 500, Message: swag.String(err.Error())})
		}
		return user.NewCreateUserCreated().WithPayload(params.Body)
	})

	// 查询用户
	api.UserGetUserByIDHandler = user.GetUserByIDHandlerFunc(func(params user.GetUserByIDParams) middleware.Responder {
		foundUser, err := dbmodels.GetUserInfo(params.UserID)
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return user.NewGetUserByIDDefault(404).WithPayload(
					&models.APIError{Code: 404, Message: swag.String(err.Error())})
			} else {
				return user.NewGetUserByIDDefault(500).WithPayload(
					&models.APIError{Code: 500, Message: swag.String(err.Error())})
			}

		}

		return user.NewGetUserByIDOK().WithPayload(foundUser)
	})

	api.TransactionAddLoanHandler = transaction.AddLoanHandlerFunc(func(params transaction.AddLoanParams) middleware.Responder {
		trans, err := dbmodels.AddTransaction(
			*(params.Body.LenderID), *(params.Body.BorrowerID), *(params.Body.Amount), "loan")
		if err != nil {
			return transaction.NewAddLoanDefault(500).WithPayload(
				&models.APIError{Code: 500, Message: swag.String(err.Error())})
		}
		ret := models.Loan{ID: trans.ID, CreatedDate: strfmt.DateTime(trans.CreatedDate),
			BorrowerID: &(trans.ToUserId), LenderID: &(trans.FromUserId), Status: trans.Status, Amount: &(trans.Amount)}
		return transaction.NewAddLoanCreated().WithPayload(&ret)

	})
	api.TransactionAddRepayHandler = transaction.AddRepayHandlerFunc(func(params transaction.AddRepayParams) middleware.Responder {
		trans, err := dbmodels.AddTransaction(*params.Body.BorrowerID, *params.Body.LenderID, *params.Body.RepayAmount, "repay")
		if err != nil {
			return transaction.NewAddRepayDefault(500).WithPayload(
				&models.APIError{Code: 500, Message: swag.String(err.Error())})
		}
		ret := models.Repayment{ID: trans.ID, CreatedDate: strfmt.DateTime(trans.CreatedDate),
			BorrowerID: &(trans.ToUserId), LenderID: &(trans.FromUserId), Status: trans.Status, RepayAmount: &(trans.Amount)}
		return transaction.NewAddRepayCreated().WithPayload(&ret)

	})

	api.TransactionGetDebtInfoHandler = transaction.GetDebtInfoHandlerFunc(func(params transaction.GetDebtInfoParams) middleware.Responder {
		lend, borrow, err := dbmodels.GetUserDebt(params.BaseUserID, params.ToUserID)
		if err != nil {
			return transaction.NewGetDebtInfoDefault(500).WithPayload(
				&models.APIError{Code: 500, Message: swag.String(err.Error())})
		}
		debt := models.Debt{Borrow: borrow, Lend: lend, BaseUserID: &(params.BaseUserID), ToUserID: &(params.ToUserID)}
		return transaction.NewGetDebtInfoOK().WithPayload(&debt)
	})

	api.TransactionGetLoanByIDHandler = transaction.GetLoanByIDHandlerFunc(func(params transaction.GetLoanByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation transaction.GetLoanByID has not yet been implemented")
	})
	api.TransactionGetRepayByIDHandler = transaction.GetRepayByIDHandlerFunc(func(params transaction.GetRepayByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation transaction.GetRepayByID has not yet been implemented")
	})

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
func configureServer(s *graceful.Server, scheme, addr string) {
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
