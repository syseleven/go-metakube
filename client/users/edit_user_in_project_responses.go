// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/go-metakube/models"
)

// EditUserInProjectReader is a Reader for the EditUserInProject structure.
type EditUserInProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EditUserInProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEditUserInProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEditUserInProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEditUserInProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEditUserInProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEditUserInProjectOK creates a EditUserInProjectOK with default headers values
func NewEditUserInProjectOK() *EditUserInProjectOK {
	return &EditUserInProjectOK{}
}

/*
EditUserInProjectOK describes a response with status code 200, with default header values.

User
*/
type EditUserInProjectOK struct {
	Payload *models.User
}

// IsSuccess returns true when this edit user in project o k response has a 2xx status code
func (o *EditUserInProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edit user in project o k response has a 3xx status code
func (o *EditUserInProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edit user in project o k response has a 4xx status code
func (o *EditUserInProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edit user in project o k response has a 5xx status code
func (o *EditUserInProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edit user in project o k response a status code equal to that given
func (o *EditUserInProjectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edit user in project o k response
func (o *EditUserInProjectOK) Code() int {
	return 200
}

func (o *EditUserInProjectOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/users/{user_id}][%d] editUserInProjectOK %s", 200, payload)
}

func (o *EditUserInProjectOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/users/{user_id}][%d] editUserInProjectOK %s", 200, payload)
}

func (o *EditUserInProjectOK) GetPayload() *models.User {
	return o.Payload
}

func (o *EditUserInProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditUserInProjectUnauthorized creates a EditUserInProjectUnauthorized with default headers values
func NewEditUserInProjectUnauthorized() *EditUserInProjectUnauthorized {
	return &EditUserInProjectUnauthorized{}
}

/*
EditUserInProjectUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type EditUserInProjectUnauthorized struct {
}

// IsSuccess returns true when this edit user in project unauthorized response has a 2xx status code
func (o *EditUserInProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edit user in project unauthorized response has a 3xx status code
func (o *EditUserInProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edit user in project unauthorized response has a 4xx status code
func (o *EditUserInProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edit user in project unauthorized response has a 5xx status code
func (o *EditUserInProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edit user in project unauthorized response a status code equal to that given
func (o *EditUserInProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edit user in project unauthorized response
func (o *EditUserInProjectUnauthorized) Code() int {
	return 401
}

func (o *EditUserInProjectUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/users/{user_id}][%d] editUserInProjectUnauthorized", 401)
}

func (o *EditUserInProjectUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/users/{user_id}][%d] editUserInProjectUnauthorized", 401)
}

func (o *EditUserInProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEditUserInProjectForbidden creates a EditUserInProjectForbidden with default headers values
func NewEditUserInProjectForbidden() *EditUserInProjectForbidden {
	return &EditUserInProjectForbidden{}
}

/*
EditUserInProjectForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type EditUserInProjectForbidden struct {
}

// IsSuccess returns true when this edit user in project forbidden response has a 2xx status code
func (o *EditUserInProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edit user in project forbidden response has a 3xx status code
func (o *EditUserInProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edit user in project forbidden response has a 4xx status code
func (o *EditUserInProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edit user in project forbidden response has a 5xx status code
func (o *EditUserInProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edit user in project forbidden response a status code equal to that given
func (o *EditUserInProjectForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edit user in project forbidden response
func (o *EditUserInProjectForbidden) Code() int {
	return 403
}

func (o *EditUserInProjectForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/users/{user_id}][%d] editUserInProjectForbidden", 403)
}

func (o *EditUserInProjectForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/users/{user_id}][%d] editUserInProjectForbidden", 403)
}

func (o *EditUserInProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEditUserInProjectDefault creates a EditUserInProjectDefault with default headers values
func NewEditUserInProjectDefault(code int) *EditUserInProjectDefault {
	return &EditUserInProjectDefault{
		_statusCode: code,
	}
}

/*
EditUserInProjectDefault describes a response with status code -1, with default header values.

errorResponse
*/
type EditUserInProjectDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this edit user in project default response has a 2xx status code
func (o *EditUserInProjectDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edit user in project default response has a 3xx status code
func (o *EditUserInProjectDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edit user in project default response has a 4xx status code
func (o *EditUserInProjectDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edit user in project default response has a 5xx status code
func (o *EditUserInProjectDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edit user in project default response a status code equal to that given
func (o *EditUserInProjectDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edit user in project default response
func (o *EditUserInProjectDefault) Code() int {
	return o._statusCode
}

func (o *EditUserInProjectDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/users/{user_id}][%d] editUserInProject default %s", o._statusCode, payload)
}

func (o *EditUserInProjectDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/users/{user_id}][%d] editUserInProject default %s", o._statusCode, payload)
}

func (o *EditUserInProjectDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *EditUserInProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
