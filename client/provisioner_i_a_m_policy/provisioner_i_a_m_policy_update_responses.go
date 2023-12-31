// Code generated by go-swagger; DO NOT EDIT.

package provisioner_i_a_m_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// ProvisionerIAMPolicyUpdateReader is a Reader for the ProvisionerIAMPolicyUpdate structure.
type ProvisionerIAMPolicyUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProvisionerIAMPolicyUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProvisionerIAMPolicyUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProvisionerIAMPolicyUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProvisionerIAMPolicyUpdateOK creates a ProvisionerIAMPolicyUpdateOK with default headers values
func NewProvisionerIAMPolicyUpdateOK() *ProvisionerIAMPolicyUpdateOK {
	return &ProvisionerIAMPolicyUpdateOK{}
}

/*
ProvisionerIAMPolicyUpdateOK describes a response with status code 200, with default header values.

A successful response.
*/
type ProvisionerIAMPolicyUpdateOK struct {
	Payload *models.ManagementclusterProvisionerUpdateProvisionerIAMPolicyResponse
}

// IsSuccess returns true when this provisioner i a m policy update o k response has a 2xx status code
func (o *ProvisionerIAMPolicyUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this provisioner i a m policy update o k response has a 3xx status code
func (o *ProvisionerIAMPolicyUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this provisioner i a m policy update o k response has a 4xx status code
func (o *ProvisionerIAMPolicyUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this provisioner i a m policy update o k response has a 5xx status code
func (o *ProvisionerIAMPolicyUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this provisioner i a m policy update o k response a status code equal to that given
func (o *ProvisionerIAMPolicyUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProvisionerIAMPolicyUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/managementclusters/{fullName.managementClusterName}/provisioners:iam/{fullName.name}][%d] provisionerIAMPolicyUpdateOK  %+v", 200, o.Payload)
}

func (o *ProvisionerIAMPolicyUpdateOK) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/managementclusters/{fullName.managementClusterName}/provisioners:iam/{fullName.name}][%d] provisionerIAMPolicyUpdateOK  %+v", 200, o.Payload)
}

func (o *ProvisionerIAMPolicyUpdateOK) GetPayload() *models.ManagementclusterProvisionerUpdateProvisionerIAMPolicyResponse {
	return o.Payload
}

func (o *ProvisionerIAMPolicyUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagementclusterProvisionerUpdateProvisionerIAMPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProvisionerIAMPolicyUpdateDefault creates a ProvisionerIAMPolicyUpdateDefault with default headers values
func NewProvisionerIAMPolicyUpdateDefault(code int) *ProvisionerIAMPolicyUpdateDefault {
	return &ProvisionerIAMPolicyUpdateDefault{
		_statusCode: code,
	}
}

/*
ProvisionerIAMPolicyUpdateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ProvisionerIAMPolicyUpdateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the provisioner i a m policy update default response
func (o *ProvisionerIAMPolicyUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this provisioner i a m policy update default response has a 2xx status code
func (o *ProvisionerIAMPolicyUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this provisioner i a m policy update default response has a 3xx status code
func (o *ProvisionerIAMPolicyUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this provisioner i a m policy update default response has a 4xx status code
func (o *ProvisionerIAMPolicyUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this provisioner i a m policy update default response has a 5xx status code
func (o *ProvisionerIAMPolicyUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this provisioner i a m policy update default response a status code equal to that given
func (o *ProvisionerIAMPolicyUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ProvisionerIAMPolicyUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/managementclusters/{fullName.managementClusterName}/provisioners:iam/{fullName.name}][%d] ProvisionerIAMPolicy_Update default  %+v", o._statusCode, o.Payload)
}

func (o *ProvisionerIAMPolicyUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/managementclusters/{fullName.managementClusterName}/provisioners:iam/{fullName.name}][%d] ProvisionerIAMPolicy_Update default  %+v", o._statusCode, o.Payload)
}

func (o *ProvisionerIAMPolicyUpdateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ProvisionerIAMPolicyUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
