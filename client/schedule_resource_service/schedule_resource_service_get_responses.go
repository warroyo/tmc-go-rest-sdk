// Code generated by go-swagger; DO NOT EDIT.

package schedule_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// ScheduleResourceServiceGetReader is a Reader for the ScheduleResourceServiceGet structure.
type ScheduleResourceServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScheduleResourceServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewScheduleResourceServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewScheduleResourceServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewScheduleResourceServiceGetOK creates a ScheduleResourceServiceGetOK with default headers values
func NewScheduleResourceServiceGetOK() *ScheduleResourceServiceGetOK {
	return &ScheduleResourceServiceGetOK{}
}

/*
ScheduleResourceServiceGetOK describes a response with status code 200, with default header values.

A successful response.
*/
type ScheduleResourceServiceGetOK struct {
	Payload *models.ClusterInspectionScheduleGetScheduleResponse
}

// IsSuccess returns true when this schedule resource service get o k response has a 2xx status code
func (o *ScheduleResourceServiceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this schedule resource service get o k response has a 3xx status code
func (o *ScheduleResourceServiceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule resource service get o k response has a 4xx status code
func (o *ScheduleResourceServiceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this schedule resource service get o k response has a 5xx status code
func (o *ScheduleResourceServiceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule resource service get o k response a status code equal to that given
func (o *ScheduleResourceServiceGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *ScheduleResourceServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.clusterName}/inspection/schedules/{fullName.name}][%d] scheduleResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *ScheduleResourceServiceGetOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.clusterName}/inspection/schedules/{fullName.name}][%d] scheduleResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *ScheduleResourceServiceGetOK) GetPayload() *models.ClusterInspectionScheduleGetScheduleResponse {
	return o.Payload
}

func (o *ScheduleResourceServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterInspectionScheduleGetScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScheduleResourceServiceGetDefault creates a ScheduleResourceServiceGetDefault with default headers values
func NewScheduleResourceServiceGetDefault(code int) *ScheduleResourceServiceGetDefault {
	return &ScheduleResourceServiceGetDefault{
		_statusCode: code,
	}
}

/*
ScheduleResourceServiceGetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ScheduleResourceServiceGetDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the schedule resource service get default response
func (o *ScheduleResourceServiceGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this schedule resource service get default response has a 2xx status code
func (o *ScheduleResourceServiceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this schedule resource service get default response has a 3xx status code
func (o *ScheduleResourceServiceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this schedule resource service get default response has a 4xx status code
func (o *ScheduleResourceServiceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this schedule resource service get default response has a 5xx status code
func (o *ScheduleResourceServiceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this schedule resource service get default response a status code equal to that given
func (o *ScheduleResourceServiceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ScheduleResourceServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.clusterName}/inspection/schedules/{fullName.name}][%d] ScheduleResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *ScheduleResourceServiceGetDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.clusterName}/inspection/schedules/{fullName.name}][%d] ScheduleResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *ScheduleResourceServiceGetDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ScheduleResourceServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
