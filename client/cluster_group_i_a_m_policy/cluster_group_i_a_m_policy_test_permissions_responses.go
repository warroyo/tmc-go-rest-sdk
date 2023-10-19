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

// ClusterGroupIAMPolicyTestPermissionsReader is a Reader for the ClusterGroupIAMPolicyTestPermissions structure.
type ClusterGroupIAMPolicyTestPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterGroupIAMPolicyTestPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterGroupIAMPolicyTestPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterGroupIAMPolicyTestPermissionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterGroupIAMPolicyTestPermissionsOK creates a ClusterGroupIAMPolicyTestPermissionsOK with default headers values
func NewClusterGroupIAMPolicyTestPermissionsOK() *ClusterGroupIAMPolicyTestPermissionsOK {
	return &ClusterGroupIAMPolicyTestPermissionsOK{}
}

/*
ClusterGroupIAMPolicyTestPermissionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ClusterGroupIAMPolicyTestPermissionsOK struct {
	Payload *models.ClustergroupTestClusterGroupIAMPermissionsResponse
}

// IsSuccess returns true when this cluster group i a m policy test permissions o k response has a 2xx status code
func (o *ClusterGroupIAMPolicyTestPermissionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster group i a m policy test permissions o k response has a 3xx status code
func (o *ClusterGroupIAMPolicyTestPermissionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster group i a m policy test permissions o k response has a 4xx status code
func (o *ClusterGroupIAMPolicyTestPermissionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster group i a m policy test permissions o k response has a 5xx status code
func (o *ClusterGroupIAMPolicyTestPermissionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster group i a m policy test permissions o k response a status code equal to that given
func (o *ClusterGroupIAMPolicyTestPermissionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ClusterGroupIAMPolicyTestPermissionsOK) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/clustergroups:iam/{fullName.name}/testPermissions][%d] clusterGroupIAMPolicyTestPermissionsOK  %+v", 200, o.Payload)
}

func (o *ClusterGroupIAMPolicyTestPermissionsOK) String() string {
	return fmt.Sprintf("[POST /v1alpha1/clustergroups:iam/{fullName.name}/testPermissions][%d] clusterGroupIAMPolicyTestPermissionsOK  %+v", 200, o.Payload)
}

func (o *ClusterGroupIAMPolicyTestPermissionsOK) GetPayload() *models.ClustergroupTestClusterGroupIAMPermissionsResponse {
	return o.Payload
}

func (o *ClusterGroupIAMPolicyTestPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClustergroupTestClusterGroupIAMPermissionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterGroupIAMPolicyTestPermissionsDefault creates a ClusterGroupIAMPolicyTestPermissionsDefault with default headers values
func NewClusterGroupIAMPolicyTestPermissionsDefault(code int) *ClusterGroupIAMPolicyTestPermissionsDefault {
	return &ClusterGroupIAMPolicyTestPermissionsDefault{
		_statusCode: code,
	}
}

/*
ClusterGroupIAMPolicyTestPermissionsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ClusterGroupIAMPolicyTestPermissionsDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the cluster group i a m policy test permissions default response
func (o *ClusterGroupIAMPolicyTestPermissionsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cluster group i a m policy test permissions default response has a 2xx status code
func (o *ClusterGroupIAMPolicyTestPermissionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster group i a m policy test permissions default response has a 3xx status code
func (o *ClusterGroupIAMPolicyTestPermissionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster group i a m policy test permissions default response has a 4xx status code
func (o *ClusterGroupIAMPolicyTestPermissionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster group i a m policy test permissions default response has a 5xx status code
func (o *ClusterGroupIAMPolicyTestPermissionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster group i a m policy test permissions default response a status code equal to that given
func (o *ClusterGroupIAMPolicyTestPermissionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ClusterGroupIAMPolicyTestPermissionsDefault) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/clustergroups:iam/{fullName.name}/testPermissions][%d] ClusterGroupIAMPolicy_TestPermissions default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterGroupIAMPolicyTestPermissionsDefault) String() string {
	return fmt.Sprintf("[POST /v1alpha1/clustergroups:iam/{fullName.name}/testPermissions][%d] ClusterGroupIAMPolicy_TestPermissions default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterGroupIAMPolicyTestPermissionsDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ClusterGroupIAMPolicyTestPermissionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}