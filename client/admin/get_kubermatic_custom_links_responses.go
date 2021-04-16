// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/go-metakube/models"
)

// GetKubermaticCustomLinksReader is a Reader for the GetKubermaticCustomLinks structure.
type GetKubermaticCustomLinksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubermaticCustomLinksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubermaticCustomLinksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetKubermaticCustomLinksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKubermaticCustomLinksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetKubermaticCustomLinksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetKubermaticCustomLinksOK creates a GetKubermaticCustomLinksOK with default headers values
func NewGetKubermaticCustomLinksOK() *GetKubermaticCustomLinksOK {
	return &GetKubermaticCustomLinksOK{}
}

/*GetKubermaticCustomLinksOK handles this case with default header values.

GlobalCustomLinks
*/
type GetKubermaticCustomLinksOK struct {
	Payload models.GlobalCustomLinks
}

func (o *GetKubermaticCustomLinksOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/settings/customlinks][%d] getKubermaticCustomLinksOK  %+v", 200, o.Payload)
}

func (o *GetKubermaticCustomLinksOK) GetPayload() models.GlobalCustomLinks {
	return o.Payload
}

func (o *GetKubermaticCustomLinksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKubermaticCustomLinksUnauthorized creates a GetKubermaticCustomLinksUnauthorized with default headers values
func NewGetKubermaticCustomLinksUnauthorized() *GetKubermaticCustomLinksUnauthorized {
	return &GetKubermaticCustomLinksUnauthorized{}
}

/*GetKubermaticCustomLinksUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type GetKubermaticCustomLinksUnauthorized struct {
}

func (o *GetKubermaticCustomLinksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/settings/customlinks][%d] getKubermaticCustomLinksUnauthorized ", 401)
}

func (o *GetKubermaticCustomLinksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubermaticCustomLinksForbidden creates a GetKubermaticCustomLinksForbidden with default headers values
func NewGetKubermaticCustomLinksForbidden() *GetKubermaticCustomLinksForbidden {
	return &GetKubermaticCustomLinksForbidden{}
}

/*GetKubermaticCustomLinksForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type GetKubermaticCustomLinksForbidden struct {
}

func (o *GetKubermaticCustomLinksForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/settings/customlinks][%d] getKubermaticCustomLinksForbidden ", 403)
}

func (o *GetKubermaticCustomLinksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubermaticCustomLinksDefault creates a GetKubermaticCustomLinksDefault with default headers values
func NewGetKubermaticCustomLinksDefault(code int) *GetKubermaticCustomLinksDefault {
	return &GetKubermaticCustomLinksDefault{
		_statusCode: code,
	}
}

/*GetKubermaticCustomLinksDefault handles this case with default header values.

errorResponse
*/
type GetKubermaticCustomLinksDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get kubermatic custom links default response
func (o *GetKubermaticCustomLinksDefault) Code() int {
	return o._statusCode
}

func (o *GetKubermaticCustomLinksDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/settings/customlinks][%d] getKubermaticCustomLinks default  %+v", o._statusCode, o.Payload)
}

func (o *GetKubermaticCustomLinksDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetKubermaticCustomLinksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
