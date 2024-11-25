// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Sys11AuthSettings sys11 auth settings
//
// swagger:model Sys11AuthSettings
type Sys11AuthSettings struct {

	// i a m authentication
	IAMAuthentication bool `json:"iamAuthentication,omitempty"`

	// +optional
	Realm string `json:"realm,omitempty"`
}

// Validate validates this sys11 auth settings
func (m *Sys11AuthSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sys11 auth settings based on context it is used
func (m *Sys11AuthSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Sys11AuthSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Sys11AuthSettings) UnmarshalBinary(b []byte) error {
	var res Sys11AuthSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
