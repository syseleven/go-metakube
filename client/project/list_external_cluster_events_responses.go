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

// ListExternalClusterEventsReader is a Reader for the ListExternalClusterEvents structure.
type ListExternalClusterEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListExternalClusterEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListExternalClusterEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListExternalClusterEventsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListExternalClusterEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListExternalClusterEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListExternalClusterEventsOK creates a ListExternalClusterEventsOK with default headers values
func NewListExternalClusterEventsOK() *ListExternalClusterEventsOK {
	return &ListExternalClusterEventsOK{}
}

/*ListExternalClusterEventsOK handles this case with default header values.

Event
*/
type ListExternalClusterEventsOK struct {
	Payload []*models.Event
}

func (o *ListExternalClusterEventsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/events][%d] listExternalClusterEventsOK  %+v", 200, o.Payload)
}

func (o *ListExternalClusterEventsOK) GetPayload() []*models.Event {
	return o.Payload
}

func (o *ListExternalClusterEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListExternalClusterEventsUnauthorized creates a ListExternalClusterEventsUnauthorized with default headers values
func NewListExternalClusterEventsUnauthorized() *ListExternalClusterEventsUnauthorized {
	return &ListExternalClusterEventsUnauthorized{}
}

/*ListExternalClusterEventsUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type ListExternalClusterEventsUnauthorized struct {
}

func (o *ListExternalClusterEventsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/events][%d] listExternalClusterEventsUnauthorized ", 401)
}

func (o *ListExternalClusterEventsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListExternalClusterEventsForbidden creates a ListExternalClusterEventsForbidden with default headers values
func NewListExternalClusterEventsForbidden() *ListExternalClusterEventsForbidden {
	return &ListExternalClusterEventsForbidden{}
}

/*ListExternalClusterEventsForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type ListExternalClusterEventsForbidden struct {
}

func (o *ListExternalClusterEventsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/events][%d] listExternalClusterEventsForbidden ", 403)
}

func (o *ListExternalClusterEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListExternalClusterEventsDefault creates a ListExternalClusterEventsDefault with default headers values
func NewListExternalClusterEventsDefault(code int) *ListExternalClusterEventsDefault {
	return &ListExternalClusterEventsDefault{
		_statusCode: code,
	}
}

/*ListExternalClusterEventsDefault handles this case with default header values.

errorResponse
*/
type ListExternalClusterEventsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list external cluster events default response
func (o *ListExternalClusterEventsDefault) Code() int {
	return o._statusCode
}

func (o *ListExternalClusterEventsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/events][%d] listExternalClusterEvents default  %+v", o._statusCode, o.Payload)
}

func (o *ListExternalClusterEventsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListExternalClusterEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
