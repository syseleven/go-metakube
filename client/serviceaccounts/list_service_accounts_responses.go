// Code generated by go-swagger; DO NOT EDIT.

package serviceaccounts

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

// ListServiceAccountsReader is a Reader for the ListServiceAccounts structure.
type ListServiceAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServiceAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServiceAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListServiceAccountsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListServiceAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListServiceAccountsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListServiceAccountsOK creates a ListServiceAccountsOK with default headers values
func NewListServiceAccountsOK() *ListServiceAccountsOK {
	return &ListServiceAccountsOK{}
}

/*
ListServiceAccountsOK describes a response with status code 200, with default header values.

ServiceAccount
*/
type ListServiceAccountsOK struct {
	Payload []*models.ServiceAccount
}

// IsSuccess returns true when this list service accounts o k response has a 2xx status code
func (o *ListServiceAccountsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list service accounts o k response has a 3xx status code
func (o *ListServiceAccountsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list service accounts o k response has a 4xx status code
func (o *ListServiceAccountsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list service accounts o k response has a 5xx status code
func (o *ListServiceAccountsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list service accounts o k response a status code equal to that given
func (o *ListServiceAccountsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list service accounts o k response
func (o *ListServiceAccountsOK) Code() int {
	return 200
}

func (o *ListServiceAccountsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/serviceaccounts][%d] listServiceAccountsOK %s", 200, payload)
}

func (o *ListServiceAccountsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/serviceaccounts][%d] listServiceAccountsOK %s", 200, payload)
}

func (o *ListServiceAccountsOK) GetPayload() []*models.ServiceAccount {
	return o.Payload
}

func (o *ListServiceAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceAccountsUnauthorized creates a ListServiceAccountsUnauthorized with default headers values
func NewListServiceAccountsUnauthorized() *ListServiceAccountsUnauthorized {
	return &ListServiceAccountsUnauthorized{}
}

/*
ListServiceAccountsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListServiceAccountsUnauthorized struct {
}

// IsSuccess returns true when this list service accounts unauthorized response has a 2xx status code
func (o *ListServiceAccountsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list service accounts unauthorized response has a 3xx status code
func (o *ListServiceAccountsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list service accounts unauthorized response has a 4xx status code
func (o *ListServiceAccountsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list service accounts unauthorized response has a 5xx status code
func (o *ListServiceAccountsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list service accounts unauthorized response a status code equal to that given
func (o *ListServiceAccountsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list service accounts unauthorized response
func (o *ListServiceAccountsUnauthorized) Code() int {
	return 401
}

func (o *ListServiceAccountsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/serviceaccounts][%d] listServiceAccountsUnauthorized", 401)
}

func (o *ListServiceAccountsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/serviceaccounts][%d] listServiceAccountsUnauthorized", 401)
}

func (o *ListServiceAccountsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListServiceAccountsForbidden creates a ListServiceAccountsForbidden with default headers values
func NewListServiceAccountsForbidden() *ListServiceAccountsForbidden {
	return &ListServiceAccountsForbidden{}
}

/*
ListServiceAccountsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListServiceAccountsForbidden struct {
}

// IsSuccess returns true when this list service accounts forbidden response has a 2xx status code
func (o *ListServiceAccountsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list service accounts forbidden response has a 3xx status code
func (o *ListServiceAccountsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list service accounts forbidden response has a 4xx status code
func (o *ListServiceAccountsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list service accounts forbidden response has a 5xx status code
func (o *ListServiceAccountsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list service accounts forbidden response a status code equal to that given
func (o *ListServiceAccountsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list service accounts forbidden response
func (o *ListServiceAccountsForbidden) Code() int {
	return 403
}

func (o *ListServiceAccountsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/serviceaccounts][%d] listServiceAccountsForbidden", 403)
}

func (o *ListServiceAccountsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/serviceaccounts][%d] listServiceAccountsForbidden", 403)
}

func (o *ListServiceAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListServiceAccountsDefault creates a ListServiceAccountsDefault with default headers values
func NewListServiceAccountsDefault(code int) *ListServiceAccountsDefault {
	return &ListServiceAccountsDefault{
		_statusCode: code,
	}
}

/*
ListServiceAccountsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListServiceAccountsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list service accounts default response has a 2xx status code
func (o *ListServiceAccountsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list service accounts default response has a 3xx status code
func (o *ListServiceAccountsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list service accounts default response has a 4xx status code
func (o *ListServiceAccountsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list service accounts default response has a 5xx status code
func (o *ListServiceAccountsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list service accounts default response a status code equal to that given
func (o *ListServiceAccountsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list service accounts default response
func (o *ListServiceAccountsDefault) Code() int {
	return o._statusCode
}

func (o *ListServiceAccountsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/serviceaccounts][%d] listServiceAccounts default %s", o._statusCode, payload)
}

func (o *ListServiceAccountsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/serviceaccounts][%d] listServiceAccounts default %s", o._statusCode, payload)
}

func (o *ListServiceAccountsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListServiceAccountsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
