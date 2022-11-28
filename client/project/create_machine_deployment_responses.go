// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/go-metakube/models"
)

// CreateMachineDeploymentReader is a Reader for the CreateMachineDeployment structure.
type CreateMachineDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMachineDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateMachineDeploymentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateMachineDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateMachineDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateMachineDeploymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMachineDeploymentCreated creates a CreateMachineDeploymentCreated with default headers values
func NewCreateMachineDeploymentCreated() *CreateMachineDeploymentCreated {
	return &CreateMachineDeploymentCreated{}
}

/*
CreateMachineDeploymentCreated describes a response with status code 201, with default header values.

NodeDeployment
*/
type CreateMachineDeploymentCreated struct {
	Payload *models.NodeDeployment
}

// IsSuccess returns true when this create machine deployment created response has a 2xx status code
func (o *CreateMachineDeploymentCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create machine deployment created response has a 3xx status code
func (o *CreateMachineDeploymentCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create machine deployment created response has a 4xx status code
func (o *CreateMachineDeploymentCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create machine deployment created response has a 5xx status code
func (o *CreateMachineDeploymentCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create machine deployment created response a status code equal to that given
func (o *CreateMachineDeploymentCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateMachineDeploymentCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments][%d] createMachineDeploymentCreated  %+v", 201, o.Payload)
}

func (o *CreateMachineDeploymentCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments][%d] createMachineDeploymentCreated  %+v", 201, o.Payload)
}

func (o *CreateMachineDeploymentCreated) GetPayload() *models.NodeDeployment {
	return o.Payload
}

func (o *CreateMachineDeploymentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeDeployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMachineDeploymentUnauthorized creates a CreateMachineDeploymentUnauthorized with default headers values
func NewCreateMachineDeploymentUnauthorized() *CreateMachineDeploymentUnauthorized {
	return &CreateMachineDeploymentUnauthorized{}
}

/*
CreateMachineDeploymentUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateMachineDeploymentUnauthorized struct {
}

// IsSuccess returns true when this create machine deployment unauthorized response has a 2xx status code
func (o *CreateMachineDeploymentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create machine deployment unauthorized response has a 3xx status code
func (o *CreateMachineDeploymentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create machine deployment unauthorized response has a 4xx status code
func (o *CreateMachineDeploymentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create machine deployment unauthorized response has a 5xx status code
func (o *CreateMachineDeploymentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create machine deployment unauthorized response a status code equal to that given
func (o *CreateMachineDeploymentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateMachineDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments][%d] createMachineDeploymentUnauthorized ", 401)
}

func (o *CreateMachineDeploymentUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments][%d] createMachineDeploymentUnauthorized ", 401)
}

func (o *CreateMachineDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateMachineDeploymentForbidden creates a CreateMachineDeploymentForbidden with default headers values
func NewCreateMachineDeploymentForbidden() *CreateMachineDeploymentForbidden {
	return &CreateMachineDeploymentForbidden{}
}

/*
CreateMachineDeploymentForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateMachineDeploymentForbidden struct {
}

// IsSuccess returns true when this create machine deployment forbidden response has a 2xx status code
func (o *CreateMachineDeploymentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create machine deployment forbidden response has a 3xx status code
func (o *CreateMachineDeploymentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create machine deployment forbidden response has a 4xx status code
func (o *CreateMachineDeploymentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create machine deployment forbidden response has a 5xx status code
func (o *CreateMachineDeploymentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create machine deployment forbidden response a status code equal to that given
func (o *CreateMachineDeploymentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateMachineDeploymentForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments][%d] createMachineDeploymentForbidden ", 403)
}

func (o *CreateMachineDeploymentForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments][%d] createMachineDeploymentForbidden ", 403)
}

func (o *CreateMachineDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateMachineDeploymentDefault creates a CreateMachineDeploymentDefault with default headers values
func NewCreateMachineDeploymentDefault(code int) *CreateMachineDeploymentDefault {
	return &CreateMachineDeploymentDefault{
		_statusCode: code,
	}
}

/*
CreateMachineDeploymentDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateMachineDeploymentDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create machine deployment default response
func (o *CreateMachineDeploymentDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create machine deployment default response has a 2xx status code
func (o *CreateMachineDeploymentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create machine deployment default response has a 3xx status code
func (o *CreateMachineDeploymentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create machine deployment default response has a 4xx status code
func (o *CreateMachineDeploymentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create machine deployment default response has a 5xx status code
func (o *CreateMachineDeploymentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create machine deployment default response a status code equal to that given
func (o *CreateMachineDeploymentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateMachineDeploymentDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments][%d] createMachineDeployment default  %+v", o._statusCode, o.Payload)
}

func (o *CreateMachineDeploymentDefault) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments][%d] createMachineDeployment default  %+v", o._statusCode, o.Payload)
}

func (o *CreateMachineDeploymentDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateMachineDeploymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
