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

// SolutionResourceServiceGetReader is a Reader for the SolutionResourceServiceGet structure.
type SolutionResourceServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SolutionResourceServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSolutionResourceServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSolutionResourceServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSolutionResourceServiceGetOK creates a SolutionResourceServiceGetOK with default headers values
func NewSolutionResourceServiceGetOK() *SolutionResourceServiceGetOK {
	return &SolutionResourceServiceGetOK{}
}

/*
SolutionResourceServiceGetOK describes a response with status code 200, with default header values.

A successful response.
*/
type SolutionResourceServiceGetOK struct {
	Payload *models.TanzupackageTapSolutionGetSolutionResponse
}

// IsSuccess returns true when this solution resource service get o k response has a 2xx status code
func (o *SolutionResourceServiceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this solution resource service get o k response has a 3xx status code
func (o *SolutionResourceServiceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this solution resource service get o k response has a 4xx status code
func (o *SolutionResourceServiceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this solution resource service get o k response has a 5xx status code
func (o *SolutionResourceServiceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this solution resource service get o k response a status code equal to that given
func (o *SolutionResourceServiceGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *SolutionResourceServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/tanzupackage/tap/solutions/{fullName.name}][%d] solutionResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *SolutionResourceServiceGetOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/tanzupackage/tap/solutions/{fullName.name}][%d] solutionResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *SolutionResourceServiceGetOK) GetPayload() *models.TanzupackageTapSolutionGetSolutionResponse {
	return o.Payload
}

func (o *SolutionResourceServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TanzupackageTapSolutionGetSolutionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSolutionResourceServiceGetDefault creates a SolutionResourceServiceGetDefault with default headers values
func NewSolutionResourceServiceGetDefault(code int) *SolutionResourceServiceGetDefault {
	return &SolutionResourceServiceGetDefault{
		_statusCode: code,
	}
}

/*
SolutionResourceServiceGetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SolutionResourceServiceGetDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the solution resource service get default response
func (o *SolutionResourceServiceGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this solution resource service get default response has a 2xx status code
func (o *SolutionResourceServiceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this solution resource service get default response has a 3xx status code
func (o *SolutionResourceServiceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this solution resource service get default response has a 4xx status code
func (o *SolutionResourceServiceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this solution resource service get default response has a 5xx status code
func (o *SolutionResourceServiceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this solution resource service get default response a status code equal to that given
func (o *SolutionResourceServiceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SolutionResourceServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/tanzupackage/tap/solutions/{fullName.name}][%d] SolutionResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *SolutionResourceServiceGetDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/tanzupackage/tap/solutions/{fullName.name}][%d] SolutionResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *SolutionResourceServiceGetDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *SolutionResourceServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
