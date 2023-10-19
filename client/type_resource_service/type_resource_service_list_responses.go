// Code generated by go-swagger; DO NOT EDIT.

package type_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// TypeResourceServiceListReader is a Reader for the TypeResourceServiceList structure.
type TypeResourceServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TypeResourceServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTypeResourceServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTypeResourceServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTypeResourceServiceListOK creates a TypeResourceServiceListOK with default headers values
func NewTypeResourceServiceListOK() *TypeResourceServiceListOK {
	return &TypeResourceServiceListOK{}
}

/*
TypeResourceServiceListOK describes a response with status code 200, with default header values.

A successful response.
*/
type TypeResourceServiceListOK struct {
	Payload *models.SettingTypeListTypesResponse
}

// IsSuccess returns true when this type resource service list o k response has a 2xx status code
func (o *TypeResourceServiceListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this type resource service list o k response has a 3xx status code
func (o *TypeResourceServiceListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this type resource service list o k response has a 4xx status code
func (o *TypeResourceServiceListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this type resource service list o k response has a 5xx status code
func (o *TypeResourceServiceListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this type resource service list o k response a status code equal to that given
func (o *TypeResourceServiceListOK) IsCode(code int) bool {
	return code == 200
}

func (o *TypeResourceServiceListOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/setting/types][%d] typeResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *TypeResourceServiceListOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/setting/types][%d] typeResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *TypeResourceServiceListOK) GetPayload() *models.SettingTypeListTypesResponse {
	return o.Payload
}

func (o *TypeResourceServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SettingTypeListTypesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTypeResourceServiceListDefault creates a TypeResourceServiceListDefault with default headers values
func NewTypeResourceServiceListDefault(code int) *TypeResourceServiceListDefault {
	return &TypeResourceServiceListDefault{
		_statusCode: code,
	}
}

/*
TypeResourceServiceListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type TypeResourceServiceListDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the type resource service list default response
func (o *TypeResourceServiceListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this type resource service list default response has a 2xx status code
func (o *TypeResourceServiceListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this type resource service list default response has a 3xx status code
func (o *TypeResourceServiceListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this type resource service list default response has a 4xx status code
func (o *TypeResourceServiceListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this type resource service list default response has a 5xx status code
func (o *TypeResourceServiceListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this type resource service list default response a status code equal to that given
func (o *TypeResourceServiceListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *TypeResourceServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/setting/types][%d] TypeResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *TypeResourceServiceListDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/setting/types][%d] TypeResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *TypeResourceServiceListDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *TypeResourceServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}