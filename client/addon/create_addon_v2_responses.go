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

// CreateAddonV2Reader is a Reader for the CreateAddonV2 structure.
type CreateAddonV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAddonV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAddonV2Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateAddonV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAddonV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateAddonV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAddonV2Created creates a CreateAddonV2Created with default headers values
func NewCreateAddonV2Created() *CreateAddonV2Created {
	return &CreateAddonV2Created{}
}

/*
CreateAddonV2Created describes a response with status code 201, with default header values.

Addon
*/
type CreateAddonV2Created struct {
	Payload *models.Addon
}

// IsSuccess returns true when this create addon v2 created response has a 2xx status code
func (o *CreateAddonV2Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create addon v2 created response has a 3xx status code
func (o *CreateAddonV2Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create addon v2 created response has a 4xx status code
func (o *CreateAddonV2Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create addon v2 created response has a 5xx status code
func (o *CreateAddonV2Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create addon v2 created response a status code equal to that given
func (o *CreateAddonV2Created) IsCode(code int) bool {
	return code == 201
}

func (o *CreateAddonV2Created) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/addons][%d] createAddonV2Created  %+v", 201, o.Payload)
}

func (o *CreateAddonV2Created) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/addons][%d] createAddonV2Created  %+v", 201, o.Payload)
}

func (o *CreateAddonV2Created) GetPayload() *models.Addon {
	return o.Payload
}

func (o *CreateAddonV2Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Addon)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAddonV2Unauthorized creates a CreateAddonV2Unauthorized with default headers values
func NewCreateAddonV2Unauthorized() *CreateAddonV2Unauthorized {
	return &CreateAddonV2Unauthorized{}
}

/*
CreateAddonV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateAddonV2Unauthorized struct {
}

// IsSuccess returns true when this create addon v2 unauthorized response has a 2xx status code
func (o *CreateAddonV2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create addon v2 unauthorized response has a 3xx status code
func (o *CreateAddonV2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create addon v2 unauthorized response has a 4xx status code
func (o *CreateAddonV2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create addon v2 unauthorized response has a 5xx status code
func (o *CreateAddonV2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create addon v2 unauthorized response a status code equal to that given
func (o *CreateAddonV2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateAddonV2Unauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/addons][%d] createAddonV2Unauthorized ", 401)
}

func (o *CreateAddonV2Unauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/addons][%d] createAddonV2Unauthorized ", 401)
}

func (o *CreateAddonV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAddonV2Forbidden creates a CreateAddonV2Forbidden with default headers values
func NewCreateAddonV2Forbidden() *CreateAddonV2Forbidden {
	return &CreateAddonV2Forbidden{}
}

/*
CreateAddonV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateAddonV2Forbidden struct {
}

// IsSuccess returns true when this create addon v2 forbidden response has a 2xx status code
func (o *CreateAddonV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create addon v2 forbidden response has a 3xx status code
func (o *CreateAddonV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create addon v2 forbidden response has a 4xx status code
func (o *CreateAddonV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create addon v2 forbidden response has a 5xx status code
func (o *CreateAddonV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create addon v2 forbidden response a status code equal to that given
func (o *CreateAddonV2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateAddonV2Forbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/addons][%d] createAddonV2Forbidden ", 403)
}

func (o *CreateAddonV2Forbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/addons][%d] createAddonV2Forbidden ", 403)
}

func (o *CreateAddonV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAddonV2Default creates a CreateAddonV2Default with default headers values
func NewCreateAddonV2Default(code int) *CreateAddonV2Default {
	return &CreateAddonV2Default{
		_statusCode: code,
	}
}

/*
CreateAddonV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type CreateAddonV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create addon v2 default response
func (o *CreateAddonV2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create addon v2 default response has a 2xx status code
func (o *CreateAddonV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create addon v2 default response has a 3xx status code
func (o *CreateAddonV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create addon v2 default response has a 4xx status code
func (o *CreateAddonV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create addon v2 default response has a 5xx status code
func (o *CreateAddonV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create addon v2 default response a status code equal to that given
func (o *CreateAddonV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateAddonV2Default) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/addons][%d] createAddonV2 default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAddonV2Default) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/addons][%d] createAddonV2 default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAddonV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateAddonV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
