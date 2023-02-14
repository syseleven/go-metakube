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

// NewListDatacentersV2Params creates a new ListDatacentersV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListDatacentersV2Params() *ListDatacentersV2Params {
	return &ListDatacentersV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewListDatacentersV2ParamsWithTimeout creates a new ListDatacentersV2Params object
// with the ability to set a timeout on a request.
func NewListDatacentersV2ParamsWithTimeout(timeout time.Duration) *ListDatacentersV2Params {
	return &ListDatacentersV2Params{
		timeout: timeout,
	}
}

// NewListDatacentersV2ParamsWithContext creates a new ListDatacentersV2Params object
// with the ability to set a context for a request.
func NewListDatacentersV2ParamsWithContext(ctx context.Context) *ListDatacentersV2Params {
	return &ListDatacentersV2Params{
		Context: ctx,
	}
}

// NewListDatacentersV2ParamsWithHTTPClient creates a new ListDatacentersV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewListDatacentersV2ParamsWithHTTPClient(client *http.Client) *ListDatacentersV2Params {
	return &ListDatacentersV2Params{
		HTTPClient: client,
	}
}

/*
ListDatacentersV2Params contains all the parameters to send to the API endpoint

	for the list datacenters v2 operation.

	Typically these are written to a http.Request.
*/
type ListDatacentersV2Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list datacenters v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDatacentersV2Params) WithDefaults() *ListDatacentersV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list datacenters v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDatacentersV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list datacenters v2 params
func (o *ListDatacentersV2Params) WithTimeout(timeout time.Duration) *ListDatacentersV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list datacenters v2 params
func (o *ListDatacentersV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list datacenters v2 params
func (o *ListDatacentersV2Params) WithContext(ctx context.Context) *ListDatacentersV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list datacenters v2 params
func (o *ListDatacentersV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list datacenters v2 params
func (o *ListDatacentersV2Params) WithHTTPClient(client *http.Client) *ListDatacentersV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list datacenters v2 params
func (o *ListDatacentersV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListDatacentersV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}