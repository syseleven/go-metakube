// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListAzureSizesNoCredentialsV2Params creates a new ListAzureSizesNoCredentialsV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAzureSizesNoCredentialsV2Params() *ListAzureSizesNoCredentialsV2Params {
	return &ListAzureSizesNoCredentialsV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAzureSizesNoCredentialsV2ParamsWithTimeout creates a new ListAzureSizesNoCredentialsV2Params object
// with the ability to set a timeout on a request.
func NewListAzureSizesNoCredentialsV2ParamsWithTimeout(timeout time.Duration) *ListAzureSizesNoCredentialsV2Params {
	return &ListAzureSizesNoCredentialsV2Params{
		timeout: timeout,
	}
}

// NewListAzureSizesNoCredentialsV2ParamsWithContext creates a new ListAzureSizesNoCredentialsV2Params object
// with the ability to set a context for a request.
func NewListAzureSizesNoCredentialsV2ParamsWithContext(ctx context.Context) *ListAzureSizesNoCredentialsV2Params {
	return &ListAzureSizesNoCredentialsV2Params{
		Context: ctx,
	}
}

// NewListAzureSizesNoCredentialsV2ParamsWithHTTPClient creates a new ListAzureSizesNoCredentialsV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewListAzureSizesNoCredentialsV2ParamsWithHTTPClient(client *http.Client) *ListAzureSizesNoCredentialsV2Params {
	return &ListAzureSizesNoCredentialsV2Params{
		HTTPClient: client,
	}
}

/*
ListAzureSizesNoCredentialsV2Params contains all the parameters to send to the API endpoint

	for the list azure sizes no credentials v2 operation.

	Typically these are written to a http.Request.
*/
type ListAzureSizesNoCredentialsV2Params struct {

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list azure sizes no credentials v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAzureSizesNoCredentialsV2Params) WithDefaults() *ListAzureSizesNoCredentialsV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list azure sizes no credentials v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAzureSizesNoCredentialsV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list azure sizes no credentials v2 params
func (o *ListAzureSizesNoCredentialsV2Params) WithTimeout(timeout time.Duration) *ListAzureSizesNoCredentialsV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list azure sizes no credentials v2 params
func (o *ListAzureSizesNoCredentialsV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list azure sizes no credentials v2 params
func (o *ListAzureSizesNoCredentialsV2Params) WithContext(ctx context.Context) *ListAzureSizesNoCredentialsV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list azure sizes no credentials v2 params
func (o *ListAzureSizesNoCredentialsV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list azure sizes no credentials v2 params
func (o *ListAzureSizesNoCredentialsV2Params) WithHTTPClient(client *http.Client) *ListAzureSizesNoCredentialsV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list azure sizes no credentials v2 params
func (o *ListAzureSizesNoCredentialsV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list azure sizes no credentials v2 params
func (o *ListAzureSizesNoCredentialsV2Params) WithClusterID(clusterID string) *ListAzureSizesNoCredentialsV2Params {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list azure sizes no credentials v2 params
func (o *ListAzureSizesNoCredentialsV2Params) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the list azure sizes no credentials v2 params
func (o *ListAzureSizesNoCredentialsV2Params) WithProjectID(projectID string) *ListAzureSizesNoCredentialsV2Params {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list azure sizes no credentials v2 params
func (o *ListAzureSizesNoCredentialsV2Params) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListAzureSizesNoCredentialsV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
