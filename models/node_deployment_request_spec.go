// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NodeDeploymentRequestSpec NodeDeploymentRequestSpec node deployment request specification
//
// swagger:model NodeDeploymentRequestSpec
type NodeDeploymentRequestSpec struct {

	// nd
	Nd *NodeDeployment `json:"nd,omitempty"`
}

// Validate validates this node deployment request spec
func (m *NodeDeploymentRequestSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNd(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeDeploymentRequestSpec) validateNd(formats strfmt.Registry) error {

	if swag.IsZero(m.Nd) { // not required
		return nil
	}

	if m.Nd != nil {
		if err := m.Nd.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nd")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeDeploymentRequestSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeDeploymentRequestSpec) UnmarshalBinary(b []byte) error {
	var res NodeDeploymentRequestSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
