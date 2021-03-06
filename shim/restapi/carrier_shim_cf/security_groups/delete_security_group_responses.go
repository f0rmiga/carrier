// Code generated by go-swagger; DO NOT EDIT.

package security_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteSecurityGroupNoContentCode is the HTTP code returned for type DeleteSecurityGroupNoContent
const DeleteSecurityGroupNoContentCode int = 204

/*DeleteSecurityGroupNoContent successful response

swagger:response deleteSecurityGroupNoContent
*/
type DeleteSecurityGroupNoContent struct {
}

// NewDeleteSecurityGroupNoContent creates DeleteSecurityGroupNoContent with default headers values
func NewDeleteSecurityGroupNoContent() *DeleteSecurityGroupNoContent {

	return &DeleteSecurityGroupNoContent{}
}

// WriteResponse to the client
func (o *DeleteSecurityGroupNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}
