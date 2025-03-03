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

// ListMachineDeploymentNodesReader is a Reader for the ListMachineDeploymentNodes structure.
type ListMachineDeploymentNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListMachineDeploymentNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListMachineDeploymentNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListMachineDeploymentNodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListMachineDeploymentNodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListMachineDeploymentNodesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListMachineDeploymentNodesOK creates a ListMachineDeploymentNodesOK with default headers values
func NewListMachineDeploymentNodesOK() *ListMachineDeploymentNodesOK {
	return &ListMachineDeploymentNodesOK{}
}

/*
ListMachineDeploymentNodesOK describes a response with status code 200, with default header values.

Node
*/
type ListMachineDeploymentNodesOK struct {
	Payload []*models.Node
}

// IsSuccess returns true when this list machine deployment nodes o k response has a 2xx status code
func (o *ListMachineDeploymentNodesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list machine deployment nodes o k response has a 3xx status code
func (o *ListMachineDeploymentNodesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list machine deployment nodes o k response has a 4xx status code
func (o *ListMachineDeploymentNodesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list machine deployment nodes o k response has a 5xx status code
func (o *ListMachineDeploymentNodesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list machine deployment nodes o k response a status code equal to that given
func (o *ListMachineDeploymentNodesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list machine deployment nodes o k response
func (o *ListMachineDeploymentNodesOK) Code() int {
	return 200
}

func (o *ListMachineDeploymentNodesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes][%d] listMachineDeploymentNodesOK %s", 200, payload)
}

func (o *ListMachineDeploymentNodesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes][%d] listMachineDeploymentNodesOK %s", 200, payload)
}

func (o *ListMachineDeploymentNodesOK) GetPayload() []*models.Node {
	return o.Payload
}

func (o *ListMachineDeploymentNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListMachineDeploymentNodesUnauthorized creates a ListMachineDeploymentNodesUnauthorized with default headers values
func NewListMachineDeploymentNodesUnauthorized() *ListMachineDeploymentNodesUnauthorized {
	return &ListMachineDeploymentNodesUnauthorized{}
}

/*
ListMachineDeploymentNodesUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListMachineDeploymentNodesUnauthorized struct {
}

// IsSuccess returns true when this list machine deployment nodes unauthorized response has a 2xx status code
func (o *ListMachineDeploymentNodesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list machine deployment nodes unauthorized response has a 3xx status code
func (o *ListMachineDeploymentNodesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list machine deployment nodes unauthorized response has a 4xx status code
func (o *ListMachineDeploymentNodesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list machine deployment nodes unauthorized response has a 5xx status code
func (o *ListMachineDeploymentNodesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list machine deployment nodes unauthorized response a status code equal to that given
func (o *ListMachineDeploymentNodesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list machine deployment nodes unauthorized response
func (o *ListMachineDeploymentNodesUnauthorized) Code() int {
	return 401
}

func (o *ListMachineDeploymentNodesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes][%d] listMachineDeploymentNodesUnauthorized", 401)
}

func (o *ListMachineDeploymentNodesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes][%d] listMachineDeploymentNodesUnauthorized", 401)
}

func (o *ListMachineDeploymentNodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListMachineDeploymentNodesForbidden creates a ListMachineDeploymentNodesForbidden with default headers values
func NewListMachineDeploymentNodesForbidden() *ListMachineDeploymentNodesForbidden {
	return &ListMachineDeploymentNodesForbidden{}
}

/*
ListMachineDeploymentNodesForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListMachineDeploymentNodesForbidden struct {
}

// IsSuccess returns true when this list machine deployment nodes forbidden response has a 2xx status code
func (o *ListMachineDeploymentNodesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list machine deployment nodes forbidden response has a 3xx status code
func (o *ListMachineDeploymentNodesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list machine deployment nodes forbidden response has a 4xx status code
func (o *ListMachineDeploymentNodesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list machine deployment nodes forbidden response has a 5xx status code
func (o *ListMachineDeploymentNodesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list machine deployment nodes forbidden response a status code equal to that given
func (o *ListMachineDeploymentNodesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list machine deployment nodes forbidden response
func (o *ListMachineDeploymentNodesForbidden) Code() int {
	return 403
}

func (o *ListMachineDeploymentNodesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes][%d] listMachineDeploymentNodesForbidden", 403)
}

func (o *ListMachineDeploymentNodesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes][%d] listMachineDeploymentNodesForbidden", 403)
}

func (o *ListMachineDeploymentNodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListMachineDeploymentNodesDefault creates a ListMachineDeploymentNodesDefault with default headers values
func NewListMachineDeploymentNodesDefault(code int) *ListMachineDeploymentNodesDefault {
	return &ListMachineDeploymentNodesDefault{
		_statusCode: code,
	}
}

/*
ListMachineDeploymentNodesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListMachineDeploymentNodesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list machine deployment nodes default response has a 2xx status code
func (o *ListMachineDeploymentNodesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list machine deployment nodes default response has a 3xx status code
func (o *ListMachineDeploymentNodesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list machine deployment nodes default response has a 4xx status code
func (o *ListMachineDeploymentNodesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list machine deployment nodes default response has a 5xx status code
func (o *ListMachineDeploymentNodesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list machine deployment nodes default response a status code equal to that given
func (o *ListMachineDeploymentNodesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list machine deployment nodes default response
func (o *ListMachineDeploymentNodesDefault) Code() int {
	return o._statusCode
}

func (o *ListMachineDeploymentNodesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes][%d] listMachineDeploymentNodes default %s", o._statusCode, payload)
}

func (o *ListMachineDeploymentNodesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes][%d] listMachineDeploymentNodes default %s", o._statusCode, payload)
}

func (o *ListMachineDeploymentNodesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListMachineDeploymentNodesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
