// Code generated by go-swagger; DO NOT EDIT.

package project

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

// CreateSSHKeyReader is a Reader for the CreateSSHKey structure.
type CreateSSHKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSSHKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateSSHKeyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateSSHKeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateSSHKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateSSHKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSSHKeyCreated creates a CreateSSHKeyCreated with default headers values
func NewCreateSSHKeyCreated() *CreateSSHKeyCreated {
	return &CreateSSHKeyCreated{}
}

/*
CreateSSHKeyCreated describes a response with status code 201, with default header values.

SSHKey
*/
type CreateSSHKeyCreated struct {
	Payload *models.SSHKey
}

// IsSuccess returns true when this create Ssh key created response has a 2xx status code
func (o *CreateSSHKeyCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create Ssh key created response has a 3xx status code
func (o *CreateSSHKeyCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Ssh key created response has a 4xx status code
func (o *CreateSSHKeyCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create Ssh key created response has a 5xx status code
func (o *CreateSSHKeyCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create Ssh key created response a status code equal to that given
func (o *CreateSSHKeyCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create Ssh key created response
func (o *CreateSSHKeyCreated) Code() int {
	return 201
}

func (o *CreateSSHKeyCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/sshkeys][%d] createSshKeyCreated %s", 201, payload)
}

func (o *CreateSSHKeyCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/sshkeys][%d] createSshKeyCreated %s", 201, payload)
}

func (o *CreateSSHKeyCreated) GetPayload() *models.SSHKey {
	return o.Payload
}

func (o *CreateSSHKeyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SSHKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSSHKeyUnauthorized creates a CreateSSHKeyUnauthorized with default headers values
func NewCreateSSHKeyUnauthorized() *CreateSSHKeyUnauthorized {
	return &CreateSSHKeyUnauthorized{}
}

/*
CreateSSHKeyUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateSSHKeyUnauthorized struct {
}

// IsSuccess returns true when this create Ssh key unauthorized response has a 2xx status code
func (o *CreateSSHKeyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create Ssh key unauthorized response has a 3xx status code
func (o *CreateSSHKeyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Ssh key unauthorized response has a 4xx status code
func (o *CreateSSHKeyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create Ssh key unauthorized response has a 5xx status code
func (o *CreateSSHKeyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create Ssh key unauthorized response a status code equal to that given
func (o *CreateSSHKeyUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create Ssh key unauthorized response
func (o *CreateSSHKeyUnauthorized) Code() int {
	return 401
}

func (o *CreateSSHKeyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/sshkeys][%d] createSshKeyUnauthorized", 401)
}

func (o *CreateSSHKeyUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/sshkeys][%d] createSshKeyUnauthorized", 401)
}

func (o *CreateSSHKeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSSHKeyForbidden creates a CreateSSHKeyForbidden with default headers values
func NewCreateSSHKeyForbidden() *CreateSSHKeyForbidden {
	return &CreateSSHKeyForbidden{}
}

/*
CreateSSHKeyForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateSSHKeyForbidden struct {
}

// IsSuccess returns true when this create Ssh key forbidden response has a 2xx status code
func (o *CreateSSHKeyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create Ssh key forbidden response has a 3xx status code
func (o *CreateSSHKeyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Ssh key forbidden response has a 4xx status code
func (o *CreateSSHKeyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create Ssh key forbidden response has a 5xx status code
func (o *CreateSSHKeyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create Ssh key forbidden response a status code equal to that given
func (o *CreateSSHKeyForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create Ssh key forbidden response
func (o *CreateSSHKeyForbidden) Code() int {
	return 403
}

func (o *CreateSSHKeyForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/sshkeys][%d] createSshKeyForbidden", 403)
}

func (o *CreateSSHKeyForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/sshkeys][%d] createSshKeyForbidden", 403)
}

func (o *CreateSSHKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSSHKeyDefault creates a CreateSSHKeyDefault with default headers values
func NewCreateSSHKeyDefault(code int) *CreateSSHKeyDefault {
	return &CreateSSHKeyDefault{
		_statusCode: code,
	}
}

/*
CreateSSHKeyDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateSSHKeyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create SSH key default response has a 2xx status code
func (o *CreateSSHKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create SSH key default response has a 3xx status code
func (o *CreateSSHKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create SSH key default response has a 4xx status code
func (o *CreateSSHKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create SSH key default response has a 5xx status code
func (o *CreateSSHKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create SSH key default response a status code equal to that given
func (o *CreateSSHKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create SSH key default response
func (o *CreateSSHKeyDefault) Code() int {
	return o._statusCode
}

func (o *CreateSSHKeyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/sshkeys][%d] createSSHKey default %s", o._statusCode, payload)
}

func (o *CreateSSHKeyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/sshkeys][%d] createSSHKey default %s", o._statusCode, payload)
}

func (o *CreateSSHKeyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateSSHKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
