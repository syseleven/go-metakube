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

// DatacenterSpecOpenstack DatacenterSpecOpenstack describes an OpenStack datacenter
//
// swagger:model DatacenterSpecOpenstack
type DatacenterSpecOpenstack struct {

	// auth URL
	AuthURL string `json:"authURL,omitempty"`

	// availability zone
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// Openstack API client timeout, as a duration string (e.g. "60s").
	// Defaults to machine-controller-internal default.
	// +optional
	ClientTimeout string `json:"clientTimeout,omitempty"`

	// Openstack compute API (nova) client timeout, as a duration string (e.g. "60s").
	// Defaults to ClientTimeout.
	// +optional
	ComputeClientTimeout string `json:"computeClientTimeout,omitempty"`

	// Used for automatic network creation
	DNSServers []string `json:"dnsServers"`

	// EnabledFlavors lists enabled flavors for the given datacenter
	// +optional
	EnabledFlavors []string `json:"enabledFlavors"`

	// EnforceFloatingIP
	// +optional
	EnforceFloatingIP bool `json:"enforceFloatingIP,omitempty"`

	// IgnoreVomluAZ
	// +optional
	IgnoreVolumeAZ bool `json:"ignoreVolumeAZ,omitempty"`

	// ManageSecurityGroups sets mapped to the "manage-security-groups" setting in the cloud config.
	// See https://kubernetes.io/docs/concepts/cluster-administration/cloud-providers/#load-balancer
	// This setting defaults to true.
	// +optional
	ManageSecurityGroups bool `json:"manageSecurityGroups,omitempty"`

	// +optional
	NodeVolumeAttachLimit uint64 `json:"nodeVolumeAttachLimit,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// TrustDevicePath gets mapped to the "trust-device-path" setting in the cloud config.
	// See https://kubernetes.io/docs/concepts/cluster-administration/cloud-providers/#block-storage
	// This setting defaults to false.
	// +optional
	TrustDevicePath bool `json:"trustDevicePath,omitempty"`

	// UseOctavia gets mapped to the "use-octavia" setting in the cloud config.
	// use-octavia is enabled by default in CCM since v1.17.0, and disabled by
	// default with the in-tree cloud provider.
	// +optional
	UseOctavia bool `json:"useOctavia,omitempty"`

	// images
	Images ImageList `json:"images,omitempty"`

	// node size requirements
	NodeSizeRequirements *OpenstackNodeSizeRequirements `json:"nodeSizeRequirements,omitempty"`
}

// Validate validates this datacenter spec openstack
func (m *DatacenterSpecOpenstack) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeSizeRequirements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatacenterSpecOpenstack) validateImages(formats strfmt.Registry) error {
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

func (m *DatacenterSpecOpenstack) validateNodeSizeRequirements(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeSizeRequirements) { // not required
		return nil
	}

	if m.NodeSizeRequirements != nil {
		if err := m.NodeSizeRequirements.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeSizeRequirements")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeSizeRequirements")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this datacenter spec openstack based on the context it is used
func (m *DatacenterSpecOpenstack) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeSizeRequirements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatacenterSpecOpenstack) contextValidateImages(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DatacenterSpecOpenstack) contextValidateNodeSizeRequirements(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeSizeRequirements != nil {
		if err := m.NodeSizeRequirements.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeSizeRequirements")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeSizeRequirements")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatacenterSpecOpenstack) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatacenterSpecOpenstack) UnmarshalBinary(b []byte) error {
	var res DatacenterSpecOpenstack
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
