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

// EtcdRestoreStatus etcd restore status
//
// swagger:model EtcdRestoreStatus
type EtcdRestoreStatus struct {

	// restore time
	RestoreTime string `json:"restoreTime,omitempty"`

	// phase
	Phase EtcdRestorePhase `json:"phase,omitempty"`
}

// Validate validates this etcd restore status
func (m *EtcdRestoreStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EtcdRestoreStatus) validatePhase(formats strfmt.Registry) error {
	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	if err := m.Phase.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("phase")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("phase")
		}
		return err
	}

	return nil
}

// ContextValidate validate this etcd restore status based on the context it is used
func (m *EtcdRestoreStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePhase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EtcdRestoreStatus) contextValidatePhase(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	if err := m.Phase.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("phase")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("phase")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EtcdRestoreStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EtcdRestoreStatus) UnmarshalBinary(b []byte) error {
	var res EtcdRestoreStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
