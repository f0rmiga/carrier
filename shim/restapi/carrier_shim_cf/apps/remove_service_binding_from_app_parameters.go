// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewRemoveServiceBindingFromAppParams creates a new RemoveServiceBindingFromAppParams object
// no default values defined in spec.
func NewRemoveServiceBindingFromAppParams() RemoveServiceBindingFromAppParams {

	return RemoveServiceBindingFromAppParams{}
}

// RemoveServiceBindingFromAppParams contains all the bound params for the remove service binding from app operation
// typically these are obtained from a http.Request
//
// swagger:parameters removeServiceBindingFromApp
type RemoveServiceBindingFromAppParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The guid parameter is used as a part of the request URL: '/v2/apps/:guid/service_bindings/:service_binding_guid'
	  Required: true
	  In: path
	*/
	GUID string
	/*The service_binding_guid parameter is used as a part of the request URL: '/v2/apps/:guid/service_bindings/:service_binding_guid'
	  Required: true
	  In: path
	*/
	ServiceBindingGUID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRemoveServiceBindingFromAppParams() beforehand.
func (o *RemoveServiceBindingFromAppParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rGUID, rhkGUID, _ := route.Params.GetOK("guid")
	if err := o.bindGUID(rGUID, rhkGUID, route.Formats); err != nil {
		res = append(res, err)
	}

	rServiceBindingGUID, rhkServiceBindingGUID, _ := route.Params.GetOK("service_binding_guid")
	if err := o.bindServiceBindingGUID(rServiceBindingGUID, rhkServiceBindingGUID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindGUID binds and validates parameter GUID from path.
func (o *RemoveServiceBindingFromAppParams) bindGUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.GUID = raw

	return nil
}

// bindServiceBindingGUID binds and validates parameter ServiceBindingGUID from path.
func (o *RemoveServiceBindingFromAppParams) bindServiceBindingGUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ServiceBindingGUID = raw

	return nil
}
