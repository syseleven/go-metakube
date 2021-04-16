// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/go-metakube/models"
)

// ListVSphereNetworksReader is a Reader for the ListVSphereNetworks structure.
type ListVSphereNetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVSphereNetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVSphereNetworksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListVSphereNetworksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListVSphereNetworksOK creates a ListVSphereNetworksOK with default headers values
func NewListVSphereNetworksOK() *ListVSphereNetworksOK {
	return &ListVSphereNetworksOK{}
}

/*ListVSphereNetworksOK handles this case with default header values.

VSphereNetwork
*/
type ListVSphereNetworksOK struct {
	Payload []*models.VSphereNetwork
}

func (o *ListVSphereNetworksOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/vsphere/networks][%d] listVSphereNetworksOK  %+v", 200, o.Payload)
}

func (o *ListVSphereNetworksOK) GetPayload() []*models.VSphereNetwork {
	return o.Payload
}

func (o *ListVSphereNetworksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVSphereNetworksDefault creates a ListVSphereNetworksDefault with default headers values
func NewListVSphereNetworksDefault(code int) *ListVSphereNetworksDefault {
	return &ListVSphereNetworksDefault{
		_statusCode: code,
	}
}

/*ListVSphereNetworksDefault handles this case with default header values.

errorResponse
*/
type ListVSphereNetworksDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list v sphere networks default response
func (o *ListVSphereNetworksDefault) Code() int {
	return o._statusCode
}

func (o *ListVSphereNetworksDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/vsphere/networks][%d] listVSphereNetworks default  %+v", o._statusCode, o.Payload)
}

func (o *ListVSphereNetworksDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListVSphereNetworksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
