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

// RestoreResourceServiceListReader is a Reader for the RestoreResourceServiceList structure.
type RestoreResourceServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestoreResourceServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestoreResourceServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRestoreResourceServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRestoreResourceServiceListOK creates a RestoreResourceServiceListOK with default headers values
func NewRestoreResourceServiceListOK() *RestoreResourceServiceListOK {
	return &RestoreResourceServiceListOK{}
}

/*
RestoreResourceServiceListOK describes a response with status code 200, with default header values.

A successful response.
*/
type RestoreResourceServiceListOK struct {
	Payload *models.ClusterDataprotectionRestoreListRestoresResponse
}

// IsSuccess returns true when this restore resource service list o k response has a 2xx status code
func (o *RestoreResourceServiceListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this restore resource service list o k response has a 3xx status code
func (o *RestoreResourceServiceListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore resource service list o k response has a 4xx status code
func (o *RestoreResourceServiceListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this restore resource service list o k response has a 5xx status code
func (o *RestoreResourceServiceListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this restore resource service list o k response a status code equal to that given
func (o *RestoreResourceServiceListOK) IsCode(code int) bool {
	return code == 200
}

func (o *RestoreResourceServiceListOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{searchScope.clusterName}/dataprotection/restores][%d] restoreResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *RestoreResourceServiceListOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{searchScope.clusterName}/dataprotection/restores][%d] restoreResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *RestoreResourceServiceListOK) GetPayload() *models.ClusterDataprotectionRestoreListRestoresResponse {
	return o.Payload
}

func (o *RestoreResourceServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterDataprotectionRestoreListRestoresResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreResourceServiceListDefault creates a RestoreResourceServiceListDefault with default headers values
func NewRestoreResourceServiceListDefault(code int) *RestoreResourceServiceListDefault {
	return &RestoreResourceServiceListDefault{
		_statusCode: code,
	}
}

/*
RestoreResourceServiceListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RestoreResourceServiceListDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the restore resource service list default response
func (o *RestoreResourceServiceListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this restore resource service list default response has a 2xx status code
func (o *RestoreResourceServiceListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this restore resource service list default response has a 3xx status code
func (o *RestoreResourceServiceListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this restore resource service list default response has a 4xx status code
func (o *RestoreResourceServiceListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this restore resource service list default response has a 5xx status code
func (o *RestoreResourceServiceListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this restore resource service list default response a status code equal to that given
func (o *RestoreResourceServiceListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *RestoreResourceServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{searchScope.clusterName}/dataprotection/restores][%d] RestoreResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *RestoreResourceServiceListDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{searchScope.clusterName}/dataprotection/restores][%d] RestoreResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *RestoreResourceServiceListDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *RestoreResourceServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
