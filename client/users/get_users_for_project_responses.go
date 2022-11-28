// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/go-metakube/models"
)

// GetUsersForProjectReader is a Reader for the GetUsersForProject structure.
type GetUsersForProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersForProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersForProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUsersForProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsersForProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetUsersForProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUsersForProjectOK creates a GetUsersForProjectOK with default headers values
func NewGetUsersForProjectOK() *GetUsersForProjectOK {
	return &GetUsersForProjectOK{}
}

/*
GetUsersForProjectOK describes a response with status code 200, with default header values.

User
*/
type GetUsersForProjectOK struct {
	Payload []*models.User
}

// IsSuccess returns true when this get users for project o k response has a 2xx status code
func (o *GetUsersForProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get users for project o k response has a 3xx status code
func (o *GetUsersForProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users for project o k response has a 4xx status code
func (o *GetUsersForProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get users for project o k response has a 5xx status code
func (o *GetUsersForProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get users for project o k response a status code equal to that given
func (o *GetUsersForProjectOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUsersForProjectOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/users][%d] getUsersForProjectOK  %+v", 200, o.Payload)
}

func (o *GetUsersForProjectOK) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/users][%d] getUsersForProjectOK  %+v", 200, o.Payload)
}

func (o *GetUsersForProjectOK) GetPayload() []*models.User {
	return o.Payload
}

func (o *GetUsersForProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersForProjectUnauthorized creates a GetUsersForProjectUnauthorized with default headers values
func NewGetUsersForProjectUnauthorized() *GetUsersForProjectUnauthorized {
	return &GetUsersForProjectUnauthorized{}
}

/*
GetUsersForProjectUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetUsersForProjectUnauthorized struct {
}

// IsSuccess returns true when this get users for project unauthorized response has a 2xx status code
func (o *GetUsersForProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users for project unauthorized response has a 3xx status code
func (o *GetUsersForProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users for project unauthorized response has a 4xx status code
func (o *GetUsersForProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users for project unauthorized response has a 5xx status code
func (o *GetUsersForProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get users for project unauthorized response a status code equal to that given
func (o *GetUsersForProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUsersForProjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/users][%d] getUsersForProjectUnauthorized ", 401)
}

func (o *GetUsersForProjectUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/users][%d] getUsersForProjectUnauthorized ", 401)
}

func (o *GetUsersForProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersForProjectForbidden creates a GetUsersForProjectForbidden with default headers values
func NewGetUsersForProjectForbidden() *GetUsersForProjectForbidden {
	return &GetUsersForProjectForbidden{}
}

/*
GetUsersForProjectForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetUsersForProjectForbidden struct {
}

// IsSuccess returns true when this get users for project forbidden response has a 2xx status code
func (o *GetUsersForProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users for project forbidden response has a 3xx status code
func (o *GetUsersForProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users for project forbidden response has a 4xx status code
func (o *GetUsersForProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users for project forbidden response has a 5xx status code
func (o *GetUsersForProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get users for project forbidden response a status code equal to that given
func (o *GetUsersForProjectForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetUsersForProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/users][%d] getUsersForProjectForbidden ", 403)
}

func (o *GetUsersForProjectForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/users][%d] getUsersForProjectForbidden ", 403)
}

func (o *GetUsersForProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersForProjectDefault creates a GetUsersForProjectDefault with default headers values
func NewGetUsersForProjectDefault(code int) *GetUsersForProjectDefault {
	return &GetUsersForProjectDefault{
		_statusCode: code,
	}
}

/*
GetUsersForProjectDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetUsersForProjectDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get users for project default response
func (o *GetUsersForProjectDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get users for project default response has a 2xx status code
func (o *GetUsersForProjectDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get users for project default response has a 3xx status code
func (o *GetUsersForProjectDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get users for project default response has a 4xx status code
func (o *GetUsersForProjectDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get users for project default response has a 5xx status code
func (o *GetUsersForProjectDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get users for project default response a status code equal to that given
func (o *GetUsersForProjectDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetUsersForProjectDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/users][%d] getUsersForProject default  %+v", o._statusCode, o.Payload)
}

func (o *GetUsersForProjectDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/users][%d] getUsersForProject default  %+v", o._statusCode, o.Payload)
}

func (o *GetUsersForProjectDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetUsersForProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
