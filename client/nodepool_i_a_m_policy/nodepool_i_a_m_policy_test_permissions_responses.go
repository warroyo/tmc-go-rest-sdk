// Code generated by go-swagger; DO NOT EDIT.

package nodepool_i_a_m_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// NodepoolIAMPolicyTestPermissionsReader is a Reader for the NodepoolIAMPolicyTestPermissions structure.
type NodepoolIAMPolicyTestPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodepoolIAMPolicyTestPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodepoolIAMPolicyTestPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNodepoolIAMPolicyTestPermissionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodepoolIAMPolicyTestPermissionsOK creates a NodepoolIAMPolicyTestPermissionsOK with default headers values
func NewNodepoolIAMPolicyTestPermissionsOK() *NodepoolIAMPolicyTestPermissionsOK {
	return &NodepoolIAMPolicyTestPermissionsOK{}
}

/*
NodepoolIAMPolicyTestPermissionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type NodepoolIAMPolicyTestPermissionsOK struct {
	Payload *models.EksclusterNodepoolTestNodepoolIAMPermissionsResponse
}

// IsSuccess returns true when this nodepool i a m policy test permissions o k response has a 2xx status code
func (o *NodepoolIAMPolicyTestPermissionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nodepool i a m policy test permissions o k response has a 3xx status code
func (o *NodepoolIAMPolicyTestPermissionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nodepool i a m policy test permissions o k response has a 4xx status code
func (o *NodepoolIAMPolicyTestPermissionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nodepool i a m policy test permissions o k response has a 5xx status code
func (o *NodepoolIAMPolicyTestPermissionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nodepool i a m policy test permissions o k response a status code equal to that given
func (o *NodepoolIAMPolicyTestPermissionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *NodepoolIAMPolicyTestPermissionsOK) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/eksclusters/{fullName.eksClusterName}/nodepools:iam/{fullName.name}/testPermissions][%d] nodepoolIAMPolicyTestPermissionsOK  %+v", 200, o.Payload)
}

func (o *NodepoolIAMPolicyTestPermissionsOK) String() string {
	return fmt.Sprintf("[POST /v1alpha1/eksclusters/{fullName.eksClusterName}/nodepools:iam/{fullName.name}/testPermissions][%d] nodepoolIAMPolicyTestPermissionsOK  %+v", 200, o.Payload)
}

func (o *NodepoolIAMPolicyTestPermissionsOK) GetPayload() *models.EksclusterNodepoolTestNodepoolIAMPermissionsResponse {
	return o.Payload
}

func (o *NodepoolIAMPolicyTestPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EksclusterNodepoolTestNodepoolIAMPermissionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodepoolIAMPolicyTestPermissionsDefault creates a NodepoolIAMPolicyTestPermissionsDefault with default headers values
func NewNodepoolIAMPolicyTestPermissionsDefault(code int) *NodepoolIAMPolicyTestPermissionsDefault {
	return &NodepoolIAMPolicyTestPermissionsDefault{
		_statusCode: code,
	}
}

/*
NodepoolIAMPolicyTestPermissionsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type NodepoolIAMPolicyTestPermissionsDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the nodepool i a m policy test permissions default response
func (o *NodepoolIAMPolicyTestPermissionsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this nodepool i a m policy test permissions default response has a 2xx status code
func (o *NodepoolIAMPolicyTestPermissionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nodepool i a m policy test permissions default response has a 3xx status code
func (o *NodepoolIAMPolicyTestPermissionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nodepool i a m policy test permissions default response has a 4xx status code
func (o *NodepoolIAMPolicyTestPermissionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nodepool i a m policy test permissions default response has a 5xx status code
func (o *NodepoolIAMPolicyTestPermissionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nodepool i a m policy test permissions default response a status code equal to that given
func (o *NodepoolIAMPolicyTestPermissionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *NodepoolIAMPolicyTestPermissionsDefault) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/eksclusters/{fullName.eksClusterName}/nodepools:iam/{fullName.name}/testPermissions][%d] NodepoolIAMPolicy_TestPermissions default  %+v", o._statusCode, o.Payload)
}

func (o *NodepoolIAMPolicyTestPermissionsDefault) String() string {
	return fmt.Sprintf("[POST /v1alpha1/eksclusters/{fullName.eksClusterName}/nodepools:iam/{fullName.name}/testPermissions][%d] NodepoolIAMPolicy_TestPermissions default  %+v", o._statusCode, o.Payload)
}

func (o *NodepoolIAMPolicyTestPermissionsDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *NodepoolIAMPolicyTestPermissionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
