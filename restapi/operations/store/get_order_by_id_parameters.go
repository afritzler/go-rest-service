package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetOrderByIDParams creates a new GetOrderByIDParams object
// with the default values initialized.
func NewGetOrderByIDParams() GetOrderByIDParams {
	var ()
	return GetOrderByIDParams{}
}

// GetOrderByIDParams contains all the bound params for the get order by Id operation
// typically these are obtained from a http.Request
//
// swagger:parameters getOrderById
type GetOrderByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*ID of pet that needs to be fetched
	  Required: true
	  Maximum: 10
	  Minimum: 1
	  In: path
	*/
	OrderID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetOrderByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rOrderID, rhkOrderID, _ := route.Params.GetOK("orderId")
	if err := o.bindOrderID(rOrderID, rhkOrderID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrderByIDParams) bindOrderID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("orderId", "path", "int64", raw)
	}
	o.OrderID = value

	if err := o.validateOrderID(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetOrderByIDParams) validateOrderID(formats strfmt.Registry) error {

	if err := validate.MinimumInt("orderId", "path", int64(o.OrderID), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("orderId", "path", int64(o.OrderID), 10, false); err != nil {
		return err
	}

	return nil
}
