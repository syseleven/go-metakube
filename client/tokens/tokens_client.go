// Code generated by go-swagger; DO NOT EDIT.

package tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new tokens API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new tokens API client with basic auth credentials.
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

// New creates a new tokens API client with a bearer token for authentication.
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
Client for tokens API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddTokenToServiceAccount(params *AddTokenToServiceAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddTokenToServiceAccountCreated, error)

	DeleteServiceAccountToken(params *DeleteServiceAccountTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteServiceAccountTokenOK, error)

	ListServiceAccountTokens(params *ListServiceAccountTokensParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListServiceAccountTokensOK, error)

	PatchServiceAccountToken(params *PatchServiceAccountTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchServiceAccountTokenOK, error)

	UpdateServiceAccountToken(params *UpdateServiceAccountTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateServiceAccountTokenOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddTokenToServiceAccount Generates a token for the given service account
*/
func (a *Client) AddTokenToServiceAccount(params *AddTokenToServiceAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddTokenToServiceAccountCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddTokenToServiceAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addTokenToServiceAccount",
		Method:             "POST",
		PathPattern:        "/api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddTokenToServiceAccountReader{formats: a.formats},
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
	success, ok := result.(*AddTokenToServiceAccountCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddTokenToServiceAccountDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteServiceAccountToken Deletes the token
*/
func (a *Client) DeleteServiceAccountToken(params *DeleteServiceAccountTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteServiceAccountTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServiceAccountTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteServiceAccountToken",
		Method:             "DELETE",
		PathPattern:        "/api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens/{token_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteServiceAccountTokenReader{formats: a.formats},
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
	success, ok := result.(*DeleteServiceAccountTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteServiceAccountTokenDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListServiceAccountTokens List tokens for the given service account
*/
func (a *Client) ListServiceAccountTokens(params *ListServiceAccountTokensParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListServiceAccountTokensOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListServiceAccountTokensParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listServiceAccountTokens",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListServiceAccountTokensReader{formats: a.formats},
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
	success, ok := result.(*ListServiceAccountTokensOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListServiceAccountTokensDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchServiceAccountToken Patches the token name
*/
func (a *Client) PatchServiceAccountToken(params *PatchServiceAccountTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchServiceAccountTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchServiceAccountTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchServiceAccountToken",
		Method:             "PATCH",
		PathPattern:        "/api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens/{token_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchServiceAccountTokenReader{formats: a.formats},
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
	success, ok := result.(*PatchServiceAccountTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchServiceAccountTokenDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateServiceAccountToken Updates and regenerates the token
*/
func (a *Client) UpdateServiceAccountToken(params *UpdateServiceAccountTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateServiceAccountTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateServiceAccountTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateServiceAccountToken",
		Method:             "PUT",
		PathPattern:        "/api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens/{token_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateServiceAccountTokenReader{formats: a.formats},
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
	success, ok := result.(*UpdateServiceAccountTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateServiceAccountTokenDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
