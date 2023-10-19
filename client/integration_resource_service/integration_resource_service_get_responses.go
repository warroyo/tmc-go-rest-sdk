// Code generated by go-swagger; DO NOT EDIT.

package integration_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// IntegrationResourceServiceGetReader is a Reader for the IntegrationResourceServiceGet structure.
type IntegrationResourceServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IntegrationResourceServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIntegrationResourceServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIntegrationResourceServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIntegrationResourceServiceGetOK creates a IntegrationResourceServiceGetOK with default headers values
func NewIntegrationResourceServiceGetOK() *IntegrationResourceServiceGetOK {
	return &IntegrationResourceServiceGetOK{}
}

/*
IntegrationResourceServiceGetOK describes a response with status code 200, with default header values.

A successful response.
*/
type IntegrationResourceServiceGetOK struct {
	Payload *models.OrganizationIntegrationGetIntegrationResponse
}

// IsSuccess returns true when this integration resource service get o k response has a 2xx status code
func (o *IntegrationResourceServiceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this integration resource service get o k response has a 3xx status code
func (o *IntegrationResourceServiceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this integration resource service get o k response has a 4xx status code
func (o *IntegrationResourceServiceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this integration resource service get o k response has a 5xx status code
func (o *IntegrationResourceServiceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this integration resource service get o k response a status code equal to that given
func (o *IntegrationResourceServiceGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *IntegrationResourceServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/organization/integrations/{fullName.name}][%d] integrationResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *IntegrationResourceServiceGetOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/organization/integrations/{fullName.name}][%d] integrationResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *IntegrationResourceServiceGetOK) GetPayload() *models.OrganizationIntegrationGetIntegrationResponse {
	return o.Payload
}

func (o *IntegrationResourceServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationIntegrationGetIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIntegrationResourceServiceGetDefault creates a IntegrationResourceServiceGetDefault with default headers values
func NewIntegrationResourceServiceGetDefault(code int) *IntegrationResourceServiceGetDefault {
	return &IntegrationResourceServiceGetDefault{
		_statusCode: code,
	}
}

/*
IntegrationResourceServiceGetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IntegrationResourceServiceGetDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the integration resource service get default response
func (o *IntegrationResourceServiceGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this integration resource service get default response has a 2xx status code
func (o *IntegrationResourceServiceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this integration resource service get default response has a 3xx status code
func (o *IntegrationResourceServiceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this integration resource service get default response has a 4xx status code
func (o *IntegrationResourceServiceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this integration resource service get default response has a 5xx status code
func (o *IntegrationResourceServiceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this integration resource service get default response a status code equal to that given
func (o *IntegrationResourceServiceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IntegrationResourceServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/organization/integrations/{fullName.name}][%d] IntegrationResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *IntegrationResourceServiceGetDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/organization/integrations/{fullName.name}][%d] IntegrationResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *IntegrationResourceServiceGetDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *IntegrationResourceServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}