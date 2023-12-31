// Code generated by go-swagger; DO NOT EDIT.

package management_cluster_i_a_m_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// ManagementClusterIAMPolicyUpdateReader is a Reader for the ManagementClusterIAMPolicyUpdate structure.
type ManagementClusterIAMPolicyUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ManagementClusterIAMPolicyUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewManagementClusterIAMPolicyUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewManagementClusterIAMPolicyUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewManagementClusterIAMPolicyUpdateOK creates a ManagementClusterIAMPolicyUpdateOK with default headers values
func NewManagementClusterIAMPolicyUpdateOK() *ManagementClusterIAMPolicyUpdateOK {
	return &ManagementClusterIAMPolicyUpdateOK{}
}

/*
ManagementClusterIAMPolicyUpdateOK describes a response with status code 200, with default header values.

A successful response.
*/
type ManagementClusterIAMPolicyUpdateOK struct {
	Payload *models.ManagementclusterUpdateManagementClusterIAMPolicyResponse
}

// IsSuccess returns true when this management cluster i a m policy update o k response has a 2xx status code
func (o *ManagementClusterIAMPolicyUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this management cluster i a m policy update o k response has a 3xx status code
func (o *ManagementClusterIAMPolicyUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this management cluster i a m policy update o k response has a 4xx status code
func (o *ManagementClusterIAMPolicyUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this management cluster i a m policy update o k response has a 5xx status code
func (o *ManagementClusterIAMPolicyUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this management cluster i a m policy update o k response a status code equal to that given
func (o *ManagementClusterIAMPolicyUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ManagementClusterIAMPolicyUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/managementclusters:iam/{fullName.name}][%d] managementClusterIAMPolicyUpdateOK  %+v", 200, o.Payload)
}

func (o *ManagementClusterIAMPolicyUpdateOK) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/managementclusters:iam/{fullName.name}][%d] managementClusterIAMPolicyUpdateOK  %+v", 200, o.Payload)
}

func (o *ManagementClusterIAMPolicyUpdateOK) GetPayload() *models.ManagementclusterUpdateManagementClusterIAMPolicyResponse {
	return o.Payload
}

func (o *ManagementClusterIAMPolicyUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagementclusterUpdateManagementClusterIAMPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManagementClusterIAMPolicyUpdateDefault creates a ManagementClusterIAMPolicyUpdateDefault with default headers values
func NewManagementClusterIAMPolicyUpdateDefault(code int) *ManagementClusterIAMPolicyUpdateDefault {
	return &ManagementClusterIAMPolicyUpdateDefault{
		_statusCode: code,
	}
}

/*
ManagementClusterIAMPolicyUpdateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ManagementClusterIAMPolicyUpdateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the management cluster i a m policy update default response
func (o *ManagementClusterIAMPolicyUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this management cluster i a m policy update default response has a 2xx status code
func (o *ManagementClusterIAMPolicyUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this management cluster i a m policy update default response has a 3xx status code
func (o *ManagementClusterIAMPolicyUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this management cluster i a m policy update default response has a 4xx status code
func (o *ManagementClusterIAMPolicyUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this management cluster i a m policy update default response has a 5xx status code
func (o *ManagementClusterIAMPolicyUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this management cluster i a m policy update default response a status code equal to that given
func (o *ManagementClusterIAMPolicyUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ManagementClusterIAMPolicyUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/managementclusters:iam/{fullName.name}][%d] ManagementClusterIAMPolicy_Update default  %+v", o._statusCode, o.Payload)
}

func (o *ManagementClusterIAMPolicyUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/managementclusters:iam/{fullName.name}][%d] ManagementClusterIAMPolicy_Update default  %+v", o._statusCode, o.Payload)
}

func (o *ManagementClusterIAMPolicyUpdateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ManagementClusterIAMPolicyUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
