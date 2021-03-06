// Code generated by go-swagger; DO NOT EDIT.

package transaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetDebtInfoHandlerFunc turns a function with the right signature into a get debt info handler
type GetDebtInfoHandlerFunc func(GetDebtInfoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDebtInfoHandlerFunc) Handle(params GetDebtInfoParams) middleware.Responder {
	return fn(params)
}

// GetDebtInfoHandler interface for that can handle valid get debt info params
type GetDebtInfoHandler interface {
	Handle(GetDebtInfoParams) middleware.Responder
}

// NewGetDebtInfo creates a new http.Handler for the get debt info operation
func NewGetDebtInfo(ctx *middleware.Context, handler GetDebtInfoHandler) *GetDebtInfo {
	return &GetDebtInfo{Context: ctx, Handler: handler}
}

/*GetDebtInfo swagger:route GET /debt/{baseUserId}/{toUserId} transaction getDebtInfo

根据基准用户ID，返回跟另外一个用户之间的债务情况

*/
type GetDebtInfo struct {
	Context *middleware.Context
	Handler GetDebtInfoHandler
}

func (o *GetDebtInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetDebtInfoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
