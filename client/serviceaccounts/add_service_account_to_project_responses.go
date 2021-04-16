// Code generated by go-swagger; DO NOT EDIT.

package serviceaccounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/go-metakube/models"
)

// AddServiceAccountToProjectReader is a Reader for the AddServiceAccountToProject structure.
type AddServiceAccountToProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddServiceAccountToProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddServiceAccountToProjectCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddServiceAccountToProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddServiceAccountToProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddServiceAccountToProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddServiceAccountToProjectCreated creates a AddServiceAccountToProjectCreated with default headers values
func NewAddServiceAccountToProjectCreated() *AddServiceAccountToProjectCreated {
	return &AddServiceAccountToProjectCreated{}
}

/*AddServiceAccountToProjectCreated handles this case with default header values.

ServiceAccount
*/
type AddServiceAccountToProjectCreated struct {
	Payload *models.ServiceAccount
}

func (o *AddServiceAccountToProjectCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/serviceaccounts][%d] addServiceAccountToProjectCreated  %+v", 201, o.Payload)
}

func (o *AddServiceAccountToProjectCreated) GetPayload() *models.ServiceAccount {
	return o.Payload
}

func (o *AddServiceAccountToProjectCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddServiceAccountToProjectUnauthorized creates a AddServiceAccountToProjectUnauthorized with default headers values
func NewAddServiceAccountToProjectUnauthorized() *AddServiceAccountToProjectUnauthorized {
	return &AddServiceAccountToProjectUnauthorized{}
}

/*AddServiceAccountToProjectUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type AddServiceAccountToProjectUnauthorized struct {
}

func (o *AddServiceAccountToProjectUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/serviceaccounts][%d] addServiceAccountToProjectUnauthorized ", 401)
}

func (o *AddServiceAccountToProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddServiceAccountToProjectForbidden creates a AddServiceAccountToProjectForbidden with default headers values
func NewAddServiceAccountToProjectForbidden() *AddServiceAccountToProjectForbidden {
	return &AddServiceAccountToProjectForbidden{}
}

/*AddServiceAccountToProjectForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type AddServiceAccountToProjectForbidden struct {
}

func (o *AddServiceAccountToProjectForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/serviceaccounts][%d] addServiceAccountToProjectForbidden ", 403)
}

func (o *AddServiceAccountToProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddServiceAccountToProjectDefault creates a AddServiceAccountToProjectDefault with default headers values
func NewAddServiceAccountToProjectDefault(code int) *AddServiceAccountToProjectDefault {
	return &AddServiceAccountToProjectDefault{
		_statusCode: code,
	}
}

/*AddServiceAccountToProjectDefault handles this case with default header values.

errorResponse
*/
type AddServiceAccountToProjectDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add service account to project default response
func (o *AddServiceAccountToProjectDefault) Code() int {
	return o._statusCode
}

func (o *AddServiceAccountToProjectDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/serviceaccounts][%d] addServiceAccountToProject default  %+v", o._statusCode, o.Payload)
}

func (o *AddServiceAccountToProjectDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddServiceAccountToProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
