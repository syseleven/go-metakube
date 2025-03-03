// Code generated by go-swagger; DO NOT EDIT.

package openstack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new openstack API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new openstack API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new openstack API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for openstack API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ListOpenstackAvailabilityZones(params *ListOpenstackAvailabilityZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackAvailabilityZonesOK, error)

	ListOpenstackAvailabilityZonesNoCredentialsV2(params *ListOpenstackAvailabilityZonesNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackAvailabilityZonesNoCredentialsV2OK, error)

	ListOpenstackImages(params *ListOpenstackImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackImagesOK, error)

	ListOpenstackImagesNoCredentialsV2(params *ListOpenstackImagesNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackImagesNoCredentialsV2OK, error)

	ListOpenstackNetworks(params *ListOpenstackNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackNetworksOK, error)

	ListOpenstackNetworksNoCredentialsV2(params *ListOpenstackNetworksNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackNetworksNoCredentialsV2OK, error)

	ListOpenstackQuotaLimits(params *ListOpenstackQuotaLimitsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackQuotaLimitsOK, error)

	ListOpenstackQuotaLimitsNoCredentialsV2(params *ListOpenstackQuotaLimitsNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackQuotaLimitsNoCredentialsV2OK, error)

	ListOpenstackSecurityGroups(params *ListOpenstackSecurityGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackSecurityGroupsOK, error)

	ListOpenstackSecurityGroupsNoCredentialsV2(params *ListOpenstackSecurityGroupsNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackSecurityGroupsNoCredentialsV2OK, error)

	ListOpenstackSizes(params *ListOpenstackSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackSizesOK, error)

	ListOpenstackSizesNoCredentialsV2(params *ListOpenstackSizesNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackSizesNoCredentialsV2OK, error)

	ListOpenstackSubnets(params *ListOpenstackSubnetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackSubnetsOK, error)

	ListOpenstackSubnetsNoCredentialsV2(params *ListOpenstackSubnetsNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackSubnetsNoCredentialsV2OK, error)

	ListOpenstackTenants(params *ListOpenstackTenantsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackTenantsOK, error)

	ListOpenstackTenantsNoCredentialsV2(params *ListOpenstackTenantsNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackTenantsNoCredentialsV2OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ListOpenstackAvailabilityZones Lists availability zones from openstack
*/
func (a *Client) ListOpenstackAvailabilityZones(params *ListOpenstackAvailabilityZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackAvailabilityZonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOpenstackAvailabilityZonesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOpenstackAvailabilityZones",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/openstack/availabilityzones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOpenstackAvailabilityZonesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOpenstackAvailabilityZonesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOpenstackAvailabilityZonesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListOpenstackAvailabilityZonesNoCredentialsV2 Lists availability zones from openstack
*/
func (a *Client) ListOpenstackAvailabilityZonesNoCredentialsV2(params *ListOpenstackAvailabilityZonesNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackAvailabilityZonesNoCredentialsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOpenstackAvailabilityZonesNoCredentialsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOpenstackAvailabilityZonesNoCredentialsV2",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/availabilityzones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOpenstackAvailabilityZonesNoCredentialsV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOpenstackAvailabilityZonesNoCredentialsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOpenstackAvailabilityZonesNoCredentialsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListOpenstackImages Lists images from openstack
*/
func (a *Client) ListOpenstackImages(params *ListOpenstackImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOpenstackImagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOpenstackImages",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/openstack/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOpenstackImagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOpenstackImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOpenstackImagesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListOpenstackImagesNoCredentialsV2 Lists images from openstack
*/
func (a *Client) ListOpenstackImagesNoCredentialsV2(params *ListOpenstackImagesNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackImagesNoCredentialsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOpenstackImagesNoCredentialsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOpenstackImagesNoCredentialsV2",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOpenstackImagesNoCredentialsV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOpenstackImagesNoCredentialsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOpenstackImagesNoCredentialsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListOpenstackNetworks Lists networks from openstack
*/
func (a *Client) ListOpenstackNetworks(params *ListOpenstackNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackNetworksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOpenstackNetworksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOpenstackNetworks",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/openstack/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOpenstackNetworksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOpenstackNetworksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOpenstackNetworksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListOpenstackNetworksNoCredentialsV2 Lists networks from openstack
*/
func (a *Client) ListOpenstackNetworksNoCredentialsV2(params *ListOpenstackNetworksNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackNetworksNoCredentialsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOpenstackNetworksNoCredentialsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOpenstackNetworksNoCredentialsV2",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOpenstackNetworksNoCredentialsV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOpenstackNetworksNoCredentialsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOpenstackNetworksNoCredentialsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListOpenstackQuotaLimits Lists quotalimits for tenant from openstack
*/
func (a *Client) ListOpenstackQuotaLimits(params *ListOpenstackQuotaLimitsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackQuotaLimitsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOpenstackQuotaLimitsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOpenstackQuotaLimits",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/openstack/quotalimits",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOpenstackQuotaLimitsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOpenstackQuotaLimitsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOpenstackQuotaLimitsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListOpenstackQuotaLimitsNoCredentialsV2 Lists quotalimits for tenant from openstack
*/
func (a *Client) ListOpenstackQuotaLimitsNoCredentialsV2(params *ListOpenstackQuotaLimitsNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackQuotaLimitsNoCredentialsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOpenstackQuotaLimitsNoCredentialsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOpenstackQuotaLimitsNoCredentialsV2",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/quotalimits",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOpenstackQuotaLimitsNoCredentialsV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOpenstackQuotaLimitsNoCredentialsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOpenstackQuotaLimitsNoCredentialsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListOpenstackSecurityGroups Lists security groups from openstack
*/
func (a *Client) ListOpenstackSecurityGroups(params *ListOpenstackSecurityGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackSecurityGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOpenstackSecurityGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOpenstackSecurityGroups",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/openstack/securitygroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOpenstackSecurityGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOpenstackSecurityGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOpenstackSecurityGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListOpenstackSecurityGroupsNoCredentialsV2 Lists security groups from openstack
*/
func (a *Client) ListOpenstackSecurityGroupsNoCredentialsV2(params *ListOpenstackSecurityGroupsNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackSecurityGroupsNoCredentialsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOpenstackSecurityGroupsNoCredentialsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOpenstackSecurityGroupsNoCredentialsV2",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/securitygroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOpenstackSecurityGroupsNoCredentialsV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOpenstackSecurityGroupsNoCredentialsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOpenstackSecurityGroupsNoCredentialsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListOpenstackSizes Lists sizes from openstack
*/
func (a *Client) ListOpenstackSizes(params *ListOpenstackSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackSizesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOpenstackSizesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOpenstackSizes",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/openstack/sizes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOpenstackSizesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOpenstackSizesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOpenstackSizesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListOpenstackSizesNoCredentialsV2 Lists sizes from openstack
*/
func (a *Client) ListOpenstackSizesNoCredentialsV2(params *ListOpenstackSizesNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackSizesNoCredentialsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOpenstackSizesNoCredentialsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOpenstackSizesNoCredentialsV2",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/sizes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOpenstackSizesNoCredentialsV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOpenstackSizesNoCredentialsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOpenstackSizesNoCredentialsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListOpenstackSubnets Lists subnets from openstack
*/
func (a *Client) ListOpenstackSubnets(params *ListOpenstackSubnetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackSubnetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOpenstackSubnetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOpenstackSubnets",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/openstack/subnets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOpenstackSubnetsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOpenstackSubnetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOpenstackSubnetsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListOpenstackSubnetsNoCredentialsV2 Lists subnets from openstack
*/
func (a *Client) ListOpenstackSubnetsNoCredentialsV2(params *ListOpenstackSubnetsNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackSubnetsNoCredentialsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOpenstackSubnetsNoCredentialsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOpenstackSubnetsNoCredentialsV2",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/subnets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOpenstackSubnetsNoCredentialsV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOpenstackSubnetsNoCredentialsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOpenstackSubnetsNoCredentialsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListOpenstackTenants Lists tenants from openstack
*/
func (a *Client) ListOpenstackTenants(params *ListOpenstackTenantsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackTenantsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOpenstackTenantsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOpenstackTenants",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/openstack/tenants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOpenstackTenantsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOpenstackTenantsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOpenstackTenantsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListOpenstackTenantsNoCredentialsV2 Lists tenants from openstack
*/
func (a *Client) ListOpenstackTenantsNoCredentialsV2(params *ListOpenstackTenantsNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOpenstackTenantsNoCredentialsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOpenstackTenantsNoCredentialsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOpenstackTenantsNoCredentialsV2",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/tenants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOpenstackTenantsNoCredentialsV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOpenstackTenantsNoCredentialsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOpenstackTenantsNoCredentialsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
