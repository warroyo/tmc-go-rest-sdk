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

// RestoreResourceServiceDeleteReader is a Reader for the RestoreResourceServiceDelete structure.
type RestoreResourceServiceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestoreResourceServiceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestoreResourceServiceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRestoreResourceServiceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRestoreResourceServiceDeleteOK creates a RestoreResourceServiceDeleteOK with default headers values
func NewRestoreResourceServiceDeleteOK() *RestoreResourceServiceDeleteOK {
	return &RestoreResourceServiceDeleteOK{}
}

/*
RestoreResourceServiceDeleteOK describes a response with status code 200, with default header values.

A successful response.
*/
type RestoreResourceServiceDeleteOK struct {
	Payload *models.ClusterDataprotectionRestoreDeleteRestoreResponse
}

// IsSuccess returns true when this restore resource service delete o k response has a 2xx status code
func (o *RestoreResourceServiceDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this restore resource service delete o k response has a 3xx status code
func (o *RestoreResourceServiceDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore resource service delete o k response has a 4xx status code
func (o *RestoreResourceServiceDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this restore resource service delete o k response has a 5xx status code
func (o *RestoreResourceServiceDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this restore resource service delete o k response a status code equal to that given
func (o *RestoreResourceServiceDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *RestoreResourceServiceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clusters/{fullName.clusterName}/dataprotection/restores/{fullName.name}][%d] restoreResourceServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *RestoreResourceServiceDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clusters/{fullName.clusterName}/dataprotection/restores/{fullName.name}][%d] restoreResourceServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *RestoreResourceServiceDeleteOK) GetPayload() *models.ClusterDataprotectionRestoreDeleteRestoreResponse {
	return o.Payload
}

func (o *RestoreResourceServiceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterDataprotectionRestoreDeleteRestoreResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreResourceServiceDeleteDefault creates a RestoreResourceServiceDeleteDefault with default headers values
func NewRestoreResourceServiceDeleteDefault(code int) *RestoreResourceServiceDeleteDefault {
	return &RestoreResourceServiceDeleteDefault{
		_statusCode: code,
	}
}

/*
RestoreResourceServiceDeleteDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RestoreResourceServiceDeleteDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the restore resource service delete default response
func (o *RestoreResourceServiceDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this restore resource service delete default response has a 2xx status code
func (o *RestoreResourceServiceDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this restore resource service delete default response has a 3xx status code
func (o *RestoreResourceServiceDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this restore resource service delete default response has a 4xx status code
func (o *RestoreResourceServiceDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this restore resource service delete default response has a 5xx status code
func (o *RestoreResourceServiceDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this restore resource service delete default response a status code equal to that given
func (o *RestoreResourceServiceDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *RestoreResourceServiceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clusters/{fullName.clusterName}/dataprotection/restores/{fullName.name}][%d] RestoreResourceService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *RestoreResourceServiceDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clusters/{fullName.clusterName}/dataprotection/restores/{fullName.name}][%d] RestoreResourceService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *RestoreResourceServiceDeleteDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *RestoreResourceServiceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
