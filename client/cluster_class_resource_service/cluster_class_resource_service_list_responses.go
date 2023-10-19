// Code generated by go-swagger; DO NOT EDIT.

package cluster_class_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// ClusterClassResourceServiceListReader is a Reader for the ClusterClassResourceServiceList structure.
type ClusterClassResourceServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterClassResourceServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterClassResourceServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterClassResourceServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterClassResourceServiceListOK creates a ClusterClassResourceServiceListOK with default headers values
func NewClusterClassResourceServiceListOK() *ClusterClassResourceServiceListOK {
	return &ClusterClassResourceServiceListOK{}
}

/*
ClusterClassResourceServiceListOK describes a response with status code 200, with default header values.

A successful response.
*/
type ClusterClassResourceServiceListOK struct {
	Payload *models.ManagementclusterProvisionerClusterclassListClusterClassesResponse
}

// IsSuccess returns true when this cluster class resource service list o k response has a 2xx status code
func (o *ClusterClassResourceServiceListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster class resource service list o k response has a 3xx status code
func (o *ClusterClassResourceServiceListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster class resource service list o k response has a 4xx status code
func (o *ClusterClassResourceServiceListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster class resource service list o k response has a 5xx status code
func (o *ClusterClassResourceServiceListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster class resource service list o k response a status code equal to that given
func (o *ClusterClassResourceServiceListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ClusterClassResourceServiceListOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/managementclusters/{searchScope.managementClusterName}/provisioners/{searchScope.provisionerName}/clusterclasses][%d] clusterClassResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *ClusterClassResourceServiceListOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/managementclusters/{searchScope.managementClusterName}/provisioners/{searchScope.provisionerName}/clusterclasses][%d] clusterClassResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *ClusterClassResourceServiceListOK) GetPayload() *models.ManagementclusterProvisionerClusterclassListClusterClassesResponse {
	return o.Payload
}

func (o *ClusterClassResourceServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagementclusterProvisionerClusterclassListClusterClassesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterClassResourceServiceListDefault creates a ClusterClassResourceServiceListDefault with default headers values
func NewClusterClassResourceServiceListDefault(code int) *ClusterClassResourceServiceListDefault {
	return &ClusterClassResourceServiceListDefault{
		_statusCode: code,
	}
}

/*
ClusterClassResourceServiceListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ClusterClassResourceServiceListDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the cluster class resource service list default response
func (o *ClusterClassResourceServiceListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cluster class resource service list default response has a 2xx status code
func (o *ClusterClassResourceServiceListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster class resource service list default response has a 3xx status code
func (o *ClusterClassResourceServiceListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster class resource service list default response has a 4xx status code
func (o *ClusterClassResourceServiceListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster class resource service list default response has a 5xx status code
func (o *ClusterClassResourceServiceListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster class resource service list default response a status code equal to that given
func (o *ClusterClassResourceServiceListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ClusterClassResourceServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/managementclusters/{searchScope.managementClusterName}/provisioners/{searchScope.provisionerName}/clusterclasses][%d] ClusterClassResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterClassResourceServiceListDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/managementclusters/{searchScope.managementClusterName}/provisioners/{searchScope.provisionerName}/clusterclasses][%d] ClusterClassResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterClassResourceServiceListDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ClusterClassResourceServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}