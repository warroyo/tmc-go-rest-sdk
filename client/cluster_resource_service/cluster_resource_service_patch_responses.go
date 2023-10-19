// Code generated by go-swagger; DO NOT EDIT.

package cluster_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// ClusterResourceServicePatchReader is a Reader for the ClusterResourceServicePatch structure.
type ClusterResourceServicePatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterResourceServicePatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterResourceServicePatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterResourceServicePatchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterResourceServicePatchOK creates a ClusterResourceServicePatchOK with default headers values
func NewClusterResourceServicePatchOK() *ClusterResourceServicePatchOK {
	return &ClusterResourceServicePatchOK{}
}

/*
ClusterResourceServicePatchOK describes a response with status code 200, with default header values.

A successful response.
*/
type ClusterResourceServicePatchOK struct {
	Payload *models.ClusterPatchClusterResponse
}

// IsSuccess returns true when this cluster resource service patch o k response has a 2xx status code
func (o *ClusterResourceServicePatchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster resource service patch o k response has a 3xx status code
func (o *ClusterResourceServicePatchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster resource service patch o k response has a 4xx status code
func (o *ClusterResourceServicePatchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster resource service patch o k response has a 5xx status code
func (o *ClusterResourceServicePatchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster resource service patch o k response a status code equal to that given
func (o *ClusterResourceServicePatchOK) IsCode(code int) bool {
	return code == 200
}

func (o *ClusterResourceServicePatchOK) Error() string {
	return fmt.Sprintf("[PATCH /v1alpha1/clusters/{fullName.name}][%d] clusterResourceServicePatchOK  %+v", 200, o.Payload)
}

func (o *ClusterResourceServicePatchOK) String() string {
	return fmt.Sprintf("[PATCH /v1alpha1/clusters/{fullName.name}][%d] clusterResourceServicePatchOK  %+v", 200, o.Payload)
}

func (o *ClusterResourceServicePatchOK) GetPayload() *models.ClusterPatchClusterResponse {
	return o.Payload
}

func (o *ClusterResourceServicePatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterPatchClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterResourceServicePatchDefault creates a ClusterResourceServicePatchDefault with default headers values
func NewClusterResourceServicePatchDefault(code int) *ClusterResourceServicePatchDefault {
	return &ClusterResourceServicePatchDefault{
		_statusCode: code,
	}
}

/*
ClusterResourceServicePatchDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ClusterResourceServicePatchDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the cluster resource service patch default response
func (o *ClusterResourceServicePatchDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cluster resource service patch default response has a 2xx status code
func (o *ClusterResourceServicePatchDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster resource service patch default response has a 3xx status code
func (o *ClusterResourceServicePatchDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster resource service patch default response has a 4xx status code
func (o *ClusterResourceServicePatchDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster resource service patch default response has a 5xx status code
func (o *ClusterResourceServicePatchDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster resource service patch default response a status code equal to that given
func (o *ClusterResourceServicePatchDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ClusterResourceServicePatchDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1alpha1/clusters/{fullName.name}][%d] ClusterResourceService_Patch default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterResourceServicePatchDefault) String() string {
	return fmt.Sprintf("[PATCH /v1alpha1/clusters/{fullName.name}][%d] ClusterResourceService_Patch default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterResourceServicePatchDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ClusterResourceServicePatchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
