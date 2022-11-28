// Code generated by go-swagger; DO NOT EDIT.

package alibaba

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

// NewListAlibabaVSwitchesNoCredentialsV2Params creates a new ListAlibabaVSwitchesNoCredentialsV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAlibabaVSwitchesNoCredentialsV2Params() *ListAlibabaVSwitchesNoCredentialsV2Params {
	return &ListAlibabaVSwitchesNoCredentialsV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAlibabaVSwitchesNoCredentialsV2ParamsWithTimeout creates a new ListAlibabaVSwitchesNoCredentialsV2Params object
// with the ability to set a timeout on a request.
func NewListAlibabaVSwitchesNoCredentialsV2ParamsWithTimeout(timeout time.Duration) *ListAlibabaVSwitchesNoCredentialsV2Params {
	return &ListAlibabaVSwitchesNoCredentialsV2Params{
		timeout: timeout,
	}
}

// NewListAlibabaVSwitchesNoCredentialsV2ParamsWithContext creates a new ListAlibabaVSwitchesNoCredentialsV2Params object
// with the ability to set a context for a request.
func NewListAlibabaVSwitchesNoCredentialsV2ParamsWithContext(ctx context.Context) *ListAlibabaVSwitchesNoCredentialsV2Params {
	return &ListAlibabaVSwitchesNoCredentialsV2Params{
		Context: ctx,
	}
}

// NewListAlibabaVSwitchesNoCredentialsV2ParamsWithHTTPClient creates a new ListAlibabaVSwitchesNoCredentialsV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewListAlibabaVSwitchesNoCredentialsV2ParamsWithHTTPClient(client *http.Client) *ListAlibabaVSwitchesNoCredentialsV2Params {
	return &ListAlibabaVSwitchesNoCredentialsV2Params{
		HTTPClient: client,
	}
}

/*
ListAlibabaVSwitchesNoCredentialsV2Params contains all the parameters to send to the API endpoint

	for the list alibaba v switches no credentials v2 operation.

	Typically these are written to a http.Request.
*/
type ListAlibabaVSwitchesNoCredentialsV2Params struct {

	// Region.
	Region *string

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list alibaba v switches no credentials v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAlibabaVSwitchesNoCredentialsV2Params) WithDefaults() *ListAlibabaVSwitchesNoCredentialsV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list alibaba v switches no credentials v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAlibabaVSwitchesNoCredentialsV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list alibaba v switches no credentials v2 params
func (o *ListAlibabaVSwitchesNoCredentialsV2Params) WithTimeout(timeout time.Duration) *ListAlibabaVSwitchesNoCredentialsV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list alibaba v switches no credentials v2 params
func (o *ListAlibabaVSwitchesNoCredentialsV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list alibaba v switches no credentials v2 params
func (o *ListAlibabaVSwitchesNoCredentialsV2Params) WithContext(ctx context.Context) *ListAlibabaVSwitchesNoCredentialsV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list alibaba v switches no credentials v2 params
func (o *ListAlibabaVSwitchesNoCredentialsV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list alibaba v switches no credentials v2 params
func (o *ListAlibabaVSwitchesNoCredentialsV2Params) WithHTTPClient(client *http.Client) *ListAlibabaVSwitchesNoCredentialsV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list alibaba v switches no credentials v2 params
func (o *ListAlibabaVSwitchesNoCredentialsV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegion adds the region to the list alibaba v switches no credentials v2 params
func (o *ListAlibabaVSwitchesNoCredentialsV2Params) WithRegion(region *string) *ListAlibabaVSwitchesNoCredentialsV2Params {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the list alibaba v switches no credentials v2 params
func (o *ListAlibabaVSwitchesNoCredentialsV2Params) SetRegion(region *string) {
	o.Region = region
}

// WithClusterID adds the clusterID to the list alibaba v switches no credentials v2 params
func (o *ListAlibabaVSwitchesNoCredentialsV2Params) WithClusterID(clusterID string) *ListAlibabaVSwitchesNoCredentialsV2Params {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list alibaba v switches no credentials v2 params
func (o *ListAlibabaVSwitchesNoCredentialsV2Params) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the list alibaba v switches no credentials v2 params
func (o *ListAlibabaVSwitchesNoCredentialsV2Params) WithProjectID(projectID string) *ListAlibabaVSwitchesNoCredentialsV2Params {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list alibaba v switches no credentials v2 params
func (o *ListAlibabaVSwitchesNoCredentialsV2Params) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListAlibabaVSwitchesNoCredentialsV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Region != nil {

		// header param Region
		if err := r.SetHeaderParam("Region", *o.Region); err != nil {
			return err
		}
	}

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
