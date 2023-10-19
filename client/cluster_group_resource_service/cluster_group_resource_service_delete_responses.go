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

// ClusterGroupResourceServiceDeleteReader is a Reader for the ClusterGroupResourceServiceDelete structure.
type ClusterGroupResourceServiceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterGroupResourceServiceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterGroupResourceServiceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterGroupResourceServiceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterGroupResourceServiceDeleteOK creates a ClusterGroupResourceServiceDeleteOK with default headers values
func NewClusterGroupResourceServiceDeleteOK() *ClusterGroupResourceServiceDeleteOK {
	return &ClusterGroupResourceServiceDeleteOK{}
}

/*
ClusterGroupResourceServiceDeleteOK describes a response with status code 200, with default header values.

A successful response.
*/
type ClusterGroupResourceServiceDeleteOK struct {
	Payload *models.ClustergroupDeleteClusterGroupResponse
}

// IsSuccess returns true when this cluster group resource service delete o k response has a 2xx status code
func (o *ClusterGroupResourceServiceDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster group resource service delete o k response has a 3xx status code
func (o *ClusterGroupResourceServiceDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster group resource service delete o k response has a 4xx status code
func (o *ClusterGroupResourceServiceDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster group resource service delete o k response has a 5xx status code
func (o *ClusterGroupResourceServiceDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster group resource service delete o k response a status code equal to that given
func (o *ClusterGroupResourceServiceDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *ClusterGroupResourceServiceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clustergroups/{fullName.name}][%d] clusterGroupResourceServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *ClusterGroupResourceServiceDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clustergroups/{fullName.name}][%d] clusterGroupResourceServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *ClusterGroupResourceServiceDeleteOK) GetPayload() *models.ClustergroupDeleteClusterGroupResponse {
	return o.Payload
}

func (o *ClusterGroupResourceServiceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClustergroupDeleteClusterGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterGroupResourceServiceDeleteDefault creates a ClusterGroupResourceServiceDeleteDefault with default headers values
func NewClusterGroupResourceServiceDeleteDefault(code int) *ClusterGroupResourceServiceDeleteDefault {
	return &ClusterGroupResourceServiceDeleteDefault{
		_statusCode: code,
	}
}

/*
ClusterGroupResourceServiceDeleteDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ClusterGroupResourceServiceDeleteDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the cluster group resource service delete default response
func (o *ClusterGroupResourceServiceDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cluster group resource service delete default response has a 2xx status code
func (o *ClusterGroupResourceServiceDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster group resource service delete default response has a 3xx status code
func (o *ClusterGroupResourceServiceDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster group resource service delete default response has a 4xx status code
func (o *ClusterGroupResourceServiceDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster group resource service delete default response has a 5xx status code
func (o *ClusterGroupResourceServiceDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster group resource service delete default response a status code equal to that given
func (o *ClusterGroupResourceServiceDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ClusterGroupResourceServiceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clustergroups/{fullName.name}][%d] ClusterGroupResourceService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterGroupResourceServiceDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clustergroups/{fullName.name}][%d] ClusterGroupResourceService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterGroupResourceServiceDeleteDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ClusterGroupResourceServiceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}