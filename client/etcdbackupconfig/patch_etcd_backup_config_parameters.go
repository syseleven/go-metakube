// Code generated by go-swagger; DO NOT EDIT.

package etcdbackupconfig

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

// NewPatchEtcdBackupConfigParams creates a new PatchEtcdBackupConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchEtcdBackupConfigParams() *PatchEtcdBackupConfigParams {
	return &PatchEtcdBackupConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchEtcdBackupConfigParamsWithTimeout creates a new PatchEtcdBackupConfigParams object
// with the ability to set a timeout on a request.
func NewPatchEtcdBackupConfigParamsWithTimeout(timeout time.Duration) *PatchEtcdBackupConfigParams {
	return &PatchEtcdBackupConfigParams{
		timeout: timeout,
	}
}

// NewPatchEtcdBackupConfigParamsWithContext creates a new PatchEtcdBackupConfigParams object
// with the ability to set a context for a request.
func NewPatchEtcdBackupConfigParamsWithContext(ctx context.Context) *PatchEtcdBackupConfigParams {
	return &PatchEtcdBackupConfigParams{
		Context: ctx,
	}
}

// NewPatchEtcdBackupConfigParamsWithHTTPClient creates a new PatchEtcdBackupConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchEtcdBackupConfigParamsWithHTTPClient(client *http.Client) *PatchEtcdBackupConfigParams {
	return &PatchEtcdBackupConfigParams{
		HTTPClient: client,
	}
}

/*
PatchEtcdBackupConfigParams contains all the parameters to send to the API endpoint

	for the patch etcd backup config operation.

	Typically these are written to a http.Request.
*/
type PatchEtcdBackupConfigParams struct {

	// Body.
	Body *models.EtcdBackupConfigSpec

	// ClusterID.
	ClusterID string

	// EbcID.
	EtcdBackupConfigID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch etcd backup config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchEtcdBackupConfigParams) WithDefaults() *PatchEtcdBackupConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch etcd backup config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchEtcdBackupConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch etcd backup config params
func (o *PatchEtcdBackupConfigParams) WithTimeout(timeout time.Duration) *PatchEtcdBackupConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch etcd backup config params
func (o *PatchEtcdBackupConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch etcd backup config params
func (o *PatchEtcdBackupConfigParams) WithContext(ctx context.Context) *PatchEtcdBackupConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch etcd backup config params
func (o *PatchEtcdBackupConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch etcd backup config params
func (o *PatchEtcdBackupConfigParams) WithHTTPClient(client *http.Client) *PatchEtcdBackupConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch etcd backup config params
func (o *PatchEtcdBackupConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch etcd backup config params
func (o *PatchEtcdBackupConfigParams) WithBody(body *models.EtcdBackupConfigSpec) *PatchEtcdBackupConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch etcd backup config params
func (o *PatchEtcdBackupConfigParams) SetBody(body *models.EtcdBackupConfigSpec) {
	o.Body = body
}

// WithClusterID adds the clusterID to the patch etcd backup config params
func (o *PatchEtcdBackupConfigParams) WithClusterID(clusterID string) *PatchEtcdBackupConfigParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the patch etcd backup config params
func (o *PatchEtcdBackupConfigParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithEtcdBackupConfigID adds the ebcID to the patch etcd backup config params
func (o *PatchEtcdBackupConfigParams) WithEtcdBackupConfigID(ebcID string) *PatchEtcdBackupConfigParams {
	o.SetEtcdBackupConfigID(ebcID)
	return o
}

// SetEtcdBackupConfigID adds the ebcId to the patch etcd backup config params
func (o *PatchEtcdBackupConfigParams) SetEtcdBackupConfigID(ebcID string) {
	o.EtcdBackupConfigID = ebcID
}

// WithProjectID adds the projectID to the patch etcd backup config params
func (o *PatchEtcdBackupConfigParams) WithProjectID(projectID string) *PatchEtcdBackupConfigParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the patch etcd backup config params
func (o *PatchEtcdBackupConfigParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchEtcdBackupConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param ebc_id
	if err := r.SetPathParam("ebc_id", o.EtcdBackupConfigID); err != nil {
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
