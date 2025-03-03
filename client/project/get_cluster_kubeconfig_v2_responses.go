// Code generated by go-swagger; DO NOT EDIT.

package project

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

// GetClusterKubeconfigV2Reader is a Reader for the GetClusterKubeconfigV2 structure.
type GetClusterKubeconfigV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterKubeconfigV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterKubeconfigV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetClusterKubeconfigV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetClusterKubeconfigV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetClusterKubeconfigV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterKubeconfigV2OK creates a GetClusterKubeconfigV2OK with default headers values
func NewGetClusterKubeconfigV2OK() *GetClusterKubeconfigV2OK {
	return &GetClusterKubeconfigV2OK{}
}

/*
GetClusterKubeconfigV2OK describes a response with status code 200, with default header values.

Kubeconfig is a clusters kubeconfig
*/
type GetClusterKubeconfigV2OK struct {
	Payload []uint8
}

// IsSuccess returns true when this get cluster kubeconfig v2 o k response has a 2xx status code
func (o *GetClusterKubeconfigV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster kubeconfig v2 o k response has a 3xx status code
func (o *GetClusterKubeconfigV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster kubeconfig v2 o k response has a 4xx status code
func (o *GetClusterKubeconfigV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster kubeconfig v2 o k response has a 5xx status code
func (o *GetClusterKubeconfigV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster kubeconfig v2 o k response a status code equal to that given
func (o *GetClusterKubeconfigV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cluster kubeconfig v2 o k response
func (o *GetClusterKubeconfigV2OK) Code() int {
	return 200
}

func (o *GetClusterKubeconfigV2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeconfig][%d] getClusterKubeconfigV2OK %s", 200, payload)
}

func (o *GetClusterKubeconfigV2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeconfig][%d] getClusterKubeconfigV2OK %s", 200, payload)
}

func (o *GetClusterKubeconfigV2OK) GetPayload() []uint8 {
	return o.Payload
}

func (o *GetClusterKubeconfigV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterKubeconfigV2Unauthorized creates a GetClusterKubeconfigV2Unauthorized with default headers values
func NewGetClusterKubeconfigV2Unauthorized() *GetClusterKubeconfigV2Unauthorized {
	return &GetClusterKubeconfigV2Unauthorized{}
}

/*
GetClusterKubeconfigV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetClusterKubeconfigV2Unauthorized struct {
}

// IsSuccess returns true when this get cluster kubeconfig v2 unauthorized response has a 2xx status code
func (o *GetClusterKubeconfigV2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster kubeconfig v2 unauthorized response has a 3xx status code
func (o *GetClusterKubeconfigV2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster kubeconfig v2 unauthorized response has a 4xx status code
func (o *GetClusterKubeconfigV2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster kubeconfig v2 unauthorized response has a 5xx status code
func (o *GetClusterKubeconfigV2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster kubeconfig v2 unauthorized response a status code equal to that given
func (o *GetClusterKubeconfigV2Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get cluster kubeconfig v2 unauthorized response
func (o *GetClusterKubeconfigV2Unauthorized) Code() int {
	return 401
}

func (o *GetClusterKubeconfigV2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeconfig][%d] getClusterKubeconfigV2Unauthorized", 401)
}

func (o *GetClusterKubeconfigV2Unauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeconfig][%d] getClusterKubeconfigV2Unauthorized", 401)
}

func (o *GetClusterKubeconfigV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterKubeconfigV2Forbidden creates a GetClusterKubeconfigV2Forbidden with default headers values
func NewGetClusterKubeconfigV2Forbidden() *GetClusterKubeconfigV2Forbidden {
	return &GetClusterKubeconfigV2Forbidden{}
}

/*
GetClusterKubeconfigV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetClusterKubeconfigV2Forbidden struct {
}

// IsSuccess returns true when this get cluster kubeconfig v2 forbidden response has a 2xx status code
func (o *GetClusterKubeconfigV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster kubeconfig v2 forbidden response has a 3xx status code
func (o *GetClusterKubeconfigV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster kubeconfig v2 forbidden response has a 4xx status code
func (o *GetClusterKubeconfigV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster kubeconfig v2 forbidden response has a 5xx status code
func (o *GetClusterKubeconfigV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster kubeconfig v2 forbidden response a status code equal to that given
func (o *GetClusterKubeconfigV2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get cluster kubeconfig v2 forbidden response
func (o *GetClusterKubeconfigV2Forbidden) Code() int {
	return 403
}

func (o *GetClusterKubeconfigV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeconfig][%d] getClusterKubeconfigV2Forbidden", 403)
}

func (o *GetClusterKubeconfigV2Forbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeconfig][%d] getClusterKubeconfigV2Forbidden", 403)
}

func (o *GetClusterKubeconfigV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterKubeconfigV2Default creates a GetClusterKubeconfigV2Default with default headers values
func NewGetClusterKubeconfigV2Default(code int) *GetClusterKubeconfigV2Default {
	return &GetClusterKubeconfigV2Default{
		_statusCode: code,
	}
}

/*
GetClusterKubeconfigV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type GetClusterKubeconfigV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get cluster kubeconfig v2 default response has a 2xx status code
func (o *GetClusterKubeconfigV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get cluster kubeconfig v2 default response has a 3xx status code
func (o *GetClusterKubeconfigV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get cluster kubeconfig v2 default response has a 4xx status code
func (o *GetClusterKubeconfigV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get cluster kubeconfig v2 default response has a 5xx status code
func (o *GetClusterKubeconfigV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get cluster kubeconfig v2 default response a status code equal to that given
func (o *GetClusterKubeconfigV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get cluster kubeconfig v2 default response
func (o *GetClusterKubeconfigV2Default) Code() int {
	return o._statusCode
}

func (o *GetClusterKubeconfigV2Default) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeconfig][%d] getClusterKubeconfigV2 default %s", o._statusCode, payload)
}

func (o *GetClusterKubeconfigV2Default) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeconfig][%d] getClusterKubeconfigV2 default %s", o._statusCode, payload)
}

func (o *GetClusterKubeconfigV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterKubeconfigV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
