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

// GetKubeLoginClusterKubeconfigV2Reader is a Reader for the GetKubeLoginClusterKubeconfigV2 structure.
type GetKubeLoginClusterKubeconfigV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubeLoginClusterKubeconfigV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubeLoginClusterKubeconfigV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetKubeLoginClusterKubeconfigV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKubeLoginClusterKubeconfigV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetKubeLoginClusterKubeconfigV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetKubeLoginClusterKubeconfigV2OK creates a GetKubeLoginClusterKubeconfigV2OK with default headers values
func NewGetKubeLoginClusterKubeconfigV2OK() *GetKubeLoginClusterKubeconfigV2OK {
	return &GetKubeLoginClusterKubeconfigV2OK{}
}

/*
GetKubeLoginClusterKubeconfigV2OK describes a response with status code 200, with default header values.

Kubeconfig is a clusters kubeconfig
*/
type GetKubeLoginClusterKubeconfigV2OK struct {
	Payload []uint8
}

// IsSuccess returns true when this get kube login cluster kubeconfig v2 o k response has a 2xx status code
func (o *GetKubeLoginClusterKubeconfigV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get kube login cluster kubeconfig v2 o k response has a 3xx status code
func (o *GetKubeLoginClusterKubeconfigV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kube login cluster kubeconfig v2 o k response has a 4xx status code
func (o *GetKubeLoginClusterKubeconfigV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kube login cluster kubeconfig v2 o k response has a 5xx status code
func (o *GetKubeLoginClusterKubeconfigV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get kube login cluster kubeconfig v2 o k response a status code equal to that given
func (o *GetKubeLoginClusterKubeconfigV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get kube login cluster kubeconfig v2 o k response
func (o *GetKubeLoginClusterKubeconfigV2OK) Code() int {
	return 200
}

func (o *GetKubeLoginClusterKubeconfigV2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeloginkubeconfig][%d] getKubeLoginClusterKubeconfigV2OK %s", 200, payload)
}

func (o *GetKubeLoginClusterKubeconfigV2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeloginkubeconfig][%d] getKubeLoginClusterKubeconfigV2OK %s", 200, payload)
}

func (o *GetKubeLoginClusterKubeconfigV2OK) GetPayload() []uint8 {
	return o.Payload
}

func (o *GetKubeLoginClusterKubeconfigV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKubeLoginClusterKubeconfigV2Unauthorized creates a GetKubeLoginClusterKubeconfigV2Unauthorized with default headers values
func NewGetKubeLoginClusterKubeconfigV2Unauthorized() *GetKubeLoginClusterKubeconfigV2Unauthorized {
	return &GetKubeLoginClusterKubeconfigV2Unauthorized{}
}

/*
GetKubeLoginClusterKubeconfigV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetKubeLoginClusterKubeconfigV2Unauthorized struct {
}

// IsSuccess returns true when this get kube login cluster kubeconfig v2 unauthorized response has a 2xx status code
func (o *GetKubeLoginClusterKubeconfigV2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kube login cluster kubeconfig v2 unauthorized response has a 3xx status code
func (o *GetKubeLoginClusterKubeconfigV2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kube login cluster kubeconfig v2 unauthorized response has a 4xx status code
func (o *GetKubeLoginClusterKubeconfigV2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kube login cluster kubeconfig v2 unauthorized response has a 5xx status code
func (o *GetKubeLoginClusterKubeconfigV2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get kube login cluster kubeconfig v2 unauthorized response a status code equal to that given
func (o *GetKubeLoginClusterKubeconfigV2Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get kube login cluster kubeconfig v2 unauthorized response
func (o *GetKubeLoginClusterKubeconfigV2Unauthorized) Code() int {
	return 401
}

func (o *GetKubeLoginClusterKubeconfigV2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeloginkubeconfig][%d] getKubeLoginClusterKubeconfigV2Unauthorized", 401)
}

func (o *GetKubeLoginClusterKubeconfigV2Unauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeloginkubeconfig][%d] getKubeLoginClusterKubeconfigV2Unauthorized", 401)
}

func (o *GetKubeLoginClusterKubeconfigV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubeLoginClusterKubeconfigV2Forbidden creates a GetKubeLoginClusterKubeconfigV2Forbidden with default headers values
func NewGetKubeLoginClusterKubeconfigV2Forbidden() *GetKubeLoginClusterKubeconfigV2Forbidden {
	return &GetKubeLoginClusterKubeconfigV2Forbidden{}
}

/*
GetKubeLoginClusterKubeconfigV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetKubeLoginClusterKubeconfigV2Forbidden struct {
}

// IsSuccess returns true when this get kube login cluster kubeconfig v2 forbidden response has a 2xx status code
func (o *GetKubeLoginClusterKubeconfigV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kube login cluster kubeconfig v2 forbidden response has a 3xx status code
func (o *GetKubeLoginClusterKubeconfigV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kube login cluster kubeconfig v2 forbidden response has a 4xx status code
func (o *GetKubeLoginClusterKubeconfigV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kube login cluster kubeconfig v2 forbidden response has a 5xx status code
func (o *GetKubeLoginClusterKubeconfigV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get kube login cluster kubeconfig v2 forbidden response a status code equal to that given
func (o *GetKubeLoginClusterKubeconfigV2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get kube login cluster kubeconfig v2 forbidden response
func (o *GetKubeLoginClusterKubeconfigV2Forbidden) Code() int {
	return 403
}

func (o *GetKubeLoginClusterKubeconfigV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeloginkubeconfig][%d] getKubeLoginClusterKubeconfigV2Forbidden", 403)
}

func (o *GetKubeLoginClusterKubeconfigV2Forbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeloginkubeconfig][%d] getKubeLoginClusterKubeconfigV2Forbidden", 403)
}

func (o *GetKubeLoginClusterKubeconfigV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubeLoginClusterKubeconfigV2Default creates a GetKubeLoginClusterKubeconfigV2Default with default headers values
func NewGetKubeLoginClusterKubeconfigV2Default(code int) *GetKubeLoginClusterKubeconfigV2Default {
	return &GetKubeLoginClusterKubeconfigV2Default{
		_statusCode: code,
	}
}

/*
GetKubeLoginClusterKubeconfigV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type GetKubeLoginClusterKubeconfigV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get kube login cluster kubeconfig v2 default response has a 2xx status code
func (o *GetKubeLoginClusterKubeconfigV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get kube login cluster kubeconfig v2 default response has a 3xx status code
func (o *GetKubeLoginClusterKubeconfigV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get kube login cluster kubeconfig v2 default response has a 4xx status code
func (o *GetKubeLoginClusterKubeconfigV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get kube login cluster kubeconfig v2 default response has a 5xx status code
func (o *GetKubeLoginClusterKubeconfigV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get kube login cluster kubeconfig v2 default response a status code equal to that given
func (o *GetKubeLoginClusterKubeconfigV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get kube login cluster kubeconfig v2 default response
func (o *GetKubeLoginClusterKubeconfigV2Default) Code() int {
	return o._statusCode
}

func (o *GetKubeLoginClusterKubeconfigV2Default) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeloginkubeconfig][%d] getKubeLoginClusterKubeconfigV2 default %s", o._statusCode, payload)
}

func (o *GetKubeLoginClusterKubeconfigV2Default) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeloginkubeconfig][%d] getKubeLoginClusterKubeconfigV2 default %s", o._statusCode, payload)
}

func (o *GetKubeLoginClusterKubeconfigV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetKubeLoginClusterKubeconfigV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
