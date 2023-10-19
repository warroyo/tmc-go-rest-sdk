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

// AksClusterResourceServicePatchReader is a Reader for the AksClusterResourceServicePatch structure.
type AksClusterResourceServicePatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AksClusterResourceServicePatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAksClusterResourceServicePatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAksClusterResourceServicePatchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAksClusterResourceServicePatchOK creates a AksClusterResourceServicePatchOK with default headers values
func NewAksClusterResourceServicePatchOK() *AksClusterResourceServicePatchOK {
	return &AksClusterResourceServicePatchOK{}
}

/*
AksClusterResourceServicePatchOK describes a response with status code 200, with default header values.

A successful response.
*/
type AksClusterResourceServicePatchOK struct {
	Payload *models.AksclusterPatchAksClusterResponse
}

// IsSuccess returns true when this aks cluster resource service patch o k response has a 2xx status code
func (o *AksClusterResourceServicePatchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aks cluster resource service patch o k response has a 3xx status code
func (o *AksClusterResourceServicePatchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aks cluster resource service patch o k response has a 4xx status code
func (o *AksClusterResourceServicePatchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aks cluster resource service patch o k response has a 5xx status code
func (o *AksClusterResourceServicePatchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aks cluster resource service patch o k response a status code equal to that given
func (o *AksClusterResourceServicePatchOK) IsCode(code int) bool {
	return code == 200
}

func (o *AksClusterResourceServicePatchOK) Error() string {
	return fmt.Sprintf("[PATCH /v1alpha1/aksclusters/{fullName.name}][%d] aksClusterResourceServicePatchOK  %+v", 200, o.Payload)
}

func (o *AksClusterResourceServicePatchOK) String() string {
	return fmt.Sprintf("[PATCH /v1alpha1/aksclusters/{fullName.name}][%d] aksClusterResourceServicePatchOK  %+v", 200, o.Payload)
}

func (o *AksClusterResourceServicePatchOK) GetPayload() *models.AksclusterPatchAksClusterResponse {
	return o.Payload
}

func (o *AksClusterResourceServicePatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AksclusterPatchAksClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAksClusterResourceServicePatchDefault creates a AksClusterResourceServicePatchDefault with default headers values
func NewAksClusterResourceServicePatchDefault(code int) *AksClusterResourceServicePatchDefault {
	return &AksClusterResourceServicePatchDefault{
		_statusCode: code,
	}
}

/*
AksClusterResourceServicePatchDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AksClusterResourceServicePatchDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the aks cluster resource service patch default response
func (o *AksClusterResourceServicePatchDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this aks cluster resource service patch default response has a 2xx status code
func (o *AksClusterResourceServicePatchDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this aks cluster resource service patch default response has a 3xx status code
func (o *AksClusterResourceServicePatchDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this aks cluster resource service patch default response has a 4xx status code
func (o *AksClusterResourceServicePatchDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this aks cluster resource service patch default response has a 5xx status code
func (o *AksClusterResourceServicePatchDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this aks cluster resource service patch default response a status code equal to that given
func (o *AksClusterResourceServicePatchDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AksClusterResourceServicePatchDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1alpha1/aksclusters/{fullName.name}][%d] AksClusterResourceService_Patch default  %+v", o._statusCode, o.Payload)
}

func (o *AksClusterResourceServicePatchDefault) String() string {
	return fmt.Sprintf("[PATCH /v1alpha1/aksclusters/{fullName.name}][%d] AksClusterResourceService_Patch default  %+v", o._statusCode, o.Payload)
}

func (o *AksClusterResourceServicePatchDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *AksClusterResourceServicePatchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
