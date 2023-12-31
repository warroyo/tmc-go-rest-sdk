// Code generated by go-swagger; DO NOT EDIT.

package install_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// InstallResourceServiceGetReader is a Reader for the InstallResourceServiceGet structure.
type InstallResourceServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstallResourceServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInstallResourceServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewInstallResourceServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInstallResourceServiceGetOK creates a InstallResourceServiceGetOK with default headers values
func NewInstallResourceServiceGetOK() *InstallResourceServiceGetOK {
	return &InstallResourceServiceGetOK{}
}

/*
InstallResourceServiceGetOK describes a response with status code 200, with default header values.

A successful response.
*/
type InstallResourceServiceGetOK struct {
	Payload *models.ClusterNamespaceTanzupackageInstallGetInstallResponse
}

// IsSuccess returns true when this install resource service get o k response has a 2xx status code
func (o *InstallResourceServiceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this install resource service get o k response has a 3xx status code
func (o *InstallResourceServiceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this install resource service get o k response has a 4xx status code
func (o *InstallResourceServiceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this install resource service get o k response has a 5xx status code
func (o *InstallResourceServiceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this install resource service get o k response a status code equal to that given
func (o *InstallResourceServiceGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *InstallResourceServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/tanzupackage/installs/{fullName.name}][%d] installResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *InstallResourceServiceGetOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/tanzupackage/installs/{fullName.name}][%d] installResourceServiceGetOK  %+v", 200, o.Payload)
}

func (o *InstallResourceServiceGetOK) GetPayload() *models.ClusterNamespaceTanzupackageInstallGetInstallResponse {
	return o.Payload
}

func (o *InstallResourceServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterNamespaceTanzupackageInstallGetInstallResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstallResourceServiceGetDefault creates a InstallResourceServiceGetDefault with default headers values
func NewInstallResourceServiceGetDefault(code int) *InstallResourceServiceGetDefault {
	return &InstallResourceServiceGetDefault{
		_statusCode: code,
	}
}

/*
InstallResourceServiceGetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type InstallResourceServiceGetDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the install resource service get default response
func (o *InstallResourceServiceGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this install resource service get default response has a 2xx status code
func (o *InstallResourceServiceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this install resource service get default response has a 3xx status code
func (o *InstallResourceServiceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this install resource service get default response has a 4xx status code
func (o *InstallResourceServiceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this install resource service get default response has a 5xx status code
func (o *InstallResourceServiceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this install resource service get default response a status code equal to that given
func (o *InstallResourceServiceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *InstallResourceServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/tanzupackage/installs/{fullName.name}][%d] InstallResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *InstallResourceServiceGetDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/tanzupackage/installs/{fullName.name}][%d] InstallResourceService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *InstallResourceServiceGetDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *InstallResourceServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
