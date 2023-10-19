// Code generated by go-swagger; DO NOT EDIT.

package cluster_i_a_m_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// ClusterIAMPolicyTestPermissionsReader is a Reader for the ClusterIAMPolicyTestPermissions structure.
type ClusterIAMPolicyTestPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterIAMPolicyTestPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterIAMPolicyTestPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterIAMPolicyTestPermissionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterIAMPolicyTestPermissionsOK creates a ClusterIAMPolicyTestPermissionsOK with default headers values
func NewClusterIAMPolicyTestPermissionsOK() *ClusterIAMPolicyTestPermissionsOK {
	return &ClusterIAMPolicyTestPermissionsOK{}
}

/*
ClusterIAMPolicyTestPermissionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ClusterIAMPolicyTestPermissionsOK struct {
	Payload *models.ClusterTestClusterIAMPermissionsResponse
}

// IsSuccess returns true when this cluster i a m policy test permissions o k response has a 2xx status code
func (o *ClusterIAMPolicyTestPermissionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster i a m policy test permissions o k response has a 3xx status code
func (o *ClusterIAMPolicyTestPermissionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster i a m policy test permissions o k response has a 4xx status code
func (o *ClusterIAMPolicyTestPermissionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster i a m policy test permissions o k response has a 5xx status code
func (o *ClusterIAMPolicyTestPermissionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster i a m policy test permissions o k response a status code equal to that given
func (o *ClusterIAMPolicyTestPermissionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ClusterIAMPolicyTestPermissionsOK) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/clusters:iam/{fullName.name}/testPermissions][%d] clusterIAMPolicyTestPermissionsOK  %+v", 200, o.Payload)
}

func (o *ClusterIAMPolicyTestPermissionsOK) String() string {
	return fmt.Sprintf("[POST /v1alpha1/clusters:iam/{fullName.name}/testPermissions][%d] clusterIAMPolicyTestPermissionsOK  %+v", 200, o.Payload)
}

func (o *ClusterIAMPolicyTestPermissionsOK) GetPayload() *models.ClusterTestClusterIAMPermissionsResponse {
	return o.Payload
}

func (o *ClusterIAMPolicyTestPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterTestClusterIAMPermissionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterIAMPolicyTestPermissionsDefault creates a ClusterIAMPolicyTestPermissionsDefault with default headers values
func NewClusterIAMPolicyTestPermissionsDefault(code int) *ClusterIAMPolicyTestPermissionsDefault {
	return &ClusterIAMPolicyTestPermissionsDefault{
		_statusCode: code,
	}
}

/*
ClusterIAMPolicyTestPermissionsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ClusterIAMPolicyTestPermissionsDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the cluster i a m policy test permissions default response
func (o *ClusterIAMPolicyTestPermissionsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cluster i a m policy test permissions default response has a 2xx status code
func (o *ClusterIAMPolicyTestPermissionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster i a m policy test permissions default response has a 3xx status code
func (o *ClusterIAMPolicyTestPermissionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster i a m policy test permissions default response has a 4xx status code
func (o *ClusterIAMPolicyTestPermissionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster i a m policy test permissions default response has a 5xx status code
func (o *ClusterIAMPolicyTestPermissionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster i a m policy test permissions default response a status code equal to that given
func (o *ClusterIAMPolicyTestPermissionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ClusterIAMPolicyTestPermissionsDefault) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/clusters:iam/{fullName.name}/testPermissions][%d] ClusterIAMPolicy_TestPermissions default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterIAMPolicyTestPermissionsDefault) String() string {
	return fmt.Sprintf("[POST /v1alpha1/clusters:iam/{fullName.name}/testPermissions][%d] ClusterIAMPolicy_TestPermissions default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterIAMPolicyTestPermissionsDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ClusterIAMPolicyTestPermissionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}