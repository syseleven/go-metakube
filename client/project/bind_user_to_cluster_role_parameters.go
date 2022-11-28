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

	"github.com/syseleven/go-metakube/models"
)

// NewBindUserToClusterRoleParams creates a new BindUserToClusterRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBindUserToClusterRoleParams() *BindUserToClusterRoleParams {
	return &BindUserToClusterRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBindUserToClusterRoleParamsWithTimeout creates a new BindUserToClusterRoleParams object
// with the ability to set a timeout on a request.
func NewBindUserToClusterRoleParamsWithTimeout(timeout time.Duration) *BindUserToClusterRoleParams {
	return &BindUserToClusterRoleParams{
		timeout: timeout,
	}
}

// NewBindUserToClusterRoleParamsWithContext creates a new BindUserToClusterRoleParams object
// with the ability to set a context for a request.
func NewBindUserToClusterRoleParamsWithContext(ctx context.Context) *BindUserToClusterRoleParams {
	return &BindUserToClusterRoleParams{
		Context: ctx,
	}
}

// NewBindUserToClusterRoleParamsWithHTTPClient creates a new BindUserToClusterRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewBindUserToClusterRoleParamsWithHTTPClient(client *http.Client) *BindUserToClusterRoleParams {
	return &BindUserToClusterRoleParams{
		HTTPClient: client,
	}
}

/*
BindUserToClusterRoleParams contains all the parameters to send to the API endpoint

	for the bind user to cluster role operation.

	Typically these are written to a http.Request.
*/
type BindUserToClusterRoleParams struct {

	// Body.
	Body *models.ClusterRoleUser

	// ClusterID.
	ClusterID string

	// Dc.
	DC string

	// ProjectID.
	ProjectID string

	// RoleID.
	RoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bind user to cluster role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BindUserToClusterRoleParams) WithDefaults() *BindUserToClusterRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bind user to cluster role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BindUserToClusterRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bind user to cluster role params
func (o *BindUserToClusterRoleParams) WithTimeout(timeout time.Duration) *BindUserToClusterRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bind user to cluster role params
func (o *BindUserToClusterRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bind user to cluster role params
func (o *BindUserToClusterRoleParams) WithContext(ctx context.Context) *BindUserToClusterRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bind user to cluster role params
func (o *BindUserToClusterRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bind user to cluster role params
func (o *BindUserToClusterRoleParams) WithHTTPClient(client *http.Client) *BindUserToClusterRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bind user to cluster role params
func (o *BindUserToClusterRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the bind user to cluster role params
func (o *BindUserToClusterRoleParams) WithBody(body *models.ClusterRoleUser) *BindUserToClusterRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the bind user to cluster role params
func (o *BindUserToClusterRoleParams) SetBody(body *models.ClusterRoleUser) {
	o.Body = body
}

// WithClusterID adds the clusterID to the bind user to cluster role params
func (o *BindUserToClusterRoleParams) WithClusterID(clusterID string) *BindUserToClusterRoleParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the bind user to cluster role params
func (o *BindUserToClusterRoleParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the bind user to cluster role params
func (o *BindUserToClusterRoleParams) WithDC(dc string) *BindUserToClusterRoleParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the bind user to cluster role params
func (o *BindUserToClusterRoleParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the bind user to cluster role params
func (o *BindUserToClusterRoleParams) WithProjectID(projectID string) *BindUserToClusterRoleParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the bind user to cluster role params
func (o *BindUserToClusterRoleParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithRoleID adds the roleID to the bind user to cluster role params
func (o *BindUserToClusterRoleParams) WithRoleID(roleID string) *BindUserToClusterRoleParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the bind user to cluster role params
func (o *BindUserToClusterRoleParams) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *BindUserToClusterRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	// path param role_id
	if err := r.SetPathParam("role_id", o.RoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
