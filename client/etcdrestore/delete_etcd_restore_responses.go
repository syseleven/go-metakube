// Code generated by go-swagger; DO NOT EDIT.

package etcdrestore

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

// DeleteEtcdRestoreReader is a Reader for the DeleteEtcdRestore structure.
type DeleteEtcdRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEtcdRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEtcdRestoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteEtcdRestoreUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteEtcdRestoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteEtcdRestoreConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteEtcdRestoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteEtcdRestoreOK creates a DeleteEtcdRestoreOK with default headers values
func NewDeleteEtcdRestoreOK() *DeleteEtcdRestoreOK {
	return &DeleteEtcdRestoreOK{}
}

/*
DeleteEtcdRestoreOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteEtcdRestoreOK struct {
}

// IsSuccess returns true when this delete etcd restore o k response has a 2xx status code
func (o *DeleteEtcdRestoreOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete etcd restore o k response has a 3xx status code
func (o *DeleteEtcdRestoreOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete etcd restore o k response has a 4xx status code
func (o *DeleteEtcdRestoreOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete etcd restore o k response has a 5xx status code
func (o *DeleteEtcdRestoreOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete etcd restore o k response a status code equal to that given
func (o *DeleteEtcdRestoreOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete etcd restore o k response
func (o *DeleteEtcdRestoreOK) Code() int {
	return 200
}

func (o *DeleteEtcdRestoreOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] deleteEtcdRestoreOK", 200)
}

func (o *DeleteEtcdRestoreOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] deleteEtcdRestoreOK", 200)
}

func (o *DeleteEtcdRestoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEtcdRestoreUnauthorized creates a DeleteEtcdRestoreUnauthorized with default headers values
func NewDeleteEtcdRestoreUnauthorized() *DeleteEtcdRestoreUnauthorized {
	return &DeleteEtcdRestoreUnauthorized{}
}

/*
DeleteEtcdRestoreUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteEtcdRestoreUnauthorized struct {
}

// IsSuccess returns true when this delete etcd restore unauthorized response has a 2xx status code
func (o *DeleteEtcdRestoreUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete etcd restore unauthorized response has a 3xx status code
func (o *DeleteEtcdRestoreUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete etcd restore unauthorized response has a 4xx status code
func (o *DeleteEtcdRestoreUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete etcd restore unauthorized response has a 5xx status code
func (o *DeleteEtcdRestoreUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete etcd restore unauthorized response a status code equal to that given
func (o *DeleteEtcdRestoreUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete etcd restore unauthorized response
func (o *DeleteEtcdRestoreUnauthorized) Code() int {
	return 401
}

func (o *DeleteEtcdRestoreUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] deleteEtcdRestoreUnauthorized", 401)
}

func (o *DeleteEtcdRestoreUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] deleteEtcdRestoreUnauthorized", 401)
}

func (o *DeleteEtcdRestoreUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEtcdRestoreForbidden creates a DeleteEtcdRestoreForbidden with default headers values
func NewDeleteEtcdRestoreForbidden() *DeleteEtcdRestoreForbidden {
	return &DeleteEtcdRestoreForbidden{}
}

/*
DeleteEtcdRestoreForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteEtcdRestoreForbidden struct {
}

// IsSuccess returns true when this delete etcd restore forbidden response has a 2xx status code
func (o *DeleteEtcdRestoreForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete etcd restore forbidden response has a 3xx status code
func (o *DeleteEtcdRestoreForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete etcd restore forbidden response has a 4xx status code
func (o *DeleteEtcdRestoreForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete etcd restore forbidden response has a 5xx status code
func (o *DeleteEtcdRestoreForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete etcd restore forbidden response a status code equal to that given
func (o *DeleteEtcdRestoreForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete etcd restore forbidden response
func (o *DeleteEtcdRestoreForbidden) Code() int {
	return 403
}

func (o *DeleteEtcdRestoreForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] deleteEtcdRestoreForbidden", 403)
}

func (o *DeleteEtcdRestoreForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] deleteEtcdRestoreForbidden", 403)
}

func (o *DeleteEtcdRestoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEtcdRestoreConflict creates a DeleteEtcdRestoreConflict with default headers values
func NewDeleteEtcdRestoreConflict() *DeleteEtcdRestoreConflict {
	return &DeleteEtcdRestoreConflict{}
}

/*
DeleteEtcdRestoreConflict describes a response with status code 409, with default header values.

errorResponse
*/
type DeleteEtcdRestoreConflict struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete etcd restore conflict response has a 2xx status code
func (o *DeleteEtcdRestoreConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete etcd restore conflict response has a 3xx status code
func (o *DeleteEtcdRestoreConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete etcd restore conflict response has a 4xx status code
func (o *DeleteEtcdRestoreConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete etcd restore conflict response has a 5xx status code
func (o *DeleteEtcdRestoreConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete etcd restore conflict response a status code equal to that given
func (o *DeleteEtcdRestoreConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the delete etcd restore conflict response
func (o *DeleteEtcdRestoreConflict) Code() int {
	return 409
}

func (o *DeleteEtcdRestoreConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] deleteEtcdRestoreConflict %s", 409, payload)
}

func (o *DeleteEtcdRestoreConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] deleteEtcdRestoreConflict %s", 409, payload)
}

func (o *DeleteEtcdRestoreConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteEtcdRestoreConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEtcdRestoreDefault creates a DeleteEtcdRestoreDefault with default headers values
func NewDeleteEtcdRestoreDefault(code int) *DeleteEtcdRestoreDefault {
	return &DeleteEtcdRestoreDefault{
		_statusCode: code,
	}
}

/*
DeleteEtcdRestoreDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteEtcdRestoreDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete etcd restore default response has a 2xx status code
func (o *DeleteEtcdRestoreDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete etcd restore default response has a 3xx status code
func (o *DeleteEtcdRestoreDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete etcd restore default response has a 4xx status code
func (o *DeleteEtcdRestoreDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete etcd restore default response has a 5xx status code
func (o *DeleteEtcdRestoreDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete etcd restore default response a status code equal to that given
func (o *DeleteEtcdRestoreDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete etcd restore default response
func (o *DeleteEtcdRestoreDefault) Code() int {
	return o._statusCode
}

func (o *DeleteEtcdRestoreDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] deleteEtcdRestore default %s", o._statusCode, payload)
}

func (o *DeleteEtcdRestoreDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] deleteEtcdRestore default %s", o._statusCode, payload)
}

func (o *DeleteEtcdRestoreDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteEtcdRestoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
