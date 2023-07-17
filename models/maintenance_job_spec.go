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

// MaintenanceJobSpec MaintenanceJobSpec specifies details of a maintenance job
//
// swagger:model MaintenanceJobSpec
type MaintenanceJobSpec struct {

	// Options are used to set certain options for the given maintenance type
	// +optional
	Options map[string]string `json:"options,omitempty"`

	// Rollback indicates whether the maintenance done should be rolled back.
	// A rollback will only be started once the maintenance has started and thus is at least at phase Running
	// +optional
	Rollback bool `json:"rollback,omitempty"`

	// Type defines the type of maintenance that should be run
	// Depending on the type of the maintenance different option may be used
	// The maintenance will fail in case of an unsupported type
	Type string `json:"type,omitempty"`

	// cluster
	Cluster *ObjectReference `json:"cluster,omitempty"`
}

// Validate validates this maintenance job spec
func (m *MaintenanceJobSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MaintenanceJobSpec) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this maintenance job spec based on the context it is used
func (m *MaintenanceJobSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MaintenanceJobSpec) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MaintenanceJobSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MaintenanceJobSpec) UnmarshalBinary(b []byte) error {
	var res MaintenanceJobSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}