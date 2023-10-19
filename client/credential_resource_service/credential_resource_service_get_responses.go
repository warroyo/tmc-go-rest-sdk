// Code generated by go-swagger; DO NOT EDIT.

package credential_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// CredentialResourceServiceGetReader is a Reader for the CredentialResourceServiceGet structure.
type CredentialResourceServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialResourceServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCredentialResourceServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCredentialResourceServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCredentialResourceServiceGetOK creates a CredentialResourceServiceGetOK with default headers values
func NewCredentialResourceServiceGetOK() *CredentialResourceServiceGetOK {
	return &CredentialResourceServiceGetOK{}
}

/*
CredentialResourceServiceGetOK describes a response with status code 200, with default header values.

A successful response.
*/
type CredentialResourceServiceGetOK struct {
	Payload *models.AccountManagementclusterProvisionerCredentialGetCredentialResponse
}

// IsSuccess returns true when this credential resource service get o k response has a 2xx status code
func (o *CredentialResourceServiceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this credential resource service get o k response has a 3xx status code
func (o *CredentialResourceServiceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this credential resource service get o k response has a 4xx status code
func (o *CredentialResourceServiceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this credential resource service get o k response has a 5xx status code
func (o *CredentialResourceServiceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this credential resource service get o k response a status code equal to that given
func (o *CredentialResourceServiceGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *CredentialResourceServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/account/managementcluster/provisioner/credentials/{fullName.name}][%d] credentialResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *CredentialResourceServiceGetOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/account/managementcluster/provisioner/credentials/{fullName.name}][%d] credentialResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *CredentialResourceServiceGetOK) GetPayload() *models.AccountManagementclusterProvisionerCredentialGetCredentialResponse {
	return o.Payload
}

func (o *CredentialResourceServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountManagementclusterProvisionerCredentialGetCredentialResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCredentialResourceServiceGetDefault creates a CredentialResourceServiceGetDefault with default headers values
func NewCredentialResourceServiceGetDefault(code int) *CredentialResourceServiceGetDefault {
	return &CredentialResourceServiceGetDefault{
		_statusCode: code,
	}
}

/*
CredentialResourceServiceGetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CredentialResourceServiceGetDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the credential resource service get default response
func (o *CredentialResourceServiceGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this credential resource service get default response has a 2xx status code
func (o *CredentialResourceServiceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this credential resource service get default response has a 3xx status code
func (o *CredentialResourceServiceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this credential resource service get default response has a 4xx status code
func (o *CredentialResourceServiceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this credential resource service get default response has a 5xx status code
func (o *CredentialResourceServiceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this credential resource service get default response a status code equal to that given
func (o *CredentialResourceServiceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CredentialResourceServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/account/managementcluster/provisioner/credentials/{fullName.name}][%d] CredentialResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *CredentialResourceServiceGetDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/account/managementcluster/provisioner/credentials/{fullName.name}][%d] CredentialResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *CredentialResourceServiceGetDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *CredentialResourceServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
