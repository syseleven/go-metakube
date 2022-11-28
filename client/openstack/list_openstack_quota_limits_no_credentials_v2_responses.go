// Code generated by go-swagger; DO NOT EDIT.

package openstack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/go-metakube/models"
)

// ListOpenstackQuotaLimitsNoCredentialsV2Reader is a Reader for the ListOpenstackQuotaLimitsNoCredentialsV2 structure.
type ListOpenstackQuotaLimitsNoCredentialsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOpenstackQuotaLimitsNoCredentialsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOpenstackQuotaLimitsNoCredentialsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListOpenstackQuotaLimitsNoCredentialsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOpenstackQuotaLimitsNoCredentialsV2OK creates a ListOpenstackQuotaLimitsNoCredentialsV2OK with default headers values
func NewListOpenstackQuotaLimitsNoCredentialsV2OK() *ListOpenstackQuotaLimitsNoCredentialsV2OK {
	return &ListOpenstackQuotaLimitsNoCredentialsV2OK{}
}

/*
ListOpenstackQuotaLimitsNoCredentialsV2OK describes a response with status code 200, with default header values.

Quotas
*/
type ListOpenstackQuotaLimitsNoCredentialsV2OK struct {
	Payload *models.Quotas
}

// IsSuccess returns true when this list openstack quota limits no credentials v2 o k response has a 2xx status code
func (o *ListOpenstackQuotaLimitsNoCredentialsV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list openstack quota limits no credentials v2 o k response has a 3xx status code
func (o *ListOpenstackQuotaLimitsNoCredentialsV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list openstack quota limits no credentials v2 o k response has a 4xx status code
func (o *ListOpenstackQuotaLimitsNoCredentialsV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list openstack quota limits no credentials v2 o k response has a 5xx status code
func (o *ListOpenstackQuotaLimitsNoCredentialsV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this list openstack quota limits no credentials v2 o k response a status code equal to that given
func (o *ListOpenstackQuotaLimitsNoCredentialsV2OK) IsCode(code int) bool {
	return code == 200
}

func (o *ListOpenstackQuotaLimitsNoCredentialsV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/clusters/{cluster_id}/providers/openstack/quotalimits][%d] listOpenstackQuotaLimitsNoCredentialsV2OK  %+v", 200, o.Payload)
}

func (o *ListOpenstackQuotaLimitsNoCredentialsV2OK) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/clusters/{cluster_id}/providers/openstack/quotalimits][%d] listOpenstackQuotaLimitsNoCredentialsV2OK  %+v", 200, o.Payload)
}

func (o *ListOpenstackQuotaLimitsNoCredentialsV2OK) GetPayload() *models.Quotas {
	return o.Payload
}

func (o *ListOpenstackQuotaLimitsNoCredentialsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Quotas)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOpenstackQuotaLimitsNoCredentialsV2Default creates a ListOpenstackQuotaLimitsNoCredentialsV2Default with default headers values
func NewListOpenstackQuotaLimitsNoCredentialsV2Default(code int) *ListOpenstackQuotaLimitsNoCredentialsV2Default {
	return &ListOpenstackQuotaLimitsNoCredentialsV2Default{
		_statusCode: code,
	}
}

/*
ListOpenstackQuotaLimitsNoCredentialsV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type ListOpenstackQuotaLimitsNoCredentialsV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list openstack quota limits no credentials v2 default response
func (o *ListOpenstackQuotaLimitsNoCredentialsV2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list openstack quota limits no credentials v2 default response has a 2xx status code
func (o *ListOpenstackQuotaLimitsNoCredentialsV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list openstack quota limits no credentials v2 default response has a 3xx status code
func (o *ListOpenstackQuotaLimitsNoCredentialsV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list openstack quota limits no credentials v2 default response has a 4xx status code
func (o *ListOpenstackQuotaLimitsNoCredentialsV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list openstack quota limits no credentials v2 default response has a 5xx status code
func (o *ListOpenstackQuotaLimitsNoCredentialsV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list openstack quota limits no credentials v2 default response a status code equal to that given
func (o *ListOpenstackQuotaLimitsNoCredentialsV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListOpenstackQuotaLimitsNoCredentialsV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/clusters/{cluster_id}/providers/openstack/quotalimits][%d] listOpenstackQuotaLimitsNoCredentialsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListOpenstackQuotaLimitsNoCredentialsV2Default) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/clusters/{cluster_id}/providers/openstack/quotalimits][%d] listOpenstackQuotaLimitsNoCredentialsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListOpenstackQuotaLimitsNoCredentialsV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListOpenstackQuotaLimitsNoCredentialsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
