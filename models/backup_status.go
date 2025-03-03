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

// BackupStatus backup status
//
// swagger:model BackupStatus
type BackupStatus struct {

	// backup finished time
	BackupFinishedTime string `json:"backupFinishedTime,omitempty"`

	// backup message
	BackupMessage string `json:"backupMessage,omitempty"`

	// backup name
	BackupName string `json:"backupName,omitempty"`

	// backup start time
	BackupStartTime string `json:"backupStartTime,omitempty"`

	// delete finished time
	DeleteFinishedTime string `json:"deleteFinishedTime,omitempty"`

	// delete job name
	DeleteJobName string `json:"deleteJobName,omitempty"`

	// delete message
	DeleteMessage string `json:"deleteMessage,omitempty"`

	// delete start time
	DeleteStartTime string `json:"deleteStartTime,omitempty"`

	// job name
	JobName string `json:"jobName,omitempty"`

	// ScheduledTime will always be set when the BackupStatus is created, so it'll never be nil
	ScheduledTime string `json:"scheduledTime,omitempty"`

	// backup phase
	BackupPhase BackupStatusPhase `json:"backupPhase,omitempty"`

	// delete phase
	DeletePhase BackupStatusPhase `json:"deletePhase,omitempty"`
}

// Validate validates this backup status
func (m *BackupStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupPhase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletePhase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupStatus) validateBackupPhase(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupPhase) { // not required
		return nil
	}

	if err := m.BackupPhase.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("backupPhase")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("backupPhase")
		}
		return err
	}

	return nil
}

func (m *BackupStatus) validateDeletePhase(formats strfmt.Registry) error {
	if swag.IsZero(m.DeletePhase) { // not required
		return nil
	}

	if err := m.DeletePhase.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deletePhase")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("deletePhase")
		}
		return err
	}

	return nil
}

// ContextValidate validate this backup status based on the context it is used
func (m *BackupStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupPhase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeletePhase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupStatus) contextValidateBackupPhase(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.BackupPhase) { // not required
		return nil
	}

	if err := m.BackupPhase.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("backupPhase")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("backupPhase")
		}
		return err
	}

	return nil
}

func (m *BackupStatus) contextValidateDeletePhase(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.DeletePhase) { // not required
		return nil
	}

	if err := m.DeletePhase.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deletePhase")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("deletePhase")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupStatus) UnmarshalBinary(b []byte) error {
	var res BackupStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
