// Code generated by go-swagger; DO NOT EDIT.

package audit_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// AuditResourceServiceCreateReader is a Reader for the AuditResourceServiceCreate structure.
type AuditResourceServiceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuditResourceServiceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuditResourceServiceCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAuditResourceServiceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuditResourceServiceCreateOK creates a AuditResourceServiceCreateOK with default headers values
func NewAuditResourceServiceCreateOK() *AuditResourceServiceCreateOK {
	return &AuditResourceServiceCreateOK{}
}

/*
AuditResourceServiceCreateOK describes a response with status code 200, with default header values.

A successful response.
*/
type AuditResourceServiceCreateOK struct {
	Payload *models.AuditCreateAuditResponse
}

// IsSuccess returns true when this audit resource service create o k response has a 2xx status code
func (o *AuditResourceServiceCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this audit resource service create o k response has a 3xx status code
func (o *AuditResourceServiceCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this audit resource service create o k response has a 4xx status code
func (o *AuditResourceServiceCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this audit resource service create o k response has a 5xx status code
func (o *AuditResourceServiceCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this audit resource service create o k response a status code equal to that given
func (o *AuditResourceServiceCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *AuditResourceServiceCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/audits][%d] auditResourceServiceCreateOK  %+v", 200, o.Payload)
}

func (o *AuditResourceServiceCreateOK) String() string {
	return fmt.Sprintf("[POST /v1alpha1/audits][%d] auditResourceServiceCreateOK  %+v", 200, o.Payload)
}

func (o *AuditResourceServiceCreateOK) GetPayload() *models.AuditCreateAuditResponse {
	return o.Payload
}

func (o *AuditResourceServiceCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuditCreateAuditResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuditResourceServiceCreateDefault creates a AuditResourceServiceCreateDefault with default headers values
func NewAuditResourceServiceCreateDefault(code int) *AuditResourceServiceCreateDefault {
	return &AuditResourceServiceCreateDefault{
		_statusCode: code,
	}
}

/*
AuditResourceServiceCreateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AuditResourceServiceCreateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the audit resource service create default response
func (o *AuditResourceServiceCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this audit resource service create default response has a 2xx status code
func (o *AuditResourceServiceCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this audit resource service create default response has a 3xx status code
func (o *AuditResourceServiceCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this audit resource service create default response has a 4xx status code
func (o *AuditResourceServiceCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this audit resource service create default response has a 5xx status code
func (o *AuditResourceServiceCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this audit resource service create default response a status code equal to that given
func (o *AuditResourceServiceCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AuditResourceServiceCreateDefault) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/audits][%d] AuditResourceService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *AuditResourceServiceCreateDefault) String() string {
	return fmt.Sprintf("[POST /v1alpha1/audits][%d] AuditResourceService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *AuditResourceServiceCreateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *AuditResourceServiceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}