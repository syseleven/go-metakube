// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddonConfigSpec AddonConfigSpec specifies configuration of addon
//
// swagger:model AddonConfigSpec
type AddonConfigSpec struct {

	// Controls that can be set for configured addon
	Controls []*AddonFormControl `json:"formSpec"`

	// Description of the configured addon, it will be displayed in the addon overview in the UI
	Description string `json:"description,omitempty"`

	// Logo of the configured addon, encoded in base64
	Logo string `json:"logo,omitempty"`

	// LogoFormat contains logo format of the configured addon, i.e. svg+xml
	LogoFormat string `json:"logoFormat,omitempty"`

	// ShortDescription of the configured addon that contains more detailed information about the addon,
	// it will be displayed in the addon details view in the UI
	ShortDescription string `json:"shortDescription,omitempty"`
}

// Validate validates this addon config spec
func (m *AddonConfigSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateControls(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddonConfigSpec) validateControls(formats strfmt.Registry) error {
	if swag.IsZero(m.Controls) { // not required
		return nil
	}

	for i := 0; i < len(m.Controls); i++ {
		if swag.IsZero(m.Controls[i]) { // not required
			continue
		}

		if m.Controls[i] != nil {
			if err := m.Controls[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("formSpec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("formSpec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this addon config spec based on the context it is used
func (m *AddonConfigSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateControls(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddonConfigSpec) contextValidateControls(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Controls); i++ {

		if m.Controls[i] != nil {
			if err := m.Controls[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("formSpec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("formSpec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddonConfigSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddonConfigSpec) UnmarshalBinary(b []byte) error {
	var res AddonConfigSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
