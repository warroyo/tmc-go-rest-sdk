// Code generated by go-swagger; DO NOT EDIT.

package effective_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// EffectiveResourceServiceListReader is a Reader for the EffectiveResourceServiceList structure.
type EffectiveResourceServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EffectiveResourceServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEffectiveResourceServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEffectiveResourceServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEffectiveResourceServiceListOK creates a EffectiveResourceServiceListOK with default headers values
func NewEffectiveResourceServiceListOK() *EffectiveResourceServiceListOK {
	return &EffectiveResourceServiceListOK{}
}

/*
EffectiveResourceServiceListOK describes a response with status code 200, with default header values.

A successful response.
*/
type EffectiveResourceServiceListOK struct {
	Payload *models.SettingEffectiveListEffectiveResponse
}

// IsSuccess returns true when this effective resource service list o k response has a 2xx status code
func (o *EffectiveResourceServiceListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this effective resource service list o k response has a 3xx status code
func (o *EffectiveResourceServiceListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this effective resource service list o k response has a 4xx status code
func (o *EffectiveResourceServiceListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this effective resource service list o k response has a 5xx status code
func (o *EffectiveResourceServiceListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this effective resource service list o k response a status code equal to that given
func (o *EffectiveResourceServiceListOK) IsCode(code int) bool {
	return code == 200
}

func (o *EffectiveResourceServiceListOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/setting/effective][%d] effectiveResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *EffectiveResourceServiceListOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/setting/effective][%d] effectiveResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *EffectiveResourceServiceListOK) GetPayload() *models.SettingEffectiveListEffectiveResponse {
	return o.Payload
}

func (o *EffectiveResourceServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SettingEffectiveListEffectiveResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEffectiveResourceServiceListDefault creates a EffectiveResourceServiceListDefault with default headers values
func NewEffectiveResourceServiceListDefault(code int) *EffectiveResourceServiceListDefault {
	return &EffectiveResourceServiceListDefault{
		_statusCode: code,
	}
}

/*
EffectiveResourceServiceListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EffectiveResourceServiceListDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the effective resource service list default response
func (o *EffectiveResourceServiceListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this effective resource service list default response has a 2xx status code
func (o *EffectiveResourceServiceListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this effective resource service list default response has a 3xx status code
func (o *EffectiveResourceServiceListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this effective resource service list default response has a 4xx status code
func (o *EffectiveResourceServiceListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this effective resource service list default response has a 5xx status code
func (o *EffectiveResourceServiceListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this effective resource service list default response a status code equal to that given
func (o *EffectiveResourceServiceListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EffectiveResourceServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/setting/effective][%d] EffectiveResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *EffectiveResourceServiceListDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/setting/effective][%d] EffectiveResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *EffectiveResourceServiceListDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *EffectiveResourceServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
