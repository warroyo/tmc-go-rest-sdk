// Code generated by go-swagger; DO NOT EDIT.

package aks_cluster_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// AksClusterResourceServiceCreateReader is a Reader for the AksClusterResourceServiceCreate structure.
type AksClusterResourceServiceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AksClusterResourceServiceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAksClusterResourceServiceCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAksClusterResourceServiceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAksClusterResourceServiceCreateOK creates a AksClusterResourceServiceCreateOK with default headers values
func NewAksClusterResourceServiceCreateOK() *AksClusterResourceServiceCreateOK {
	return &AksClusterResourceServiceCreateOK{}
}

/*
AksClusterResourceServiceCreateOK describes a response with status code 200, with default header values.

A successful response.
*/
type AksClusterResourceServiceCreateOK struct {
	Payload *models.AksclusterCreateAksClusterResponse
}

// IsSuccess returns true when this aks cluster resource service create o k response has a 2xx status code
func (o *AksClusterResourceServiceCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aks cluster resource service create o k response has a 3xx status code
func (o *AksClusterResourceServiceCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aks cluster resource service create o k response has a 4xx status code
func (o *AksClusterResourceServiceCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aks cluster resource service create o k response has a 5xx status code
func (o *AksClusterResourceServiceCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aks cluster resource service create o k response a status code equal to that given
func (o *AksClusterResourceServiceCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *AksClusterResourceServiceCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/aksclusters][%d] aksClusterResourceServiceCreateOK  %+v", 200, o.Payload)
}

func (o *AksClusterResourceServiceCreateOK) String() string {
	return fmt.Sprintf("[POST /v1alpha1/aksclusters][%d] aksClusterResourceServiceCreateOK  %+v", 200, o.Payload)
}

func (o *AksClusterResourceServiceCreateOK) GetPayload() *models.AksclusterCreateAksClusterResponse {
	return o.Payload
}

func (o *AksClusterResourceServiceCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AksclusterCreateAksClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAksClusterResourceServiceCreateDefault creates a AksClusterResourceServiceCreateDefault with default headers values
func NewAksClusterResourceServiceCreateDefault(code int) *AksClusterResourceServiceCreateDefault {
	return &AksClusterResourceServiceCreateDefault{
		_statusCode: code,
	}
}

/*
AksClusterResourceServiceCreateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AksClusterResourceServiceCreateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the aks cluster resource service create default response
func (o *AksClusterResourceServiceCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this aks cluster resource service create default response has a 2xx status code
func (o *AksClusterResourceServiceCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this aks cluster resource service create default response has a 3xx status code
func (o *AksClusterResourceServiceCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this aks cluster resource service create default response has a 4xx status code
func (o *AksClusterResourceServiceCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this aks cluster resource service create default response has a 5xx status code
func (o *AksClusterResourceServiceCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this aks cluster resource service create default response a status code equal to that given
func (o *AksClusterResourceServiceCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AksClusterResourceServiceCreateDefault) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/aksclusters][%d] AksClusterResourceService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *AksClusterResourceServiceCreateDefault) String() string {
	return fmt.Sprintf("[POST /v1alpha1/aksclusters][%d] AksClusterResourceService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *AksClusterResourceServiceCreateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *AksClusterResourceServiceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
