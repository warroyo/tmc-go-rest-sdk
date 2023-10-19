// Code generated by go-swagger; DO NOT EDIT.

package restore_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// RestoreResourceServiceGetReader is a Reader for the RestoreResourceServiceGet structure.
type RestoreResourceServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestoreResourceServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestoreResourceServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRestoreResourceServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRestoreResourceServiceGetOK creates a RestoreResourceServiceGetOK with default headers values
func NewRestoreResourceServiceGetOK() *RestoreResourceServiceGetOK {
	return &RestoreResourceServiceGetOK{}
}

/*
RestoreResourceServiceGetOK describes a response with status code 200, with default header values.

A successful response.
*/
type RestoreResourceServiceGetOK struct {
	Payload *models.ClusterDataprotectionRestoreGetRestoreResponse
}

// IsSuccess returns true when this restore resource service get o k response has a 2xx status code
func (o *RestoreResourceServiceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this restore resource service get o k response has a 3xx status code
func (o *RestoreResourceServiceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore resource service get o k response has a 4xx status code
func (o *RestoreResourceServiceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this restore resource service get o k response has a 5xx status code
func (o *RestoreResourceServiceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this restore resource service get o k response a status code equal to that given
func (o *RestoreResourceServiceGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *RestoreResourceServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.clusterName}/dataprotection/restores/{fullName.name}][%d] restoreResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *RestoreResourceServiceGetOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.clusterName}/dataprotection/restores/{fullName.name}][%d] restoreResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *RestoreResourceServiceGetOK) GetPayload() *models.ClusterDataprotectionRestoreGetRestoreResponse {
	return o.Payload
}

func (o *RestoreResourceServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterDataprotectionRestoreGetRestoreResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreResourceServiceGetDefault creates a RestoreResourceServiceGetDefault with default headers values
func NewRestoreResourceServiceGetDefault(code int) *RestoreResourceServiceGetDefault {
	return &RestoreResourceServiceGetDefault{
		_statusCode: code,
	}
}

/*
RestoreResourceServiceGetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RestoreResourceServiceGetDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the restore resource service get default response
func (o *RestoreResourceServiceGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this restore resource service get default response has a 2xx status code
func (o *RestoreResourceServiceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this restore resource service get default response has a 3xx status code
func (o *RestoreResourceServiceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this restore resource service get default response has a 4xx status code
func (o *RestoreResourceServiceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this restore resource service get default response has a 5xx status code
func (o *RestoreResourceServiceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this restore resource service get default response a status code equal to that given
func (o *RestoreResourceServiceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *RestoreResourceServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.clusterName}/dataprotection/restores/{fullName.name}][%d] RestoreResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *RestoreResourceServiceGetDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.clusterName}/dataprotection/restores/{fullName.name}][%d] RestoreResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *RestoreResourceServiceGetDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *RestoreResourceServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
