// Code generated by go-swagger; DO NOT EDIT.

package core_cli_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// CoreCliServiceGetReader is a Reader for the CoreCliServiceGet structure.
type CoreCliServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CoreCliServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCoreCliServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCoreCliServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCoreCliServiceGetOK creates a CoreCliServiceGetOK with default headers values
func NewCoreCliServiceGetOK() *CoreCliServiceGetOK {
	return &CoreCliServiceGetOK{}
}

/*
CoreCliServiceGetOK describes a response with status code 200, with default header values.

A successful response.
*/
type CoreCliServiceGetOK struct {
	Payload *models.SystemBinariesCorecliGetCoreCliResponse
}

// IsSuccess returns true when this core cli service get o k response has a 2xx status code
func (o *CoreCliServiceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this core cli service get o k response has a 3xx status code
func (o *CoreCliServiceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this core cli service get o k response has a 4xx status code
func (o *CoreCliServiceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this core cli service get o k response has a 5xx status code
func (o *CoreCliServiceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this core cli service get o k response a status code equal to that given
func (o *CoreCliServiceGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *CoreCliServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/system/binaries/corecli][%d] coreCliServiceGetOK  %+v", 200, o.Payload)
}

func (o *CoreCliServiceGetOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/system/binaries/corecli][%d] coreCliServiceGetOK  %+v", 200, o.Payload)
}

func (o *CoreCliServiceGetOK) GetPayload() *models.SystemBinariesCorecliGetCoreCliResponse {
	return o.Payload
}

func (o *CoreCliServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemBinariesCorecliGetCoreCliResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCoreCliServiceGetDefault creates a CoreCliServiceGetDefault with default headers values
func NewCoreCliServiceGetDefault(code int) *CoreCliServiceGetDefault {
	return &CoreCliServiceGetDefault{
		_statusCode: code,
	}
}

/*
CoreCliServiceGetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CoreCliServiceGetDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the core cli service get default response
func (o *CoreCliServiceGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this core cli service get default response has a 2xx status code
func (o *CoreCliServiceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this core cli service get default response has a 3xx status code
func (o *CoreCliServiceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this core cli service get default response has a 4xx status code
func (o *CoreCliServiceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this core cli service get default response has a 5xx status code
func (o *CoreCliServiceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this core cli service get default response a status code equal to that given
func (o *CoreCliServiceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CoreCliServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/system/binaries/corecli][%d] CoreCliService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *CoreCliServiceGetDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/system/binaries/corecli][%d] CoreCliService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *CoreCliServiceGetDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *CoreCliServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
