// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewListNodeDeploymentNodesEventsParams creates a new ListNodeDeploymentNodesEventsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListNodeDeploymentNodesEventsParams() *ListNodeDeploymentNodesEventsParams {
	return &ListNodeDeploymentNodesEventsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListNodeDeploymentNodesEventsParamsWithTimeout creates a new ListNodeDeploymentNodesEventsParams object
// with the ability to set a timeout on a request.
func NewListNodeDeploymentNodesEventsParamsWithTimeout(timeout time.Duration) *ListNodeDeploymentNodesEventsParams {
	return &ListNodeDeploymentNodesEventsParams{
		timeout: timeout,
	}
}

// NewListNodeDeploymentNodesEventsParamsWithContext creates a new ListNodeDeploymentNodesEventsParams object
// with the ability to set a context for a request.
func NewListNodeDeploymentNodesEventsParamsWithContext(ctx context.Context) *ListNodeDeploymentNodesEventsParams {
	return &ListNodeDeploymentNodesEventsParams{
		Context: ctx,
	}
}

// NewListNodeDeploymentNodesEventsParamsWithHTTPClient creates a new ListNodeDeploymentNodesEventsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListNodeDeploymentNodesEventsParamsWithHTTPClient(client *http.Client) *ListNodeDeploymentNodesEventsParams {
	return &ListNodeDeploymentNodesEventsParams{
		HTTPClient: client,
	}
}

/*
ListNodeDeploymentNodesEventsParams contains all the parameters to send to the API endpoint

	for the list node deployment nodes events operation.

	Typically these are written to a http.Request.
*/
type ListNodeDeploymentNodesEventsParams struct {

	// ClusterID.
	ClusterID string

	// Dc.
	DC string

	// NodedeploymentID.
	NodeDeploymentID string

	// ProjectID.
	ProjectID string

	// Type.
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list node deployment nodes events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNodeDeploymentNodesEventsParams) WithDefaults() *ListNodeDeploymentNodesEventsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list node deployment nodes events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNodeDeploymentNodesEventsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list node deployment nodes events params
func (o *ListNodeDeploymentNodesEventsParams) WithTimeout(timeout time.Duration) *ListNodeDeploymentNodesEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list node deployment nodes events params
func (o *ListNodeDeploymentNodesEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list node deployment nodes events params
func (o *ListNodeDeploymentNodesEventsParams) WithContext(ctx context.Context) *ListNodeDeploymentNodesEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list node deployment nodes events params
func (o *ListNodeDeploymentNodesEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list node deployment nodes events params
func (o *ListNodeDeploymentNodesEventsParams) WithHTTPClient(client *http.Client) *ListNodeDeploymentNodesEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list node deployment nodes events params
func (o *ListNodeDeploymentNodesEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list node deployment nodes events params
func (o *ListNodeDeploymentNodesEventsParams) WithClusterID(clusterID string) *ListNodeDeploymentNodesEventsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list node deployment nodes events params
func (o *ListNodeDeploymentNodesEventsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the list node deployment nodes events params
func (o *ListNodeDeploymentNodesEventsParams) WithDC(dc string) *ListNodeDeploymentNodesEventsParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the list node deployment nodes events params
func (o *ListNodeDeploymentNodesEventsParams) SetDC(dc string) {
	o.DC = dc
}

// WithNodeDeploymentID adds the nodedeploymentID to the list node deployment nodes events params
func (o *ListNodeDeploymentNodesEventsParams) WithNodeDeploymentID(nodedeploymentID string) *ListNodeDeploymentNodesEventsParams {
	o.SetNodeDeploymentID(nodedeploymentID)
	return o
}

// SetNodeDeploymentID adds the nodedeploymentId to the list node deployment nodes events params
func (o *ListNodeDeploymentNodesEventsParams) SetNodeDeploymentID(nodedeploymentID string) {
	o.NodeDeploymentID = nodedeploymentID
}

// WithProjectID adds the projectID to the list node deployment nodes events params
func (o *ListNodeDeploymentNodesEventsParams) WithProjectID(projectID string) *ListNodeDeploymentNodesEventsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list node deployment nodes events params
func (o *ListNodeDeploymentNodesEventsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithType adds the typeVar to the list node deployment nodes events params
func (o *ListNodeDeploymentNodesEventsParams) WithType(typeVar *string) *ListNodeDeploymentNodesEventsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the list node deployment nodes events params
func (o *ListNodeDeploymentNodesEventsParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *ListNodeDeploymentNodesEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
		return err
	}

	// path param nodedeployment_id
	if err := r.SetPathParam("nodedeployment_id", o.NodeDeploymentID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
