// Code generated by go-swagger; DO NOT EDIT.

package kustomization_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// KustomizationResourceServiceUpdateReader is a Reader for the KustomizationResourceServiceUpdate structure.
type KustomizationResourceServiceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KustomizationResourceServiceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKustomizationResourceServiceUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKustomizationResourceServiceUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKustomizationResourceServiceUpdateOK creates a KustomizationResourceServiceUpdateOK with default headers values
func NewKustomizationResourceServiceUpdateOK() *KustomizationResourceServiceUpdateOK {
	return &KustomizationResourceServiceUpdateOK{}
}

/*
KustomizationResourceServiceUpdateOK describes a response with status code 200, with default header values.

A successful response.
*/
type KustomizationResourceServiceUpdateOK struct {
	Payload *models.ClusterNamespaceFluxcdKustomizationUpdateKustomizationResponse
}

// IsSuccess returns true when this kustomization resource service update o k response has a 2xx status code
func (o *KustomizationResourceServiceUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kustomization resource service update o k response has a 3xx status code
func (o *KustomizationResourceServiceUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kustomization resource service update o k response has a 4xx status code
func (o *KustomizationResourceServiceUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kustomization resource service update o k response has a 5xx status code
func (o *KustomizationResourceServiceUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kustomization resource service update o k response a status code equal to that given
func (o *KustomizationResourceServiceUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *KustomizationResourceServiceUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/clusters/{kustomization.fullName.clusterName}/namespaces/{kustomization.fullName.namespaceName}/fluxcd/kustomizations/{kustomization.fullName.name}][%d] kustomizationResourceServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *KustomizationResourceServiceUpdateOK) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/clusters/{kustomization.fullName.clusterName}/namespaces/{kustomization.fullName.namespaceName}/fluxcd/kustomizations/{kustomization.fullName.name}][%d] kustomizationResourceServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *KustomizationResourceServiceUpdateOK) GetPayload() *models.ClusterNamespaceFluxcdKustomizationUpdateKustomizationResponse {
	return o.Payload
}

func (o *KustomizationResourceServiceUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterNamespaceFluxcdKustomizationUpdateKustomizationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKustomizationResourceServiceUpdateDefault creates a KustomizationResourceServiceUpdateDefault with default headers values
func NewKustomizationResourceServiceUpdateDefault(code int) *KustomizationResourceServiceUpdateDefault {
	return &KustomizationResourceServiceUpdateDefault{
		_statusCode: code,
	}
}

/*
KustomizationResourceServiceUpdateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type KustomizationResourceServiceUpdateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the kustomization resource service update default response
func (o *KustomizationResourceServiceUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this kustomization resource service update default response has a 2xx status code
func (o *KustomizationResourceServiceUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this kustomization resource service update default response has a 3xx status code
func (o *KustomizationResourceServiceUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this kustomization resource service update default response has a 4xx status code
func (o *KustomizationResourceServiceUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this kustomization resource service update default response has a 5xx status code
func (o *KustomizationResourceServiceUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this kustomization resource service update default response a status code equal to that given
func (o *KustomizationResourceServiceUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *KustomizationResourceServiceUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/clusters/{kustomization.fullName.clusterName}/namespaces/{kustomization.fullName.namespaceName}/fluxcd/kustomizations/{kustomization.fullName.name}][%d] KustomizationResourceService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *KustomizationResourceServiceUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/clusters/{kustomization.fullName.clusterName}/namespaces/{kustomization.fullName.namespaceName}/fluxcd/kustomizations/{kustomization.fullName.name}][%d] KustomizationResourceService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *KustomizationResourceServiceUpdateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *KustomizationResourceServiceUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
