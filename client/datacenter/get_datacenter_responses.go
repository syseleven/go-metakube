// Code generated by go-swagger; DO NOT EDIT.

package datacenter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/go-metakube/models"
)

// GetDatacenterReader is a Reader for the GetDatacenter structure.
type GetDatacenterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatacenterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDatacenterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDatacenterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDatacenterOK creates a GetDatacenterOK with default headers values
func NewGetDatacenterOK() *GetDatacenterOK {
	return &GetDatacenterOK{}
}

/*GetDatacenterOK handles this case with default header values.

Datacenter
*/
type GetDatacenterOK struct {
	Payload *models.Datacenter
}

func (o *GetDatacenterOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/dc/{dc}][%d] getDatacenterOK  %+v", 200, o.Payload)
}

func (o *GetDatacenterOK) GetPayload() *models.Datacenter {
	return o.Payload
}

func (o *GetDatacenterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Datacenter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatacenterDefault creates a GetDatacenterDefault with default headers values
func NewGetDatacenterDefault(code int) *GetDatacenterDefault {
	return &GetDatacenterDefault{
		_statusCode: code,
	}
}

/*GetDatacenterDefault handles this case with default header values.

errorResponse
*/
type GetDatacenterDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get datacenter default response
func (o *GetDatacenterDefault) Code() int {
	return o._statusCode
}

func (o *GetDatacenterDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/dc/{dc}][%d] getDatacenter default  %+v", o._statusCode, o.Payload)
}

func (o *GetDatacenterDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDatacenterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
