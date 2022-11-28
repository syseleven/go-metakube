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

// ListGCPZonesNoCredentialsReader is a Reader for the ListGCPZonesNoCredentials structure.
type ListGCPZonesNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGCPZonesNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGCPZonesNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGCPZonesNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGCPZonesNoCredentialsOK creates a ListGCPZonesNoCredentialsOK with default headers values
func NewListGCPZonesNoCredentialsOK() *ListGCPZonesNoCredentialsOK {
	return &ListGCPZonesNoCredentialsOK{}
}

/*
ListGCPZonesNoCredentialsOK describes a response with status code 200, with default header values.

GCPZoneList
*/
type ListGCPZonesNoCredentialsOK struct {
	Payload models.GCPZoneList
}

// IsSuccess returns true when this list g c p zones no credentials o k response has a 2xx status code
func (o *ListGCPZonesNoCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list g c p zones no credentials o k response has a 3xx status code
func (o *ListGCPZonesNoCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list g c p zones no credentials o k response has a 4xx status code
func (o *ListGCPZonesNoCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list g c p zones no credentials o k response has a 5xx status code
func (o *ListGCPZonesNoCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list g c p zones no credentials o k response a status code equal to that given
func (o *ListGCPZonesNoCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListGCPZonesNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/zones][%d] listGCPZonesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListGCPZonesNoCredentialsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/zones][%d] listGCPZonesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListGCPZonesNoCredentialsOK) GetPayload() models.GCPZoneList {
	return o.Payload
}

func (o *ListGCPZonesNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGCPZonesNoCredentialsDefault creates a ListGCPZonesNoCredentialsDefault with default headers values
func NewListGCPZonesNoCredentialsDefault(code int) *ListGCPZonesNoCredentialsDefault {
	return &ListGCPZonesNoCredentialsDefault{
		_statusCode: code,
	}
}

/*
ListGCPZonesNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListGCPZonesNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list g c p zones no credentials default response
func (o *ListGCPZonesNoCredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list g c p zones no credentials default response has a 2xx status code
func (o *ListGCPZonesNoCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list g c p zones no credentials default response has a 3xx status code
func (o *ListGCPZonesNoCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list g c p zones no credentials default response has a 4xx status code
func (o *ListGCPZonesNoCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list g c p zones no credentials default response has a 5xx status code
func (o *ListGCPZonesNoCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list g c p zones no credentials default response a status code equal to that given
func (o *ListGCPZonesNoCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListGCPZonesNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/zones][%d] listGCPZonesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListGCPZonesNoCredentialsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/zones][%d] listGCPZonesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListGCPZonesNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListGCPZonesNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
