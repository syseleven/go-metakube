// Code generated by go-swagger; DO NOT EDIT.

package datacenter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/go-metakube/models"
)

// ListDatacentersV2Reader is a Reader for the ListDatacentersV2 structure.
type ListDatacentersV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDatacentersV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDatacentersV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListDatacentersV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDatacentersV2OK creates a ListDatacentersV2OK with default headers values
func NewListDatacentersV2OK() *ListDatacentersV2OK {
	return &ListDatacentersV2OK{}
}

/*
ListDatacentersV2OK describes a response with status code 200, with default header values.

Datacenter
*/
type ListDatacentersV2OK struct {
	Payload []*models.Datacenter
}

// IsSuccess returns true when this list datacenters v2 o k response has a 2xx status code
func (o *ListDatacentersV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list datacenters v2 o k response has a 3xx status code
func (o *ListDatacentersV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list datacenters v2 o k response has a 4xx status code
func (o *ListDatacentersV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list datacenters v2 o k response has a 5xx status code
func (o *ListDatacentersV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this list datacenters v2 o k response a status code equal to that given
func (o *ListDatacentersV2OK) IsCode(code int) bool {
	return code == 200
}

func (o *ListDatacentersV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/datacenters][%d] listDatacentersV2OK  %+v", 200, o.Payload)
}

func (o *ListDatacentersV2OK) String() string {
	return fmt.Sprintf("[GET /api/v2/datacenters][%d] listDatacentersV2OK  %+v", 200, o.Payload)
}

func (o *ListDatacentersV2OK) GetPayload() []*models.Datacenter {
	return o.Payload
}

func (o *ListDatacentersV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDatacentersV2Default creates a ListDatacentersV2Default with default headers values
func NewListDatacentersV2Default(code int) *ListDatacentersV2Default {
	return &ListDatacentersV2Default{
		_statusCode: code,
	}
}

/*
ListDatacentersV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type ListDatacentersV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list datacenters v2 default response
func (o *ListDatacentersV2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list datacenters v2 default response has a 2xx status code
func (o *ListDatacentersV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list datacenters v2 default response has a 3xx status code
func (o *ListDatacentersV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list datacenters v2 default response has a 4xx status code
func (o *ListDatacentersV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list datacenters v2 default response has a 5xx status code
func (o *ListDatacentersV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list datacenters v2 default response a status code equal to that given
func (o *ListDatacentersV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListDatacentersV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/datacenters][%d] listDatacentersV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListDatacentersV2Default) String() string {
	return fmt.Sprintf("[GET /api/v2/datacenters][%d] listDatacentersV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListDatacentersV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListDatacentersV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}