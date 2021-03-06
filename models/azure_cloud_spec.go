// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureCloudSpec AzureCloudSpec specifies access credentials to Azure cloud.
//
// swagger:model AzureCloudSpec
type AzureCloudSpec struct {

	// availability set
	AvailabilitySet string `json:"availabilitySet,omitempty"`

	// client ID
	ClientID string `json:"clientID,omitempty"`

	// client secret
	ClientSecret string `json:"clientSecret,omitempty"`

	// openstack billing tenant
	OpenstackBillingTenant string `json:"openstackBillingTenant,omitempty"`

	// resource group
	ResourceGroup string `json:"resourceGroup,omitempty"`

	// route table name
	RouteTableName string `json:"routeTable,omitempty"`

	// security group
	SecurityGroup string `json:"securityGroup,omitempty"`

	// subnet name
	SubnetName string `json:"subnet,omitempty"`

	// subscription ID
	SubscriptionID string `json:"subscriptionID,omitempty"`

	// tenant ID
	TenantID string `json:"tenantID,omitempty"`

	// v net name
	VNetName string `json:"vnet,omitempty"`

	// v net resource group
	VNetResourceGroup string `json:"vnetResourceGroup,omitempty"`

	// credentials reference
	CredentialsReference GlobalSecretKeySelector `json:"credentialsReference,omitempty"`

	// load balancer s k u
	LoadBalancerSKU LBSKU `json:"loadBalancerSKU,omitempty"`
}

// Validate validates this azure cloud spec
func (m *AzureCloudSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialsReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoadBalancerSKU(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureCloudSpec) validateCredentialsReference(formats strfmt.Registry) error {

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

func (m *AzureCloudSpec) validateLoadBalancerSKU(formats strfmt.Registry) error {

	if swag.IsZero(m.LoadBalancerSKU) { // not required
		return nil
	}

	if err := m.LoadBalancerSKU.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("loadBalancerSKU")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureCloudSpec) UnmarshalBinary(b []byte) error {
	var res AzureCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
