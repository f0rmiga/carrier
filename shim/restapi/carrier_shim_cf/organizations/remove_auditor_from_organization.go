// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RemoveAuditorFromOrganizationHandlerFunc turns a function with the right signature into a remove auditor from organization handler
type RemoveAuditorFromOrganizationHandlerFunc func(RemoveAuditorFromOrganizationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RemoveAuditorFromOrganizationHandlerFunc) Handle(params RemoveAuditorFromOrganizationParams) middleware.Responder {
	return fn(params)
}

// RemoveAuditorFromOrganizationHandler interface for that can handle valid remove auditor from organization params
type RemoveAuditorFromOrganizationHandler interface {
	Handle(RemoveAuditorFromOrganizationParams) middleware.Responder
}

// NewRemoveAuditorFromOrganization creates a new http.Handler for the remove auditor from organization operation
func NewRemoveAuditorFromOrganization(ctx *middleware.Context, handler RemoveAuditorFromOrganizationHandler) *RemoveAuditorFromOrganization {
	return &RemoveAuditorFromOrganization{Context: ctx, Handler: handler}
}

/*RemoveAuditorFromOrganization swagger:route DELETE /organizations/{guid}/auditors/{auditor_guid} organizations removeAuditorFromOrganization

Remove Auditor from the Organization

curl --insecure -i %s/v2/organizations/{guid}/auditors/{auditor_guid} -X DELETE -H 'Authorization: %s'

*/
type RemoveAuditorFromOrganization struct {
	Context *middleware.Context
	Handler RemoveAuditorFromOrganizationHandler
}

func (o *RemoveAuditorFromOrganization) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRemoveAuditorFromOrganizationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
