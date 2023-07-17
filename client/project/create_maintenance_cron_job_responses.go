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

// CreateMaintenanceCronJobReader is a Reader for the CreateMaintenanceCronJob structure.
type CreateMaintenanceCronJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMaintenanceCronJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateMaintenanceCronJobCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateMaintenanceCronJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateMaintenanceCronJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateMaintenanceCronJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMaintenanceCronJobCreated creates a CreateMaintenanceCronJobCreated with default headers values
func NewCreateMaintenanceCronJobCreated() *CreateMaintenanceCronJobCreated {
	return &CreateMaintenanceCronJobCreated{}
}

/*
CreateMaintenanceCronJobCreated describes a response with status code 201, with default header values.

MaintenanceCronJob
*/
type CreateMaintenanceCronJobCreated struct {
	Payload *models.MaintenanceCronJob
}

// IsSuccess returns true when this create maintenance cron job created response has a 2xx status code
func (o *CreateMaintenanceCronJobCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create maintenance cron job created response has a 3xx status code
func (o *CreateMaintenanceCronJobCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create maintenance cron job created response has a 4xx status code
func (o *CreateMaintenanceCronJobCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create maintenance cron job created response has a 5xx status code
func (o *CreateMaintenanceCronJobCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create maintenance cron job created response a status code equal to that given
func (o *CreateMaintenanceCronJobCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create maintenance cron job created response
func (o *CreateMaintenanceCronJobCreated) Code() int {
	return 201
}

func (o *CreateMaintenanceCronJobCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/maintenancecronjobs][%d] createMaintenanceCronJobCreated  %+v", 201, o.Payload)
}

func (o *CreateMaintenanceCronJobCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/maintenancecronjobs][%d] createMaintenanceCronJobCreated  %+v", 201, o.Payload)
}

func (o *CreateMaintenanceCronJobCreated) GetPayload() *models.MaintenanceCronJob {
	return o.Payload
}

func (o *CreateMaintenanceCronJobCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MaintenanceCronJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMaintenanceCronJobUnauthorized creates a CreateMaintenanceCronJobUnauthorized with default headers values
func NewCreateMaintenanceCronJobUnauthorized() *CreateMaintenanceCronJobUnauthorized {
	return &CreateMaintenanceCronJobUnauthorized{}
}

/*
CreateMaintenanceCronJobUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateMaintenanceCronJobUnauthorized struct {
}

// IsSuccess returns true when this create maintenance cron job unauthorized response has a 2xx status code
func (o *CreateMaintenanceCronJobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create maintenance cron job unauthorized response has a 3xx status code
func (o *CreateMaintenanceCronJobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create maintenance cron job unauthorized response has a 4xx status code
func (o *CreateMaintenanceCronJobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create maintenance cron job unauthorized response has a 5xx status code
func (o *CreateMaintenanceCronJobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create maintenance cron job unauthorized response a status code equal to that given
func (o *CreateMaintenanceCronJobUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create maintenance cron job unauthorized response
func (o *CreateMaintenanceCronJobUnauthorized) Code() int {
	return 401
}

func (o *CreateMaintenanceCronJobUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/maintenancecronjobs][%d] createMaintenanceCronJobUnauthorized ", 401)
}

func (o *CreateMaintenanceCronJobUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/maintenancecronjobs][%d] createMaintenanceCronJobUnauthorized ", 401)
}

func (o *CreateMaintenanceCronJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateMaintenanceCronJobForbidden creates a CreateMaintenanceCronJobForbidden with default headers values
func NewCreateMaintenanceCronJobForbidden() *CreateMaintenanceCronJobForbidden {
	return &CreateMaintenanceCronJobForbidden{}
}

/*
CreateMaintenanceCronJobForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateMaintenanceCronJobForbidden struct {
}

// IsSuccess returns true when this create maintenance cron job forbidden response has a 2xx status code
func (o *CreateMaintenanceCronJobForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create maintenance cron job forbidden response has a 3xx status code
func (o *CreateMaintenanceCronJobForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create maintenance cron job forbidden response has a 4xx status code
func (o *CreateMaintenanceCronJobForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create maintenance cron job forbidden response has a 5xx status code
func (o *CreateMaintenanceCronJobForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create maintenance cron job forbidden response a status code equal to that given
func (o *CreateMaintenanceCronJobForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create maintenance cron job forbidden response
func (o *CreateMaintenanceCronJobForbidden) Code() int {
	return 403
}

func (o *CreateMaintenanceCronJobForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/maintenancecronjobs][%d] createMaintenanceCronJobForbidden ", 403)
}

func (o *CreateMaintenanceCronJobForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/maintenancecronjobs][%d] createMaintenanceCronJobForbidden ", 403)
}

func (o *CreateMaintenanceCronJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateMaintenanceCronJobDefault creates a CreateMaintenanceCronJobDefault with default headers values
func NewCreateMaintenanceCronJobDefault(code int) *CreateMaintenanceCronJobDefault {
	return &CreateMaintenanceCronJobDefault{
		_statusCode: code,
	}
}

/*
CreateMaintenanceCronJobDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateMaintenanceCronJobDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create maintenance cron job default response has a 2xx status code
func (o *CreateMaintenanceCronJobDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create maintenance cron job default response has a 3xx status code
func (o *CreateMaintenanceCronJobDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create maintenance cron job default response has a 4xx status code
func (o *CreateMaintenanceCronJobDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create maintenance cron job default response has a 5xx status code
func (o *CreateMaintenanceCronJobDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create maintenance cron job default response a status code equal to that given
func (o *CreateMaintenanceCronJobDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create maintenance cron job default response
func (o *CreateMaintenanceCronJobDefault) Code() int {
	return o._statusCode
}

func (o *CreateMaintenanceCronJobDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/maintenancecronjobs][%d] createMaintenanceCronJob default  %+v", o._statusCode, o.Payload)
}

func (o *CreateMaintenanceCronJobDefault) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/maintenancecronjobs][%d] createMaintenanceCronJob default  %+v", o._statusCode, o.Payload)
}

func (o *CreateMaintenanceCronJobDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateMaintenanceCronJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}