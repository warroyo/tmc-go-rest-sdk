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

// ProvisionerResourceServiceDeleteReader is a Reader for the ProvisionerResourceServiceDelete structure.
type ProvisionerResourceServiceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProvisionerResourceServiceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProvisionerResourceServiceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProvisionerResourceServiceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProvisionerResourceServiceDeleteOK creates a ProvisionerResourceServiceDeleteOK with default headers values
func NewProvisionerResourceServiceDeleteOK() *ProvisionerResourceServiceDeleteOK {
	return &ProvisionerResourceServiceDeleteOK{}
}

/*
ProvisionerResourceServiceDeleteOK describes a response with status code 200, with default header values.

A successful response.
*/
type ProvisionerResourceServiceDeleteOK struct {
	Payload *models.ManagementclusterProvisionerDeleteProvisionerResponse
}

// IsSuccess returns true when this provisioner resource service delete o k response has a 2xx status code
func (o *ProvisionerResourceServiceDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this provisioner resource service delete o k response has a 3xx status code
func (o *ProvisionerResourceServiceDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this provisioner resource service delete o k response has a 4xx status code
func (o *ProvisionerResourceServiceDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this provisioner resource service delete o k response has a 5xx status code
func (o *ProvisionerResourceServiceDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this provisioner resource service delete o k response a status code equal to that given
func (o *ProvisionerResourceServiceDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProvisionerResourceServiceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1alpha1/managementclusters/{fullName.managementClusterName}/provisioners/{fullName.name}][%d] provisionerResourceServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *ProvisionerResourceServiceDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /v1alpha1/managementclusters/{fullName.managementClusterName}/provisioners/{fullName.name}][%d] provisionerResourceServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *ProvisionerResourceServiceDeleteOK) GetPayload() *models.ManagementclusterProvisionerDeleteProvisionerResponse {
	return o.Payload
}

func (o *ProvisionerResourceServiceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagementclusterProvisionerDeleteProvisionerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProvisionerResourceServiceDeleteDefault creates a ProvisionerResourceServiceDeleteDefault with default headers values
func NewProvisionerResourceServiceDeleteDefault(code int) *ProvisionerResourceServiceDeleteDefault {
	return &ProvisionerResourceServiceDeleteDefault{
		_statusCode: code,
	}
}

/*
ProvisionerResourceServiceDeleteDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ProvisionerResourceServiceDeleteDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the provisioner resource service delete default response
func (o *ProvisionerResourceServiceDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this provisioner resource service delete default response has a 2xx status code
func (o *ProvisionerResourceServiceDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this provisioner resource service delete default response has a 3xx status code
func (o *ProvisionerResourceServiceDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this provisioner resource service delete default response has a 4xx status code
func (o *ProvisionerResourceServiceDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this provisioner resource service delete default response has a 5xx status code
func (o *ProvisionerResourceServiceDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this provisioner resource service delete default response a status code equal to that given
func (o *ProvisionerResourceServiceDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ProvisionerResourceServiceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1alpha1/managementclusters/{fullName.managementClusterName}/provisioners/{fullName.name}][%d] ProvisionerResourceService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *ProvisionerResourceServiceDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /v1alpha1/managementclusters/{fullName.managementClusterName}/provisioners/{fullName.name}][%d] ProvisionerResourceService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *ProvisionerResourceServiceDeleteDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ProvisionerResourceServiceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
