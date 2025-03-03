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

// DatacenterSpecAWS DatacenterSpecAWS describes an AWS datacenter
//
// swagger:model DatacenterSpecAWS
type DatacenterSpecAWS struct {

	// The AWS region to use, e.g. "us-east-1". For a list of available regions, see
	// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html
	Region string `json:"region,omitempty"`

	// images
	Images ImageList `json:"images,omitempty"`
}

// Validate validates this datacenter spec a w s
func (m *DatacenterSpecAWS) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatacenterSpecAWS) validateImages(formats strfmt.Registry) error {
	if swag.IsZero(m.Images) { // not required
		return nil
	}

	if m.Images != nil {
		if err := m.Images.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("images")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("images")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this datacenter spec a w s based on the context it is used
func (m *DatacenterSpecAWS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatacenterSpecAWS) contextValidateImages(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Images) { // not required
		return nil
	}

	if err := m.Images.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("images")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("images")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatacenterSpecAWS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatacenterSpecAWS) UnmarshalBinary(b []byte) error {
	var res DatacenterSpecAWS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
