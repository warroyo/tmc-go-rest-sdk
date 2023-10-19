// Code generated by go-swagger; DO NOT EDIT.

package credential_i_a_m_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// CredentialIAMPolicyGetReader is a Reader for the CredentialIAMPolicyGet structure.
type CredentialIAMPolicyGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialIAMPolicyGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCredentialIAMPolicyGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCredentialIAMPolicyGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCredentialIAMPolicyGetOK creates a CredentialIAMPolicyGetOK with default headers values
func NewCredentialIAMPolicyGetOK() *CredentialIAMPolicyGetOK {
	return &CredentialIAMPolicyGetOK{}
}

/*
CredentialIAMPolicyGetOK describes a response with status code 200, with default header values.

A successful response.
*/
type CredentialIAMPolicyGetOK struct {
	Payload *models.AccountManagementclusterProvisionerCredentialGetCredentialIAMPolicyResponse
}

// IsSuccess returns true when this credential i a m policy get o k response has a 2xx status code
func (o *CredentialIAMPolicyGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this credential i a m policy get o k response has a 3xx status code
func (o *CredentialIAMPolicyGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this credential i a m policy get o k response has a 4xx status code
func (o *CredentialIAMPolicyGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this credential i a m policy get o k response has a 5xx status code
func (o *CredentialIAMPolicyGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this credential i a m policy get o k response a status code equal to that given
func (o *CredentialIAMPolicyGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *CredentialIAMPolicyGetOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/account/managementcluster/provisioner/credentials:iam/{fullName.name}][%d] credentialIAMPolicyGetOK  %+v", 200, o.Payload)
}

func (o *CredentialIAMPolicyGetOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/account/managementcluster/provisioner/credentials:iam/{fullName.name}][%d] credentialIAMPolicyGetOK  %+v", 200, o.Payload)
}

func (o *CredentialIAMPolicyGetOK) GetPayload() *models.AccountManagementclusterProvisionerCredentialGetCredentialIAMPolicyResponse {
	return o.Payload
}

func (o *CredentialIAMPolicyGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountManagementclusterProvisionerCredentialGetCredentialIAMPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCredentialIAMPolicyGetDefault creates a CredentialIAMPolicyGetDefault with default headers values
func NewCredentialIAMPolicyGetDefault(code int) *CredentialIAMPolicyGetDefault {
	return &CredentialIAMPolicyGetDefault{
		_statusCode: code,
	}
}

/*
CredentialIAMPolicyGetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CredentialIAMPolicyGetDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the credential i a m policy get default response
func (o *CredentialIAMPolicyGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this credential i a m policy get default response has a 2xx status code
func (o *CredentialIAMPolicyGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this credential i a m policy get default response has a 3xx status code
func (o *CredentialIAMPolicyGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this credential i a m policy get default response has a 4xx status code
func (o *CredentialIAMPolicyGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this credential i a m policy get default response has a 5xx status code
func (o *CredentialIAMPolicyGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this credential i a m policy get default response a status code equal to that given
func (o *CredentialIAMPolicyGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CredentialIAMPolicyGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/account/managementcluster/provisioner/credentials:iam/{fullName.name}][%d] CredentialIAMPolicy_Get default  %+v", o._statusCode, o.Payload)
}

func (o *CredentialIAMPolicyGetDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/account/managementcluster/provisioner/credentials:iam/{fullName.name}][%d] CredentialIAMPolicy_Get default  %+v", o._statusCode, o.Payload)
}

func (o *CredentialIAMPolicyGetDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *CredentialIAMPolicyGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
