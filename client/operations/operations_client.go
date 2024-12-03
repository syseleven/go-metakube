// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new operations API client with basic auth credentials.
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

// New creates a new operations API client with a bearer token for authentication.
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
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateOIDCKubeconfig(params *CreateOIDCKubeconfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOIDCKubeconfigOK, error)

	ListSystemLabels(params *ListSystemLabelsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSystemLabelsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	CreateOIDCKubeconfig Starts OIDC flow and generates kubeconfig, the generated config

contains OIDC provider authentication info
*/
func (a *Client) CreateOIDCKubeconfig(params *CreateOIDCKubeconfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOIDCKubeconfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOIDCKubeconfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createOIDCKubeconfig",
		Method:             "GET",
		PathPattern:        "/api/v1/kubeconfig",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateOIDCKubeconfigReader{formats: a.formats},
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
	success, ok := result.(*CreateOIDCKubeconfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateOIDCKubeconfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListSystemLabels List restricted system labels
*/
func (a *Client) ListSystemLabels(params *ListSystemLabelsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSystemLabelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSystemLabelsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listSystemLabels",
		Method:             "GET",
		PathPattern:        "/api/v1/labels/system",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSystemLabelsReader{formats: a.formats},
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
	success, ok := result.(*ListSystemLabelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListSystemLabelsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
