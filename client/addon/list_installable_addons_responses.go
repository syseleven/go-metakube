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

// ListInstallableAddonsReader is a Reader for the ListInstallableAddons structure.
type ListInstallableAddonsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListInstallableAddonsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListInstallableAddonsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListInstallableAddonsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListInstallableAddonsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListInstallableAddonsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListInstallableAddonsOK creates a ListInstallableAddonsOK with default headers values
func NewListInstallableAddonsOK() *ListInstallableAddonsOK {
	return &ListInstallableAddonsOK{}
}

/*ListInstallableAddonsOK handles this case with default header values.

AccessibleAddons
*/
type ListInstallableAddonsOK struct {
	Payload models.AccessibleAddons
}

func (o *ListInstallableAddonsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/installableaddons][%d] listInstallableAddonsOK  %+v", 200, o.Payload)
}

func (o *ListInstallableAddonsOK) GetPayload() models.AccessibleAddons {
	return o.Payload
}

func (o *ListInstallableAddonsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListInstallableAddonsUnauthorized creates a ListInstallableAddonsUnauthorized with default headers values
func NewListInstallableAddonsUnauthorized() *ListInstallableAddonsUnauthorized {
	return &ListInstallableAddonsUnauthorized{}
}

/*ListInstallableAddonsUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type ListInstallableAddonsUnauthorized struct {
}

func (o *ListInstallableAddonsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/installableaddons][%d] listInstallableAddonsUnauthorized ", 401)
}

func (o *ListInstallableAddonsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListInstallableAddonsForbidden creates a ListInstallableAddonsForbidden with default headers values
func NewListInstallableAddonsForbidden() *ListInstallableAddonsForbidden {
	return &ListInstallableAddonsForbidden{}
}

/*ListInstallableAddonsForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type ListInstallableAddonsForbidden struct {
}

func (o *ListInstallableAddonsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/installableaddons][%d] listInstallableAddonsForbidden ", 403)
}

func (o *ListInstallableAddonsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListInstallableAddonsDefault creates a ListInstallableAddonsDefault with default headers values
func NewListInstallableAddonsDefault(code int) *ListInstallableAddonsDefault {
	return &ListInstallableAddonsDefault{
		_statusCode: code,
	}
}

/*ListInstallableAddonsDefault handles this case with default header values.

errorResponse
*/
type ListInstallableAddonsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list installable addons default response
func (o *ListInstallableAddonsDefault) Code() int {
	return o._statusCode
}

func (o *ListInstallableAddonsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/installableaddons][%d] listInstallableAddons default  %+v", o._statusCode, o.Payload)
}

func (o *ListInstallableAddonsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListInstallableAddonsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
