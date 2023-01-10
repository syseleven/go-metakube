// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CleanupOptions cleanup options
//
// swagger:model CleanupOptions
type CleanupOptions struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// enforced
	Enforced bool `json:"enforced,omitempty"`
}

// Validate validates this cleanup options
func (m *CleanupOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cleanup options based on context it is used
func (m *CleanupOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CleanupOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CleanupOptions) UnmarshalBinary(b []byte) error {
	var res CleanupOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
