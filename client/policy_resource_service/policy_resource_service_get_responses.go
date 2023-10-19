// Code generated by go-swagger; DO NOT EDIT.

package policy_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// PolicyResourceServiceGetReader is a Reader for the PolicyResourceServiceGet structure.
type PolicyResourceServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PolicyResourceServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPolicyResourceServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPolicyResourceServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPolicyResourceServiceGetOK creates a PolicyResourceServiceGetOK with default headers values
func NewPolicyResourceServiceGetOK() *PolicyResourceServiceGetOK {
	return &PolicyResourceServiceGetOK{}
}

/*
PolicyResourceServiceGetOK describes a response with status code 200, with default header values.

A successful response.
*/
type PolicyResourceServiceGetOK struct {
	Payload *models.WorkspacePolicyGetPolicyResponse
}

// IsSuccess returns true when this policy resource service get o k response has a 2xx status code
func (o *PolicyResourceServiceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this policy resource service get o k response has a 3xx status code
func (o *PolicyResourceServiceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this policy resource service get o k response has a 4xx status code
func (o *PolicyResourceServiceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this policy resource service get o k response has a 5xx status code
func (o *PolicyResourceServiceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this policy resource service get o k response a status code equal to that given
func (o *PolicyResourceServiceGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *PolicyResourceServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/workspaces/{fullName.workspaceName}/policies/{fullName.name}][%d] policyResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *PolicyResourceServiceGetOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/workspaces/{fullName.workspaceName}/policies/{fullName.name}][%d] policyResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *PolicyResourceServiceGetOK) GetPayload() *models.WorkspacePolicyGetPolicyResponse {
	return o.Payload
}

func (o *PolicyResourceServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkspacePolicyGetPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPolicyResourceServiceGetDefault creates a PolicyResourceServiceGetDefault with default headers values
func NewPolicyResourceServiceGetDefault(code int) *PolicyResourceServiceGetDefault {
	return &PolicyResourceServiceGetDefault{
		_statusCode: code,
	}
}

/*
PolicyResourceServiceGetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PolicyResourceServiceGetDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the policy resource service get default response
func (o *PolicyResourceServiceGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this policy resource service get default response has a 2xx status code
func (o *PolicyResourceServiceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this policy resource service get default response has a 3xx status code
func (o *PolicyResourceServiceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this policy resource service get default response has a 4xx status code
func (o *PolicyResourceServiceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this policy resource service get default response has a 5xx status code
func (o *PolicyResourceServiceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this policy resource service get default response a status code equal to that given
func (o *PolicyResourceServiceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PolicyResourceServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/workspaces/{fullName.workspaceName}/policies/{fullName.name}][%d] PolicyResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *PolicyResourceServiceGetDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/workspaces/{fullName.workspaceName}/policies/{fullName.name}][%d] PolicyResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *PolicyResourceServiceGetDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *PolicyResourceServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
