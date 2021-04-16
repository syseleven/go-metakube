// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/go-metakube/models"
)

// GetClusterKubeconfigV2Reader is a Reader for the GetClusterKubeconfigV2 structure.
type GetClusterKubeconfigV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterKubeconfigV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterKubeconfigV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetClusterKubeconfigV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetClusterKubeconfigV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetClusterKubeconfigV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterKubeconfigV2OK creates a GetClusterKubeconfigV2OK with default headers values
func NewGetClusterKubeconfigV2OK() *GetClusterKubeconfigV2OK {
	return &GetClusterKubeconfigV2OK{}
}

/*GetClusterKubeconfigV2OK handles this case with default header values.

Kubeconfig is a clusters kubeconfig
*/
type GetClusterKubeconfigV2OK struct {
	Payload []uint8
}

func (o *GetClusterKubeconfigV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeconfig][%d] getClusterKubeconfigV2OK  %+v", 200, o.Payload)
}

func (o *GetClusterKubeconfigV2OK) GetPayload() []uint8 {
	return o.Payload
}

func (o *GetClusterKubeconfigV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterKubeconfigV2Unauthorized creates a GetClusterKubeconfigV2Unauthorized with default headers values
func NewGetClusterKubeconfigV2Unauthorized() *GetClusterKubeconfigV2Unauthorized {
	return &GetClusterKubeconfigV2Unauthorized{}
}

/*GetClusterKubeconfigV2Unauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type GetClusterKubeconfigV2Unauthorized struct {
}

func (o *GetClusterKubeconfigV2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeconfig][%d] getClusterKubeconfigV2Unauthorized ", 401)
}

func (o *GetClusterKubeconfigV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterKubeconfigV2Forbidden creates a GetClusterKubeconfigV2Forbidden with default headers values
func NewGetClusterKubeconfigV2Forbidden() *GetClusterKubeconfigV2Forbidden {
	return &GetClusterKubeconfigV2Forbidden{}
}

/*GetClusterKubeconfigV2Forbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type GetClusterKubeconfigV2Forbidden struct {
}

func (o *GetClusterKubeconfigV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeconfig][%d] getClusterKubeconfigV2Forbidden ", 403)
}

func (o *GetClusterKubeconfigV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterKubeconfigV2Default creates a GetClusterKubeconfigV2Default with default headers values
func NewGetClusterKubeconfigV2Default(code int) *GetClusterKubeconfigV2Default {
	return &GetClusterKubeconfigV2Default{
		_statusCode: code,
	}
}

/*GetClusterKubeconfigV2Default handles this case with default header values.

errorResponse
*/
type GetClusterKubeconfigV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get cluster kubeconfig v2 default response
func (o *GetClusterKubeconfigV2Default) Code() int {
	return o._statusCode
}

func (o *GetClusterKubeconfigV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeconfig][%d] getClusterKubeconfigV2 default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterKubeconfigV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterKubeconfigV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
