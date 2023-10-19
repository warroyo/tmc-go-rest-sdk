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

// IntegrationResourceServiceListReader is a Reader for the IntegrationResourceServiceList structure.
type IntegrationResourceServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IntegrationResourceServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIntegrationResourceServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIntegrationResourceServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIntegrationResourceServiceListOK creates a IntegrationResourceServiceListOK with default headers values
func NewIntegrationResourceServiceListOK() *IntegrationResourceServiceListOK {
	return &IntegrationResourceServiceListOK{}
}

/*
IntegrationResourceServiceListOK describes a response with status code 200, with default header values.

A successful response.
*/
type IntegrationResourceServiceListOK struct {
	Payload *models.OrganizationIntegrationListIntegrationsResponse
}

// IsSuccess returns true when this integration resource service list o k response has a 2xx status code
func (o *IntegrationResourceServiceListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this integration resource service list o k response has a 3xx status code
func (o *IntegrationResourceServiceListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this integration resource service list o k response has a 4xx status code
func (o *IntegrationResourceServiceListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this integration resource service list o k response has a 5xx status code
func (o *IntegrationResourceServiceListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this integration resource service list o k response a status code equal to that given
func (o *IntegrationResourceServiceListOK) IsCode(code int) bool {
	return code == 200
}

func (o *IntegrationResourceServiceListOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/organization/integrations][%d] integrationResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *IntegrationResourceServiceListOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/organization/integrations][%d] integrationResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *IntegrationResourceServiceListOK) GetPayload() *models.OrganizationIntegrationListIntegrationsResponse {
	return o.Payload
}

func (o *IntegrationResourceServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationIntegrationListIntegrationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIntegrationResourceServiceListDefault creates a IntegrationResourceServiceListDefault with default headers values
func NewIntegrationResourceServiceListDefault(code int) *IntegrationResourceServiceListDefault {
	return &IntegrationResourceServiceListDefault{
		_statusCode: code,
	}
}

/*
IntegrationResourceServiceListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IntegrationResourceServiceListDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the integration resource service list default response
func (o *IntegrationResourceServiceListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this integration resource service list default response has a 2xx status code
func (o *IntegrationResourceServiceListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this integration resource service list default response has a 3xx status code
func (o *IntegrationResourceServiceListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this integration resource service list default response has a 4xx status code
func (o *IntegrationResourceServiceListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this integration resource service list default response has a 5xx status code
func (o *IntegrationResourceServiceListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this integration resource service list default response a status code equal to that given
func (o *IntegrationResourceServiceListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IntegrationResourceServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/organization/integrations][%d] IntegrationResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *IntegrationResourceServiceListDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/organization/integrations][%d] IntegrationResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *IntegrationResourceServiceListDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *IntegrationResourceServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
