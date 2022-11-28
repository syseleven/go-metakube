// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenstackNetwork OpenstackNetwork is the object representing a openstack network.
//
// swagger:model OpenstackNetwork
type OpenstackNetwork struct {

	// External set if network is the external network
	External bool `json:"external,omitempty"`

	// Id uniquely identifies the current network
	ID string `json:"id,omitempty"`

	// Name is the name of the network
	Name string `json:"name,omitempty"`
}

// Validate validates this openstack network
func (m *OpenstackNetwork) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openstack network based on context it is used
func (m *OpenstackNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenstackNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenstackNetwork) UnmarshalBinary(b []byte) error {
	var res OpenstackNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
