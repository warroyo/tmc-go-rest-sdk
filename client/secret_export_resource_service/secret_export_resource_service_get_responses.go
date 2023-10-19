// Code generated by go-swagger; DO NOT EDIT.

package secret_export_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// SecretExportResourceServiceGetReader is a Reader for the SecretExportResourceServiceGet structure.
type SecretExportResourceServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretExportResourceServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretExportResourceServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecretExportResourceServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecretExportResourceServiceGetOK creates a SecretExportResourceServiceGetOK with default headers values
func NewSecretExportResourceServiceGetOK() *SecretExportResourceServiceGetOK {
	return &SecretExportResourceServiceGetOK{}
}

/*
SecretExportResourceServiceGetOK describes a response with status code 200, with default header values.

A successful response.
*/
type SecretExportResourceServiceGetOK struct {
	Payload *models.ClusterNamespaceSecretexportGetSecretExportResponse
}

// IsSuccess returns true when this secret export resource service get o k response has a 2xx status code
func (o *SecretExportResourceServiceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this secret export resource service get o k response has a 3xx status code
func (o *SecretExportResourceServiceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secret export resource service get o k response has a 4xx status code
func (o *SecretExportResourceServiceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this secret export resource service get o k response has a 5xx status code
func (o *SecretExportResourceServiceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this secret export resource service get o k response a status code equal to that given
func (o *SecretExportResourceServiceGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *SecretExportResourceServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/secretexports/{fullName.name}][%d] secretExportResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *SecretExportResourceServiceGetOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/secretexports/{fullName.name}][%d] secretExportResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *SecretExportResourceServiceGetOK) GetPayload() *models.ClusterNamespaceSecretexportGetSecretExportResponse {
	return o.Payload
}

func (o *SecretExportResourceServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterNamespaceSecretexportGetSecretExportResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretExportResourceServiceGetDefault creates a SecretExportResourceServiceGetDefault with default headers values
func NewSecretExportResourceServiceGetDefault(code int) *SecretExportResourceServiceGetDefault {
	return &SecretExportResourceServiceGetDefault{
		_statusCode: code,
	}
}

/*
SecretExportResourceServiceGetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SecretExportResourceServiceGetDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the secret export resource service get default response
func (o *SecretExportResourceServiceGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this secret export resource service get default response has a 2xx status code
func (o *SecretExportResourceServiceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this secret export resource service get default response has a 3xx status code
func (o *SecretExportResourceServiceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this secret export resource service get default response has a 4xx status code
func (o *SecretExportResourceServiceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this secret export resource service get default response has a 5xx status code
func (o *SecretExportResourceServiceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this secret export resource service get default response a status code equal to that given
func (o *SecretExportResourceServiceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SecretExportResourceServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/secretexports/{fullName.name}][%d] SecretExportResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *SecretExportResourceServiceGetDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/secretexports/{fullName.name}][%d] SecretExportResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *SecretExportResourceServiceGetDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *SecretExportResourceServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}