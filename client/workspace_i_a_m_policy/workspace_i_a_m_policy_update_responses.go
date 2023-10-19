// Code generated by go-swagger; DO NOT EDIT.

package workspace_i_a_m_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// WorkspaceIAMPolicyUpdateReader is a Reader for the WorkspaceIAMPolicyUpdate structure.
type WorkspaceIAMPolicyUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkspaceIAMPolicyUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkspaceIAMPolicyUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWorkspaceIAMPolicyUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWorkspaceIAMPolicyUpdateOK creates a WorkspaceIAMPolicyUpdateOK with default headers values
func NewWorkspaceIAMPolicyUpdateOK() *WorkspaceIAMPolicyUpdateOK {
	return &WorkspaceIAMPolicyUpdateOK{}
}

/*
WorkspaceIAMPolicyUpdateOK describes a response with status code 200, with default header values.

A successful response.
*/
type WorkspaceIAMPolicyUpdateOK struct {
	Payload *models.WorkspaceUpdateWorkspaceIAMPolicyResponse
}

// IsSuccess returns true when this workspace i a m policy update o k response has a 2xx status code
func (o *WorkspaceIAMPolicyUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workspace i a m policy update o k response has a 3xx status code
func (o *WorkspaceIAMPolicyUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workspace i a m policy update o k response has a 4xx status code
func (o *WorkspaceIAMPolicyUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this workspace i a m policy update o k response has a 5xx status code
func (o *WorkspaceIAMPolicyUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this workspace i a m policy update o k response a status code equal to that given
func (o *WorkspaceIAMPolicyUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *WorkspaceIAMPolicyUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/workspaces:iam/{fullName.name}][%d] workspaceIAMPolicyUpdateOK  %+v", 200, o.Payload)
}

func (o *WorkspaceIAMPolicyUpdateOK) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/workspaces:iam/{fullName.name}][%d] workspaceIAMPolicyUpdateOK  %+v", 200, o.Payload)
}

func (o *WorkspaceIAMPolicyUpdateOK) GetPayload() *models.WorkspaceUpdateWorkspaceIAMPolicyResponse {
	return o.Payload
}

func (o *WorkspaceIAMPolicyUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkspaceUpdateWorkspaceIAMPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWorkspaceIAMPolicyUpdateDefault creates a WorkspaceIAMPolicyUpdateDefault with default headers values
func NewWorkspaceIAMPolicyUpdateDefault(code int) *WorkspaceIAMPolicyUpdateDefault {
	return &WorkspaceIAMPolicyUpdateDefault{
		_statusCode: code,
	}
}

/*
WorkspaceIAMPolicyUpdateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WorkspaceIAMPolicyUpdateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the workspace i a m policy update default response
func (o *WorkspaceIAMPolicyUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this workspace i a m policy update default response has a 2xx status code
func (o *WorkspaceIAMPolicyUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this workspace i a m policy update default response has a 3xx status code
func (o *WorkspaceIAMPolicyUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this workspace i a m policy update default response has a 4xx status code
func (o *WorkspaceIAMPolicyUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this workspace i a m policy update default response has a 5xx status code
func (o *WorkspaceIAMPolicyUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this workspace i a m policy update default response a status code equal to that given
func (o *WorkspaceIAMPolicyUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WorkspaceIAMPolicyUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/workspaces:iam/{fullName.name}][%d] WorkspaceIAMPolicy_Update default  %+v", o._statusCode, o.Payload)
}

func (o *WorkspaceIAMPolicyUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/workspaces:iam/{fullName.name}][%d] WorkspaceIAMPolicy_Update default  %+v", o._statusCode, o.Payload)
}

func (o *WorkspaceIAMPolicyUpdateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WorkspaceIAMPolicyUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}