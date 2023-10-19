// Code generated by go-swagger; DO NOT EDIT.

package source_secret_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// SourceSecretResourceServiceGetReader is a Reader for the SourceSecretResourceServiceGet structure.
type SourceSecretResourceServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SourceSecretResourceServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSourceSecretResourceServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSourceSecretResourceServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSourceSecretResourceServiceGetOK creates a SourceSecretResourceServiceGetOK with default headers values
func NewSourceSecretResourceServiceGetOK() *SourceSecretResourceServiceGetOK {
	return &SourceSecretResourceServiceGetOK{}
}

/*
SourceSecretResourceServiceGetOK describes a response with status code 200, with default header values.

A successful response.
*/
type SourceSecretResourceServiceGetOK struct {
	Payload *models.ClusterFluxcdSourcesecretGetSourceSecretResponse
}

// IsSuccess returns true when this source secret resource service get o k response has a 2xx status code
func (o *SourceSecretResourceServiceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this source secret resource service get o k response has a 3xx status code
func (o *SourceSecretResourceServiceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this source secret resource service get o k response has a 4xx status code
func (o *SourceSecretResourceServiceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this source secret resource service get o k response has a 5xx status code
func (o *SourceSecretResourceServiceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this source secret resource service get o k response a status code equal to that given
func (o *SourceSecretResourceServiceGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *SourceSecretResourceServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.clusterName}/fluxcd/sourcesecrets/{fullName.name}][%d] sourceSecretResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *SourceSecretResourceServiceGetOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.clusterName}/fluxcd/sourcesecrets/{fullName.name}][%d] sourceSecretResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *SourceSecretResourceServiceGetOK) GetPayload() *models.ClusterFluxcdSourcesecretGetSourceSecretResponse {
	return o.Payload
}

func (o *SourceSecretResourceServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterFluxcdSourcesecretGetSourceSecretResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSourceSecretResourceServiceGetDefault creates a SourceSecretResourceServiceGetDefault with default headers values
func NewSourceSecretResourceServiceGetDefault(code int) *SourceSecretResourceServiceGetDefault {
	return &SourceSecretResourceServiceGetDefault{
		_statusCode: code,
	}
}

/*
SourceSecretResourceServiceGetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SourceSecretResourceServiceGetDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the source secret resource service get default response
func (o *SourceSecretResourceServiceGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this source secret resource service get default response has a 2xx status code
func (o *SourceSecretResourceServiceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this source secret resource service get default response has a 3xx status code
func (o *SourceSecretResourceServiceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this source secret resource service get default response has a 4xx status code
func (o *SourceSecretResourceServiceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this source secret resource service get default response has a 5xx status code
func (o *SourceSecretResourceServiceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this source secret resource service get default response a status code equal to that given
func (o *SourceSecretResourceServiceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SourceSecretResourceServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.clusterName}/fluxcd/sourcesecrets/{fullName.name}][%d] SourceSecretResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *SourceSecretResourceServiceGetDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.clusterName}/fluxcd/sourcesecrets/{fullName.name}][%d] SourceSecretResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *SourceSecretResourceServiceGetDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *SourceSecretResourceServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
