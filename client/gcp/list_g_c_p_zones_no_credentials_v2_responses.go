// Code generated by go-swagger; DO NOT EDIT.

package gcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/go-metakube/models"
)

// ListGCPZonesNoCredentialsV2Reader is a Reader for the ListGCPZonesNoCredentialsV2 structure.
type ListGCPZonesNoCredentialsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGCPZonesNoCredentialsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGCPZonesNoCredentialsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGCPZonesNoCredentialsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGCPZonesNoCredentialsV2OK creates a ListGCPZonesNoCredentialsV2OK with default headers values
func NewListGCPZonesNoCredentialsV2OK() *ListGCPZonesNoCredentialsV2OK {
	return &ListGCPZonesNoCredentialsV2OK{}
}

/*
ListGCPZonesNoCredentialsV2OK describes a response with status code 200, with default header values.

GCPZoneList
*/
type ListGCPZonesNoCredentialsV2OK struct {
	Payload models.GCPZoneList
}

// IsSuccess returns true when this list g c p zones no credentials v2 o k response has a 2xx status code
func (o *ListGCPZonesNoCredentialsV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list g c p zones no credentials v2 o k response has a 3xx status code
func (o *ListGCPZonesNoCredentialsV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list g c p zones no credentials v2 o k response has a 4xx status code
func (o *ListGCPZonesNoCredentialsV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list g c p zones no credentials v2 o k response has a 5xx status code
func (o *ListGCPZonesNoCredentialsV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this list g c p zones no credentials v2 o k response a status code equal to that given
func (o *ListGCPZonesNoCredentialsV2OK) IsCode(code int) bool {
	return code == 200
}

func (o *ListGCPZonesNoCredentialsV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/gcp/zones][%d] listGCPZonesNoCredentialsV2OK  %+v", 200, o.Payload)
}

func (o *ListGCPZonesNoCredentialsV2OK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/gcp/zones][%d] listGCPZonesNoCredentialsV2OK  %+v", 200, o.Payload)
}

func (o *ListGCPZonesNoCredentialsV2OK) GetPayload() models.GCPZoneList {
	return o.Payload
}

func (o *ListGCPZonesNoCredentialsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGCPZonesNoCredentialsV2Default creates a ListGCPZonesNoCredentialsV2Default with default headers values
func NewListGCPZonesNoCredentialsV2Default(code int) *ListGCPZonesNoCredentialsV2Default {
	return &ListGCPZonesNoCredentialsV2Default{
		_statusCode: code,
	}
}

/*
ListGCPZonesNoCredentialsV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type ListGCPZonesNoCredentialsV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list g c p zones no credentials v2 default response
func (o *ListGCPZonesNoCredentialsV2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list g c p zones no credentials v2 default response has a 2xx status code
func (o *ListGCPZonesNoCredentialsV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list g c p zones no credentials v2 default response has a 3xx status code
func (o *ListGCPZonesNoCredentialsV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list g c p zones no credentials v2 default response has a 4xx status code
func (o *ListGCPZonesNoCredentialsV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list g c p zones no credentials v2 default response has a 5xx status code
func (o *ListGCPZonesNoCredentialsV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list g c p zones no credentials v2 default response a status code equal to that given
func (o *ListGCPZonesNoCredentialsV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListGCPZonesNoCredentialsV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/gcp/zones][%d] listGCPZonesNoCredentialsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListGCPZonesNoCredentialsV2Default) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/gcp/zones][%d] listGCPZonesNoCredentialsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListGCPZonesNoCredentialsV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListGCPZonesNoCredentialsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
