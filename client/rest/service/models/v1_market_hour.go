// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1MarketHour v1 market hour
//
// swagger:model v1MarketHour
type V1MarketHour struct {

	// next end time
	NextEndTime string `json:"nextEndTime,omitempty"`

	// next start time
	NextStartTime string `json:"nextStartTime,omitempty"`

	// period
	Period string `json:"period,omitempty"`

	// round interval
	RoundInterval string `json:"roundInterval,omitempty"`
}

// Validate validates this v1 market hour
func (m *V1MarketHour) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 market hour based on context it is used
func (m *V1MarketHour) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MarketHour) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MarketHour) UnmarshalBinary(b []byte) error {
	var res V1MarketHour
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
