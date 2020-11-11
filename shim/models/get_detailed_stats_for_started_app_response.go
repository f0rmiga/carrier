// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetDetailedStatsForStartedAppResponse get detailed stats for started app response
//
// swagger:model getDetailedStatsForStartedAppResponse
type GetDetailedStatsForStartedAppResponse struct {

	// The state
	State string `json:"state,omitempty"`

	// The stats
	Stats GenericObject `json:"stats,omitempty"`
}

// Validate validates this get detailed stats for started app response
func (m *GetDetailedStatsForStartedAppResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetDetailedStatsForStartedAppResponse) validateStats(formats strfmt.Registry) error {

	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if err := m.Stats.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("stats")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetDetailedStatsForStartedAppResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetDetailedStatsForStartedAppResponse) UnmarshalBinary(b []byte) error {
	var res GetDetailedStatsForStartedAppResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}