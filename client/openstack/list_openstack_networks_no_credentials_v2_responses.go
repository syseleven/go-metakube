// Code generated by go-swagger; DO NOT EDIT.

package openstack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/go-metakube/models"
)

// ListOpenstackNetworksNoCredentialsV2Reader is a Reader for the ListOpenstackNetworksNoCredentialsV2 structure.
type ListOpenstackNetworksNoCredentialsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOpenstackNetworksNoCredentialsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOpenstackNetworksNoCredentialsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListOpenstackNetworksNoCredentialsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOpenstackNetworksNoCredentialsV2OK creates a ListOpenstackNetworksNoCredentialsV2OK with default headers values
func NewListOpenstackNetworksNoCredentialsV2OK() *ListOpenstackNetworksNoCredentialsV2OK {
	return &ListOpenstackNetworksNoCredentialsV2OK{}
}

/*ListOpenstackNetworksNoCredentialsV2OK handles this case with default header values.

OpenstackNetwork
*/
type ListOpenstackNetworksNoCredentialsV2OK struct {
	Payload []*models.OpenstackNetwork
}

func (o *ListOpenstackNetworksNoCredentialsV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/networks][%d] listOpenstackNetworksNoCredentialsV2OK  %+v", 200, o.Payload)
}

func (o *ListOpenstackNetworksNoCredentialsV2OK) GetPayload() []*models.OpenstackNetwork {
	return o.Payload
}

func (o *ListOpenstackNetworksNoCredentialsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOpenstackNetworksNoCredentialsV2Default creates a ListOpenstackNetworksNoCredentialsV2Default with default headers values
func NewListOpenstackNetworksNoCredentialsV2Default(code int) *ListOpenstackNetworksNoCredentialsV2Default {
	return &ListOpenstackNetworksNoCredentialsV2Default{
		_statusCode: code,
	}
}

/*ListOpenstackNetworksNoCredentialsV2Default handles this case with default header values.

errorResponse
*/
type ListOpenstackNetworksNoCredentialsV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list openstack networks no credentials v2 default response
func (o *ListOpenstackNetworksNoCredentialsV2Default) Code() int {
	return o._statusCode
}

func (o *ListOpenstackNetworksNoCredentialsV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/networks][%d] listOpenstackNetworksNoCredentialsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListOpenstackNetworksNoCredentialsV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListOpenstackNetworksNoCredentialsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
