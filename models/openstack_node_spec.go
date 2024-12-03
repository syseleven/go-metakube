// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OpenstackNodeSpec OpenstackNodeSpec openstack node settings
//
// swagger:model OpenstackNodeSpec
type OpenstackNodeSpec struct {

	// if not set, the default AZ from the Datacenter spec will be used
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// instance flavor
	// Required: true
	Flavor *string `json:"flavor"`

	// image to use
	// Required: true
	Image *string `json:"image"`

	// Period of time to check for instance ready status, i.e. 10s/1m
	InstanceReadyCheckPeriod string `json:"instanceReadyCheckPeriod,omitempty"`

	// Max time to wait for the instance to be ready, i.e. 10s/1m
	InstanceReadyCheckTimeout string `json:"instanceReadyCheckTimeout,omitempty"`

	// if set, the rootDisk will be a volume. If not, the rootDisk will be on ephemeral storage and its size will be derived from the flavor
	RootDiskSizeGB *int64 `json:"diskSize,omitempty"`

	// Openstack server group ID of the instance
	// Defaults to the corresponding setting from the cluster spec
	ServerGroupID string `json:"serverGroupID,omitempty"`

	// Additional metadata to set
	Tags map[string]string `json:"tags,omitempty"`

	// Defines whether floating ip should be used
	UseFloatingIP *bool `json:"useFloatingIP,omitempty"`
}

// Validate validates this openstack node spec
func (m *OpenstackNodeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlavor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenstackNodeSpec) validateFlavor(formats strfmt.Registry) error {

	if err := validate.Required("flavor", "body", m.Flavor); err != nil {
		return err
	}

	return nil
}

func (m *OpenstackNodeSpec) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this openstack node spec based on context it is used
func (m *OpenstackNodeSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenstackNodeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenstackNodeSpec) UnmarshalBinary(b []byte) error {
	var res OpenstackNodeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
