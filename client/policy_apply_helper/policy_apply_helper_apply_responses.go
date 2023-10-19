// Code generated by go-swagger; DO NOT EDIT.

package policy_apply_helper

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// PolicyApplyHelperApplyReader is a Reader for the PolicyApplyHelperApply structure.
type PolicyApplyHelperApplyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PolicyApplyHelperApplyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPolicyApplyHelperApplyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPolicyApplyHelperApplyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPolicyApplyHelperApplyOK creates a PolicyApplyHelperApplyOK with default headers values
func NewPolicyApplyHelperApplyOK() *PolicyApplyHelperApplyOK {
	return &PolicyApplyHelperApplyOK{}
}

/*
PolicyApplyHelperApplyOK describes a response with status code 200, with default header values.

A successful response.
*/
type PolicyApplyHelperApplyOK struct {
	Payload *models.WorkspacePolicyApplyPolicyResponse
}

// IsSuccess returns true when this policy apply helper apply o k response has a 2xx status code
func (o *PolicyApplyHelperApplyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this policy apply helper apply o k response has a 3xx status code
func (o *PolicyApplyHelperApplyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this policy apply helper apply o k response has a 4xx status code
func (o *PolicyApplyHelperApplyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this policy apply helper apply o k response has a 5xx status code
func (o *PolicyApplyHelperApplyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this policy apply helper apply o k response a status code equal to that given
func (o *PolicyApplyHelperApplyOK) IsCode(code int) bool {
	return code == 200
}

func (o *PolicyApplyHelperApplyOK) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/workspaces/{policy.fullName.workspaceName}/policies:apply][%d] policyApplyHelperApplyOK  %+v", 200, o.Payload)
}

func (o *PolicyApplyHelperApplyOK) String() string {
	return fmt.Sprintf("[POST /v1alpha1/workspaces/{policy.fullName.workspaceName}/policies:apply][%d] policyApplyHelperApplyOK  %+v", 200, o.Payload)
}

func (o *PolicyApplyHelperApplyOK) GetPayload() *models.WorkspacePolicyApplyPolicyResponse {
	return o.Payload
}

func (o *PolicyApplyHelperApplyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkspacePolicyApplyPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPolicyApplyHelperApplyDefault creates a PolicyApplyHelperApplyDefault with default headers values
func NewPolicyApplyHelperApplyDefault(code int) *PolicyApplyHelperApplyDefault {
	return &PolicyApplyHelperApplyDefault{
		_statusCode: code,
	}
}

/*
PolicyApplyHelperApplyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PolicyApplyHelperApplyDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the policy apply helper apply default response
func (o *PolicyApplyHelperApplyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this policy apply helper apply default response has a 2xx status code
func (o *PolicyApplyHelperApplyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this policy apply helper apply default response has a 3xx status code
func (o *PolicyApplyHelperApplyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this policy apply helper apply default response has a 4xx status code
func (o *PolicyApplyHelperApplyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this policy apply helper apply default response has a 5xx status code
func (o *PolicyApplyHelperApplyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this policy apply helper apply default response a status code equal to that given
func (o *PolicyApplyHelperApplyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PolicyApplyHelperApplyDefault) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/workspaces/{policy.fullName.workspaceName}/policies:apply][%d] PolicyApplyHelper_Apply default  %+v", o._statusCode, o.Payload)
}

func (o *PolicyApplyHelperApplyDefault) String() string {
	return fmt.Sprintf("[POST /v1alpha1/workspaces/{policy.fullName.workspaceName}/policies:apply][%d] PolicyApplyHelper_Apply default  %+v", o._statusCode, o.Payload)
}

func (o *PolicyApplyHelperApplyDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *PolicyApplyHelperApplyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}