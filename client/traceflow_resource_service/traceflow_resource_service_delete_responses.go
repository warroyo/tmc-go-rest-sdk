// Code generated by go-swagger; DO NOT EDIT.

package traceflow_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// TraceflowResourceServiceDeleteReader is a Reader for the TraceflowResourceServiceDelete structure.
type TraceflowResourceServiceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TraceflowResourceServiceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTraceflowResourceServiceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTraceflowResourceServiceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTraceflowResourceServiceDeleteOK creates a TraceflowResourceServiceDeleteOK with default headers values
func NewTraceflowResourceServiceDeleteOK() *TraceflowResourceServiceDeleteOK {
	return &TraceflowResourceServiceDeleteOK{}
}

/*
TraceflowResourceServiceDeleteOK describes a response with status code 200, with default header values.

A successful response.
*/
type TraceflowResourceServiceDeleteOK struct {
	Payload *models.ClusterNetworkAntreaTraceflowDeleteTraceflowResponse
}

// IsSuccess returns true when this traceflow resource service delete o k response has a 2xx status code
func (o *TraceflowResourceServiceDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this traceflow resource service delete o k response has a 3xx status code
func (o *TraceflowResourceServiceDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this traceflow resource service delete o k response has a 4xx status code
func (o *TraceflowResourceServiceDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this traceflow resource service delete o k response has a 5xx status code
func (o *TraceflowResourceServiceDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this traceflow resource service delete o k response a status code equal to that given
func (o *TraceflowResourceServiceDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *TraceflowResourceServiceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clusters/{fullName.clusterName}/network/antrea/traceflows/{fullName.name}][%d] traceflowResourceServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *TraceflowResourceServiceDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clusters/{fullName.clusterName}/network/antrea/traceflows/{fullName.name}][%d] traceflowResourceServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *TraceflowResourceServiceDeleteOK) GetPayload() *models.ClusterNetworkAntreaTraceflowDeleteTraceflowResponse {
	return o.Payload
}

func (o *TraceflowResourceServiceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterNetworkAntreaTraceflowDeleteTraceflowResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTraceflowResourceServiceDeleteDefault creates a TraceflowResourceServiceDeleteDefault with default headers values
func NewTraceflowResourceServiceDeleteDefault(code int) *TraceflowResourceServiceDeleteDefault {
	return &TraceflowResourceServiceDeleteDefault{
		_statusCode: code,
	}
}

/*
TraceflowResourceServiceDeleteDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type TraceflowResourceServiceDeleteDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the traceflow resource service delete default response
func (o *TraceflowResourceServiceDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this traceflow resource service delete default response has a 2xx status code
func (o *TraceflowResourceServiceDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this traceflow resource service delete default response has a 3xx status code
func (o *TraceflowResourceServiceDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this traceflow resource service delete default response has a 4xx status code
func (o *TraceflowResourceServiceDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this traceflow resource service delete default response has a 5xx status code
func (o *TraceflowResourceServiceDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this traceflow resource service delete default response a status code equal to that given
func (o *TraceflowResourceServiceDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *TraceflowResourceServiceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clusters/{fullName.clusterName}/network/antrea/traceflows/{fullName.name}][%d] TraceflowResourceService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *TraceflowResourceServiceDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clusters/{fullName.clusterName}/network/antrea/traceflows/{fullName.name}][%d] TraceflowResourceService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *TraceflowResourceServiceDeleteDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *TraceflowResourceServiceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
