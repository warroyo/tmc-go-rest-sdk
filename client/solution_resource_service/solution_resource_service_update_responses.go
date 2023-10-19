// Code generated by go-swagger; DO NOT EDIT.

package solution_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// SolutionResourceServiceUpdateReader is a Reader for the SolutionResourceServiceUpdate structure.
type SolutionResourceServiceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SolutionResourceServiceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSolutionResourceServiceUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSolutionResourceServiceUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSolutionResourceServiceUpdateOK creates a SolutionResourceServiceUpdateOK with default headers values
func NewSolutionResourceServiceUpdateOK() *SolutionResourceServiceUpdateOK {
	return &SolutionResourceServiceUpdateOK{}
}

/*
SolutionResourceServiceUpdateOK describes a response with status code 200, with default header values.

A successful response.
*/
type SolutionResourceServiceUpdateOK struct {
	Payload *models.TanzupackageTapSolutionUpdateSolutionResponse
}

// IsSuccess returns true when this solution resource service update o k response has a 2xx status code
func (o *SolutionResourceServiceUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this solution resource service update o k response has a 3xx status code
func (o *SolutionResourceServiceUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this solution resource service update o k response has a 4xx status code
func (o *SolutionResourceServiceUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this solution resource service update o k response has a 5xx status code
func (o *SolutionResourceServiceUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this solution resource service update o k response a status code equal to that given
func (o *SolutionResourceServiceUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *SolutionResourceServiceUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/tanzupackage/tap/solutions/{solution.fullName.name}][%d] solutionResourceServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *SolutionResourceServiceUpdateOK) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/tanzupackage/tap/solutions/{solution.fullName.name}][%d] solutionResourceServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *SolutionResourceServiceUpdateOK) GetPayload() *models.TanzupackageTapSolutionUpdateSolutionResponse {
	return o.Payload
}

func (o *SolutionResourceServiceUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TanzupackageTapSolutionUpdateSolutionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSolutionResourceServiceUpdateDefault creates a SolutionResourceServiceUpdateDefault with default headers values
func NewSolutionResourceServiceUpdateDefault(code int) *SolutionResourceServiceUpdateDefault {
	return &SolutionResourceServiceUpdateDefault{
		_statusCode: code,
	}
}

/*
SolutionResourceServiceUpdateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SolutionResourceServiceUpdateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the solution resource service update default response
func (o *SolutionResourceServiceUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this solution resource service update default response has a 2xx status code
func (o *SolutionResourceServiceUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this solution resource service update default response has a 3xx status code
func (o *SolutionResourceServiceUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this solution resource service update default response has a 4xx status code
func (o *SolutionResourceServiceUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this solution resource service update default response has a 5xx status code
func (o *SolutionResourceServiceUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this solution resource service update default response a status code equal to that given
func (o *SolutionResourceServiceUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SolutionResourceServiceUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/tanzupackage/tap/solutions/{solution.fullName.name}][%d] SolutionResourceService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *SolutionResourceServiceUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/tanzupackage/tap/solutions/{solution.fullName.name}][%d] SolutionResourceService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *SolutionResourceServiceUpdateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *SolutionResourceServiceUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}