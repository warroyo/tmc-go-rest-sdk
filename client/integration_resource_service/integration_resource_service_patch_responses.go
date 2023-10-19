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

// IntegrationResourceServicePatchReader is a Reader for the IntegrationResourceServicePatch structure.
type IntegrationResourceServicePatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IntegrationResourceServicePatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIntegrationResourceServicePatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIntegrationResourceServicePatchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIntegrationResourceServicePatchOK creates a IntegrationResourceServicePatchOK with default headers values
func NewIntegrationResourceServicePatchOK() *IntegrationResourceServicePatchOK {
	return &IntegrationResourceServicePatchOK{}
}

/*
IntegrationResourceServicePatchOK describes a response with status code 200, with default header values.

A successful response.
*/
type IntegrationResourceServicePatchOK struct {
	Payload *models.ClusterIntegrationPatchIntegrationResponse
}

// IsSuccess returns true when this integration resource service patch o k response has a 2xx status code
func (o *IntegrationResourceServicePatchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this integration resource service patch o k response has a 3xx status code
func (o *IntegrationResourceServicePatchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this integration resource service patch o k response has a 4xx status code
func (o *IntegrationResourceServicePatchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this integration resource service patch o k response has a 5xx status code
func (o *IntegrationResourceServicePatchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this integration resource service patch o k response a status code equal to that given
func (o *IntegrationResourceServicePatchOK) IsCode(code int) bool {
	return code == 200
}

func (o *IntegrationResourceServicePatchOK) Error() string {
	return fmt.Sprintf("[PATCH /v1alpha1/clusters/{fullName.clusterName}/integrations/{fullName.name}][%d] integrationResourceServicePatchOK  %+v", 200, o.Payload)
}

func (o *IntegrationResourceServicePatchOK) String() string {
	return fmt.Sprintf("[PATCH /v1alpha1/clusters/{fullName.clusterName}/integrations/{fullName.name}][%d] integrationResourceServicePatchOK  %+v", 200, o.Payload)
}

func (o *IntegrationResourceServicePatchOK) GetPayload() *models.ClusterIntegrationPatchIntegrationResponse {
	return o.Payload
}

func (o *IntegrationResourceServicePatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterIntegrationPatchIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIntegrationResourceServicePatchDefault creates a IntegrationResourceServicePatchDefault with default headers values
func NewIntegrationResourceServicePatchDefault(code int) *IntegrationResourceServicePatchDefault {
	return &IntegrationResourceServicePatchDefault{
		_statusCode: code,
	}
}

/*
IntegrationResourceServicePatchDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IntegrationResourceServicePatchDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the integration resource service patch default response
func (o *IntegrationResourceServicePatchDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this integration resource service patch default response has a 2xx status code
func (o *IntegrationResourceServicePatchDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this integration resource service patch default response has a 3xx status code
func (o *IntegrationResourceServicePatchDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this integration resource service patch default response has a 4xx status code
func (o *IntegrationResourceServicePatchDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this integration resource service patch default response has a 5xx status code
func (o *IntegrationResourceServicePatchDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this integration resource service patch default response a status code equal to that given
func (o *IntegrationResourceServicePatchDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IntegrationResourceServicePatchDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1alpha1/clusters/{fullName.clusterName}/integrations/{fullName.name}][%d] IntegrationResourceService_Patch default  %+v", o._statusCode, o.Payload)
}

func (o *IntegrationResourceServicePatchDefault) String() string {
	return fmt.Sprintf("[PATCH /v1alpha1/clusters/{fullName.clusterName}/integrations/{fullName.name}][%d] IntegrationResourceService_Patch default  %+v", o._statusCode, o.Payload)
}

func (o *IntegrationResourceServicePatchDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *IntegrationResourceServicePatchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}