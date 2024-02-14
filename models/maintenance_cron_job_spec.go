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

// MaintenanceCronJobSpec MaintenanceCronJobSpec represents an object holding the configuration for maintenance cron jobs.
//
// swagger:model MaintenanceCronJobSpec
type MaintenanceCronJobSpec struct {

	// Schedule is a cron expression defining when to perform the maintenance.
	Schedule string `json:"schedule,omitempty"`

	// maintenance job template
	MaintenanceJobTemplate *MaintenanceJobTemplate `json:"maintenanceJobTemplate,omitempty"`
}

// Validate validates this maintenance cron job spec
func (m *MaintenanceCronJobSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaintenanceJobTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MaintenanceCronJobSpec) validateMaintenanceJobTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.MaintenanceJobTemplate) { // not required
		return nil
	}

	if m.MaintenanceJobTemplate != nil {
		if err := m.MaintenanceJobTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maintenanceJobTemplate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("maintenanceJobTemplate")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this maintenance cron job spec based on the context it is used
func (m *MaintenanceCronJobSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMaintenanceJobTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MaintenanceCronJobSpec) contextValidateMaintenanceJobTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.MaintenanceJobTemplate != nil {
		if err := m.MaintenanceJobTemplate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maintenanceJobTemplate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("maintenanceJobTemplate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MaintenanceCronJobSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MaintenanceCronJobSpec) UnmarshalBinary(b []byte) error {
	var res MaintenanceCronJobSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
