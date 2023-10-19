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

// InstallResourceServiceDeleteReader is a Reader for the InstallResourceServiceDelete structure.
type InstallResourceServiceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstallResourceServiceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInstallResourceServiceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewInstallResourceServiceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInstallResourceServiceDeleteOK creates a InstallResourceServiceDeleteOK with default headers values
func NewInstallResourceServiceDeleteOK() *InstallResourceServiceDeleteOK {
	return &InstallResourceServiceDeleteOK{}
}

/*
InstallResourceServiceDeleteOK describes a response with status code 200, with default header values.

A successful response.
*/
type InstallResourceServiceDeleteOK struct {
	Payload *models.ClusterNamespaceTanzupackageInstallDeleteInstallResponse
}

// IsSuccess returns true when this install resource service delete o k response has a 2xx status code
func (o *InstallResourceServiceDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this install resource service delete o k response has a 3xx status code
func (o *InstallResourceServiceDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this install resource service delete o k response has a 4xx status code
func (o *InstallResourceServiceDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this install resource service delete o k response has a 5xx status code
func (o *InstallResourceServiceDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this install resource service delete o k response a status code equal to that given
func (o *InstallResourceServiceDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *InstallResourceServiceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/tanzupackage/installs/{fullName.name}][%d] installResourceServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *InstallResourceServiceDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/tanzupackage/installs/{fullName.name}][%d] installResourceServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *InstallResourceServiceDeleteOK) GetPayload() *models.ClusterNamespaceTanzupackageInstallDeleteInstallResponse {
	return o.Payload
}

func (o *InstallResourceServiceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterNamespaceTanzupackageInstallDeleteInstallResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstallResourceServiceDeleteDefault creates a InstallResourceServiceDeleteDefault with default headers values
func NewInstallResourceServiceDeleteDefault(code int) *InstallResourceServiceDeleteDefault {
	return &InstallResourceServiceDeleteDefault{
		_statusCode: code,
	}
}

/*
InstallResourceServiceDeleteDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type InstallResourceServiceDeleteDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the install resource service delete default response
func (o *InstallResourceServiceDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this install resource service delete default response has a 2xx status code
func (o *InstallResourceServiceDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this install resource service delete default response has a 3xx status code
func (o *InstallResourceServiceDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this install resource service delete default response has a 4xx status code
func (o *InstallResourceServiceDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this install resource service delete default response has a 5xx status code
func (o *InstallResourceServiceDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this install resource service delete default response a status code equal to that given
func (o *InstallResourceServiceDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *InstallResourceServiceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/tanzupackage/installs/{fullName.name}][%d] InstallResourceService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *InstallResourceServiceDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/tanzupackage/installs/{fullName.name}][%d] InstallResourceService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *InstallResourceServiceDeleteDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *InstallResourceServiceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
