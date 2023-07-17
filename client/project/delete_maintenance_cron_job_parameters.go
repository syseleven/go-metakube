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

// NewDeleteMaintenanceCronJobParams creates a new DeleteMaintenanceCronJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMaintenanceCronJobParams() *DeleteMaintenanceCronJobParams {
	return &DeleteMaintenanceCronJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMaintenanceCronJobParamsWithTimeout creates a new DeleteMaintenanceCronJobParams object
// with the ability to set a timeout on a request.
func NewDeleteMaintenanceCronJobParamsWithTimeout(timeout time.Duration) *DeleteMaintenanceCronJobParams {
	return &DeleteMaintenanceCronJobParams{
		timeout: timeout,
	}
}

// NewDeleteMaintenanceCronJobParamsWithContext creates a new DeleteMaintenanceCronJobParams object
// with the ability to set a context for a request.
func NewDeleteMaintenanceCronJobParamsWithContext(ctx context.Context) *DeleteMaintenanceCronJobParams {
	return &DeleteMaintenanceCronJobParams{
		Context: ctx,
	}
}

// NewDeleteMaintenanceCronJobParamsWithHTTPClient creates a new DeleteMaintenanceCronJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMaintenanceCronJobParamsWithHTTPClient(client *http.Client) *DeleteMaintenanceCronJobParams {
	return &DeleteMaintenanceCronJobParams{
		HTTPClient: client,
	}
}

/*
DeleteMaintenanceCronJobParams contains all the parameters to send to the API endpoint

	for the delete maintenance cron job operation.

	Typically these are written to a http.Request.
*/
type DeleteMaintenanceCronJobParams struct {

	// ClusterID.
	ClusterID string

	// MaintenancecronjobID.
	MaintenanceCronJobID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete maintenance cron job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMaintenanceCronJobParams) WithDefaults() *DeleteMaintenanceCronJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete maintenance cron job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMaintenanceCronJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete maintenance cron job params
func (o *DeleteMaintenanceCronJobParams) WithTimeout(timeout time.Duration) *DeleteMaintenanceCronJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete maintenance cron job params
func (o *DeleteMaintenanceCronJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete maintenance cron job params
func (o *DeleteMaintenanceCronJobParams) WithContext(ctx context.Context) *DeleteMaintenanceCronJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete maintenance cron job params
func (o *DeleteMaintenanceCronJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete maintenance cron job params
func (o *DeleteMaintenanceCronJobParams) WithHTTPClient(client *http.Client) *DeleteMaintenanceCronJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete maintenance cron job params
func (o *DeleteMaintenanceCronJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the delete maintenance cron job params
func (o *DeleteMaintenanceCronJobParams) WithClusterID(clusterID string) *DeleteMaintenanceCronJobParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete maintenance cron job params
func (o *DeleteMaintenanceCronJobParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithMaintenanceCronJobID adds the maintenancecronjobID to the delete maintenance cron job params
func (o *DeleteMaintenanceCronJobParams) WithMaintenanceCronJobID(maintenancecronjobID string) *DeleteMaintenanceCronJobParams {
	o.SetMaintenanceCronJobID(maintenancecronjobID)
	return o
}

// SetMaintenanceCronJobID adds the maintenancecronjobId to the delete maintenance cron job params
func (o *DeleteMaintenanceCronJobParams) SetMaintenanceCronJobID(maintenancecronjobID string) {
	o.MaintenanceCronJobID = maintenancecronjobID
}

// WithProjectID adds the projectID to the delete maintenance cron job params
func (o *DeleteMaintenanceCronJobParams) WithProjectID(projectID string) *DeleteMaintenanceCronJobParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete maintenance cron job params
func (o *DeleteMaintenanceCronJobParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMaintenanceCronJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param maintenancecronjob_id
	if err := r.SetPathParam("maintenancecronjob_id", o.MaintenanceCronJobID); err != nil {
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