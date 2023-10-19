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

// CredentialResourceServicePatchReader is a Reader for the CredentialResourceServicePatch structure.
type CredentialResourceServicePatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialResourceServicePatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCredentialResourceServicePatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCredentialResourceServicePatchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCredentialResourceServicePatchOK creates a CredentialResourceServicePatchOK with default headers values
func NewCredentialResourceServicePatchOK() *CredentialResourceServicePatchOK {
	return &CredentialResourceServicePatchOK{}
}

/*
CredentialResourceServicePatchOK describes a response with status code 200, with default header values.

A successful response.
*/
type CredentialResourceServicePatchOK struct {
	Payload *models.AccountManagementclusterProvisionerCredentialPatchCredentialResponse
}

// IsSuccess returns true when this credential resource service patch o k response has a 2xx status code
func (o *CredentialResourceServicePatchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this credential resource service patch o k response has a 3xx status code
func (o *CredentialResourceServicePatchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this credential resource service patch o k response has a 4xx status code
func (o *CredentialResourceServicePatchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this credential resource service patch o k response has a 5xx status code
func (o *CredentialResourceServicePatchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this credential resource service patch o k response a status code equal to that given
func (o *CredentialResourceServicePatchOK) IsCode(code int) bool {
	return code == 200
}

func (o *CredentialResourceServicePatchOK) Error() string {
	return fmt.Sprintf("[PATCH /v1alpha1/account/managementcluster/provisioner/credentials/{fullName.name}][%d] credentialResourceServicePatchOK  %+v", 200, o.Payload)
}

func (o *CredentialResourceServicePatchOK) String() string {
	return fmt.Sprintf("[PATCH /v1alpha1/account/managementcluster/provisioner/credentials/{fullName.name}][%d] credentialResourceServicePatchOK  %+v", 200, o.Payload)
}

func (o *CredentialResourceServicePatchOK) GetPayload() *models.AccountManagementclusterProvisionerCredentialPatchCredentialResponse {
	return o.Payload
}

func (o *CredentialResourceServicePatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountManagementclusterProvisionerCredentialPatchCredentialResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCredentialResourceServicePatchDefault creates a CredentialResourceServicePatchDefault with default headers values
func NewCredentialResourceServicePatchDefault(code int) *CredentialResourceServicePatchDefault {
	return &CredentialResourceServicePatchDefault{
		_statusCode: code,
	}
}

/*
CredentialResourceServicePatchDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CredentialResourceServicePatchDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the credential resource service patch default response
func (o *CredentialResourceServicePatchDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this credential resource service patch default response has a 2xx status code
func (o *CredentialResourceServicePatchDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this credential resource service patch default response has a 3xx status code
func (o *CredentialResourceServicePatchDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this credential resource service patch default response has a 4xx status code
func (o *CredentialResourceServicePatchDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this credential resource service patch default response has a 5xx status code
func (o *CredentialResourceServicePatchDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this credential resource service patch default response a status code equal to that given
func (o *CredentialResourceServicePatchDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CredentialResourceServicePatchDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1alpha1/account/managementcluster/provisioner/credentials/{fullName.name}][%d] CredentialResourceService_Patch default  %+v", o._statusCode, o.Payload)
}

func (o *CredentialResourceServicePatchDefault) String() string {
	return fmt.Sprintf("[PATCH /v1alpha1/account/managementcluster/provisioner/credentials/{fullName.name}][%d] CredentialResourceService_Patch default  %+v", o._statusCode, o.Payload)
}

func (o *CredentialResourceServicePatchDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *CredentialResourceServicePatchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
