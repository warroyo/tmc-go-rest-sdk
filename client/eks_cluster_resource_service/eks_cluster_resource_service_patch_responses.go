// Code generated by go-swagger; DO NOT EDIT.

package eks_cluster_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// EksClusterResourceServicePatchReader is a Reader for the EksClusterResourceServicePatch structure.
type EksClusterResourceServicePatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EksClusterResourceServicePatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEksClusterResourceServicePatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEksClusterResourceServicePatchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEksClusterResourceServicePatchOK creates a EksClusterResourceServicePatchOK with default headers values
func NewEksClusterResourceServicePatchOK() *EksClusterResourceServicePatchOK {
	return &EksClusterResourceServicePatchOK{}
}

/*
EksClusterResourceServicePatchOK describes a response with status code 200, with default header values.

A successful response.
*/
type EksClusterResourceServicePatchOK struct {
	Payload *models.EksclusterPatchEksClusterResponse
}

// IsSuccess returns true when this eks cluster resource service patch o k response has a 2xx status code
func (o *EksClusterResourceServicePatchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this eks cluster resource service patch o k response has a 3xx status code
func (o *EksClusterResourceServicePatchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this eks cluster resource service patch o k response has a 4xx status code
func (o *EksClusterResourceServicePatchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this eks cluster resource service patch o k response has a 5xx status code
func (o *EksClusterResourceServicePatchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this eks cluster resource service patch o k response a status code equal to that given
func (o *EksClusterResourceServicePatchOK) IsCode(code int) bool {
	return code == 200
}

func (o *EksClusterResourceServicePatchOK) Error() string {
	return fmt.Sprintf("[PATCH /v1alpha1/eksclusters/{fullName.name}][%d] eksClusterResourceServicePatchOK  %+v", 200, o.Payload)
}

func (o *EksClusterResourceServicePatchOK) String() string {
	return fmt.Sprintf("[PATCH /v1alpha1/eksclusters/{fullName.name}][%d] eksClusterResourceServicePatchOK  %+v", 200, o.Payload)
}

func (o *EksClusterResourceServicePatchOK) GetPayload() *models.EksclusterPatchEksClusterResponse {
	return o.Payload
}

func (o *EksClusterResourceServicePatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EksclusterPatchEksClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEksClusterResourceServicePatchDefault creates a EksClusterResourceServicePatchDefault with default headers values
func NewEksClusterResourceServicePatchDefault(code int) *EksClusterResourceServicePatchDefault {
	return &EksClusterResourceServicePatchDefault{
		_statusCode: code,
	}
}

/*
EksClusterResourceServicePatchDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EksClusterResourceServicePatchDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the eks cluster resource service patch default response
func (o *EksClusterResourceServicePatchDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this eks cluster resource service patch default response has a 2xx status code
func (o *EksClusterResourceServicePatchDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this eks cluster resource service patch default response has a 3xx status code
func (o *EksClusterResourceServicePatchDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this eks cluster resource service patch default response has a 4xx status code
func (o *EksClusterResourceServicePatchDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this eks cluster resource service patch default response has a 5xx status code
func (o *EksClusterResourceServicePatchDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this eks cluster resource service patch default response a status code equal to that given
func (o *EksClusterResourceServicePatchDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EksClusterResourceServicePatchDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1alpha1/eksclusters/{fullName.name}][%d] EksClusterResourceService_Patch default  %+v", o._statusCode, o.Payload)
}

func (o *EksClusterResourceServicePatchDefault) String() string {
	return fmt.Sprintf("[PATCH /v1alpha1/eksclusters/{fullName.name}][%d] EksClusterResourceService_Patch default  %+v", o._statusCode, o.Payload)
}

func (o *EksClusterResourceServicePatchDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *EksClusterResourceServicePatchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
