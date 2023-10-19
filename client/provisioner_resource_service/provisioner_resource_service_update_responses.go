// Code generated by go-swagger; DO NOT EDIT.

package provisioner_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// ProvisionerResourceServiceUpdateReader is a Reader for the ProvisionerResourceServiceUpdate structure.
type ProvisionerResourceServiceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProvisionerResourceServiceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProvisionerResourceServiceUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProvisionerResourceServiceUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProvisionerResourceServiceUpdateOK creates a ProvisionerResourceServiceUpdateOK with default headers values
func NewProvisionerResourceServiceUpdateOK() *ProvisionerResourceServiceUpdateOK {
	return &ProvisionerResourceServiceUpdateOK{}
}

/*
ProvisionerResourceServiceUpdateOK describes a response with status code 200, with default header values.

A successful response.
*/
type ProvisionerResourceServiceUpdateOK struct {
	Payload *models.ManagementclusterProvisionerUpdateProvisionerResponse
}

// IsSuccess returns true when this provisioner resource service update o k response has a 2xx status code
func (o *ProvisionerResourceServiceUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this provisioner resource service update o k response has a 3xx status code
func (o *ProvisionerResourceServiceUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this provisioner resource service update o k response has a 4xx status code
func (o *ProvisionerResourceServiceUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this provisioner resource service update o k response has a 5xx status code
func (o *ProvisionerResourceServiceUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this provisioner resource service update o k response a status code equal to that given
func (o *ProvisionerResourceServiceUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProvisionerResourceServiceUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/managementclusters/{provisioner.fullName.managementClusterName}/provisioners/{provisioner.fullName.name}][%d] provisionerResourceServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *ProvisionerResourceServiceUpdateOK) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/managementclusters/{provisioner.fullName.managementClusterName}/provisioners/{provisioner.fullName.name}][%d] provisionerResourceServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *ProvisionerResourceServiceUpdateOK) GetPayload() *models.ManagementclusterProvisionerUpdateProvisionerResponse {
	return o.Payload
}

func (o *ProvisionerResourceServiceUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagementclusterProvisionerUpdateProvisionerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProvisionerResourceServiceUpdateDefault creates a ProvisionerResourceServiceUpdateDefault with default headers values
func NewProvisionerResourceServiceUpdateDefault(code int) *ProvisionerResourceServiceUpdateDefault {
	return &ProvisionerResourceServiceUpdateDefault{
		_statusCode: code,
	}
}

/*
ProvisionerResourceServiceUpdateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ProvisionerResourceServiceUpdateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the provisioner resource service update default response
func (o *ProvisionerResourceServiceUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this provisioner resource service update default response has a 2xx status code
func (o *ProvisionerResourceServiceUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this provisioner resource service update default response has a 3xx status code
func (o *ProvisionerResourceServiceUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this provisioner resource service update default response has a 4xx status code
func (o *ProvisionerResourceServiceUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this provisioner resource service update default response has a 5xx status code
func (o *ProvisionerResourceServiceUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this provisioner resource service update default response a status code equal to that given
func (o *ProvisionerResourceServiceUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ProvisionerResourceServiceUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/managementclusters/{provisioner.fullName.managementClusterName}/provisioners/{provisioner.fullName.name}][%d] ProvisionerResourceService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *ProvisionerResourceServiceUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/managementclusters/{provisioner.fullName.managementClusterName}/provisioners/{provisioner.fullName.name}][%d] ProvisionerResourceService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *ProvisionerResourceServiceUpdateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ProvisionerResourceServiceUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
