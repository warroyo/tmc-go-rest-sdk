// Code generated by go-swagger; DO NOT EDIT.

package cluster_group_i_a_m_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// ClusterGroupIAMPolicyGetReader is a Reader for the ClusterGroupIAMPolicyGet structure.
type ClusterGroupIAMPolicyGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterGroupIAMPolicyGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterGroupIAMPolicyGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterGroupIAMPolicyGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterGroupIAMPolicyGetOK creates a ClusterGroupIAMPolicyGetOK with default headers values
func NewClusterGroupIAMPolicyGetOK() *ClusterGroupIAMPolicyGetOK {
	return &ClusterGroupIAMPolicyGetOK{}
}

/*
ClusterGroupIAMPolicyGetOK describes a response with status code 200, with default header values.

A successful response.
*/
type ClusterGroupIAMPolicyGetOK struct {
	Payload *models.ClustergroupGetClusterGroupIAMPolicyResponse
}

// IsSuccess returns true when this cluster group i a m policy get o k response has a 2xx status code
func (o *ClusterGroupIAMPolicyGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster group i a m policy get o k response has a 3xx status code
func (o *ClusterGroupIAMPolicyGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster group i a m policy get o k response has a 4xx status code
func (o *ClusterGroupIAMPolicyGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster group i a m policy get o k response has a 5xx status code
func (o *ClusterGroupIAMPolicyGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster group i a m policy get o k response a status code equal to that given
func (o *ClusterGroupIAMPolicyGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *ClusterGroupIAMPolicyGetOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clustergroups:iam/{fullName.name}][%d] clusterGroupIAMPolicyGetOK  %+v", 200, o.Payload)
}

func (o *ClusterGroupIAMPolicyGetOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clustergroups:iam/{fullName.name}][%d] clusterGroupIAMPolicyGetOK  %+v", 200, o.Payload)
}

func (o *ClusterGroupIAMPolicyGetOK) GetPayload() *models.ClustergroupGetClusterGroupIAMPolicyResponse {
	return o.Payload
}

func (o *ClusterGroupIAMPolicyGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClustergroupGetClusterGroupIAMPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterGroupIAMPolicyGetDefault creates a ClusterGroupIAMPolicyGetDefault with default headers values
func NewClusterGroupIAMPolicyGetDefault(code int) *ClusterGroupIAMPolicyGetDefault {
	return &ClusterGroupIAMPolicyGetDefault{
		_statusCode: code,
	}
}

/*
ClusterGroupIAMPolicyGetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ClusterGroupIAMPolicyGetDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the cluster group i a m policy get default response
func (o *ClusterGroupIAMPolicyGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cluster group i a m policy get default response has a 2xx status code
func (o *ClusterGroupIAMPolicyGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster group i a m policy get default response has a 3xx status code
func (o *ClusterGroupIAMPolicyGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster group i a m policy get default response has a 4xx status code
func (o *ClusterGroupIAMPolicyGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster group i a m policy get default response has a 5xx status code
func (o *ClusterGroupIAMPolicyGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster group i a m policy get default response a status code equal to that given
func (o *ClusterGroupIAMPolicyGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ClusterGroupIAMPolicyGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clustergroups:iam/{fullName.name}][%d] ClusterGroupIAMPolicy_Get default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterGroupIAMPolicyGetDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clustergroups:iam/{fullName.name}][%d] ClusterGroupIAMPolicy_Get default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterGroupIAMPolicyGetDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ClusterGroupIAMPolicyGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
