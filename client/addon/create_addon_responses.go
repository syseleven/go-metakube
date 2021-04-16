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

// CreateAddonReader is a Reader for the CreateAddon structure.
type CreateAddonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAddonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAddonCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateAddonUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAddonForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateAddonDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAddonCreated creates a CreateAddonCreated with default headers values
func NewCreateAddonCreated() *CreateAddonCreated {
	return &CreateAddonCreated{}
}

/*CreateAddonCreated handles this case with default header values.

Addon
*/
type CreateAddonCreated struct {
	Payload *models.Addon
}

func (o *CreateAddonCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons][%d] createAddonCreated  %+v", 201, o.Payload)
}

func (o *CreateAddonCreated) GetPayload() *models.Addon {
	return o.Payload
}

func (o *CreateAddonCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Addon)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAddonUnauthorized creates a CreateAddonUnauthorized with default headers values
func NewCreateAddonUnauthorized() *CreateAddonUnauthorized {
	return &CreateAddonUnauthorized{}
}

/*CreateAddonUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type CreateAddonUnauthorized struct {
}

func (o *CreateAddonUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons][%d] createAddonUnauthorized ", 401)
}

func (o *CreateAddonUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAddonForbidden creates a CreateAddonForbidden with default headers values
func NewCreateAddonForbidden() *CreateAddonForbidden {
	return &CreateAddonForbidden{}
}

/*CreateAddonForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type CreateAddonForbidden struct {
}

func (o *CreateAddonForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons][%d] createAddonForbidden ", 403)
}

func (o *CreateAddonForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAddonDefault creates a CreateAddonDefault with default headers values
func NewCreateAddonDefault(code int) *CreateAddonDefault {
	return &CreateAddonDefault{
		_statusCode: code,
	}
}

/*CreateAddonDefault handles this case with default header values.

errorResponse
*/
type CreateAddonDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create addon default response
func (o *CreateAddonDefault) Code() int {
	return o._statusCode
}

func (o *CreateAddonDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons][%d] createAddon default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAddonDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateAddonDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
