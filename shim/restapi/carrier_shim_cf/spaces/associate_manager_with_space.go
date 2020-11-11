// Code generated by go-swagger; DO NOT EDIT.

package spaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AssociateManagerWithSpaceHandlerFunc turns a function with the right signature into a associate manager with space handler
type AssociateManagerWithSpaceHandlerFunc func(AssociateManagerWithSpaceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AssociateManagerWithSpaceHandlerFunc) Handle(params AssociateManagerWithSpaceParams) middleware.Responder {
	return fn(params)
}

// AssociateManagerWithSpaceHandler interface for that can handle valid associate manager with space params
type AssociateManagerWithSpaceHandler interface {
	Handle(AssociateManagerWithSpaceParams) middleware.Responder
}

// NewAssociateManagerWithSpace creates a new http.Handler for the associate manager with space operation
func NewAssociateManagerWithSpace(ctx *middleware.Context, handler AssociateManagerWithSpaceHandler) *AssociateManagerWithSpace {
	return &AssociateManagerWithSpace{Context: ctx, Handler: handler}
}

/*AssociateManagerWithSpace swagger:route PUT /spaces/{guid}/managers/{manager_guid} spaces associateManagerWithSpace

Associate Manager with the Space

curl --insecure -i %s/v2/spaces/{guid}/managers/{manager_guid} -X PUT -H 'Authorization: %s'

*/
type AssociateManagerWithSpace struct {
	Context *middleware.Context
	Handler AssociateManagerWithSpaceHandler
}

func (o *AssociateManagerWithSpace) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAssociateManagerWithSpaceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}