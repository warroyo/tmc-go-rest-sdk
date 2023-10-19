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

// ScheduleResourceServiceCreateReader is a Reader for the ScheduleResourceServiceCreate structure.
type ScheduleResourceServiceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScheduleResourceServiceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewScheduleResourceServiceCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewScheduleResourceServiceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewScheduleResourceServiceCreateOK creates a ScheduleResourceServiceCreateOK with default headers values
func NewScheduleResourceServiceCreateOK() *ScheduleResourceServiceCreateOK {
	return &ScheduleResourceServiceCreateOK{}
}

/*
ScheduleResourceServiceCreateOK describes a response with status code 200, with default header values.

A successful response.
*/
type ScheduleResourceServiceCreateOK struct {
	Payload *models.ClusterInspectionScheduleCreateScheduleResponse
}

// IsSuccess returns true when this schedule resource service create o k response has a 2xx status code
func (o *ScheduleResourceServiceCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this schedule resource service create o k response has a 3xx status code
func (o *ScheduleResourceServiceCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule resource service create o k response has a 4xx status code
func (o *ScheduleResourceServiceCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this schedule resource service create o k response has a 5xx status code
func (o *ScheduleResourceServiceCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule resource service create o k response a status code equal to that given
func (o *ScheduleResourceServiceCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ScheduleResourceServiceCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/clusters/{schedule.fullName.clusterName}/inspection/schedules][%d] scheduleResourceServiceCreateOK  %+v", 200, o.Payload)
}

func (o *ScheduleResourceServiceCreateOK) String() string {
	return fmt.Sprintf("[POST /v1alpha1/clusters/{schedule.fullName.clusterName}/inspection/schedules][%d] scheduleResourceServiceCreateOK  %+v", 200, o.Payload)
}

func (o *ScheduleResourceServiceCreateOK) GetPayload() *models.ClusterInspectionScheduleCreateScheduleResponse {
	return o.Payload
}

func (o *ScheduleResourceServiceCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterInspectionScheduleCreateScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScheduleResourceServiceCreateDefault creates a ScheduleResourceServiceCreateDefault with default headers values
func NewScheduleResourceServiceCreateDefault(code int) *ScheduleResourceServiceCreateDefault {
	return &ScheduleResourceServiceCreateDefault{
		_statusCode: code,
	}
}

/*
ScheduleResourceServiceCreateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ScheduleResourceServiceCreateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the schedule resource service create default response
func (o *ScheduleResourceServiceCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this schedule resource service create default response has a 2xx status code
func (o *ScheduleResourceServiceCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this schedule resource service create default response has a 3xx status code
func (o *ScheduleResourceServiceCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this schedule resource service create default response has a 4xx status code
func (o *ScheduleResourceServiceCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this schedule resource service create default response has a 5xx status code
func (o *ScheduleResourceServiceCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this schedule resource service create default response a status code equal to that given
func (o *ScheduleResourceServiceCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ScheduleResourceServiceCreateDefault) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/clusters/{schedule.fullName.clusterName}/inspection/schedules][%d] ScheduleResourceService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *ScheduleResourceServiceCreateDefault) String() string {
	return fmt.Sprintf("[POST /v1alpha1/clusters/{schedule.fullName.clusterName}/inspection/schedules][%d] ScheduleResourceService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *ScheduleResourceServiceCreateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ScheduleResourceServiceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
