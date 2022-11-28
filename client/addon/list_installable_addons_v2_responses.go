// Code generated by go-swagger; DO NOT EDIT.

package addon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/go-metakube/models"
)

// ListInstallableAddonsV2Reader is a Reader for the ListInstallableAddonsV2 structure.
type ListInstallableAddonsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListInstallableAddonsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListInstallableAddonsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListInstallableAddonsV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListInstallableAddonsV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListInstallableAddonsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListInstallableAddonsV2OK creates a ListInstallableAddonsV2OK with default headers values
func NewListInstallableAddonsV2OK() *ListInstallableAddonsV2OK {
	return &ListInstallableAddonsV2OK{}
}

/*
ListInstallableAddonsV2OK describes a response with status code 200, with default header values.

AccessibleAddons
*/
type ListInstallableAddonsV2OK struct {
	Payload models.AccessibleAddons
}

// IsSuccess returns true when this list installable addons v2 o k response has a 2xx status code
func (o *ListInstallableAddonsV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list installable addons v2 o k response has a 3xx status code
func (o *ListInstallableAddonsV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list installable addons v2 o k response has a 4xx status code
func (o *ListInstallableAddonsV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list installable addons v2 o k response has a 5xx status code
func (o *ListInstallableAddonsV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this list installable addons v2 o k response a status code equal to that given
func (o *ListInstallableAddonsV2OK) IsCode(code int) bool {
	return code == 200
}

func (o *ListInstallableAddonsV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/installableaddons][%d] listInstallableAddonsV2OK  %+v", 200, o.Payload)
}

func (o *ListInstallableAddonsV2OK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/installableaddons][%d] listInstallableAddonsV2OK  %+v", 200, o.Payload)
}

func (o *ListInstallableAddonsV2OK) GetPayload() models.AccessibleAddons {
	return o.Payload
}

func (o *ListInstallableAddonsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListInstallableAddonsV2Unauthorized creates a ListInstallableAddonsV2Unauthorized with default headers values
func NewListInstallableAddonsV2Unauthorized() *ListInstallableAddonsV2Unauthorized {
	return &ListInstallableAddonsV2Unauthorized{}
}

/*
ListInstallableAddonsV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListInstallableAddonsV2Unauthorized struct {
}

// IsSuccess returns true when this list installable addons v2 unauthorized response has a 2xx status code
func (o *ListInstallableAddonsV2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list installable addons v2 unauthorized response has a 3xx status code
func (o *ListInstallableAddonsV2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list installable addons v2 unauthorized response has a 4xx status code
func (o *ListInstallableAddonsV2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list installable addons v2 unauthorized response has a 5xx status code
func (o *ListInstallableAddonsV2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list installable addons v2 unauthorized response a status code equal to that given
func (o *ListInstallableAddonsV2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListInstallableAddonsV2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/installableaddons][%d] listInstallableAddonsV2Unauthorized ", 401)
}

func (o *ListInstallableAddonsV2Unauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/installableaddons][%d] listInstallableAddonsV2Unauthorized ", 401)
}

func (o *ListInstallableAddonsV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListInstallableAddonsV2Forbidden creates a ListInstallableAddonsV2Forbidden with default headers values
func NewListInstallableAddonsV2Forbidden() *ListInstallableAddonsV2Forbidden {
	return &ListInstallableAddonsV2Forbidden{}
}

/*
ListInstallableAddonsV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListInstallableAddonsV2Forbidden struct {
}

// IsSuccess returns true when this list installable addons v2 forbidden response has a 2xx status code
func (o *ListInstallableAddonsV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list installable addons v2 forbidden response has a 3xx status code
func (o *ListInstallableAddonsV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list installable addons v2 forbidden response has a 4xx status code
func (o *ListInstallableAddonsV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list installable addons v2 forbidden response has a 5xx status code
func (o *ListInstallableAddonsV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list installable addons v2 forbidden response a status code equal to that given
func (o *ListInstallableAddonsV2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListInstallableAddonsV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/installableaddons][%d] listInstallableAddonsV2Forbidden ", 403)
}

func (o *ListInstallableAddonsV2Forbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/installableaddons][%d] listInstallableAddonsV2Forbidden ", 403)
}

func (o *ListInstallableAddonsV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListInstallableAddonsV2Default creates a ListInstallableAddonsV2Default with default headers values
func NewListInstallableAddonsV2Default(code int) *ListInstallableAddonsV2Default {
	return &ListInstallableAddonsV2Default{
		_statusCode: code,
	}
}

/*
ListInstallableAddonsV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type ListInstallableAddonsV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list installable addons v2 default response
func (o *ListInstallableAddonsV2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list installable addons v2 default response has a 2xx status code
func (o *ListInstallableAddonsV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list installable addons v2 default response has a 3xx status code
func (o *ListInstallableAddonsV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list installable addons v2 default response has a 4xx status code
func (o *ListInstallableAddonsV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list installable addons v2 default response has a 5xx status code
func (o *ListInstallableAddonsV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list installable addons v2 default response a status code equal to that given
func (o *ListInstallableAddonsV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListInstallableAddonsV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/installableaddons][%d] listInstallableAddonsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListInstallableAddonsV2Default) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/installableaddons][%d] listInstallableAddonsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListInstallableAddonsV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListInstallableAddonsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
