// Code generated by go-swagger; DO NOT EDIT.

package cluster_group_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// ClusterGroupResourceServiceGetReader is a Reader for the ClusterGroupResourceServiceGet structure.
type ClusterGroupResourceServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterGroupResourceServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterGroupResourceServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterGroupResourceServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterGroupResourceServiceGetOK creates a ClusterGroupResourceServiceGetOK with default headers values
func NewClusterGroupResourceServiceGetOK() *ClusterGroupResourceServiceGetOK {
	return &ClusterGroupResourceServiceGetOK{}
}

/*
ClusterGroupResourceServiceGetOK describes a response with status code 200, with default header values.

A successful response.
*/
type ClusterGroupResourceServiceGetOK struct {
	Payload *models.ClustergroupGetClusterGroupResponse
}

// IsSuccess returns true when this cluster group resource service get o k response has a 2xx status code
func (o *ClusterGroupResourceServiceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster group resource service get o k response has a 3xx status code
func (o *ClusterGroupResourceServiceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster group resource service get o k response has a 4xx status code
func (o *ClusterGroupResourceServiceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster group resource service get o k response has a 5xx status code
func (o *ClusterGroupResourceServiceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster group resource service get o k response a status code equal to that given
func (o *ClusterGroupResourceServiceGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *ClusterGroupResourceServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clustergroups/{fullName.name}][%d] clusterGroupResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *ClusterGroupResourceServiceGetOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clustergroups/{fullName.name}][%d] clusterGroupResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *ClusterGroupResourceServiceGetOK) GetPayload() *models.ClustergroupGetClusterGroupResponse {
	return o.Payload
}

func (o *ClusterGroupResourceServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClustergroupGetClusterGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterGroupResourceServiceGetDefault creates a ClusterGroupResourceServiceGetDefault with default headers values
func NewClusterGroupResourceServiceGetDefault(code int) *ClusterGroupResourceServiceGetDefault {
	return &ClusterGroupResourceServiceGetDefault{
		_statusCode: code,
	}
}

/*
ClusterGroupResourceServiceGetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ClusterGroupResourceServiceGetDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the cluster group resource service get default response
func (o *ClusterGroupResourceServiceGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cluster group resource service get default response has a 2xx status code
func (o *ClusterGroupResourceServiceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster group resource service get default response has a 3xx status code
func (o *ClusterGroupResourceServiceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster group resource service get default response has a 4xx status code
func (o *ClusterGroupResourceServiceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster group resource service get default response has a 5xx status code
func (o *ClusterGroupResourceServiceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster group resource service get default response a status code equal to that given
func (o *ClusterGroupResourceServiceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ClusterGroupResourceServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clustergroups/{fullName.name}][%d] ClusterGroupResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterGroupResourceServiceGetDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clustergroups/{fullName.name}][%d] ClusterGroupResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterGroupResourceServiceGetDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ClusterGroupResourceServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}