// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// ListSystemLabelsReader is a Reader for the ListSystemLabels structure.
type ListSystemLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSystemLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSystemLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListSystemLabelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListSystemLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListSystemLabelsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSystemLabelsOK creates a ListSystemLabelsOK with default headers values
func NewListSystemLabelsOK() *ListSystemLabelsOK {
	return &ListSystemLabelsOK{}
}

/*
ListSystemLabelsOK describes a response with status code 200, with default header values.

ResourceLabelMap
*/
type ListSystemLabelsOK struct {
	Payload models.ResourceLabelMap
}

// IsSuccess returns true when this list system labels o k response has a 2xx status code
func (o *ListSystemLabelsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list system labels o k response has a 3xx status code
func (o *ListSystemLabelsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list system labels o k response has a 4xx status code
func (o *ListSystemLabelsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list system labels o k response has a 5xx status code
func (o *ListSystemLabelsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list system labels o k response a status code equal to that given
func (o *ListSystemLabelsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list system labels o k response
func (o *ListSystemLabelsOK) Code() int {
	return 200
}

func (o *ListSystemLabelsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/labels/system][%d] listSystemLabelsOK %s", 200, payload)
}

func (o *ListSystemLabelsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/labels/system][%d] listSystemLabelsOK %s", 200, payload)
}

func (o *ListSystemLabelsOK) GetPayload() models.ResourceLabelMap {
	return o.Payload
}

func (o *ListSystemLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSystemLabelsUnauthorized creates a ListSystemLabelsUnauthorized with default headers values
func NewListSystemLabelsUnauthorized() *ListSystemLabelsUnauthorized {
	return &ListSystemLabelsUnauthorized{}
}

/*
ListSystemLabelsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListSystemLabelsUnauthorized struct {
}

// IsSuccess returns true when this list system labels unauthorized response has a 2xx status code
func (o *ListSystemLabelsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list system labels unauthorized response has a 3xx status code
func (o *ListSystemLabelsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list system labels unauthorized response has a 4xx status code
func (o *ListSystemLabelsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list system labels unauthorized response has a 5xx status code
func (o *ListSystemLabelsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list system labels unauthorized response a status code equal to that given
func (o *ListSystemLabelsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list system labels unauthorized response
func (o *ListSystemLabelsUnauthorized) Code() int {
	return 401
}

func (o *ListSystemLabelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/labels/system][%d] listSystemLabelsUnauthorized", 401)
}

func (o *ListSystemLabelsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/labels/system][%d] listSystemLabelsUnauthorized", 401)
}

func (o *ListSystemLabelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListSystemLabelsForbidden creates a ListSystemLabelsForbidden with default headers values
func NewListSystemLabelsForbidden() *ListSystemLabelsForbidden {
	return &ListSystemLabelsForbidden{}
}

/*
ListSystemLabelsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListSystemLabelsForbidden struct {
}

// IsSuccess returns true when this list system labels forbidden response has a 2xx status code
func (o *ListSystemLabelsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list system labels forbidden response has a 3xx status code
func (o *ListSystemLabelsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list system labels forbidden response has a 4xx status code
func (o *ListSystemLabelsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list system labels forbidden response has a 5xx status code
func (o *ListSystemLabelsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list system labels forbidden response a status code equal to that given
func (o *ListSystemLabelsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list system labels forbidden response
func (o *ListSystemLabelsForbidden) Code() int {
	return 403
}

func (o *ListSystemLabelsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/labels/system][%d] listSystemLabelsForbidden", 403)
}

func (o *ListSystemLabelsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/labels/system][%d] listSystemLabelsForbidden", 403)
}

func (o *ListSystemLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListSystemLabelsDefault creates a ListSystemLabelsDefault with default headers values
func NewListSystemLabelsDefault(code int) *ListSystemLabelsDefault {
	return &ListSystemLabelsDefault{
		_statusCode: code,
	}
}

/*
ListSystemLabelsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListSystemLabelsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list system labels default response has a 2xx status code
func (o *ListSystemLabelsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list system labels default response has a 3xx status code
func (o *ListSystemLabelsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list system labels default response has a 4xx status code
func (o *ListSystemLabelsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list system labels default response has a 5xx status code
func (o *ListSystemLabelsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list system labels default response a status code equal to that given
func (o *ListSystemLabelsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list system labels default response
func (o *ListSystemLabelsDefault) Code() int {
	return o._statusCode
}

func (o *ListSystemLabelsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/labels/system][%d] listSystemLabels default %s", o._statusCode, payload)
}

func (o *ListSystemLabelsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/labels/system][%d] listSystemLabels default %s", o._statusCode, payload)
}

func (o *ListSystemLabelsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListSystemLabelsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
