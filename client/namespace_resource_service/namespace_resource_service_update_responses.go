// Code generated by go-swagger; DO NOT EDIT.

package namespace_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// NamespaceResourceServiceUpdateReader is a Reader for the NamespaceResourceServiceUpdate structure.
type NamespaceResourceServiceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NamespaceResourceServiceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNamespaceResourceServiceUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNamespaceResourceServiceUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNamespaceResourceServiceUpdateOK creates a NamespaceResourceServiceUpdateOK with default headers values
func NewNamespaceResourceServiceUpdateOK() *NamespaceResourceServiceUpdateOK {
	return &NamespaceResourceServiceUpdateOK{}
}

/*
NamespaceResourceServiceUpdateOK describes a response with status code 200, with default header values.

A successful response.
*/
type NamespaceResourceServiceUpdateOK struct {
	Payload *models.ClusterNamespaceUpdateNamespaceResponse
}

// IsSuccess returns true when this namespace resource service update o k response has a 2xx status code
func (o *NamespaceResourceServiceUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this namespace resource service update o k response has a 3xx status code
func (o *NamespaceResourceServiceUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this namespace resource service update o k response has a 4xx status code
func (o *NamespaceResourceServiceUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this namespace resource service update o k response has a 5xx status code
func (o *NamespaceResourceServiceUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this namespace resource service update o k response a status code equal to that given
func (o *NamespaceResourceServiceUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *NamespaceResourceServiceUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/clusters/{namespace.fullName.clusterName}/namespaces/{namespace.fullName.name}][%d] namespaceResourceServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *NamespaceResourceServiceUpdateOK) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/clusters/{namespace.fullName.clusterName}/namespaces/{namespace.fullName.name}][%d] namespaceResourceServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *NamespaceResourceServiceUpdateOK) GetPayload() *models.ClusterNamespaceUpdateNamespaceResponse {
	return o.Payload
}

func (o *NamespaceResourceServiceUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterNamespaceUpdateNamespaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNamespaceResourceServiceUpdateDefault creates a NamespaceResourceServiceUpdateDefault with default headers values
func NewNamespaceResourceServiceUpdateDefault(code int) *NamespaceResourceServiceUpdateDefault {
	return &NamespaceResourceServiceUpdateDefault{
		_statusCode: code,
	}
}

/*
NamespaceResourceServiceUpdateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type NamespaceResourceServiceUpdateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the namespace resource service update default response
func (o *NamespaceResourceServiceUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this namespace resource service update default response has a 2xx status code
func (o *NamespaceResourceServiceUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this namespace resource service update default response has a 3xx status code
func (o *NamespaceResourceServiceUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this namespace resource service update default response has a 4xx status code
func (o *NamespaceResourceServiceUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this namespace resource service update default response has a 5xx status code
func (o *NamespaceResourceServiceUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this namespace resource service update default response a status code equal to that given
func (o *NamespaceResourceServiceUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *NamespaceResourceServiceUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/clusters/{namespace.fullName.clusterName}/namespaces/{namespace.fullName.name}][%d] NamespaceResourceService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *NamespaceResourceServiceUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/clusters/{namespace.fullName.clusterName}/namespaces/{namespace.fullName.name}][%d] NamespaceResourceService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *NamespaceResourceServiceUpdateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *NamespaceResourceServiceUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}