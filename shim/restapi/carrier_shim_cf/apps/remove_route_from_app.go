// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RemoveRouteFromAppHandlerFunc turns a function with the right signature into a remove route from app handler
type RemoveRouteFromAppHandlerFunc func(RemoveRouteFromAppParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RemoveRouteFromAppHandlerFunc) Handle(params RemoveRouteFromAppParams) middleware.Responder {
	return fn(params)
}

// RemoveRouteFromAppHandler interface for that can handle valid remove route from app params
type RemoveRouteFromAppHandler interface {
	Handle(RemoveRouteFromAppParams) middleware.Responder
}

// NewRemoveRouteFromApp creates a new http.Handler for the remove route from app operation
func NewRemoveRouteFromApp(ctx *middleware.Context, handler RemoveRouteFromAppHandler) *RemoveRouteFromApp {
	return &RemoveRouteFromApp{Context: ctx, Handler: handler}
}

/*RemoveRouteFromApp swagger:route DELETE /apps/{guid}/routes/{route_guid} apps removeRouteFromApp

Remove Route from the App

curl --insecure -i %s/v2/apps/{guid}/routes/{route_guid} -X DELETE -H 'Authorization: %s'

*/
type RemoveRouteFromApp struct {
	Context *middleware.Context
	Handler RemoveRouteFromAppHandler
}

func (o *RemoveRouteFromApp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRemoveRouteFromAppParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}