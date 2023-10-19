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

// ClusterGroupResourceServiceListReader is a Reader for the ClusterGroupResourceServiceList structure.
type ClusterGroupResourceServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterGroupResourceServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterGroupResourceServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterGroupResourceServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterGroupResourceServiceListOK creates a ClusterGroupResourceServiceListOK with default headers values
func NewClusterGroupResourceServiceListOK() *ClusterGroupResourceServiceListOK {
	return &ClusterGroupResourceServiceListOK{}
}

/*
ClusterGroupResourceServiceListOK describes a response with status code 200, with default header values.

A successful response.
*/
type ClusterGroupResourceServiceListOK struct {
	Payload *models.ClustergroupListClusterGroupsResponse
}

// IsSuccess returns true when this cluster group resource service list o k response has a 2xx status code
func (o *ClusterGroupResourceServiceListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster group resource service list o k response has a 3xx status code
func (o *ClusterGroupResourceServiceListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster group resource service list o k response has a 4xx status code
func (o *ClusterGroupResourceServiceListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster group resource service list o k response has a 5xx status code
func (o *ClusterGroupResourceServiceListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster group resource service list o k response a status code equal to that given
func (o *ClusterGroupResourceServiceListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ClusterGroupResourceServiceListOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clustergroups][%d] clusterGroupResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *ClusterGroupResourceServiceListOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clustergroups][%d] clusterGroupResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *ClusterGroupResourceServiceListOK) GetPayload() *models.ClustergroupListClusterGroupsResponse {
	return o.Payload
}

func (o *ClusterGroupResourceServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClustergroupListClusterGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterGroupResourceServiceListDefault creates a ClusterGroupResourceServiceListDefault with default headers values
func NewClusterGroupResourceServiceListDefault(code int) *ClusterGroupResourceServiceListDefault {
	return &ClusterGroupResourceServiceListDefault{
		_statusCode: code,
	}
}

/*
ClusterGroupResourceServiceListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ClusterGroupResourceServiceListDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the cluster group resource service list default response
func (o *ClusterGroupResourceServiceListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cluster group resource service list default response has a 2xx status code
func (o *ClusterGroupResourceServiceListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster group resource service list default response has a 3xx status code
func (o *ClusterGroupResourceServiceListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster group resource service list default response has a 4xx status code
func (o *ClusterGroupResourceServiceListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster group resource service list default response has a 5xx status code
func (o *ClusterGroupResourceServiceListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster group resource service list default response a status code equal to that given
func (o *ClusterGroupResourceServiceListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ClusterGroupResourceServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clustergroups][%d] ClusterGroupResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterGroupResourceServiceListDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clustergroups][%d] ClusterGroupResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterGroupResourceServiceListDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ClusterGroupResourceServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}