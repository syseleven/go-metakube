// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/go-metakube/models"
)

// ListAWSSizesNoCredentialsV2Reader is a Reader for the ListAWSSizesNoCredentialsV2 structure.
type ListAWSSizesNoCredentialsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAWSSizesNoCredentialsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAWSSizesNoCredentialsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAWSSizesNoCredentialsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAWSSizesNoCredentialsV2OK creates a ListAWSSizesNoCredentialsV2OK with default headers values
func NewListAWSSizesNoCredentialsV2OK() *ListAWSSizesNoCredentialsV2OK {
	return &ListAWSSizesNoCredentialsV2OK{}
}

/*ListAWSSizesNoCredentialsV2OK handles this case with default header values.

AWSSizeList
*/
type ListAWSSizesNoCredentialsV2OK struct {
	Payload models.AWSSizeList
}

func (o *ListAWSSizesNoCredentialsV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/aws/sizes][%d] listAWSSizesNoCredentialsV2OK  %+v", 200, o.Payload)
}

func (o *ListAWSSizesNoCredentialsV2OK) GetPayload() models.AWSSizeList {
	return o.Payload
}

func (o *ListAWSSizesNoCredentialsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAWSSizesNoCredentialsV2Default creates a ListAWSSizesNoCredentialsV2Default with default headers values
func NewListAWSSizesNoCredentialsV2Default(code int) *ListAWSSizesNoCredentialsV2Default {
	return &ListAWSSizesNoCredentialsV2Default{
		_statusCode: code,
	}
}

/*ListAWSSizesNoCredentialsV2Default handles this case with default header values.

errorResponse
*/
type ListAWSSizesNoCredentialsV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list a w s sizes no credentials v2 default response
func (o *ListAWSSizesNoCredentialsV2Default) Code() int {
	return o._statusCode
}

func (o *ListAWSSizesNoCredentialsV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/aws/sizes][%d] listAWSSizesNoCredentialsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListAWSSizesNoCredentialsV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAWSSizesNoCredentialsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
