// Code generated by go-swagger; DO NOT EDIT.

package seed

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

// GetSeedSettingsReader is a Reader for the GetSeedSettings structure.
type GetSeedSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSeedSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSeedSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSeedSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSeedSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSeedSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSeedSettingsOK creates a GetSeedSettingsOK with default headers values
func NewGetSeedSettingsOK() *GetSeedSettingsOK {
	return &GetSeedSettingsOK{}
}

/*
GetSeedSettingsOK describes a response with status code 200, with default header values.

SeedSettings
*/
type GetSeedSettingsOK struct {
	Payload *models.SeedSettings
}

// IsSuccess returns true when this get seed settings o k response has a 2xx status code
func (o *GetSeedSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get seed settings o k response has a 3xx status code
func (o *GetSeedSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get seed settings o k response has a 4xx status code
func (o *GetSeedSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get seed settings o k response has a 5xx status code
func (o *GetSeedSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get seed settings o k response a status code equal to that given
func (o *GetSeedSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get seed settings o k response
func (o *GetSeedSettingsOK) Code() int {
	return 200
}

func (o *GetSeedSettingsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/settings][%d] getSeedSettingsOK %s", 200, payload)
}

func (o *GetSeedSettingsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/settings][%d] getSeedSettingsOK %s", 200, payload)
}

func (o *GetSeedSettingsOK) GetPayload() *models.SeedSettings {
	return o.Payload
}

func (o *GetSeedSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SeedSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSeedSettingsUnauthorized creates a GetSeedSettingsUnauthorized with default headers values
func NewGetSeedSettingsUnauthorized() *GetSeedSettingsUnauthorized {
	return &GetSeedSettingsUnauthorized{}
}

/*
GetSeedSettingsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetSeedSettingsUnauthorized struct {
}

// IsSuccess returns true when this get seed settings unauthorized response has a 2xx status code
func (o *GetSeedSettingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get seed settings unauthorized response has a 3xx status code
func (o *GetSeedSettingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get seed settings unauthorized response has a 4xx status code
func (o *GetSeedSettingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get seed settings unauthorized response has a 5xx status code
func (o *GetSeedSettingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get seed settings unauthorized response a status code equal to that given
func (o *GetSeedSettingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get seed settings unauthorized response
func (o *GetSeedSettingsUnauthorized) Code() int {
	return 401
}

func (o *GetSeedSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/settings][%d] getSeedSettingsUnauthorized", 401)
}

func (o *GetSeedSettingsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/settings][%d] getSeedSettingsUnauthorized", 401)
}

func (o *GetSeedSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSeedSettingsForbidden creates a GetSeedSettingsForbidden with default headers values
func NewGetSeedSettingsForbidden() *GetSeedSettingsForbidden {
	return &GetSeedSettingsForbidden{}
}

/*
GetSeedSettingsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetSeedSettingsForbidden struct {
}

// IsSuccess returns true when this get seed settings forbidden response has a 2xx status code
func (o *GetSeedSettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get seed settings forbidden response has a 3xx status code
func (o *GetSeedSettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get seed settings forbidden response has a 4xx status code
func (o *GetSeedSettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get seed settings forbidden response has a 5xx status code
func (o *GetSeedSettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get seed settings forbidden response a status code equal to that given
func (o *GetSeedSettingsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get seed settings forbidden response
func (o *GetSeedSettingsForbidden) Code() int {
	return 403
}

func (o *GetSeedSettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/settings][%d] getSeedSettingsForbidden", 403)
}

func (o *GetSeedSettingsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/settings][%d] getSeedSettingsForbidden", 403)
}

func (o *GetSeedSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSeedSettingsDefault creates a GetSeedSettingsDefault with default headers values
func NewGetSeedSettingsDefault(code int) *GetSeedSettingsDefault {
	return &GetSeedSettingsDefault{
		_statusCode: code,
	}
}

/*
GetSeedSettingsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetSeedSettingsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get seed settings default response has a 2xx status code
func (o *GetSeedSettingsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get seed settings default response has a 3xx status code
func (o *GetSeedSettingsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get seed settings default response has a 4xx status code
func (o *GetSeedSettingsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get seed settings default response has a 5xx status code
func (o *GetSeedSettingsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get seed settings default response a status code equal to that given
func (o *GetSeedSettingsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get seed settings default response
func (o *GetSeedSettingsDefault) Code() int {
	return o._statusCode
}

func (o *GetSeedSettingsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/settings][%d] getSeedSettings default %s", o._statusCode, payload)
}

func (o *GetSeedSettingsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/seeds/{seed_name}/settings][%d] getSeedSettings default %s", o._statusCode, payload)
}

func (o *GetSeedSettingsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSeedSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
