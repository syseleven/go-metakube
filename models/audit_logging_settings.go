// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuditLoggingSettings audit logging settings
//
// swagger:model AuditLoggingSettings
type AuditLoggingSettings struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this audit logging settings
func (m *AuditLoggingSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuditLoggingSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditLoggingSettings) UnmarshalBinary(b []byte) error {
	var res AuditLoggingSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
