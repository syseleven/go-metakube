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

// ListEtcdBackupConfigReader is a Reader for the ListEtcdBackupConfig structure.
type ListEtcdBackupConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEtcdBackupConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEtcdBackupConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListEtcdBackupConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListEtcdBackupConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListEtcdBackupConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListEtcdBackupConfigOK creates a ListEtcdBackupConfigOK with default headers values
func NewListEtcdBackupConfigOK() *ListEtcdBackupConfigOK {
	return &ListEtcdBackupConfigOK{}
}

/*
ListEtcdBackupConfigOK describes a response with status code 200, with default header values.

EtcdBackupConfig
*/
type ListEtcdBackupConfigOK struct {
	Payload []*models.EtcdBackupConfig
}

// IsSuccess returns true when this list etcd backup config o k response has a 2xx status code
func (o *ListEtcdBackupConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list etcd backup config o k response has a 3xx status code
func (o *ListEtcdBackupConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list etcd backup config o k response has a 4xx status code
func (o *ListEtcdBackupConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list etcd backup config o k response has a 5xx status code
func (o *ListEtcdBackupConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list etcd backup config o k response a status code equal to that given
func (o *ListEtcdBackupConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list etcd backup config o k response
func (o *ListEtcdBackupConfigOK) Code() int {
	return 200
}

func (o *ListEtcdBackupConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs][%d] listEtcdBackupConfigOK %s", 200, payload)
}

func (o *ListEtcdBackupConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs][%d] listEtcdBackupConfigOK %s", 200, payload)
}

func (o *ListEtcdBackupConfigOK) GetPayload() []*models.EtcdBackupConfig {
	return o.Payload
}

func (o *ListEtcdBackupConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEtcdBackupConfigUnauthorized creates a ListEtcdBackupConfigUnauthorized with default headers values
func NewListEtcdBackupConfigUnauthorized() *ListEtcdBackupConfigUnauthorized {
	return &ListEtcdBackupConfigUnauthorized{}
}

/*
ListEtcdBackupConfigUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListEtcdBackupConfigUnauthorized struct {
}

// IsSuccess returns true when this list etcd backup config unauthorized response has a 2xx status code
func (o *ListEtcdBackupConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list etcd backup config unauthorized response has a 3xx status code
func (o *ListEtcdBackupConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list etcd backup config unauthorized response has a 4xx status code
func (o *ListEtcdBackupConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list etcd backup config unauthorized response has a 5xx status code
func (o *ListEtcdBackupConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list etcd backup config unauthorized response a status code equal to that given
func (o *ListEtcdBackupConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list etcd backup config unauthorized response
func (o *ListEtcdBackupConfigUnauthorized) Code() int {
	return 401
}

func (o *ListEtcdBackupConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs][%d] listEtcdBackupConfigUnauthorized", 401)
}

func (o *ListEtcdBackupConfigUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs][%d] listEtcdBackupConfigUnauthorized", 401)
}

func (o *ListEtcdBackupConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEtcdBackupConfigForbidden creates a ListEtcdBackupConfigForbidden with default headers values
func NewListEtcdBackupConfigForbidden() *ListEtcdBackupConfigForbidden {
	return &ListEtcdBackupConfigForbidden{}
}

/*
ListEtcdBackupConfigForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListEtcdBackupConfigForbidden struct {
}

// IsSuccess returns true when this list etcd backup config forbidden response has a 2xx status code
func (o *ListEtcdBackupConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list etcd backup config forbidden response has a 3xx status code
func (o *ListEtcdBackupConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list etcd backup config forbidden response has a 4xx status code
func (o *ListEtcdBackupConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list etcd backup config forbidden response has a 5xx status code
func (o *ListEtcdBackupConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list etcd backup config forbidden response a status code equal to that given
func (o *ListEtcdBackupConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list etcd backup config forbidden response
func (o *ListEtcdBackupConfigForbidden) Code() int {
	return 403
}

func (o *ListEtcdBackupConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs][%d] listEtcdBackupConfigForbidden", 403)
}

func (o *ListEtcdBackupConfigForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs][%d] listEtcdBackupConfigForbidden", 403)
}

func (o *ListEtcdBackupConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEtcdBackupConfigDefault creates a ListEtcdBackupConfigDefault with default headers values
func NewListEtcdBackupConfigDefault(code int) *ListEtcdBackupConfigDefault {
	return &ListEtcdBackupConfigDefault{
		_statusCode: code,
	}
}

/*
ListEtcdBackupConfigDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListEtcdBackupConfigDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list etcd backup config default response has a 2xx status code
func (o *ListEtcdBackupConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list etcd backup config default response has a 3xx status code
func (o *ListEtcdBackupConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list etcd backup config default response has a 4xx status code
func (o *ListEtcdBackupConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list etcd backup config default response has a 5xx status code
func (o *ListEtcdBackupConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list etcd backup config default response a status code equal to that given
func (o *ListEtcdBackupConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list etcd backup config default response
func (o *ListEtcdBackupConfigDefault) Code() int {
	return o._statusCode
}

func (o *ListEtcdBackupConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs][%d] listEtcdBackupConfig default %s", o._statusCode, payload)
}

func (o *ListEtcdBackupConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs][%d] listEtcdBackupConfig default %s", o._statusCode, payload)
}

func (o *ListEtcdBackupConfigDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListEtcdBackupConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
