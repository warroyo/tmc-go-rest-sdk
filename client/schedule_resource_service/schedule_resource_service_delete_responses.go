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

// ScheduleResourceServiceDeleteReader is a Reader for the ScheduleResourceServiceDelete structure.
type ScheduleResourceServiceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScheduleResourceServiceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewScheduleResourceServiceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewScheduleResourceServiceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewScheduleResourceServiceDeleteOK creates a ScheduleResourceServiceDeleteOK with default headers values
func NewScheduleResourceServiceDeleteOK() *ScheduleResourceServiceDeleteOK {
	return &ScheduleResourceServiceDeleteOK{}
}

/*
ScheduleResourceServiceDeleteOK describes a response with status code 200, with default header values.

A successful response.
*/
type ScheduleResourceServiceDeleteOK struct {
	Payload *models.ClusterInspectionScheduleDeleteScheduleResponse
}

// IsSuccess returns true when this schedule resource service delete o k response has a 2xx status code
func (o *ScheduleResourceServiceDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this schedule resource service delete o k response has a 3xx status code
func (o *ScheduleResourceServiceDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule resource service delete o k response has a 4xx status code
func (o *ScheduleResourceServiceDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this schedule resource service delete o k response has a 5xx status code
func (o *ScheduleResourceServiceDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule resource service delete o k response a status code equal to that given
func (o *ScheduleResourceServiceDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *ScheduleResourceServiceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clusters/{fullName.clusterName}/inspection/schedules/{fullName.name}][%d] scheduleResourceServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *ScheduleResourceServiceDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clusters/{fullName.clusterName}/inspection/schedules/{fullName.name}][%d] scheduleResourceServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *ScheduleResourceServiceDeleteOK) GetPayload() *models.ClusterInspectionScheduleDeleteScheduleResponse {
	return o.Payload
}

func (o *ScheduleResourceServiceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterInspectionScheduleDeleteScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScheduleResourceServiceDeleteDefault creates a ScheduleResourceServiceDeleteDefault with default headers values
func NewScheduleResourceServiceDeleteDefault(code int) *ScheduleResourceServiceDeleteDefault {
	return &ScheduleResourceServiceDeleteDefault{
		_statusCode: code,
	}
}

/*
ScheduleResourceServiceDeleteDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ScheduleResourceServiceDeleteDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the schedule resource service delete default response
func (o *ScheduleResourceServiceDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this schedule resource service delete default response has a 2xx status code
func (o *ScheduleResourceServiceDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this schedule resource service delete default response has a 3xx status code
func (o *ScheduleResourceServiceDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this schedule resource service delete default response has a 4xx status code
func (o *ScheduleResourceServiceDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this schedule resource service delete default response has a 5xx status code
func (o *ScheduleResourceServiceDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this schedule resource service delete default response a status code equal to that given
func (o *ScheduleResourceServiceDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ScheduleResourceServiceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clusters/{fullName.clusterName}/inspection/schedules/{fullName.name}][%d] ScheduleResourceService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *ScheduleResourceServiceDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clusters/{fullName.clusterName}/inspection/schedules/{fullName.name}][%d] ScheduleResourceService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *ScheduleResourceServiceDeleteDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ScheduleResourceServiceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
