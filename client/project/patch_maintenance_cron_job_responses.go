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

// PatchMaintenanceCronJobReader is a Reader for the PatchMaintenanceCronJob structure.
type PatchMaintenanceCronJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchMaintenanceCronJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchMaintenanceCronJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchMaintenanceCronJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchMaintenanceCronJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchMaintenanceCronJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchMaintenanceCronJobOK creates a PatchMaintenanceCronJobOK with default headers values
func NewPatchMaintenanceCronJobOK() *PatchMaintenanceCronJobOK {
	return &PatchMaintenanceCronJobOK{}
}

/*
PatchMaintenanceCronJobOK describes a response with status code 200, with default header values.

MaintenanceCronJob
*/
type PatchMaintenanceCronJobOK struct {
	Payload *models.MaintenanceCronJob
}

// IsSuccess returns true when this patch maintenance cron job o k response has a 2xx status code
func (o *PatchMaintenanceCronJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch maintenance cron job o k response has a 3xx status code
func (o *PatchMaintenanceCronJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch maintenance cron job o k response has a 4xx status code
func (o *PatchMaintenanceCronJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch maintenance cron job o k response has a 5xx status code
func (o *PatchMaintenanceCronJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch maintenance cron job o k response a status code equal to that given
func (o *PatchMaintenanceCronJobOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch maintenance cron job o k response
func (o *PatchMaintenanceCronJobOK) Code() int {
	return 200
}

func (o *PatchMaintenanceCronJobOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/maintenancecronjobs/{maintenancecronjob_id}][%d] patchMaintenanceCronJobOK  %+v", 200, o.Payload)
}

func (o *PatchMaintenanceCronJobOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/maintenancecronjobs/{maintenancecronjob_id}][%d] patchMaintenanceCronJobOK  %+v", 200, o.Payload)
}

func (o *PatchMaintenanceCronJobOK) GetPayload() *models.MaintenanceCronJob {
	return o.Payload
}

func (o *PatchMaintenanceCronJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MaintenanceCronJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMaintenanceCronJobUnauthorized creates a PatchMaintenanceCronJobUnauthorized with default headers values
func NewPatchMaintenanceCronJobUnauthorized() *PatchMaintenanceCronJobUnauthorized {
	return &PatchMaintenanceCronJobUnauthorized{}
}

/*
PatchMaintenanceCronJobUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PatchMaintenanceCronJobUnauthorized struct {
}

// IsSuccess returns true when this patch maintenance cron job unauthorized response has a 2xx status code
func (o *PatchMaintenanceCronJobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch maintenance cron job unauthorized response has a 3xx status code
func (o *PatchMaintenanceCronJobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch maintenance cron job unauthorized response has a 4xx status code
func (o *PatchMaintenanceCronJobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch maintenance cron job unauthorized response has a 5xx status code
func (o *PatchMaintenanceCronJobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch maintenance cron job unauthorized response a status code equal to that given
func (o *PatchMaintenanceCronJobUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the patch maintenance cron job unauthorized response
func (o *PatchMaintenanceCronJobUnauthorized) Code() int {
	return 401
}

func (o *PatchMaintenanceCronJobUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/maintenancecronjobs/{maintenancecronjob_id}][%d] patchMaintenanceCronJobUnauthorized ", 401)
}

func (o *PatchMaintenanceCronJobUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/maintenancecronjobs/{maintenancecronjob_id}][%d] patchMaintenanceCronJobUnauthorized ", 401)
}

func (o *PatchMaintenanceCronJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchMaintenanceCronJobForbidden creates a PatchMaintenanceCronJobForbidden with default headers values
func NewPatchMaintenanceCronJobForbidden() *PatchMaintenanceCronJobForbidden {
	return &PatchMaintenanceCronJobForbidden{}
}

/*
PatchMaintenanceCronJobForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type PatchMaintenanceCronJobForbidden struct {
}

// IsSuccess returns true when this patch maintenance cron job forbidden response has a 2xx status code
func (o *PatchMaintenanceCronJobForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch maintenance cron job forbidden response has a 3xx status code
func (o *PatchMaintenanceCronJobForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch maintenance cron job forbidden response has a 4xx status code
func (o *PatchMaintenanceCronJobForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch maintenance cron job forbidden response has a 5xx status code
func (o *PatchMaintenanceCronJobForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch maintenance cron job forbidden response a status code equal to that given
func (o *PatchMaintenanceCronJobForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the patch maintenance cron job forbidden response
func (o *PatchMaintenanceCronJobForbidden) Code() int {
	return 403
}

func (o *PatchMaintenanceCronJobForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/maintenancecronjobs/{maintenancecronjob_id}][%d] patchMaintenanceCronJobForbidden ", 403)
}

func (o *PatchMaintenanceCronJobForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/maintenancecronjobs/{maintenancecronjob_id}][%d] patchMaintenanceCronJobForbidden ", 403)
}

func (o *PatchMaintenanceCronJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchMaintenanceCronJobDefault creates a PatchMaintenanceCronJobDefault with default headers values
func NewPatchMaintenanceCronJobDefault(code int) *PatchMaintenanceCronJobDefault {
	return &PatchMaintenanceCronJobDefault{
		_statusCode: code,
	}
}

/*
PatchMaintenanceCronJobDefault describes a response with status code -1, with default header values.

errorResponse
*/
type PatchMaintenanceCronJobDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this patch maintenance cron job default response has a 2xx status code
func (o *PatchMaintenanceCronJobDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch maintenance cron job default response has a 3xx status code
func (o *PatchMaintenanceCronJobDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch maintenance cron job default response has a 4xx status code
func (o *PatchMaintenanceCronJobDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch maintenance cron job default response has a 5xx status code
func (o *PatchMaintenanceCronJobDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch maintenance cron job default response a status code equal to that given
func (o *PatchMaintenanceCronJobDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch maintenance cron job default response
func (o *PatchMaintenanceCronJobDefault) Code() int {
	return o._statusCode
}

func (o *PatchMaintenanceCronJobDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/maintenancecronjobs/{maintenancecronjob_id}][%d] patchMaintenanceCronJob default  %+v", o._statusCode, o.Payload)
}

func (o *PatchMaintenanceCronJobDefault) String() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/maintenancecronjobs/{maintenancecronjob_id}][%d] patchMaintenanceCronJob default  %+v", o._statusCode, o.Payload)
}

func (o *PatchMaintenanceCronJobDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchMaintenanceCronJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
