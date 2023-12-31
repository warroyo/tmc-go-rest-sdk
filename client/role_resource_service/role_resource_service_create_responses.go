// Code generated by go-swagger; DO NOT EDIT.

package role_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// RoleResourceServiceCreateReader is a Reader for the RoleResourceServiceCreate structure.
type RoleResourceServiceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RoleResourceServiceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRoleResourceServiceCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRoleResourceServiceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRoleResourceServiceCreateOK creates a RoleResourceServiceCreateOK with default headers values
func NewRoleResourceServiceCreateOK() *RoleResourceServiceCreateOK {
	return &RoleResourceServiceCreateOK{}
}

/*
RoleResourceServiceCreateOK describes a response with status code 200, with default header values.

A successful response.
*/
type RoleResourceServiceCreateOK struct {
	Payload *models.IamRoleCreateRoleResponse
}

// IsSuccess returns true when this role resource service create o k response has a 2xx status code
func (o *RoleResourceServiceCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this role resource service create o k response has a 3xx status code
func (o *RoleResourceServiceCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this role resource service create o k response has a 4xx status code
func (o *RoleResourceServiceCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this role resource service create o k response has a 5xx status code
func (o *RoleResourceServiceCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this role resource service create o k response a status code equal to that given
func (o *RoleResourceServiceCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *RoleResourceServiceCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/iam/roles][%d] roleResourceServiceCreateOK  %+v", 200, o.Payload)
}

func (o *RoleResourceServiceCreateOK) String() string {
	return fmt.Sprintf("[POST /v1alpha1/iam/roles][%d] roleResourceServiceCreateOK  %+v", 200, o.Payload)
}

func (o *RoleResourceServiceCreateOK) GetPayload() *models.IamRoleCreateRoleResponse {
	return o.Payload
}

func (o *RoleResourceServiceCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IamRoleCreateRoleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRoleResourceServiceCreateDefault creates a RoleResourceServiceCreateDefault with default headers values
func NewRoleResourceServiceCreateDefault(code int) *RoleResourceServiceCreateDefault {
	return &RoleResourceServiceCreateDefault{
		_statusCode: code,
	}
}

/*
RoleResourceServiceCreateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RoleResourceServiceCreateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the role resource service create default response
func (o *RoleResourceServiceCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this role resource service create default response has a 2xx status code
func (o *RoleResourceServiceCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this role resource service create default response has a 3xx status code
func (o *RoleResourceServiceCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this role resource service create default response has a 4xx status code
func (o *RoleResourceServiceCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this role resource service create default response has a 5xx status code
func (o *RoleResourceServiceCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this role resource service create default response a status code equal to that given
func (o *RoleResourceServiceCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *RoleResourceServiceCreateDefault) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/iam/roles][%d] RoleResourceService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *RoleResourceServiceCreateDefault) String() string {
	return fmt.Sprintf("[POST /v1alpha1/iam/roles][%d] RoleResourceService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *RoleResourceServiceCreateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *RoleResourceServiceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
