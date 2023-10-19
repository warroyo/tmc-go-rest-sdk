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

// InstallResourceServiceUpdateReader is a Reader for the InstallResourceServiceUpdate structure.
type InstallResourceServiceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstallResourceServiceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInstallResourceServiceUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewInstallResourceServiceUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInstallResourceServiceUpdateOK creates a InstallResourceServiceUpdateOK with default headers values
func NewInstallResourceServiceUpdateOK() *InstallResourceServiceUpdateOK {
	return &InstallResourceServiceUpdateOK{}
}

/*
InstallResourceServiceUpdateOK describes a response with status code 200, with default header values.

A successful response.
*/
type InstallResourceServiceUpdateOK struct {
	Payload *models.ClusterNamespaceTanzupackageInstallUpdateInstallResponse
}

// IsSuccess returns true when this install resource service update o k response has a 2xx status code
func (o *InstallResourceServiceUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this install resource service update o k response has a 3xx status code
func (o *InstallResourceServiceUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this install resource service update o k response has a 4xx status code
func (o *InstallResourceServiceUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this install resource service update o k response has a 5xx status code
func (o *InstallResourceServiceUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this install resource service update o k response a status code equal to that given
func (o *InstallResourceServiceUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *InstallResourceServiceUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/clusters/{install.fullName.clusterName}/namespaces/{install.fullName.namespaceName}/tanzupackage/installs/{install.fullName.name}][%d] installResourceServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *InstallResourceServiceUpdateOK) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/clusters/{install.fullName.clusterName}/namespaces/{install.fullName.namespaceName}/tanzupackage/installs/{install.fullName.name}][%d] installResourceServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *InstallResourceServiceUpdateOK) GetPayload() *models.ClusterNamespaceTanzupackageInstallUpdateInstallResponse {
	return o.Payload
}

func (o *InstallResourceServiceUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterNamespaceTanzupackageInstallUpdateInstallResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstallResourceServiceUpdateDefault creates a InstallResourceServiceUpdateDefault with default headers values
func NewInstallResourceServiceUpdateDefault(code int) *InstallResourceServiceUpdateDefault {
	return &InstallResourceServiceUpdateDefault{
		_statusCode: code,
	}
}

/*
InstallResourceServiceUpdateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type InstallResourceServiceUpdateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the install resource service update default response
func (o *InstallResourceServiceUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this install resource service update default response has a 2xx status code
func (o *InstallResourceServiceUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this install resource service update default response has a 3xx status code
func (o *InstallResourceServiceUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this install resource service update default response has a 4xx status code
func (o *InstallResourceServiceUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this install resource service update default response has a 5xx status code
func (o *InstallResourceServiceUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this install resource service update default response a status code equal to that given
func (o *InstallResourceServiceUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *InstallResourceServiceUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/clusters/{install.fullName.clusterName}/namespaces/{install.fullName.namespaceName}/tanzupackage/installs/{install.fullName.name}][%d] InstallResourceService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *InstallResourceServiceUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/clusters/{install.fullName.clusterName}/namespaces/{install.fullName.namespaceName}/tanzupackage/installs/{install.fullName.name}][%d] InstallResourceService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *InstallResourceServiceUpdateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *InstallResourceServiceUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}