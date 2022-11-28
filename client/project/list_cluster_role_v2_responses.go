// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/go-metakube/models"
)

// ListClusterRoleV2Reader is a Reader for the ListClusterRoleV2 structure.
type ListClusterRoleV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClusterRoleV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListClusterRoleV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListClusterRoleV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListClusterRoleV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListClusterRoleV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListClusterRoleV2OK creates a ListClusterRoleV2OK with default headers values
func NewListClusterRoleV2OK() *ListClusterRoleV2OK {
	return &ListClusterRoleV2OK{}
}

/*
ListClusterRoleV2OK describes a response with status code 200, with default header values.

ClusterRole
*/
type ListClusterRoleV2OK struct {
	Payload []*models.ClusterRole
}

// IsSuccess returns true when this list cluster role v2 o k response has a 2xx status code
func (o *ListClusterRoleV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list cluster role v2 o k response has a 3xx status code
func (o *ListClusterRoleV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list cluster role v2 o k response has a 4xx status code
func (o *ListClusterRoleV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list cluster role v2 o k response has a 5xx status code
func (o *ListClusterRoleV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this list cluster role v2 o k response a status code equal to that given
func (o *ListClusterRoleV2OK) IsCode(code int) bool {
	return code == 200
}

func (o *ListClusterRoleV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles][%d] listClusterRoleV2OK  %+v", 200, o.Payload)
}

func (o *ListClusterRoleV2OK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles][%d] listClusterRoleV2OK  %+v", 200, o.Payload)
}

func (o *ListClusterRoleV2OK) GetPayload() []*models.ClusterRole {
	return o.Payload
}

func (o *ListClusterRoleV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClusterRoleV2Unauthorized creates a ListClusterRoleV2Unauthorized with default headers values
func NewListClusterRoleV2Unauthorized() *ListClusterRoleV2Unauthorized {
	return &ListClusterRoleV2Unauthorized{}
}

/*
ListClusterRoleV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListClusterRoleV2Unauthorized struct {
}

// IsSuccess returns true when this list cluster role v2 unauthorized response has a 2xx status code
func (o *ListClusterRoleV2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list cluster role v2 unauthorized response has a 3xx status code
func (o *ListClusterRoleV2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list cluster role v2 unauthorized response has a 4xx status code
func (o *ListClusterRoleV2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list cluster role v2 unauthorized response has a 5xx status code
func (o *ListClusterRoleV2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list cluster role v2 unauthorized response a status code equal to that given
func (o *ListClusterRoleV2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListClusterRoleV2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles][%d] listClusterRoleV2Unauthorized ", 401)
}

func (o *ListClusterRoleV2Unauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles][%d] listClusterRoleV2Unauthorized ", 401)
}

func (o *ListClusterRoleV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClusterRoleV2Forbidden creates a ListClusterRoleV2Forbidden with default headers values
func NewListClusterRoleV2Forbidden() *ListClusterRoleV2Forbidden {
	return &ListClusterRoleV2Forbidden{}
}

/*
ListClusterRoleV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListClusterRoleV2Forbidden struct {
}

// IsSuccess returns true when this list cluster role v2 forbidden response has a 2xx status code
func (o *ListClusterRoleV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list cluster role v2 forbidden response has a 3xx status code
func (o *ListClusterRoleV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list cluster role v2 forbidden response has a 4xx status code
func (o *ListClusterRoleV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list cluster role v2 forbidden response has a 5xx status code
func (o *ListClusterRoleV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list cluster role v2 forbidden response a status code equal to that given
func (o *ListClusterRoleV2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListClusterRoleV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles][%d] listClusterRoleV2Forbidden ", 403)
}

func (o *ListClusterRoleV2Forbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles][%d] listClusterRoleV2Forbidden ", 403)
}

func (o *ListClusterRoleV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClusterRoleV2Default creates a ListClusterRoleV2Default with default headers values
func NewListClusterRoleV2Default(code int) *ListClusterRoleV2Default {
	return &ListClusterRoleV2Default{
		_statusCode: code,
	}
}

/*
ListClusterRoleV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type ListClusterRoleV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list cluster role v2 default response
func (o *ListClusterRoleV2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list cluster role v2 default response has a 2xx status code
func (o *ListClusterRoleV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list cluster role v2 default response has a 3xx status code
func (o *ListClusterRoleV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list cluster role v2 default response has a 4xx status code
func (o *ListClusterRoleV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list cluster role v2 default response has a 5xx status code
func (o *ListClusterRoleV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list cluster role v2 default response a status code equal to that given
func (o *ListClusterRoleV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListClusterRoleV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles][%d] listClusterRoleV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListClusterRoleV2Default) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles][%d] listClusterRoleV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListClusterRoleV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListClusterRoleV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
