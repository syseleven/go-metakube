// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/go-metakube/models"
)

// ListAzureVnetsReader is a Reader for the ListAzureVnets structure.
type ListAzureVnetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAzureVnetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAzureVnetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAzureVnetsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAzureVnetsOK creates a ListAzureVnetsOK with default headers values
func NewListAzureVnetsOK() *ListAzureVnetsOK {
	return &ListAzureVnetsOK{}
}

/*
ListAzureVnetsOK describes a response with status code 200, with default header values.

AzureVirtualNetworksList
*/
type ListAzureVnetsOK struct {
	Payload *models.AzureVirtualNetworksList
}

// IsSuccess returns true when this list azure vnets o k response has a 2xx status code
func (o *ListAzureVnetsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list azure vnets o k response has a 3xx status code
func (o *ListAzureVnetsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list azure vnets o k response has a 4xx status code
func (o *ListAzureVnetsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list azure vnets o k response has a 5xx status code
func (o *ListAzureVnetsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list azure vnets o k response a status code equal to that given
func (o *ListAzureVnetsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListAzureVnetsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/azure/vnets][%d] listAzureVnetsOK  %+v", 200, o.Payload)
}

func (o *ListAzureVnetsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/azure/vnets][%d] listAzureVnetsOK  %+v", 200, o.Payload)
}

func (o *ListAzureVnetsOK) GetPayload() *models.AzureVirtualNetworksList {
	return o.Payload
}

func (o *ListAzureVnetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureVirtualNetworksList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAzureVnetsDefault creates a ListAzureVnetsDefault with default headers values
func NewListAzureVnetsDefault(code int) *ListAzureVnetsDefault {
	return &ListAzureVnetsDefault{
		_statusCode: code,
	}
}

/*
ListAzureVnetsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListAzureVnetsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list azure vnets default response
func (o *ListAzureVnetsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list azure vnets default response has a 2xx status code
func (o *ListAzureVnetsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list azure vnets default response has a 3xx status code
func (o *ListAzureVnetsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list azure vnets default response has a 4xx status code
func (o *ListAzureVnetsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list azure vnets default response has a 5xx status code
func (o *ListAzureVnetsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list azure vnets default response a status code equal to that given
func (o *ListAzureVnetsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListAzureVnetsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/azure/vnets][%d] listAzureVnets default  %+v", o._statusCode, o.Payload)
}

func (o *ListAzureVnetsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/azure/vnets][%d] listAzureVnets default  %+v", o._statusCode, o.Payload)
}

func (o *ListAzureVnetsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAzureVnetsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
