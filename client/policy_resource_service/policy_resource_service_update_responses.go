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

// PolicyResourceServiceUpdateReader is a Reader for the PolicyResourceServiceUpdate structure.
type PolicyResourceServiceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PolicyResourceServiceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPolicyResourceServiceUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPolicyResourceServiceUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPolicyResourceServiceUpdateOK creates a PolicyResourceServiceUpdateOK with default headers values
func NewPolicyResourceServiceUpdateOK() *PolicyResourceServiceUpdateOK {
	return &PolicyResourceServiceUpdateOK{}
}

/*
PolicyResourceServiceUpdateOK describes a response with status code 200, with default header values.

A successful response.
*/
type PolicyResourceServiceUpdateOK struct {
	Payload *models.WorkspacePolicyUpdatePolicyResponse
}

// IsSuccess returns true when this policy resource service update o k response has a 2xx status code
func (o *PolicyResourceServiceUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this policy resource service update o k response has a 3xx status code
func (o *PolicyResourceServiceUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this policy resource service update o k response has a 4xx status code
func (o *PolicyResourceServiceUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this policy resource service update o k response has a 5xx status code
func (o *PolicyResourceServiceUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this policy resource service update o k response a status code equal to that given
func (o *PolicyResourceServiceUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *PolicyResourceServiceUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/workspaces/{policy.fullName.workspaceName}/policies/{policy.fullName.name}][%d] policyResourceServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *PolicyResourceServiceUpdateOK) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/workspaces/{policy.fullName.workspaceName}/policies/{policy.fullName.name}][%d] policyResourceServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *PolicyResourceServiceUpdateOK) GetPayload() *models.WorkspacePolicyUpdatePolicyResponse {
	return o.Payload
}

func (o *PolicyResourceServiceUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkspacePolicyUpdatePolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPolicyResourceServiceUpdateDefault creates a PolicyResourceServiceUpdateDefault with default headers values
func NewPolicyResourceServiceUpdateDefault(code int) *PolicyResourceServiceUpdateDefault {
	return &PolicyResourceServiceUpdateDefault{
		_statusCode: code,
	}
}

/*
PolicyResourceServiceUpdateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PolicyResourceServiceUpdateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the policy resource service update default response
func (o *PolicyResourceServiceUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this policy resource service update default response has a 2xx status code
func (o *PolicyResourceServiceUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this policy resource service update default response has a 3xx status code
func (o *PolicyResourceServiceUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this policy resource service update default response has a 4xx status code
func (o *PolicyResourceServiceUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this policy resource service update default response has a 5xx status code
func (o *PolicyResourceServiceUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this policy resource service update default response a status code equal to that given
func (o *PolicyResourceServiceUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PolicyResourceServiceUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/workspaces/{policy.fullName.workspaceName}/policies/{policy.fullName.name}][%d] PolicyResourceService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *PolicyResourceServiceUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/workspaces/{policy.fullName.workspaceName}/policies/{policy.fullName.name}][%d] PolicyResourceService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *PolicyResourceServiceUpdateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *PolicyResourceServiceUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
