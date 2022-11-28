// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MasterVersion MasterVersion describes a version of the master components
//
// swagger:model MasterVersion
type MasterVersion struct {

	// default
	Default bool `json:"default,omitempty"`

	// If true, then given version control plane version is not compatible
	// with one of the kubelets inside cluster and shouldn't be used.
	RestrictedByKubeletVersion bool `json:"restrictedByKubeletVersion,omitempty"`

	// version
	Version Version `json:"version,omitempty"`
}

// Validate validates this master version
func (m *MasterVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this master version based on context it is used
func (m *MasterVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MasterVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MasterVersion) UnmarshalBinary(b []byte) error {
	var res MasterVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
