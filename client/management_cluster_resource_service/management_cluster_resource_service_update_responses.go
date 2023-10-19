// Code generated by go-swagger; DO NOT EDIT.

package management_cluster_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// ManagementClusterResourceServiceUpdateReader is a Reader for the ManagementClusterResourceServiceUpdate structure.
type ManagementClusterResourceServiceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ManagementClusterResourceServiceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewManagementClusterResourceServiceUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewManagementClusterResourceServiceUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewManagementClusterResourceServiceUpdateOK creates a ManagementClusterResourceServiceUpdateOK with default headers values
func NewManagementClusterResourceServiceUpdateOK() *ManagementClusterResourceServiceUpdateOK {
	return &ManagementClusterResourceServiceUpdateOK{}
}

/*
ManagementClusterResourceServiceUpdateOK describes a response with status code 200, with default header values.

A successful response.
*/
type ManagementClusterResourceServiceUpdateOK struct {
	Payload *models.ManagementclusterUpdateManagementClusterResponse
}

// IsSuccess returns true when this management cluster resource service update o k response has a 2xx status code
func (o *ManagementClusterResourceServiceUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this management cluster resource service update o k response has a 3xx status code
func (o *ManagementClusterResourceServiceUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this management cluster resource service update o k response has a 4xx status code
func (o *ManagementClusterResourceServiceUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this management cluster resource service update o k response has a 5xx status code
func (o *ManagementClusterResourceServiceUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this management cluster resource service update o k response a status code equal to that given
func (o *ManagementClusterResourceServiceUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ManagementClusterResourceServiceUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/managementclusters/{managementCluster.fullName.name}][%d] managementClusterResourceServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *ManagementClusterResourceServiceUpdateOK) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/managementclusters/{managementCluster.fullName.name}][%d] managementClusterResourceServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *ManagementClusterResourceServiceUpdateOK) GetPayload() *models.ManagementclusterUpdateManagementClusterResponse {
	return o.Payload
}

func (o *ManagementClusterResourceServiceUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagementclusterUpdateManagementClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManagementClusterResourceServiceUpdateDefault creates a ManagementClusterResourceServiceUpdateDefault with default headers values
func NewManagementClusterResourceServiceUpdateDefault(code int) *ManagementClusterResourceServiceUpdateDefault {
	return &ManagementClusterResourceServiceUpdateDefault{
		_statusCode: code,
	}
}

/*
ManagementClusterResourceServiceUpdateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ManagementClusterResourceServiceUpdateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the management cluster resource service update default response
func (o *ManagementClusterResourceServiceUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this management cluster resource service update default response has a 2xx status code
func (o *ManagementClusterResourceServiceUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this management cluster resource service update default response has a 3xx status code
func (o *ManagementClusterResourceServiceUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this management cluster resource service update default response has a 4xx status code
func (o *ManagementClusterResourceServiceUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this management cluster resource service update default response has a 5xx status code
func (o *ManagementClusterResourceServiceUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this management cluster resource service update default response a status code equal to that given
func (o *ManagementClusterResourceServiceUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ManagementClusterResourceServiceUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/managementclusters/{managementCluster.fullName.name}][%d] ManagementClusterResourceService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *ManagementClusterResourceServiceUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/managementclusters/{managementCluster.fullName.name}][%d] ManagementClusterResourceService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *ManagementClusterResourceServiceUpdateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ManagementClusterResourceServiceUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
