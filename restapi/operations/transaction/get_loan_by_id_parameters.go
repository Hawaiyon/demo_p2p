// Code generated by go-swagger; DO NOT EDIT.

package transaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLoanByIDParams creates a new GetLoanByIDParams object
// with the default values initialized.
func NewGetLoanByIDParams() GetLoanByIDParams {
	var ()
	return GetLoanByIDParams{}
}

// GetLoanByIDParams contains all the bound params for the get loan by Id operation
// typically these are obtained from a http.Request
//
// swagger:parameters getLoanById
type GetLoanByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*ID of transaction
	  Required: true
	  In: path
	*/
	TransactionID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetLoanByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rTransactionID, rhkTransactionID, _ := route.Params.GetOK("transactionId")
	if err := o.bindTransactionID(rTransactionID, rhkTransactionID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLoanByIDParams) bindTransactionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("transactionId", "path", "int64", raw)
	}
	o.TransactionID = value

	return nil
}