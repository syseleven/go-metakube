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

// ListEtcdRestoreReader is a Reader for the ListEtcdRestore structure.
type ListEtcdRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEtcdRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEtcdRestoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListEtcdRestoreUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListEtcdRestoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListEtcdRestoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListEtcdRestoreOK creates a ListEtcdRestoreOK with default headers values
func NewListEtcdRestoreOK() *ListEtcdRestoreOK {
	return &ListEtcdRestoreOK{}
}

/*
ListEtcdRestoreOK describes a response with status code 200, with default header values.

EtcdRestore
*/
type ListEtcdRestoreOK struct {
	Payload []*models.EtcdRestore
}

// IsSuccess returns true when this list etcd restore o k response has a 2xx status code
func (o *ListEtcdRestoreOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list etcd restore o k response has a 3xx status code
func (o *ListEtcdRestoreOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list etcd restore o k response has a 4xx status code
func (o *ListEtcdRestoreOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list etcd restore o k response has a 5xx status code
func (o *ListEtcdRestoreOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list etcd restore o k response a status code equal to that given
func (o *ListEtcdRestoreOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListEtcdRestoreOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores][%d] listEtcdRestoreOK  %+v", 200, o.Payload)
}

func (o *ListEtcdRestoreOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores][%d] listEtcdRestoreOK  %+v", 200, o.Payload)
}

func (o *ListEtcdRestoreOK) GetPayload() []*models.EtcdRestore {
	return o.Payload
}

func (o *ListEtcdRestoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEtcdRestoreUnauthorized creates a ListEtcdRestoreUnauthorized with default headers values
func NewListEtcdRestoreUnauthorized() *ListEtcdRestoreUnauthorized {
	return &ListEtcdRestoreUnauthorized{}
}

/*
ListEtcdRestoreUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListEtcdRestoreUnauthorized struct {
}

// IsSuccess returns true when this list etcd restore unauthorized response has a 2xx status code
func (o *ListEtcdRestoreUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list etcd restore unauthorized response has a 3xx status code
func (o *ListEtcdRestoreUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list etcd restore unauthorized response has a 4xx status code
func (o *ListEtcdRestoreUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list etcd restore unauthorized response has a 5xx status code
func (o *ListEtcdRestoreUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list etcd restore unauthorized response a status code equal to that given
func (o *ListEtcdRestoreUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListEtcdRestoreUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores][%d] listEtcdRestoreUnauthorized ", 401)
}

func (o *ListEtcdRestoreUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores][%d] listEtcdRestoreUnauthorized ", 401)
}

func (o *ListEtcdRestoreUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEtcdRestoreForbidden creates a ListEtcdRestoreForbidden with default headers values
func NewListEtcdRestoreForbidden() *ListEtcdRestoreForbidden {
	return &ListEtcdRestoreForbidden{}
}

/*
ListEtcdRestoreForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListEtcdRestoreForbidden struct {
}

// IsSuccess returns true when this list etcd restore forbidden response has a 2xx status code
func (o *ListEtcdRestoreForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list etcd restore forbidden response has a 3xx status code
func (o *ListEtcdRestoreForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list etcd restore forbidden response has a 4xx status code
func (o *ListEtcdRestoreForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list etcd restore forbidden response has a 5xx status code
func (o *ListEtcdRestoreForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list etcd restore forbidden response a status code equal to that given
func (o *ListEtcdRestoreForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListEtcdRestoreForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores][%d] listEtcdRestoreForbidden ", 403)
}

func (o *ListEtcdRestoreForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores][%d] listEtcdRestoreForbidden ", 403)
}

func (o *ListEtcdRestoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEtcdRestoreDefault creates a ListEtcdRestoreDefault with default headers values
func NewListEtcdRestoreDefault(code int) *ListEtcdRestoreDefault {
	return &ListEtcdRestoreDefault{
		_statusCode: code,
	}
}

/*
ListEtcdRestoreDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListEtcdRestoreDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list etcd restore default response
func (o *ListEtcdRestoreDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list etcd restore default response has a 2xx status code
func (o *ListEtcdRestoreDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list etcd restore default response has a 3xx status code
func (o *ListEtcdRestoreDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list etcd restore default response has a 4xx status code
func (o *ListEtcdRestoreDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list etcd restore default response has a 5xx status code
func (o *ListEtcdRestoreDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list etcd restore default response a status code equal to that given
func (o *ListEtcdRestoreDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListEtcdRestoreDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores][%d] listEtcdRestore default  %+v", o._statusCode, o.Payload)
}

func (o *ListEtcdRestoreDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores][%d] listEtcdRestore default  %+v", o._statusCode, o.Payload)
}

func (o *ListEtcdRestoreDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListEtcdRestoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
