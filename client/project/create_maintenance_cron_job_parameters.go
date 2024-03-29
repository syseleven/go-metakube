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

// NewCreateMaintenanceCronJobParams creates a new CreateMaintenanceCronJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateMaintenanceCronJobParams() *CreateMaintenanceCronJobParams {
	return &CreateMaintenanceCronJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMaintenanceCronJobParamsWithTimeout creates a new CreateMaintenanceCronJobParams object
// with the ability to set a timeout on a request.
func NewCreateMaintenanceCronJobParamsWithTimeout(timeout time.Duration) *CreateMaintenanceCronJobParams {
	return &CreateMaintenanceCronJobParams{
		timeout: timeout,
	}
}

// NewCreateMaintenanceCronJobParamsWithContext creates a new CreateMaintenanceCronJobParams object
// with the ability to set a context for a request.
func NewCreateMaintenanceCronJobParamsWithContext(ctx context.Context) *CreateMaintenanceCronJobParams {
	return &CreateMaintenanceCronJobParams{
		Context: ctx,
	}
}

// NewCreateMaintenanceCronJobParamsWithHTTPClient creates a new CreateMaintenanceCronJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateMaintenanceCronJobParamsWithHTTPClient(client *http.Client) *CreateMaintenanceCronJobParams {
	return &CreateMaintenanceCronJobParams{
		HTTPClient: client,
	}
}

/*
CreateMaintenanceCronJobParams contains all the parameters to send to the API endpoint

	for the create maintenance cron job operation.

	Typically these are written to a http.Request.
*/
type CreateMaintenanceCronJobParams struct {

	// Body.
	Body *models.MaintenanceCronJob

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create maintenance cron job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMaintenanceCronJobParams) WithDefaults() *CreateMaintenanceCronJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create maintenance cron job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMaintenanceCronJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create maintenance cron job params
func (o *CreateMaintenanceCronJobParams) WithTimeout(timeout time.Duration) *CreateMaintenanceCronJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create maintenance cron job params
func (o *CreateMaintenanceCronJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create maintenance cron job params
func (o *CreateMaintenanceCronJobParams) WithContext(ctx context.Context) *CreateMaintenanceCronJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create maintenance cron job params
func (o *CreateMaintenanceCronJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create maintenance cron job params
func (o *CreateMaintenanceCronJobParams) WithHTTPClient(client *http.Client) *CreateMaintenanceCronJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create maintenance cron job params
func (o *CreateMaintenanceCronJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create maintenance cron job params
func (o *CreateMaintenanceCronJobParams) WithBody(body *models.MaintenanceCronJob) *CreateMaintenanceCronJobParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create maintenance cron job params
func (o *CreateMaintenanceCronJobParams) SetBody(body *models.MaintenanceCronJob) {
	o.Body = body
}

// WithClusterID adds the clusterID to the create maintenance cron job params
func (o *CreateMaintenanceCronJobParams) WithClusterID(clusterID string) *CreateMaintenanceCronJobParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the create maintenance cron job params
func (o *CreateMaintenanceCronJobParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the create maintenance cron job params
func (o *CreateMaintenanceCronJobParams) WithProjectID(projectID string) *CreateMaintenanceCronJobParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create maintenance cron job params
func (o *CreateMaintenanceCronJobParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMaintenanceCronJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
