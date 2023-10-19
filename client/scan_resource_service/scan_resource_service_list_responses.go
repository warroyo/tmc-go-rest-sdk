// Code generated by go-swagger; DO NOT EDIT.

package scan_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// ScanResourceServiceListReader is a Reader for the ScanResourceServiceList structure.
type ScanResourceServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScanResourceServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewScanResourceServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewScanResourceServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewScanResourceServiceListOK creates a ScanResourceServiceListOK with default headers values
func NewScanResourceServiceListOK() *ScanResourceServiceListOK {
	return &ScanResourceServiceListOK{}
}

/*
ScanResourceServiceListOK describes a response with status code 200, with default header values.

A successful response.
*/
type ScanResourceServiceListOK struct {
	Payload *models.ClusterInspectionScanListScansResponse
}

// IsSuccess returns true when this scan resource service list o k response has a 2xx status code
func (o *ScanResourceServiceListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this scan resource service list o k response has a 3xx status code
func (o *ScanResourceServiceListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this scan resource service list o k response has a 4xx status code
func (o *ScanResourceServiceListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this scan resource service list o k response has a 5xx status code
func (o *ScanResourceServiceListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this scan resource service list o k response a status code equal to that given
func (o *ScanResourceServiceListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ScanResourceServiceListOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{searchScope.clusterName}/inspection/scans][%d] scanResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *ScanResourceServiceListOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{searchScope.clusterName}/inspection/scans][%d] scanResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *ScanResourceServiceListOK) GetPayload() *models.ClusterInspectionScanListScansResponse {
	return o.Payload
}

func (o *ScanResourceServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterInspectionScanListScansResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScanResourceServiceListDefault creates a ScanResourceServiceListDefault with default headers values
func NewScanResourceServiceListDefault(code int) *ScanResourceServiceListDefault {
	return &ScanResourceServiceListDefault{
		_statusCode: code,
	}
}

/*
ScanResourceServiceListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ScanResourceServiceListDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the scan resource service list default response
func (o *ScanResourceServiceListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this scan resource service list default response has a 2xx status code
func (o *ScanResourceServiceListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this scan resource service list default response has a 3xx status code
func (o *ScanResourceServiceListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this scan resource service list default response has a 4xx status code
func (o *ScanResourceServiceListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this scan resource service list default response has a 5xx status code
func (o *ScanResourceServiceListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this scan resource service list default response a status code equal to that given
func (o *ScanResourceServiceListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ScanResourceServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{searchScope.clusterName}/inspection/scans][%d] ScanResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *ScanResourceServiceListDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{searchScope.clusterName}/inspection/scans][%d] ScanResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *ScanResourceServiceListDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ScanResourceServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}