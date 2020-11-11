// Code generated by go-swagger; DO NOT EDIT.

package domains_deprecated

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteDomainDeprecatedHandlerFunc turns a function with the right signature into a delete domain deprecated handler
type DeleteDomainDeprecatedHandlerFunc func(DeleteDomainDeprecatedParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteDomainDeprecatedHandlerFunc) Handle(params DeleteDomainDeprecatedParams) middleware.Responder {
	return fn(params)
}

// DeleteDomainDeprecatedHandler interface for that can handle valid delete domain deprecated params
type DeleteDomainDeprecatedHandler interface {
	Handle(DeleteDomainDeprecatedParams) middleware.Responder
}

// NewDeleteDomainDeprecated creates a new http.Handler for the delete domain deprecated operation
func NewDeleteDomainDeprecated(ctx *middleware.Context, handler DeleteDomainDeprecatedHandler) *DeleteDomainDeprecated {
	return &DeleteDomainDeprecated{Context: ctx, Handler: handler}
}

/*DeleteDomainDeprecated swagger:route DELETE /domains/{guid} domainsDeprecated deleteDomainDeprecated

Delete a Particular Domain (deprecated)

curl --insecure -i %s/v2/domains/{guid} -X DELETE -H 'Authorization: %s'

*/
type DeleteDomainDeprecated struct {
	Context *middleware.Context
	Handler DeleteDomainDeprecatedHandler
}

func (o *DeleteDomainDeprecated) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteDomainDeprecatedParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}