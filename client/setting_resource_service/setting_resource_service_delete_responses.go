// Code generated by go-swagger; DO NOT EDIT.

package setting_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// SettingResourceServiceDeleteReader is a Reader for the SettingResourceServiceDelete structure.
type SettingResourceServiceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SettingResourceServiceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSettingResourceServiceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSettingResourceServiceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSettingResourceServiceDeleteOK creates a SettingResourceServiceDeleteOK with default headers values
func NewSettingResourceServiceDeleteOK() *SettingResourceServiceDeleteOK {
	return &SettingResourceServiceDeleteOK{}
}

/*
SettingResourceServiceDeleteOK describes a response with status code 200, with default header values.

A successful response.
*/
type SettingResourceServiceDeleteOK struct {
	Payload *models.OrganizationSettingDeleteSettingResponse
}

// IsSuccess returns true when this setting resource service delete o k response has a 2xx status code
func (o *SettingResourceServiceDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this setting resource service delete o k response has a 3xx status code
func (o *SettingResourceServiceDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this setting resource service delete o k response has a 4xx status code
func (o *SettingResourceServiceDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this setting resource service delete o k response has a 5xx status code
func (o *SettingResourceServiceDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this setting resource service delete o k response a status code equal to that given
func (o *SettingResourceServiceDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *SettingResourceServiceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1alpha1/organization/settings/{fullName.name}][%d] settingResourceServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *SettingResourceServiceDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /v1alpha1/organization/settings/{fullName.name}][%d] settingResourceServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *SettingResourceServiceDeleteOK) GetPayload() *models.OrganizationSettingDeleteSettingResponse {
	return o.Payload
}

func (o *SettingResourceServiceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationSettingDeleteSettingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSettingResourceServiceDeleteDefault creates a SettingResourceServiceDeleteDefault with default headers values
func NewSettingResourceServiceDeleteDefault(code int) *SettingResourceServiceDeleteDefault {
	return &SettingResourceServiceDeleteDefault{
		_statusCode: code,
	}
}

/*
SettingResourceServiceDeleteDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SettingResourceServiceDeleteDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the setting resource service delete default response
func (o *SettingResourceServiceDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this setting resource service delete default response has a 2xx status code
func (o *SettingResourceServiceDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this setting resource service delete default response has a 3xx status code
func (o *SettingResourceServiceDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this setting resource service delete default response has a 4xx status code
func (o *SettingResourceServiceDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this setting resource service delete default response has a 5xx status code
func (o *SettingResourceServiceDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this setting resource service delete default response a status code equal to that given
func (o *SettingResourceServiceDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SettingResourceServiceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1alpha1/organization/settings/{fullName.name}][%d] SettingResourceService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *SettingResourceServiceDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /v1alpha1/organization/settings/{fullName.name}][%d] SettingResourceService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *SettingResourceServiceDeleteDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *SettingResourceServiceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
