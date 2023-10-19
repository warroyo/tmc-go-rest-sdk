// Code generated by go-swagger; DO NOT EDIT.

package version_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// VersionResourceServiceListReader is a Reader for the VersionResourceServiceList structure.
type VersionResourceServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VersionResourceServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVersionResourceServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVersionResourceServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVersionResourceServiceListOK creates a VersionResourceServiceListOK with default headers values
func NewVersionResourceServiceListOK() *VersionResourceServiceListOK {
	return &VersionResourceServiceListOK{}
}

/*
VersionResourceServiceListOK describes a response with status code 200, with default header values.

A successful response.
*/
type VersionResourceServiceListOK struct {
	Payload *models.PolicyTypeRecipeVersionListVersionsResponse
}

// IsSuccess returns true when this version resource service list o k response has a 2xx status code
func (o *VersionResourceServiceListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this version resource service list o k response has a 3xx status code
func (o *VersionResourceServiceListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this version resource service list o k response has a 4xx status code
func (o *VersionResourceServiceListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this version resource service list o k response has a 5xx status code
func (o *VersionResourceServiceListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this version resource service list o k response a status code equal to that given
func (o *VersionResourceServiceListOK) IsCode(code int) bool {
	return code == 200
}

func (o *VersionResourceServiceListOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/policy/types/{searchScope.typeName}/recipes/{searchScope.recipeName}/versions][%d] versionResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *VersionResourceServiceListOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/policy/types/{searchScope.typeName}/recipes/{searchScope.recipeName}/versions][%d] versionResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *VersionResourceServiceListOK) GetPayload() *models.PolicyTypeRecipeVersionListVersionsResponse {
	return o.Payload
}

func (o *VersionResourceServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyTypeRecipeVersionListVersionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVersionResourceServiceListDefault creates a VersionResourceServiceListDefault with default headers values
func NewVersionResourceServiceListDefault(code int) *VersionResourceServiceListDefault {
	return &VersionResourceServiceListDefault{
		_statusCode: code,
	}
}

/*
VersionResourceServiceListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type VersionResourceServiceListDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the version resource service list default response
func (o *VersionResourceServiceListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this version resource service list default response has a 2xx status code
func (o *VersionResourceServiceListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this version resource service list default response has a 3xx status code
func (o *VersionResourceServiceListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this version resource service list default response has a 4xx status code
func (o *VersionResourceServiceListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this version resource service list default response has a 5xx status code
func (o *VersionResourceServiceListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this version resource service list default response a status code equal to that given
func (o *VersionResourceServiceListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VersionResourceServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/policy/types/{searchScope.typeName}/recipes/{searchScope.recipeName}/versions][%d] VersionResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *VersionResourceServiceListDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/policy/types/{searchScope.typeName}/recipes/{searchScope.recipeName}/versions][%d] VersionResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *VersionResourceServiceListDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *VersionResourceServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}