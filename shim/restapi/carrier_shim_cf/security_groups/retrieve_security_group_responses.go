// Code generated by go-swagger; DO NOT EDIT.

package security_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// RetrieveSecurityGroupOKCode is the HTTP code returned for type RetrieveSecurityGroupOK
const RetrieveSecurityGroupOKCode int = 200

/*RetrieveSecurityGroupOK successful response

swagger:response retrieveSecurityGroupOK
*/
type RetrieveSecurityGroupOK struct {

	/*
	  In: Body
	*/
	Payload *models.RetrieveSecurityGroupResponse `json:"body,omitempty"`
}

// NewRetrieveSecurityGroupOK creates RetrieveSecurityGroupOK with default headers values
func NewRetrieveSecurityGroupOK() *RetrieveSecurityGroupOK {

	return &RetrieveSecurityGroupOK{}
}

// WithPayload adds the payload to the retrieve security group o k response
func (o *RetrieveSecurityGroupOK) WithPayload(payload *models.RetrieveSecurityGroupResponse) *RetrieveSecurityGroupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve security group o k response
func (o *RetrieveSecurityGroupOK) SetPayload(payload *models.RetrieveSecurityGroupResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveSecurityGroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
