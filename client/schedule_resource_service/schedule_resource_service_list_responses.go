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

// ScheduleResourceServiceListReader is a Reader for the ScheduleResourceServiceList structure.
type ScheduleResourceServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScheduleResourceServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewScheduleResourceServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewScheduleResourceServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewScheduleResourceServiceListOK creates a ScheduleResourceServiceListOK with default headers values
func NewScheduleResourceServiceListOK() *ScheduleResourceServiceListOK {
	return &ScheduleResourceServiceListOK{}
}

/*
ScheduleResourceServiceListOK describes a response with status code 200, with default header values.

A successful response.
*/
type ScheduleResourceServiceListOK struct {
	Payload *models.ClusterInspectionScheduleListSchedulesResponse
}

// IsSuccess returns true when this schedule resource service list o k response has a 2xx status code
func (o *ScheduleResourceServiceListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this schedule resource service list o k response has a 3xx status code
func (o *ScheduleResourceServiceListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule resource service list o k response has a 4xx status code
func (o *ScheduleResourceServiceListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this schedule resource service list o k response has a 5xx status code
func (o *ScheduleResourceServiceListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule resource service list o k response a status code equal to that given
func (o *ScheduleResourceServiceListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ScheduleResourceServiceListOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{searchScope.clusterName}/inspection/schedules][%d] scheduleResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *ScheduleResourceServiceListOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{searchScope.clusterName}/inspection/schedules][%d] scheduleResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *ScheduleResourceServiceListOK) GetPayload() *models.ClusterInspectionScheduleListSchedulesResponse {
	return o.Payload
}

func (o *ScheduleResourceServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterInspectionScheduleListSchedulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScheduleResourceServiceListDefault creates a ScheduleResourceServiceListDefault with default headers values
func NewScheduleResourceServiceListDefault(code int) *ScheduleResourceServiceListDefault {
	return &ScheduleResourceServiceListDefault{
		_statusCode: code,
	}
}

/*
ScheduleResourceServiceListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ScheduleResourceServiceListDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the schedule resource service list default response
func (o *ScheduleResourceServiceListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this schedule resource service list default response has a 2xx status code
func (o *ScheduleResourceServiceListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this schedule resource service list default response has a 3xx status code
func (o *ScheduleResourceServiceListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this schedule resource service list default response has a 4xx status code
func (o *ScheduleResourceServiceListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this schedule resource service list default response has a 5xx status code
func (o *ScheduleResourceServiceListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this schedule resource service list default response a status code equal to that given
func (o *ScheduleResourceServiceListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ScheduleResourceServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{searchScope.clusterName}/inspection/schedules][%d] ScheduleResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *ScheduleResourceServiceListDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{searchScope.clusterName}/inspection/schedules][%d] ScheduleResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *ScheduleResourceServiceListDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ScheduleResourceServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
