// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterMetrics ClusterMetrics defines a metric for the given cluster
//
// swagger:model ClusterMetrics
type ClusterMetrics struct {

	// name
	Name string `json:"name,omitempty"`

	// control plane
	ControlPlane *ControlPlaneMetrics `json:"controlPlane,omitempty"`

	// nodes
	Nodes *NodesMetric `json:"nodes,omitempty"`
}

// Validate validates this cluster metrics
func (m *ClusterMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateControlPlane(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterMetrics) validateControlPlane(formats strfmt.Registry) error {
	if swag.IsZero(m.ControlPlane) { // not required
		return nil
	}

	if m.ControlPlane != nil {
		if err := m.ControlPlane.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("controlPlane")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("controlPlane")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterMetrics) validateNodes(formats strfmt.Registry) error {
	if swag.IsZero(m.Nodes) { // not required
		return nil
	}

	if m.Nodes != nil {
		if err := m.Nodes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodes")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster metrics based on the context it is used
func (m *ClusterMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateControlPlane(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterMetrics) contextValidateControlPlane(ctx context.Context, formats strfmt.Registry) error {

	if m.ControlPlane != nil {
		if err := m.ControlPlane.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("controlPlane")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("controlPlane")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterMetrics) contextValidateNodes(ctx context.Context, formats strfmt.Registry) error {

	if m.Nodes != nil {
		if err := m.Nodes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterMetrics) UnmarshalBinary(b []byte) error {
	var res ClusterMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
