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

// BackupCredentials BackupCredentials contains credentials for etcd backups
//
// swagger:model BackupCredentials
type BackupCredentials struct {

	// s3
	S3 *S3BackupCredentials `json:"s3,omitempty"`
}

// Validate validates this backup credentials
func (m *BackupCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateS3(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupCredentials) validateS3(formats strfmt.Registry) error {
	if swag.IsZero(m.S3) { // not required
		return nil
	}

	if m.S3 != nil {
		if err := m.S3.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("s3")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this backup credentials based on the context it is used
func (m *BackupCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateS3(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupCredentials) contextValidateS3(ctx context.Context, formats strfmt.Registry) error {

	if m.S3 != nil {

		if swag.IsZero(m.S3) { // not required
			return nil
		}

		if err := m.S3.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("s3")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupCredentials) UnmarshalBinary(b []byte) error {
	var res BackupCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
