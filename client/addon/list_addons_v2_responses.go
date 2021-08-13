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

// ListAddonsV2Reader is a Reader for the ListAddonsV2 structure.
type ListAddonsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAddonsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAddonsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAddonsV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAddonsV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListAddonsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAddonsV2OK creates a ListAddonsV2OK with default headers values
func NewListAddonsV2OK() *ListAddonsV2OK {
	return &ListAddonsV2OK{}
}

/*ListAddonsV2OK handles this case with default header values.

Addon
*/
type ListAddonsV2OK struct {
	Payload []*models.Addon
}

func (o *ListAddonsV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/addons][%d] listAddonsV2OK  %+v", 200, o.Payload)
}

func (o *ListAddonsV2OK) GetPayload() []*models.Addon {
	return o.Payload
}

func (o *ListAddonsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAddonsV2Unauthorized creates a ListAddonsV2Unauthorized with default headers values
func NewListAddonsV2Unauthorized() *ListAddonsV2Unauthorized {
	return &ListAddonsV2Unauthorized{}
}

/*ListAddonsV2Unauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type ListAddonsV2Unauthorized struct {
}

func (o *ListAddonsV2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/addons][%d] listAddonsV2Unauthorized ", 401)
}

func (o *ListAddonsV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAddonsV2Forbidden creates a ListAddonsV2Forbidden with default headers values
func NewListAddonsV2Forbidden() *ListAddonsV2Forbidden {
	return &ListAddonsV2Forbidden{}
}

/*ListAddonsV2Forbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type ListAddonsV2Forbidden struct {
}

func (o *ListAddonsV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/addons][%d] listAddonsV2Forbidden ", 403)
}

func (o *ListAddonsV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAddonsV2Default creates a ListAddonsV2Default with default headers values
func NewListAddonsV2Default(code int) *ListAddonsV2Default {
	return &ListAddonsV2Default{
		_statusCode: code,
	}
}

/*ListAddonsV2Default handles this case with default header values.

errorResponse
*/
type ListAddonsV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list addons v2 default response
func (o *ListAddonsV2Default) Code() int {
	return o._statusCode
}

func (o *ListAddonsV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/addons][%d] listAddonsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListAddonsV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAddonsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}