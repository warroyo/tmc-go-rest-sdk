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

// ClusterGroupIAMPolicyUpdateReader is a Reader for the ClusterGroupIAMPolicyUpdate structure.
type ClusterGroupIAMPolicyUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterGroupIAMPolicyUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterGroupIAMPolicyUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterGroupIAMPolicyUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterGroupIAMPolicyUpdateOK creates a ClusterGroupIAMPolicyUpdateOK with default headers values
func NewClusterGroupIAMPolicyUpdateOK() *ClusterGroupIAMPolicyUpdateOK {
	return &ClusterGroupIAMPolicyUpdateOK{}
}

/*
ClusterGroupIAMPolicyUpdateOK describes a response with status code 200, with default header values.

A successful response.
*/
type ClusterGroupIAMPolicyUpdateOK struct {
	Payload *models.ClustergroupUpdateClusterGroupIAMPolicyResponse
}

// IsSuccess returns true when this cluster group i a m policy update o k response has a 2xx status code
func (o *ClusterGroupIAMPolicyUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster group i a m policy update o k response has a 3xx status code
func (o *ClusterGroupIAMPolicyUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster group i a m policy update o k response has a 4xx status code
func (o *ClusterGroupIAMPolicyUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster group i a m policy update o k response has a 5xx status code
func (o *ClusterGroupIAMPolicyUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster group i a m policy update o k response a status code equal to that given
func (o *ClusterGroupIAMPolicyUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ClusterGroupIAMPolicyUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/clustergroups:iam/{fullName.name}][%d] clusterGroupIAMPolicyUpdateOK  %+v", 200, o.Payload)
}

func (o *ClusterGroupIAMPolicyUpdateOK) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/clustergroups:iam/{fullName.name}][%d] clusterGroupIAMPolicyUpdateOK  %+v", 200, o.Payload)
}

func (o *ClusterGroupIAMPolicyUpdateOK) GetPayload() *models.ClustergroupUpdateClusterGroupIAMPolicyResponse {
	return o.Payload
}

func (o *ClusterGroupIAMPolicyUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClustergroupUpdateClusterGroupIAMPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterGroupIAMPolicyUpdateDefault creates a ClusterGroupIAMPolicyUpdateDefault with default headers values
func NewClusterGroupIAMPolicyUpdateDefault(code int) *ClusterGroupIAMPolicyUpdateDefault {
	return &ClusterGroupIAMPolicyUpdateDefault{
		_statusCode: code,
	}
}

/*
ClusterGroupIAMPolicyUpdateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ClusterGroupIAMPolicyUpdateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the cluster group i a m policy update default response
func (o *ClusterGroupIAMPolicyUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cluster group i a m policy update default response has a 2xx status code
func (o *ClusterGroupIAMPolicyUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster group i a m policy update default response has a 3xx status code
func (o *ClusterGroupIAMPolicyUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster group i a m policy update default response has a 4xx status code
func (o *ClusterGroupIAMPolicyUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster group i a m policy update default response has a 5xx status code
func (o *ClusterGroupIAMPolicyUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster group i a m policy update default response a status code equal to that given
func (o *ClusterGroupIAMPolicyUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ClusterGroupIAMPolicyUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/clustergroups:iam/{fullName.name}][%d] ClusterGroupIAMPolicy_Update default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterGroupIAMPolicyUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/clustergroups:iam/{fullName.name}][%d] ClusterGroupIAMPolicy_Update default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterGroupIAMPolicyUpdateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ClusterGroupIAMPolicyUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}