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

// NamespaceResourceServiceCreateReader is a Reader for the NamespaceResourceServiceCreate structure.
type NamespaceResourceServiceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NamespaceResourceServiceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNamespaceResourceServiceCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNamespaceResourceServiceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNamespaceResourceServiceCreateOK creates a NamespaceResourceServiceCreateOK with default headers values
func NewNamespaceResourceServiceCreateOK() *NamespaceResourceServiceCreateOK {
	return &NamespaceResourceServiceCreateOK{}
}

/*
NamespaceResourceServiceCreateOK describes a response with status code 200, with default header values.

A successful response.
*/
type NamespaceResourceServiceCreateOK struct {
	Payload *models.ClusterNamespaceCreateNamespaceResponse
}

// IsSuccess returns true when this namespace resource service create o k response has a 2xx status code
func (o *NamespaceResourceServiceCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this namespace resource service create o k response has a 3xx status code
func (o *NamespaceResourceServiceCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this namespace resource service create o k response has a 4xx status code
func (o *NamespaceResourceServiceCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this namespace resource service create o k response has a 5xx status code
func (o *NamespaceResourceServiceCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this namespace resource service create o k response a status code equal to that given
func (o *NamespaceResourceServiceCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *NamespaceResourceServiceCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/clusters/{namespace.fullName.clusterName}/namespaces][%d] namespaceResourceServiceCreateOK  %+v", 200, o.Payload)
}

func (o *NamespaceResourceServiceCreateOK) String() string {
	return fmt.Sprintf("[POST /v1alpha1/clusters/{namespace.fullName.clusterName}/namespaces][%d] namespaceResourceServiceCreateOK  %+v", 200, o.Payload)
}

func (o *NamespaceResourceServiceCreateOK) GetPayload() *models.ClusterNamespaceCreateNamespaceResponse {
	return o.Payload
}

func (o *NamespaceResourceServiceCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterNamespaceCreateNamespaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNamespaceResourceServiceCreateDefault creates a NamespaceResourceServiceCreateDefault with default headers values
func NewNamespaceResourceServiceCreateDefault(code int) *NamespaceResourceServiceCreateDefault {
	return &NamespaceResourceServiceCreateDefault{
		_statusCode: code,
	}
}

/*
NamespaceResourceServiceCreateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type NamespaceResourceServiceCreateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the namespace resource service create default response
func (o *NamespaceResourceServiceCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this namespace resource service create default response has a 2xx status code
func (o *NamespaceResourceServiceCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this namespace resource service create default response has a 3xx status code
func (o *NamespaceResourceServiceCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this namespace resource service create default response has a 4xx status code
func (o *NamespaceResourceServiceCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this namespace resource service create default response has a 5xx status code
func (o *NamespaceResourceServiceCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this namespace resource service create default response a status code equal to that given
func (o *NamespaceResourceServiceCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *NamespaceResourceServiceCreateDefault) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/clusters/{namespace.fullName.clusterName}/namespaces][%d] NamespaceResourceService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *NamespaceResourceServiceCreateDefault) String() string {
	return fmt.Sprintf("[POST /v1alpha1/clusters/{namespace.fullName.clusterName}/namespaces][%d] NamespaceResourceService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *NamespaceResourceServiceCreateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *NamespaceResourceServiceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}