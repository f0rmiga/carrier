// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RemoveManagerFromOrganizationResponse remove manager from organization response
//
// swagger:model removeManagerFromOrganizationResponse
type RemoveManagerFromOrganizationResponse struct {

	// The app Events Url
	AppEventsURL string `json:"app_events_url,omitempty"`

	// The auditors Url
	AuditorsURL string `json:"auditors_url,omitempty"`

	// The billing Enabled
	BillingEnabled bool `json:"billing_enabled,omitempty"`

	// The billing Managers Url
	BillingManagersURL string `json:"billing_managers_url,omitempty"`

	// The domains Url
	DomainsURL string `json:"domains_url,omitempty"`

	// The managers Url
	ManagersURL string `json:"managers_url,omitempty"`

	// The name
	Name string `json:"name,omitempty"`

	// The private Domains Url
	PrivateDomainsURL string `json:"private_domains_url,omitempty"`

	// The quota Definition Guid
	QuotaDefinitionGUID string `json:"quota_definition_guid,omitempty"`

	// The quota Definition Url
	QuotaDefinitionURL string `json:"quota_definition_url,omitempty"`

	// The space Quota Definitions Url
	SpaceQuotaDefinitionsURL string `json:"space_quota_definitions_url,omitempty"`

	// The spaces Url
	SpacesURL string `json:"spaces_url,omitempty"`

	// The status
	Status string `json:"status,omitempty"`

	// The users Url
	UsersURL string `json:"users_url,omitempty"`
}

// Validate validates this remove manager from organization response
func (m *RemoveManagerFromOrganizationResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RemoveManagerFromOrganizationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoveManagerFromOrganizationResponse) UnmarshalBinary(b []byte) error {
	var res RemoveManagerFromOrganizationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
