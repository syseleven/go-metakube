// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenstackCloudSpec OpenstackCloudSpec specifies access data to an OpenStack cloud.
//
// swagger:model OpenstackCloudSpec
type OpenstackCloudSpec struct {

	// application credential ID
	ApplicationCredentialID string `json:"applicationCredentialID,omitempty"`

	// application credential secret
	ApplicationCredentialSecret string `json:"applicationCredentialSecret,omitempty"`

	// domain
	Domain string `json:"domain,omitempty"`

	// FloatingIPPool holds the name of the public network
	// The public network is reachable from the outside world
	// and should provide the pool of IP addresses to choose from.
	//
	// When specified, all worker nodes will receive a public ip from this floating ip pool
	//
	// Note that the network is external if the "External" field is set to true
	FloatingIPPool string `json:"floatingIpPool,omitempty"`

	// Network holds the name of the internal network
	// When specified, all worker nodes will be attached to this network. If not specified, a network, subnet & router will be created
	//
	// Note that the network is internal if the "External" field is set to false
	Network string `json:"network,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// router ID
	RouterID string `json:"routerID,omitempty"`

	// security groups
	SecurityGroups string `json:"securityGroups,omitempty"`

	// ServerGroupID used as schedule hint shared between all machines in the cluster,
	// When not specified, soft-anti-affinity server group will be automatically created
	ServerGroupID string `json:"serverGroupID,omitempty"`

	// subnet c ID r
	SubnetCIDR string `json:"subnetCIDR,omitempty"`

	// subnet ID
	SubnetID string `json:"subnetID,omitempty"`

	// tenant
	Tenant string `json:"tenant,omitempty"`

	// tenant ID
	TenantID string `json:"tenantID,omitempty"`

	// Used internally during cluster creation
	Token string `json:"token,omitempty"`

	// Whether or not to use Octavia for LoadBalancer type of Service
	// implementation instead of using Neutron-LBaaS.
	// Attention:Openstack CCM use Octavia as default load balancer
	// implementation since v1.17.0
	//
	// Takes precedence over the 'use_octavia' flag provided at datacenter
	// level if both are specified.
	// +optional
	UseOctavia bool `json:"useOctavia"`

	// use token
	UseToken bool `json:"useToken"`

	// username
	Username string `json:"username,omitempty"`

	// credentials reference
	CredentialsReference GlobalSecretKeySelector `json:"credentialsReference,omitempty"`
}

// Validate validates this openstack cloud spec
func (m *OpenstackCloudSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialsReference(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenstackCloudSpec) validateCredentialsReference(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialsReference) { // not required
		return nil
	}

	if err := m.CredentialsReference.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("credentialsReference")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenstackCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenstackCloudSpec) UnmarshalBinary(b []byte) error {
	var res OpenstackCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
