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

// ClusterHealth ClusterHealth stores health information about the cluster's components.
//
// swagger:model ClusterHealth
type ClusterHealth struct {

	// apiserver
	Apiserver HealthStatus `json:"apiserver,omitempty"`

	// cloud provider infrastructure
	CloudProviderInfrastructure HealthStatus `json:"cloudProviderInfrastructure,omitempty"`

	// controller
	Controller HealthStatus `json:"controller,omitempty"`

	// etcd
	Etcd HealthStatus `json:"etcd,omitempty"`

	// machine controller
	MachineController HealthStatus `json:"machineController,omitempty"`

	// scheduler
	Scheduler HealthStatus `json:"scheduler,omitempty"`

	// user cluster controller manager
	UserClusterControllerManager HealthStatus `json:"userClusterControllerManager,omitempty"`
}

// Validate validates this cluster health
func (m *ClusterHealth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApiserver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudProviderInfrastructure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateController(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEtcd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineController(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduler(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserClusterControllerManager(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterHealth) validateApiserver(formats strfmt.Registry) error {
	if swag.IsZero(m.Apiserver) { // not required
		return nil
	}

	if err := m.Apiserver.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("apiserver")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("apiserver")
		}
		return err
	}

	return nil
}

func (m *ClusterHealth) validateCloudProviderInfrastructure(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudProviderInfrastructure) { // not required
		return nil
	}

	if err := m.CloudProviderInfrastructure.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cloudProviderInfrastructure")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("cloudProviderInfrastructure")
		}
		return err
	}

	return nil
}

func (m *ClusterHealth) validateController(formats strfmt.Registry) error {
	if swag.IsZero(m.Controller) { // not required
		return nil
	}

	if err := m.Controller.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("controller")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("controller")
		}
		return err
	}

	return nil
}

func (m *ClusterHealth) validateEtcd(formats strfmt.Registry) error {
	if swag.IsZero(m.Etcd) { // not required
		return nil
	}

	if err := m.Etcd.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("etcd")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("etcd")
		}
		return err
	}

	return nil
}

func (m *ClusterHealth) validateMachineController(formats strfmt.Registry) error {
	if swag.IsZero(m.MachineController) { // not required
		return nil
	}

	if err := m.MachineController.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("machineController")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("machineController")
		}
		return err
	}

	return nil
}

func (m *ClusterHealth) validateScheduler(formats strfmt.Registry) error {
	if swag.IsZero(m.Scheduler) { // not required
		return nil
	}

	if err := m.Scheduler.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scheduler")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("scheduler")
		}
		return err
	}

	return nil
}

func (m *ClusterHealth) validateUserClusterControllerManager(formats strfmt.Registry) error {
	if swag.IsZero(m.UserClusterControllerManager) { // not required
		return nil
	}

	if err := m.UserClusterControllerManager.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("userClusterControllerManager")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("userClusterControllerManager")
		}
		return err
	}

	return nil
}

// ContextValidate validate this cluster health based on the context it is used
func (m *ClusterHealth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApiserver(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCloudProviderInfrastructure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateController(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEtcd(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMachineController(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScheduler(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserClusterControllerManager(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterHealth) contextValidateApiserver(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Apiserver.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("apiserver")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("apiserver")
		}
		return err
	}

	return nil
}

func (m *ClusterHealth) contextValidateCloudProviderInfrastructure(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CloudProviderInfrastructure.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cloudProviderInfrastructure")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("cloudProviderInfrastructure")
		}
		return err
	}

	return nil
}

func (m *ClusterHealth) contextValidateController(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Controller.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("controller")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("controller")
		}
		return err
	}

	return nil
}

func (m *ClusterHealth) contextValidateEtcd(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Etcd.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("etcd")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("etcd")
		}
		return err
	}

	return nil
}

func (m *ClusterHealth) contextValidateMachineController(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MachineController.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("machineController")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("machineController")
		}
		return err
	}

	return nil
}

func (m *ClusterHealth) contextValidateScheduler(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Scheduler.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scheduler")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("scheduler")
		}
		return err
	}

	return nil
}

func (m *ClusterHealth) contextValidateUserClusterControllerManager(ctx context.Context, formats strfmt.Registry) error {

	if err := m.UserClusterControllerManager.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("userClusterControllerManager")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("userClusterControllerManager")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterHealth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterHealth) UnmarshalBinary(b []byte) error {
	var res ClusterHealth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
