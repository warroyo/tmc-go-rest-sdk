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

// RoleResourceServicePatchReader is a Reader for the RoleResourceServicePatch structure.
type RoleResourceServicePatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RoleResourceServicePatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRoleResourceServicePatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRoleResourceServicePatchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRoleResourceServicePatchOK creates a RoleResourceServicePatchOK with default headers values
func NewRoleResourceServicePatchOK() *RoleResourceServicePatchOK {
	return &RoleResourceServicePatchOK{}
}

/*
RoleResourceServicePatchOK describes a response with status code 200, with default header values.

A successful response.
*/
type RoleResourceServicePatchOK struct {
	Payload *models.IamRolePatchRoleResponse
}

// IsSuccess returns true when this role resource service patch o k response has a 2xx status code
func (o *RoleResourceServicePatchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this role resource service patch o k response has a 3xx status code
func (o *RoleResourceServicePatchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this role resource service patch o k response has a 4xx status code
func (o *RoleResourceServicePatchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this role resource service patch o k response has a 5xx status code
func (o *RoleResourceServicePatchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this role resource service patch o k response a status code equal to that given
func (o *RoleResourceServicePatchOK) IsCode(code int) bool {
	return code == 200
}

func (o *RoleResourceServicePatchOK) Error() string {
	return fmt.Sprintf("[PATCH /v1alpha1/iam/roles/{fullName.name}][%d] roleResourceServicePatchOK  %+v", 200, o.Payload)
}

func (o *RoleResourceServicePatchOK) String() string {
	return fmt.Sprintf("[PATCH /v1alpha1/iam/roles/{fullName.name}][%d] roleResourceServicePatchOK  %+v", 200, o.Payload)
}

func (o *RoleResourceServicePatchOK) GetPayload() *models.IamRolePatchRoleResponse {
	return o.Payload
}

func (o *RoleResourceServicePatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IamRolePatchRoleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRoleResourceServicePatchDefault creates a RoleResourceServicePatchDefault with default headers values
func NewRoleResourceServicePatchDefault(code int) *RoleResourceServicePatchDefault {
	return &RoleResourceServicePatchDefault{
		_statusCode: code,
	}
}

/*
RoleResourceServicePatchDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RoleResourceServicePatchDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the role resource service patch default response
func (o *RoleResourceServicePatchDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this role resource service patch default response has a 2xx status code
func (o *RoleResourceServicePatchDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this role resource service patch default response has a 3xx status code
func (o *RoleResourceServicePatchDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this role resource service patch default response has a 4xx status code
func (o *RoleResourceServicePatchDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this role resource service patch default response has a 5xx status code
func (o *RoleResourceServicePatchDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this role resource service patch default response a status code equal to that given
func (o *RoleResourceServicePatchDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *RoleResourceServicePatchDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1alpha1/iam/roles/{fullName.name}][%d] RoleResourceService_Patch default  %+v", o._statusCode, o.Payload)
}

func (o *RoleResourceServicePatchDefault) String() string {
	return fmt.Sprintf("[PATCH /v1alpha1/iam/roles/{fullName.name}][%d] RoleResourceService_Patch default  %+v", o._statusCode, o.Payload)
}

func (o *RoleResourceServicePatchDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *RoleResourceServicePatchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
