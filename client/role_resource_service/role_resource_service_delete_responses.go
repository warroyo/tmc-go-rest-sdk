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

// RoleResourceServiceDeleteReader is a Reader for the RoleResourceServiceDelete structure.
type RoleResourceServiceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RoleResourceServiceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRoleResourceServiceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRoleResourceServiceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRoleResourceServiceDeleteOK creates a RoleResourceServiceDeleteOK with default headers values
func NewRoleResourceServiceDeleteOK() *RoleResourceServiceDeleteOK {
	return &RoleResourceServiceDeleteOK{}
}

/*
RoleResourceServiceDeleteOK describes a response with status code 200, with default header values.

A successful response.
*/
type RoleResourceServiceDeleteOK struct {
	Payload *models.IamRoleDeleteRoleResponse
}

// IsSuccess returns true when this role resource service delete o k response has a 2xx status code
func (o *RoleResourceServiceDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this role resource service delete o k response has a 3xx status code
func (o *RoleResourceServiceDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this role resource service delete o k response has a 4xx status code
func (o *RoleResourceServiceDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this role resource service delete o k response has a 5xx status code
func (o *RoleResourceServiceDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this role resource service delete o k response a status code equal to that given
func (o *RoleResourceServiceDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *RoleResourceServiceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1alpha1/iam/roles/{fullName.name}][%d] roleResourceServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *RoleResourceServiceDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /v1alpha1/iam/roles/{fullName.name}][%d] roleResourceServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *RoleResourceServiceDeleteOK) GetPayload() *models.IamRoleDeleteRoleResponse {
	return o.Payload
}

func (o *RoleResourceServiceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IamRoleDeleteRoleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRoleResourceServiceDeleteDefault creates a RoleResourceServiceDeleteDefault with default headers values
func NewRoleResourceServiceDeleteDefault(code int) *RoleResourceServiceDeleteDefault {
	return &RoleResourceServiceDeleteDefault{
		_statusCode: code,
	}
}

/*
RoleResourceServiceDeleteDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RoleResourceServiceDeleteDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the role resource service delete default response
func (o *RoleResourceServiceDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this role resource service delete default response has a 2xx status code
func (o *RoleResourceServiceDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this role resource service delete default response has a 3xx status code
func (o *RoleResourceServiceDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this role resource service delete default response has a 4xx status code
func (o *RoleResourceServiceDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this role resource service delete default response has a 5xx status code
func (o *RoleResourceServiceDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this role resource service delete default response a status code equal to that given
func (o *RoleResourceServiceDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *RoleResourceServiceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1alpha1/iam/roles/{fullName.name}][%d] RoleResourceService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *RoleResourceServiceDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /v1alpha1/iam/roles/{fullName.name}][%d] RoleResourceService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *RoleResourceServiceDeleteDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *RoleResourceServiceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
