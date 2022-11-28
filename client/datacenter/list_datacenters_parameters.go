// Code generated by go-swagger; DO NOT EDIT.

package datacenter

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

// NewListDatacentersParams creates a new ListDatacentersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListDatacentersParams() *ListDatacentersParams {
	return &ListDatacentersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListDatacentersParamsWithTimeout creates a new ListDatacentersParams object
// with the ability to set a timeout on a request.
func NewListDatacentersParamsWithTimeout(timeout time.Duration) *ListDatacentersParams {
	return &ListDatacentersParams{
		timeout: timeout,
	}
}

// NewListDatacentersParamsWithContext creates a new ListDatacentersParams object
// with the ability to set a context for a request.
func NewListDatacentersParamsWithContext(ctx context.Context) *ListDatacentersParams {
	return &ListDatacentersParams{
		Context: ctx,
	}
}

// NewListDatacentersParamsWithHTTPClient creates a new ListDatacentersParams object
// with the ability to set a custom HTTPClient for a request.
func NewListDatacentersParamsWithHTTPClient(client *http.Client) *ListDatacentersParams {
	return &ListDatacentersParams{
		HTTPClient: client,
	}
}

/*
ListDatacentersParams contains all the parameters to send to the API endpoint

	for the list datacenters operation.

	Typically these are written to a http.Request.
*/
type ListDatacentersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list datacenters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDatacentersParams) WithDefaults() *ListDatacentersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list datacenters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDatacentersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list datacenters params
func (o *ListDatacentersParams) WithTimeout(timeout time.Duration) *ListDatacentersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list datacenters params
func (o *ListDatacentersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list datacenters params
func (o *ListDatacentersParams) WithContext(ctx context.Context) *ListDatacentersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list datacenters params
func (o *ListDatacentersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list datacenters params
func (o *ListDatacentersParams) WithHTTPClient(client *http.Client) *ListDatacentersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list datacenters params
func (o *ListDatacentersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListDatacentersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
