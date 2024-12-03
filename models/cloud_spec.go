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

// CloudSpec CloudSpec mutually stores access data to a cloud provider.
//
// swagger:model CloudSpec
type CloudSpec struct {

	// DatacenterName where the users 'cloud' lives in.
	DatacenterName string `json:"dc,omitempty"`

	// aws
	Aws *AWSCloudSpec `json:"aws,omitempty"`

	// azure
	Azure *AzureCloudSpec `json:"azure,omitempty"`

	// fake
	Fake *FakeCloudSpec `json:"fake,omitempty"`

	// openstack
	Openstack *OpenstackCloudSpec `json:"openstack,omitempty"`
}

// Validate validates this cloud spec
func (m *CloudSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFake(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenstack(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudSpec) validateAws(formats strfmt.Registry) error {
	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *CloudSpec) validateAzure(formats strfmt.Registry) error {
	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *CloudSpec) validateFake(formats strfmt.Registry) error {
	if swag.IsZero(m.Fake) { // not required
		return nil
	}

	if m.Fake != nil {
		if err := m.Fake.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fake")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fake")
			}
			return err
		}
	}

	return nil
}

func (m *CloudSpec) validateOpenstack(formats strfmt.Registry) error {
	if swag.IsZero(m.Openstack) { // not required
		return nil
	}

	if m.Openstack != nil {
		if err := m.Openstack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openstack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("openstack")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cloud spec based on the context it is used
func (m *CloudSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAws(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFake(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpenstack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudSpec) contextValidateAws(ctx context.Context, formats strfmt.Registry) error {

	if m.Aws != nil {

		if swag.IsZero(m.Aws) { // not required
			return nil
		}

		if err := m.Aws.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *CloudSpec) contextValidateAzure(ctx context.Context, formats strfmt.Registry) error {

	if m.Azure != nil {

		if swag.IsZero(m.Azure) { // not required
			return nil
		}

		if err := m.Azure.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *CloudSpec) contextValidateFake(ctx context.Context, formats strfmt.Registry) error {

	if m.Fake != nil {

		if swag.IsZero(m.Fake) { // not required
			return nil
		}

		if err := m.Fake.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fake")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fake")
			}
			return err
		}
	}

	return nil
}

func (m *CloudSpec) contextValidateOpenstack(ctx context.Context, formats strfmt.Registry) error {

	if m.Openstack != nil {

		if swag.IsZero(m.Openstack) { // not required
			return nil
		}

		if err := m.Openstack.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openstack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("openstack")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudSpec) UnmarshalBinary(b []byte) error {
	var res CloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
