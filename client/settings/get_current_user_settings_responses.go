// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/go-metakube/models"
)

// GetCurrentUserSettingsReader is a Reader for the GetCurrentUserSettings structure.
type GetCurrentUserSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentUserSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrentUserSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCurrentUserSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCurrentUserSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCurrentUserSettingsOK creates a GetCurrentUserSettingsOK with default headers values
func NewGetCurrentUserSettingsOK() *GetCurrentUserSettingsOK {
	return &GetCurrentUserSettingsOK{}
}

/*
GetCurrentUserSettingsOK describes a response with status code 200, with default header values.

UserSettings
*/
type GetCurrentUserSettingsOK struct {
	Payload *models.UserSettings
}

// IsSuccess returns true when this get current user settings o k response has a 2xx status code
func (o *GetCurrentUserSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get current user settings o k response has a 3xx status code
func (o *GetCurrentUserSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current user settings o k response has a 4xx status code
func (o *GetCurrentUserSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get current user settings o k response has a 5xx status code
func (o *GetCurrentUserSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get current user settings o k response a status code equal to that given
func (o *GetCurrentUserSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get current user settings o k response
func (o *GetCurrentUserSettingsOK) Code() int {
	return 200
}

func (o *GetCurrentUserSettingsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/me/settings][%d] getCurrentUserSettingsOK %s", 200, payload)
}

func (o *GetCurrentUserSettingsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/me/settings][%d] getCurrentUserSettingsOK %s", 200, payload)
}

func (o *GetCurrentUserSettingsOK) GetPayload() *models.UserSettings {
	return o.Payload
}

func (o *GetCurrentUserSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentUserSettingsUnauthorized creates a GetCurrentUserSettingsUnauthorized with default headers values
func NewGetCurrentUserSettingsUnauthorized() *GetCurrentUserSettingsUnauthorized {
	return &GetCurrentUserSettingsUnauthorized{}
}

/*
GetCurrentUserSettingsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetCurrentUserSettingsUnauthorized struct {
}

// IsSuccess returns true when this get current user settings unauthorized response has a 2xx status code
func (o *GetCurrentUserSettingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get current user settings unauthorized response has a 3xx status code
func (o *GetCurrentUserSettingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current user settings unauthorized response has a 4xx status code
func (o *GetCurrentUserSettingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get current user settings unauthorized response has a 5xx status code
func (o *GetCurrentUserSettingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get current user settings unauthorized response a status code equal to that given
func (o *GetCurrentUserSettingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get current user settings unauthorized response
func (o *GetCurrentUserSettingsUnauthorized) Code() int {
	return 401
}

func (o *GetCurrentUserSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/me/settings][%d] getCurrentUserSettingsUnauthorized", 401)
}

func (o *GetCurrentUserSettingsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/me/settings][%d] getCurrentUserSettingsUnauthorized", 401)
}

func (o *GetCurrentUserSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCurrentUserSettingsDefault creates a GetCurrentUserSettingsDefault with default headers values
func NewGetCurrentUserSettingsDefault(code int) *GetCurrentUserSettingsDefault {
	return &GetCurrentUserSettingsDefault{
		_statusCode: code,
	}
}

/*
GetCurrentUserSettingsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetCurrentUserSettingsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get current user settings default response has a 2xx status code
func (o *GetCurrentUserSettingsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get current user settings default response has a 3xx status code
func (o *GetCurrentUserSettingsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get current user settings default response has a 4xx status code
func (o *GetCurrentUserSettingsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get current user settings default response has a 5xx status code
func (o *GetCurrentUserSettingsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get current user settings default response a status code equal to that given
func (o *GetCurrentUserSettingsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get current user settings default response
func (o *GetCurrentUserSettingsDefault) Code() int {
	return o._statusCode
}

func (o *GetCurrentUserSettingsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/me/settings][%d] getCurrentUserSettings default %s", o._statusCode, payload)
}

func (o *GetCurrentUserSettingsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/me/settings][%d] getCurrentUserSettings default %s", o._statusCode, payload)
}

func (o *GetCurrentUserSettingsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetCurrentUserSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
