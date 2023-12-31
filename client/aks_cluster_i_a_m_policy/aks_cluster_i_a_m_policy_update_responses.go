// Code generated by go-swagger; DO NOT EDIT.

package aks_cluster_i_a_m_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// AksClusterIAMPolicyUpdateReader is a Reader for the AksClusterIAMPolicyUpdate structure.
type AksClusterIAMPolicyUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AksClusterIAMPolicyUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAksClusterIAMPolicyUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAksClusterIAMPolicyUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAksClusterIAMPolicyUpdateOK creates a AksClusterIAMPolicyUpdateOK with default headers values
func NewAksClusterIAMPolicyUpdateOK() *AksClusterIAMPolicyUpdateOK {
	return &AksClusterIAMPolicyUpdateOK{}
}

/*
AksClusterIAMPolicyUpdateOK describes a response with status code 200, with default header values.

A successful response.
*/
type AksClusterIAMPolicyUpdateOK struct {
	Payload *models.AksclusterUpdateAksClusterIAMPolicyResponse
}

// IsSuccess returns true when this aks cluster i a m policy update o k response has a 2xx status code
func (o *AksClusterIAMPolicyUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aks cluster i a m policy update o k response has a 3xx status code
func (o *AksClusterIAMPolicyUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aks cluster i a m policy update o k response has a 4xx status code
func (o *AksClusterIAMPolicyUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aks cluster i a m policy update o k response has a 5xx status code
func (o *AksClusterIAMPolicyUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aks cluster i a m policy update o k response a status code equal to that given
func (o *AksClusterIAMPolicyUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *AksClusterIAMPolicyUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/aksclusters:iam/{fullName.name}][%d] aksClusterIAMPolicyUpdateOK  %+v", 200, o.Payload)
}

func (o *AksClusterIAMPolicyUpdateOK) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/aksclusters:iam/{fullName.name}][%d] aksClusterIAMPolicyUpdateOK  %+v", 200, o.Payload)
}

func (o *AksClusterIAMPolicyUpdateOK) GetPayload() *models.AksclusterUpdateAksClusterIAMPolicyResponse {
	return o.Payload
}

func (o *AksClusterIAMPolicyUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AksclusterUpdateAksClusterIAMPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAksClusterIAMPolicyUpdateDefault creates a AksClusterIAMPolicyUpdateDefault with default headers values
func NewAksClusterIAMPolicyUpdateDefault(code int) *AksClusterIAMPolicyUpdateDefault {
	return &AksClusterIAMPolicyUpdateDefault{
		_statusCode: code,
	}
}

/*
AksClusterIAMPolicyUpdateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AksClusterIAMPolicyUpdateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the aks cluster i a m policy update default response
func (o *AksClusterIAMPolicyUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this aks cluster i a m policy update default response has a 2xx status code
func (o *AksClusterIAMPolicyUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this aks cluster i a m policy update default response has a 3xx status code
func (o *AksClusterIAMPolicyUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this aks cluster i a m policy update default response has a 4xx status code
func (o *AksClusterIAMPolicyUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this aks cluster i a m policy update default response has a 5xx status code
func (o *AksClusterIAMPolicyUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this aks cluster i a m policy update default response a status code equal to that given
func (o *AksClusterIAMPolicyUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AksClusterIAMPolicyUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/aksclusters:iam/{fullName.name}][%d] AksClusterIAMPolicy_Update default  %+v", o._statusCode, o.Payload)
}

func (o *AksClusterIAMPolicyUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/aksclusters:iam/{fullName.name}][%d] AksClusterIAMPolicy_Update default  %+v", o._statusCode, o.Payload)
}

func (o *AksClusterIAMPolicyUpdateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *AksClusterIAMPolicyUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
