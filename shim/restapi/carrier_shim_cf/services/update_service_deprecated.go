// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateServiceDeprecatedHandlerFunc turns a function with the right signature into a update service deprecated handler
type UpdateServiceDeprecatedHandlerFunc func(UpdateServiceDeprecatedParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateServiceDeprecatedHandlerFunc) Handle(params UpdateServiceDeprecatedParams) middleware.Responder {
	return fn(params)
}

// UpdateServiceDeprecatedHandler interface for that can handle valid update service deprecated params
type UpdateServiceDeprecatedHandler interface {
	Handle(UpdateServiceDeprecatedParams) middleware.Responder
}

// NewUpdateServiceDeprecated creates a new http.Handler for the update service deprecated operation
func NewUpdateServiceDeprecated(ctx *middleware.Context, handler UpdateServiceDeprecatedHandler) *UpdateServiceDeprecated {
	return &UpdateServiceDeprecated{Context: ctx, Handler: handler}
}

/*UpdateServiceDeprecated swagger:route PUT /services services updateServiceDeprecated

Updating a Service (deprecated)

curl --insecure -i %s/v2/services -X PUT -H 'Authorization: %s' -d '%s'

*/
type UpdateServiceDeprecated struct {
	Context *middleware.Context
	Handler UpdateServiceDeprecatedHandler
}

func (o *UpdateServiceDeprecated) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateServiceDeprecatedParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}