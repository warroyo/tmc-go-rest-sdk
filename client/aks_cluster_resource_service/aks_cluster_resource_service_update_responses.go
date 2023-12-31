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

// AksClusterResourceServiceUpdateReader is a Reader for the AksClusterResourceServiceUpdate structure.
type AksClusterResourceServiceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AksClusterResourceServiceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAksClusterResourceServiceUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAksClusterResourceServiceUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAksClusterResourceServiceUpdateOK creates a AksClusterResourceServiceUpdateOK with default headers values
func NewAksClusterResourceServiceUpdateOK() *AksClusterResourceServiceUpdateOK {
	return &AksClusterResourceServiceUpdateOK{}
}

/*
AksClusterResourceServiceUpdateOK describes a response with status code 200, with default header values.

A successful response.
*/
type AksClusterResourceServiceUpdateOK struct {
	Payload *models.AksclusterUpdateAksClusterResponse
}

// IsSuccess returns true when this aks cluster resource service update o k response has a 2xx status code
func (o *AksClusterResourceServiceUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aks cluster resource service update o k response has a 3xx status code
func (o *AksClusterResourceServiceUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aks cluster resource service update o k response has a 4xx status code
func (o *AksClusterResourceServiceUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aks cluster resource service update o k response has a 5xx status code
func (o *AksClusterResourceServiceUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aks cluster resource service update o k response a status code equal to that given
func (o *AksClusterResourceServiceUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *AksClusterResourceServiceUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/aksclusters/{aksCluster.fullName.name}][%d] aksClusterResourceServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *AksClusterResourceServiceUpdateOK) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/aksclusters/{aksCluster.fullName.name}][%d] aksClusterResourceServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *AksClusterResourceServiceUpdateOK) GetPayload() *models.AksclusterUpdateAksClusterResponse {
	return o.Payload
}

func (o *AksClusterResourceServiceUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AksclusterUpdateAksClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAksClusterResourceServiceUpdateDefault creates a AksClusterResourceServiceUpdateDefault with default headers values
func NewAksClusterResourceServiceUpdateDefault(code int) *AksClusterResourceServiceUpdateDefault {
	return &AksClusterResourceServiceUpdateDefault{
		_statusCode: code,
	}
}

/*
AksClusterResourceServiceUpdateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AksClusterResourceServiceUpdateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the aks cluster resource service update default response
func (o *AksClusterResourceServiceUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this aks cluster resource service update default response has a 2xx status code
func (o *AksClusterResourceServiceUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this aks cluster resource service update default response has a 3xx status code
func (o *AksClusterResourceServiceUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this aks cluster resource service update default response has a 4xx status code
func (o *AksClusterResourceServiceUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this aks cluster resource service update default response has a 5xx status code
func (o *AksClusterResourceServiceUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this aks cluster resource service update default response a status code equal to that given
func (o *AksClusterResourceServiceUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AksClusterResourceServiceUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/aksclusters/{aksCluster.fullName.name}][%d] AksClusterResourceService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *AksClusterResourceServiceUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/aksclusters/{aksCluster.fullName.name}][%d] AksClusterResourceService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *AksClusterResourceServiceUpdateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *AksClusterResourceServiceUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
