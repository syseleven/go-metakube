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

// OperatingSystemSpec OperatingSystemSpec represents the collection of os specific settings. Only one must be set at a time.
//
// swagger:model OperatingSystemSpec
type OperatingSystemSpec struct {

	// flatcar
	Flatcar *FlatcarSpec `json:"flatcar,omitempty"`

	// ubuntu
	Ubuntu *UbuntuSpec `json:"ubuntu,omitempty"`
}

// Validate validates this operating system spec
func (m *OperatingSystemSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlatcar(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUbuntu(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OperatingSystemSpec) validateFlatcar(formats strfmt.Registry) error {
	if swag.IsZero(m.Flatcar) { // not required
		return nil
	}

	if m.Flatcar != nil {
		if err := m.Flatcar.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flatcar")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flatcar")
			}
			return err
		}
	}

	return nil
}

func (m *OperatingSystemSpec) validateUbuntu(formats strfmt.Registry) error {
	if swag.IsZero(m.Ubuntu) { // not required
		return nil
	}

	if m.Ubuntu != nil {
		if err := m.Ubuntu.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ubuntu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ubuntu")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this operating system spec based on the context it is used
func (m *OperatingSystemSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFlatcar(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUbuntu(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OperatingSystemSpec) contextValidateFlatcar(ctx context.Context, formats strfmt.Registry) error {

	if m.Flatcar != nil {

		if swag.IsZero(m.Flatcar) { // not required
			return nil
		}

		if err := m.Flatcar.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flatcar")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flatcar")
			}
			return err
		}
	}

	return nil
}

func (m *OperatingSystemSpec) contextValidateUbuntu(ctx context.Context, formats strfmt.Registry) error {

	if m.Ubuntu != nil {

		if swag.IsZero(m.Ubuntu) { // not required
			return nil
		}

		if err := m.Ubuntu.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ubuntu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ubuntu")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OperatingSystemSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperatingSystemSpec) UnmarshalBinary(b []byte) error {
	var res OperatingSystemSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
