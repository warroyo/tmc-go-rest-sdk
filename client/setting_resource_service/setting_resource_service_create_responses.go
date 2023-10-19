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

// SettingResourceServiceCreateReader is a Reader for the SettingResourceServiceCreate structure.
type SettingResourceServiceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SettingResourceServiceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSettingResourceServiceCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSettingResourceServiceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSettingResourceServiceCreateOK creates a SettingResourceServiceCreateOK with default headers values
func NewSettingResourceServiceCreateOK() *SettingResourceServiceCreateOK {
	return &SettingResourceServiceCreateOK{}
}

/*
SettingResourceServiceCreateOK describes a response with status code 200, with default header values.

A successful response.
*/
type SettingResourceServiceCreateOK struct {
	Payload *models.OrganizationSettingCreateSettingResponse
}

// IsSuccess returns true when this setting resource service create o k response has a 2xx status code
func (o *SettingResourceServiceCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this setting resource service create o k response has a 3xx status code
func (o *SettingResourceServiceCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this setting resource service create o k response has a 4xx status code
func (o *SettingResourceServiceCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this setting resource service create o k response has a 5xx status code
func (o *SettingResourceServiceCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this setting resource service create o k response a status code equal to that given
func (o *SettingResourceServiceCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *SettingResourceServiceCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/organization/settings][%d] settingResourceServiceCreateOK  %+v", 200, o.Payload)
}

func (o *SettingResourceServiceCreateOK) String() string {
	return fmt.Sprintf("[POST /v1alpha1/organization/settings][%d] settingResourceServiceCreateOK  %+v", 200, o.Payload)
}

func (o *SettingResourceServiceCreateOK) GetPayload() *models.OrganizationSettingCreateSettingResponse {
	return o.Payload
}

func (o *SettingResourceServiceCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationSettingCreateSettingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSettingResourceServiceCreateDefault creates a SettingResourceServiceCreateDefault with default headers values
func NewSettingResourceServiceCreateDefault(code int) *SettingResourceServiceCreateDefault {
	return &SettingResourceServiceCreateDefault{
		_statusCode: code,
	}
}

/*
SettingResourceServiceCreateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SettingResourceServiceCreateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the setting resource service create default response
func (o *SettingResourceServiceCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this setting resource service create default response has a 2xx status code
func (o *SettingResourceServiceCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this setting resource service create default response has a 3xx status code
func (o *SettingResourceServiceCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this setting resource service create default response has a 4xx status code
func (o *SettingResourceServiceCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this setting resource service create default response has a 5xx status code
func (o *SettingResourceServiceCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this setting resource service create default response a status code equal to that given
func (o *SettingResourceServiceCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SettingResourceServiceCreateDefault) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/organization/settings][%d] SettingResourceService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *SettingResourceServiceCreateDefault) String() string {
	return fmt.Sprintf("[POST /v1alpha1/organization/settings][%d] SettingResourceService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *SettingResourceServiceCreateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *SettingResourceServiceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}