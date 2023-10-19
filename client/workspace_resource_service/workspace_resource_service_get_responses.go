// Code generated by go-swagger; DO NOT EDIT.

package workspace_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// WorkspaceResourceServiceGetReader is a Reader for the WorkspaceResourceServiceGet structure.
type WorkspaceResourceServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkspaceResourceServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkspaceResourceServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWorkspaceResourceServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWorkspaceResourceServiceGetOK creates a WorkspaceResourceServiceGetOK with default headers values
func NewWorkspaceResourceServiceGetOK() *WorkspaceResourceServiceGetOK {
	return &WorkspaceResourceServiceGetOK{}
}

/*
WorkspaceResourceServiceGetOK describes a response with status code 200, with default header values.

A successful response.
*/
type WorkspaceResourceServiceGetOK struct {
	Payload *models.WorkspaceGetWorkspaceResponse
}

// IsSuccess returns true when this workspace resource service get o k response has a 2xx status code
func (o *WorkspaceResourceServiceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workspace resource service get o k response has a 3xx status code
func (o *WorkspaceResourceServiceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workspace resource service get o k response has a 4xx status code
func (o *WorkspaceResourceServiceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this workspace resource service get o k response has a 5xx status code
func (o *WorkspaceResourceServiceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this workspace resource service get o k response a status code equal to that given
func (o *WorkspaceResourceServiceGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *WorkspaceResourceServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/workspaces/{fullName.name}][%d] workspaceResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *WorkspaceResourceServiceGetOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/workspaces/{fullName.name}][%d] workspaceResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *WorkspaceResourceServiceGetOK) GetPayload() *models.WorkspaceGetWorkspaceResponse {
	return o.Payload
}

func (o *WorkspaceResourceServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkspaceGetWorkspaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWorkspaceResourceServiceGetDefault creates a WorkspaceResourceServiceGetDefault with default headers values
func NewWorkspaceResourceServiceGetDefault(code int) *WorkspaceResourceServiceGetDefault {
	return &WorkspaceResourceServiceGetDefault{
		_statusCode: code,
	}
}

/*
WorkspaceResourceServiceGetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WorkspaceResourceServiceGetDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the workspace resource service get default response
func (o *WorkspaceResourceServiceGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this workspace resource service get default response has a 2xx status code
func (o *WorkspaceResourceServiceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this workspace resource service get default response has a 3xx status code
func (o *WorkspaceResourceServiceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this workspace resource service get default response has a 4xx status code
func (o *WorkspaceResourceServiceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this workspace resource service get default response has a 5xx status code
func (o *WorkspaceResourceServiceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this workspace resource service get default response a status code equal to that given
func (o *WorkspaceResourceServiceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WorkspaceResourceServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/workspaces/{fullName.name}][%d] WorkspaceResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *WorkspaceResourceServiceGetDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/workspaces/{fullName.name}][%d] WorkspaceResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *WorkspaceResourceServiceGetDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WorkspaceResourceServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}