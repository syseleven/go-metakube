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

// ListProjectEtcdBackupConfigReader is a Reader for the ListProjectEtcdBackupConfig structure.
type ListProjectEtcdBackupConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectEtcdBackupConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectEtcdBackupConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListProjectEtcdBackupConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListProjectEtcdBackupConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListProjectEtcdBackupConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectEtcdBackupConfigOK creates a ListProjectEtcdBackupConfigOK with default headers values
func NewListProjectEtcdBackupConfigOK() *ListProjectEtcdBackupConfigOK {
	return &ListProjectEtcdBackupConfigOK{}
}

/*
ListProjectEtcdBackupConfigOK describes a response with status code 200, with default header values.

EtcdBackupConfig
*/
type ListProjectEtcdBackupConfigOK struct {
	Payload []*models.EtcdBackupConfig
}

// IsSuccess returns true when this list project etcd backup config o k response has a 2xx status code
func (o *ListProjectEtcdBackupConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project etcd backup config o k response has a 3xx status code
func (o *ListProjectEtcdBackupConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project etcd backup config o k response has a 4xx status code
func (o *ListProjectEtcdBackupConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project etcd backup config o k response has a 5xx status code
func (o *ListProjectEtcdBackupConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project etcd backup config o k response a status code equal to that given
func (o *ListProjectEtcdBackupConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list project etcd backup config o k response
func (o *ListProjectEtcdBackupConfigOK) Code() int {
	return 200
}

func (o *ListProjectEtcdBackupConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/etcdbackupconfigs][%d] listProjectEtcdBackupConfigOK %s", 200, payload)
}

func (o *ListProjectEtcdBackupConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/etcdbackupconfigs][%d] listProjectEtcdBackupConfigOK %s", 200, payload)
}

func (o *ListProjectEtcdBackupConfigOK) GetPayload() []*models.EtcdBackupConfig {
	return o.Payload
}

func (o *ListProjectEtcdBackupConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectEtcdBackupConfigUnauthorized creates a ListProjectEtcdBackupConfigUnauthorized with default headers values
func NewListProjectEtcdBackupConfigUnauthorized() *ListProjectEtcdBackupConfigUnauthorized {
	return &ListProjectEtcdBackupConfigUnauthorized{}
}

/*
ListProjectEtcdBackupConfigUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListProjectEtcdBackupConfigUnauthorized struct {
}

// IsSuccess returns true when this list project etcd backup config unauthorized response has a 2xx status code
func (o *ListProjectEtcdBackupConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project etcd backup config unauthorized response has a 3xx status code
func (o *ListProjectEtcdBackupConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project etcd backup config unauthorized response has a 4xx status code
func (o *ListProjectEtcdBackupConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project etcd backup config unauthorized response has a 5xx status code
func (o *ListProjectEtcdBackupConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list project etcd backup config unauthorized response a status code equal to that given
func (o *ListProjectEtcdBackupConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list project etcd backup config unauthorized response
func (o *ListProjectEtcdBackupConfigUnauthorized) Code() int {
	return 401
}

func (o *ListProjectEtcdBackupConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/etcdbackupconfigs][%d] listProjectEtcdBackupConfigUnauthorized", 401)
}

func (o *ListProjectEtcdBackupConfigUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/etcdbackupconfigs][%d] listProjectEtcdBackupConfigUnauthorized", 401)
}

func (o *ListProjectEtcdBackupConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectEtcdBackupConfigForbidden creates a ListProjectEtcdBackupConfigForbidden with default headers values
func NewListProjectEtcdBackupConfigForbidden() *ListProjectEtcdBackupConfigForbidden {
	return &ListProjectEtcdBackupConfigForbidden{}
}

/*
ListProjectEtcdBackupConfigForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListProjectEtcdBackupConfigForbidden struct {
}

// IsSuccess returns true when this list project etcd backup config forbidden response has a 2xx status code
func (o *ListProjectEtcdBackupConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project etcd backup config forbidden response has a 3xx status code
func (o *ListProjectEtcdBackupConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project etcd backup config forbidden response has a 4xx status code
func (o *ListProjectEtcdBackupConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project etcd backup config forbidden response has a 5xx status code
func (o *ListProjectEtcdBackupConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list project etcd backup config forbidden response a status code equal to that given
func (o *ListProjectEtcdBackupConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list project etcd backup config forbidden response
func (o *ListProjectEtcdBackupConfigForbidden) Code() int {
	return 403
}

func (o *ListProjectEtcdBackupConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/etcdbackupconfigs][%d] listProjectEtcdBackupConfigForbidden", 403)
}

func (o *ListProjectEtcdBackupConfigForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/etcdbackupconfigs][%d] listProjectEtcdBackupConfigForbidden", 403)
}

func (o *ListProjectEtcdBackupConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListProjectEtcdBackupConfigDefault creates a ListProjectEtcdBackupConfigDefault with default headers values
func NewListProjectEtcdBackupConfigDefault(code int) *ListProjectEtcdBackupConfigDefault {
	return &ListProjectEtcdBackupConfigDefault{
		_statusCode: code,
	}
}

/*
ListProjectEtcdBackupConfigDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectEtcdBackupConfigDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list project etcd backup config default response has a 2xx status code
func (o *ListProjectEtcdBackupConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project etcd backup config default response has a 3xx status code
func (o *ListProjectEtcdBackupConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project etcd backup config default response has a 4xx status code
func (o *ListProjectEtcdBackupConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project etcd backup config default response has a 5xx status code
func (o *ListProjectEtcdBackupConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project etcd backup config default response a status code equal to that given
func (o *ListProjectEtcdBackupConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list project etcd backup config default response
func (o *ListProjectEtcdBackupConfigDefault) Code() int {
	return o._statusCode
}

func (o *ListProjectEtcdBackupConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/etcdbackupconfigs][%d] listProjectEtcdBackupConfig default %s", o._statusCode, payload)
}

func (o *ListProjectEtcdBackupConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/etcdbackupconfigs][%d] listProjectEtcdBackupConfig default %s", o._statusCode, payload)
}

func (o *ListProjectEtcdBackupConfigDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectEtcdBackupConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
