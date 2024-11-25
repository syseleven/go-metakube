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

// GetEtcdRestoreReader is a Reader for the GetEtcdRestore structure.
type GetEtcdRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEtcdRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEtcdRestoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEtcdRestoreUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEtcdRestoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetEtcdRestoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEtcdRestoreOK creates a GetEtcdRestoreOK with default headers values
func NewGetEtcdRestoreOK() *GetEtcdRestoreOK {
	return &GetEtcdRestoreOK{}
}

/*
GetEtcdRestoreOK describes a response with status code 200, with default header values.

EtcdRestore
*/
type GetEtcdRestoreOK struct {
	Payload *models.EtcdRestore
}

// IsSuccess returns true when this get etcd restore o k response has a 2xx status code
func (o *GetEtcdRestoreOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get etcd restore o k response has a 3xx status code
func (o *GetEtcdRestoreOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get etcd restore o k response has a 4xx status code
func (o *GetEtcdRestoreOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get etcd restore o k response has a 5xx status code
func (o *GetEtcdRestoreOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get etcd restore o k response a status code equal to that given
func (o *GetEtcdRestoreOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get etcd restore o k response
func (o *GetEtcdRestoreOK) Code() int {
	return 200
}

func (o *GetEtcdRestoreOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] getEtcdRestoreOK %s", 200, payload)
}

func (o *GetEtcdRestoreOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] getEtcdRestoreOK %s", 200, payload)
}

func (o *GetEtcdRestoreOK) GetPayload() *models.EtcdRestore {
	return o.Payload
}

func (o *GetEtcdRestoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EtcdRestore)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEtcdRestoreUnauthorized creates a GetEtcdRestoreUnauthorized with default headers values
func NewGetEtcdRestoreUnauthorized() *GetEtcdRestoreUnauthorized {
	return &GetEtcdRestoreUnauthorized{}
}

/*
GetEtcdRestoreUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetEtcdRestoreUnauthorized struct {
}

// IsSuccess returns true when this get etcd restore unauthorized response has a 2xx status code
func (o *GetEtcdRestoreUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get etcd restore unauthorized response has a 3xx status code
func (o *GetEtcdRestoreUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get etcd restore unauthorized response has a 4xx status code
func (o *GetEtcdRestoreUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get etcd restore unauthorized response has a 5xx status code
func (o *GetEtcdRestoreUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get etcd restore unauthorized response a status code equal to that given
func (o *GetEtcdRestoreUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get etcd restore unauthorized response
func (o *GetEtcdRestoreUnauthorized) Code() int {
	return 401
}

func (o *GetEtcdRestoreUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] getEtcdRestoreUnauthorized", 401)
}

func (o *GetEtcdRestoreUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] getEtcdRestoreUnauthorized", 401)
}

func (o *GetEtcdRestoreUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEtcdRestoreForbidden creates a GetEtcdRestoreForbidden with default headers values
func NewGetEtcdRestoreForbidden() *GetEtcdRestoreForbidden {
	return &GetEtcdRestoreForbidden{}
}

/*
GetEtcdRestoreForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetEtcdRestoreForbidden struct {
}

// IsSuccess returns true when this get etcd restore forbidden response has a 2xx status code
func (o *GetEtcdRestoreForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get etcd restore forbidden response has a 3xx status code
func (o *GetEtcdRestoreForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get etcd restore forbidden response has a 4xx status code
func (o *GetEtcdRestoreForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get etcd restore forbidden response has a 5xx status code
func (o *GetEtcdRestoreForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get etcd restore forbidden response a status code equal to that given
func (o *GetEtcdRestoreForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get etcd restore forbidden response
func (o *GetEtcdRestoreForbidden) Code() int {
	return 403
}

func (o *GetEtcdRestoreForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] getEtcdRestoreForbidden", 403)
}

func (o *GetEtcdRestoreForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] getEtcdRestoreForbidden", 403)
}

func (o *GetEtcdRestoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEtcdRestoreDefault creates a GetEtcdRestoreDefault with default headers values
func NewGetEtcdRestoreDefault(code int) *GetEtcdRestoreDefault {
	return &GetEtcdRestoreDefault{
		_statusCode: code,
	}
}

/*
GetEtcdRestoreDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetEtcdRestoreDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get etcd restore default response has a 2xx status code
func (o *GetEtcdRestoreDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get etcd restore default response has a 3xx status code
func (o *GetEtcdRestoreDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get etcd restore default response has a 4xx status code
func (o *GetEtcdRestoreDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get etcd restore default response has a 5xx status code
func (o *GetEtcdRestoreDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get etcd restore default response a status code equal to that given
func (o *GetEtcdRestoreDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get etcd restore default response
func (o *GetEtcdRestoreDefault) Code() int {
	return o._statusCode
}

func (o *GetEtcdRestoreDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] getEtcdRestore default %s", o._statusCode, payload)
}

func (o *GetEtcdRestoreDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}][%d] getEtcdRestore default %s", o._statusCode, payload)
}

func (o *GetEtcdRestoreDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetEtcdRestoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
