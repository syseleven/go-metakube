// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddonFormControl AddonFormControl specifies addon form control
//
// swagger:model AddonFormControl
type AddonFormControl struct {

	// DisplayName is visible in the UI
	DisplayName string `json:"displayName,omitempty"`

	// HelpText is visible in the UI next to the control
	HelpText string `json:"helpText,omitempty"`

	// InternalName is used internally to save in the addon object
	InternalName string `json:"internalName,omitempty"`

	// Required indicates if the control has to be set
	Required bool `json:"required,omitempty"`

	// Type of displayed control
	Type string `json:"type,omitempty"`
}

// Validate validates this addon form control
func (m *AddonFormControl) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this addon form control based on context it is used
func (m *AddonFormControl) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AddonFormControl) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddonFormControl) UnmarshalBinary(b []byte) error {
	var res AddonFormControl
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
