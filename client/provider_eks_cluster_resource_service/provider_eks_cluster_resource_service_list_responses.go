// Code generated by go-swagger; DO NOT EDIT.

package provider_eks_cluster_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// ProviderEksClusterResourceServiceListReader is a Reader for the ProviderEksClusterResourceServiceList structure.
type ProviderEksClusterResourceServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProviderEksClusterResourceServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProviderEksClusterResourceServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProviderEksClusterResourceServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProviderEksClusterResourceServiceListOK creates a ProviderEksClusterResourceServiceListOK with default headers values
func NewProviderEksClusterResourceServiceListOK() *ProviderEksClusterResourceServiceListOK {
	return &ProviderEksClusterResourceServiceListOK{}
}

/*
ProviderEksClusterResourceServiceListOK describes a response with status code 200, with default header values.

A successful response.
*/
type ProviderEksClusterResourceServiceListOK struct {
	Payload *models.ManageEksProvidereksclusterListProviderEksClustersResponse
}

// IsSuccess returns true when this provider eks cluster resource service list o k response has a 2xx status code
func (o *ProviderEksClusterResourceServiceListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this provider eks cluster resource service list o k response has a 3xx status code
func (o *ProviderEksClusterResourceServiceListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this provider eks cluster resource service list o k response has a 4xx status code
func (o *ProviderEksClusterResourceServiceListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this provider eks cluster resource service list o k response has a 5xx status code
func (o *ProviderEksClusterResourceServiceListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this provider eks cluster resource service list o k response a status code equal to that given
func (o *ProviderEksClusterResourceServiceListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProviderEksClusterResourceServiceListOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/manage/providereksclusters][%d] providerEksClusterResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *ProviderEksClusterResourceServiceListOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/manage/providereksclusters][%d] providerEksClusterResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *ProviderEksClusterResourceServiceListOK) GetPayload() *models.ManageEksProvidereksclusterListProviderEksClustersResponse {
	return o.Payload
}

func (o *ProviderEksClusterResourceServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManageEksProvidereksclusterListProviderEksClustersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProviderEksClusterResourceServiceListDefault creates a ProviderEksClusterResourceServiceListDefault with default headers values
func NewProviderEksClusterResourceServiceListDefault(code int) *ProviderEksClusterResourceServiceListDefault {
	return &ProviderEksClusterResourceServiceListDefault{
		_statusCode: code,
	}
}

/*
ProviderEksClusterResourceServiceListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ProviderEksClusterResourceServiceListDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the provider eks cluster resource service list default response
func (o *ProviderEksClusterResourceServiceListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this provider eks cluster resource service list default response has a 2xx status code
func (o *ProviderEksClusterResourceServiceListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this provider eks cluster resource service list default response has a 3xx status code
func (o *ProviderEksClusterResourceServiceListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this provider eks cluster resource service list default response has a 4xx status code
func (o *ProviderEksClusterResourceServiceListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this provider eks cluster resource service list default response has a 5xx status code
func (o *ProviderEksClusterResourceServiceListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this provider eks cluster resource service list default response a status code equal to that given
func (o *ProviderEksClusterResourceServiceListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ProviderEksClusterResourceServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/manage/providereksclusters][%d] ProviderEksClusterResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *ProviderEksClusterResourceServiceListDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/manage/providereksclusters][%d] ProviderEksClusterResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *ProviderEksClusterResourceServiceListDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ProviderEksClusterResourceServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
