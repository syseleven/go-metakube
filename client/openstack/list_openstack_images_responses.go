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

// ListOpenstackImagesReader is a Reader for the ListOpenstackImages structure.
type ListOpenstackImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOpenstackImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOpenstackImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListOpenstackImagesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOpenstackImagesOK creates a ListOpenstackImagesOK with default headers values
func NewListOpenstackImagesOK() *ListOpenstackImagesOK {
	return &ListOpenstackImagesOK{}
}

/*
ListOpenstackImagesOK describes a response with status code 200, with default header values.

Image
*/
type ListOpenstackImagesOK struct {
	Payload []*models.Image
}

// IsSuccess returns true when this list openstack images o k response has a 2xx status code
func (o *ListOpenstackImagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list openstack images o k response has a 3xx status code
func (o *ListOpenstackImagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list openstack images o k response has a 4xx status code
func (o *ListOpenstackImagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list openstack images o k response has a 5xx status code
func (o *ListOpenstackImagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list openstack images o k response a status code equal to that given
func (o *ListOpenstackImagesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListOpenstackImagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/openstack/images][%d] listOpenstackImagesOK  %+v", 200, o.Payload)
}

func (o *ListOpenstackImagesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/openstack/images][%d] listOpenstackImagesOK  %+v", 200, o.Payload)
}

func (o *ListOpenstackImagesOK) GetPayload() []*models.Image {
	return o.Payload
}

func (o *ListOpenstackImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOpenstackImagesDefault creates a ListOpenstackImagesDefault with default headers values
func NewListOpenstackImagesDefault(code int) *ListOpenstackImagesDefault {
	return &ListOpenstackImagesDefault{
		_statusCode: code,
	}
}

/*
ListOpenstackImagesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListOpenstackImagesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list openstack images default response
func (o *ListOpenstackImagesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list openstack images default response has a 2xx status code
func (o *ListOpenstackImagesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list openstack images default response has a 3xx status code
func (o *ListOpenstackImagesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list openstack images default response has a 4xx status code
func (o *ListOpenstackImagesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list openstack images default response has a 5xx status code
func (o *ListOpenstackImagesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list openstack images default response a status code equal to that given
func (o *ListOpenstackImagesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListOpenstackImagesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/openstack/images][%d] listOpenstackImages default  %+v", o._statusCode, o.Payload)
}

func (o *ListOpenstackImagesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/openstack/images][%d] listOpenstackImages default  %+v", o._statusCode, o.Payload)
}

func (o *ListOpenstackImagesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListOpenstackImagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
