// Code generated by go-swagger; DO NOT EDIT.

package etcdrestore

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/go-metakube/models"
)

// ListProjectEtcdRestoreReader is a Reader for the ListProjectEtcdRestore structure.
type ListProjectEtcdRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectEtcdRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectEtcdRestoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListProjectEtcdRestoreUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListProjectEtcdRestoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListProjectEtcdRestoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectEtcdRestoreOK creates a ListProjectEtcdRestoreOK with default headers values
func NewListProjectEtcdRestoreOK() *ListProjectEtcdRestoreOK {
	return &ListProjectEtcdRestoreOK{}
}

/*
ListProjectEtcdRestoreOK describes a response with status code 200, with default header values.

EtcdRestore
*/
type ListProjectEtcdRestoreOK struct {
	Payload []*models.EtcdRestore
}

// IsSuccess returns true when this list project etcd restore o k response has a 2xx status code
func (o *ListProjectEtcdRestoreOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project etcd restore o k response has a 3xx status code
func (o *ListProjectEtcdRestoreOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project etcd restore o k response has a 4xx status code
func (o *ListProjectEtcdRestoreOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project etcd restore o k response has a 5xx status code
func (o *ListProjectEtcdRestoreOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project etcd restore o k response a status code equal to that given
func (o *ListProjectEtcdRestoreOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListProjectEtcdRestoreOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/etcdrestores][%d] listProjectEtcdRestoreOK  %+v", 200, o.Payload)
}

func (o *ListProjectEtcdRestoreOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/etcdrestores][%d] listProjectEtcdRestoreOK  %+v", 200, o.Payload)
}

func (o *ListProjectEtcdRestoreOK) GetPayload() []*models.EtcdRestore {
	return o.Payload
}

func (o *ListProjectEtcdRestoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectEtcdRestoreUnauthorized creates a ListProjectEtcdRestoreUnauthorized with default headers values
func NewListProjectEtcdRestoreUnauthorized() *ListProjectEtcdRestoreUnauthorized {
	return &ListProjectEtcdRestoreUnauthorized{}
}

/*
ListProjectEtcdRestoreUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListProjectEtcdRestoreUnauthorized struct {
}

// IsSuccess returns true when this list project etcd restore unauthorized response has a 2xx status code
func (o *ListProjectEtcdRestoreUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project etcd restore unauthorized response has a 3xx status code
func (o *ListProjectEtcdRestoreUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project etcd restore unauthorized response has a 4xx status code
func (o *ListProjectEtcdRestoreUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project etcd restore unauthorized response has a 5xx status code
func (o *ListProjectEtcdRestoreUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list project etcd restore unauthorized response a status code equal to that given
func (o *ListProjectEtcdRestoreUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListProjectEtcdRestoreUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/etcdrestores][%d] listProjectEtcdRestoreUnauthorized ", 401)
}

func (o *ListProjectEtcdRestoreUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/etcdrestores][%d] listProjectEtcdRestoreUnauthorized ", 401)
}

func (o *ListProjectEtcdRestoreUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectEtcdRestoreForbidden creates a ListProjectEtcdRestoreForbidden with default headers values
func NewListProjectEtcdRestoreForbidden() *ListProjectEtcdRestoreForbidden {
	return &ListProjectEtcdRestoreForbidden{}
}

/*
ListProjectEtcdRestoreForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListProjectEtcdRestoreForbidden struct {
}

// IsSuccess returns true when this list project etcd restore forbidden response has a 2xx status code
func (o *ListProjectEtcdRestoreForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project etcd restore forbidden response has a 3xx status code
func (o *ListProjectEtcdRestoreForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project etcd restore forbidden response has a 4xx status code
func (o *ListProjectEtcdRestoreForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project etcd restore forbidden response has a 5xx status code
func (o *ListProjectEtcdRestoreForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list project etcd restore forbidden response a status code equal to that given
func (o *ListProjectEtcdRestoreForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListProjectEtcdRestoreForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/etcdrestores][%d] listProjectEtcdRestoreForbidden ", 403)
}

func (o *ListProjectEtcdRestoreForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/etcdrestores][%d] listProjectEtcdRestoreForbidden ", 403)
}

func (o *ListProjectEtcdRestoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectEtcdRestoreDefault creates a ListProjectEtcdRestoreDefault with default headers values
func NewListProjectEtcdRestoreDefault(code int) *ListProjectEtcdRestoreDefault {
	return &ListProjectEtcdRestoreDefault{
		_statusCode: code,
	}
}

/*
ListProjectEtcdRestoreDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectEtcdRestoreDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list project etcd restore default response
func (o *ListProjectEtcdRestoreDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list project etcd restore default response has a 2xx status code
func (o *ListProjectEtcdRestoreDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project etcd restore default response has a 3xx status code
func (o *ListProjectEtcdRestoreDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project etcd restore default response has a 4xx status code
func (o *ListProjectEtcdRestoreDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project etcd restore default response has a 5xx status code
func (o *ListProjectEtcdRestoreDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project etcd restore default response a status code equal to that given
func (o *ListProjectEtcdRestoreDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListProjectEtcdRestoreDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/etcdrestores][%d] listProjectEtcdRestore default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectEtcdRestoreDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/etcdrestores][%d] listProjectEtcdRestore default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectEtcdRestoreDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectEtcdRestoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
