// Code generated by go-swagger; DO NOT EDIT.

package etcdbackupconfig

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

// GetEtcdBackupConfigReader is a Reader for the GetEtcdBackupConfig structure.
type GetEtcdBackupConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEtcdBackupConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEtcdBackupConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEtcdBackupConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEtcdBackupConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetEtcdBackupConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEtcdBackupConfigOK creates a GetEtcdBackupConfigOK with default headers values
func NewGetEtcdBackupConfigOK() *GetEtcdBackupConfigOK {
	return &GetEtcdBackupConfigOK{}
}

/*
GetEtcdBackupConfigOK describes a response with status code 200, with default header values.

EtcdBackupConfig
*/
type GetEtcdBackupConfigOK struct {
	Payload *models.EtcdBackupConfig
}

// IsSuccess returns true when this get etcd backup config o k response has a 2xx status code
func (o *GetEtcdBackupConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get etcd backup config o k response has a 3xx status code
func (o *GetEtcdBackupConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get etcd backup config o k response has a 4xx status code
func (o *GetEtcdBackupConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get etcd backup config o k response has a 5xx status code
func (o *GetEtcdBackupConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get etcd backup config o k response a status code equal to that given
func (o *GetEtcdBackupConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get etcd backup config o k response
func (o *GetEtcdBackupConfigOK) Code() int {
	return 200
}

func (o *GetEtcdBackupConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}][%d] getEtcdBackupConfigOK %s", 200, payload)
}

func (o *GetEtcdBackupConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}][%d] getEtcdBackupConfigOK %s", 200, payload)
}

func (o *GetEtcdBackupConfigOK) GetPayload() *models.EtcdBackupConfig {
	return o.Payload
}

func (o *GetEtcdBackupConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EtcdBackupConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEtcdBackupConfigUnauthorized creates a GetEtcdBackupConfigUnauthorized with default headers values
func NewGetEtcdBackupConfigUnauthorized() *GetEtcdBackupConfigUnauthorized {
	return &GetEtcdBackupConfigUnauthorized{}
}

/*
GetEtcdBackupConfigUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetEtcdBackupConfigUnauthorized struct {
}

// IsSuccess returns true when this get etcd backup config unauthorized response has a 2xx status code
func (o *GetEtcdBackupConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get etcd backup config unauthorized response has a 3xx status code
func (o *GetEtcdBackupConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get etcd backup config unauthorized response has a 4xx status code
func (o *GetEtcdBackupConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get etcd backup config unauthorized response has a 5xx status code
func (o *GetEtcdBackupConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get etcd backup config unauthorized response a status code equal to that given
func (o *GetEtcdBackupConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get etcd backup config unauthorized response
func (o *GetEtcdBackupConfigUnauthorized) Code() int {
	return 401
}

func (o *GetEtcdBackupConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}][%d] getEtcdBackupConfigUnauthorized", 401)
}

func (o *GetEtcdBackupConfigUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}][%d] getEtcdBackupConfigUnauthorized", 401)
}

func (o *GetEtcdBackupConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEtcdBackupConfigForbidden creates a GetEtcdBackupConfigForbidden with default headers values
func NewGetEtcdBackupConfigForbidden() *GetEtcdBackupConfigForbidden {
	return &GetEtcdBackupConfigForbidden{}
}

/*
GetEtcdBackupConfigForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetEtcdBackupConfigForbidden struct {
}

// IsSuccess returns true when this get etcd backup config forbidden response has a 2xx status code
func (o *GetEtcdBackupConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get etcd backup config forbidden response has a 3xx status code
func (o *GetEtcdBackupConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get etcd backup config forbidden response has a 4xx status code
func (o *GetEtcdBackupConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get etcd backup config forbidden response has a 5xx status code
func (o *GetEtcdBackupConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get etcd backup config forbidden response a status code equal to that given
func (o *GetEtcdBackupConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get etcd backup config forbidden response
func (o *GetEtcdBackupConfigForbidden) Code() int {
	return 403
}

func (o *GetEtcdBackupConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}][%d] getEtcdBackupConfigForbidden", 403)
}

func (o *GetEtcdBackupConfigForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}][%d] getEtcdBackupConfigForbidden", 403)
}

func (o *GetEtcdBackupConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEtcdBackupConfigDefault creates a GetEtcdBackupConfigDefault with default headers values
func NewGetEtcdBackupConfigDefault(code int) *GetEtcdBackupConfigDefault {
	return &GetEtcdBackupConfigDefault{
		_statusCode: code,
	}
}

/*
GetEtcdBackupConfigDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetEtcdBackupConfigDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get etcd backup config default response has a 2xx status code
func (o *GetEtcdBackupConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get etcd backup config default response has a 3xx status code
func (o *GetEtcdBackupConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get etcd backup config default response has a 4xx status code
func (o *GetEtcdBackupConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get etcd backup config default response has a 5xx status code
func (o *GetEtcdBackupConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get etcd backup config default response a status code equal to that given
func (o *GetEtcdBackupConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get etcd backup config default response
func (o *GetEtcdBackupConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetEtcdBackupConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}][%d] getEtcdBackupConfig default %s", o._statusCode, payload)
}

func (o *GetEtcdBackupConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}][%d] getEtcdBackupConfig default %s", o._statusCode, payload)
}

func (o *GetEtcdBackupConfigDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetEtcdBackupConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
