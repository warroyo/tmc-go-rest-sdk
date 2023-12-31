// Code generated by go-swagger; DO NOT EDIT.

package usage_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// UsageServiceGetClusterLevelUsageReader is a Reader for the UsageServiceGetClusterLevelUsage structure.
type UsageServiceGetClusterLevelUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsageServiceGetClusterLevelUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsageServiceGetClusterLevelUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsageServiceGetClusterLevelUsageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsageServiceGetClusterLevelUsageOK creates a UsageServiceGetClusterLevelUsageOK with default headers values
func NewUsageServiceGetClusterLevelUsageOK() *UsageServiceGetClusterLevelUsageOK {
	return &UsageServiceGetClusterLevelUsageOK{}
}

/*
UsageServiceGetClusterLevelUsageOK describes a response with status code 200, with default header values.

A successful response.
*/
type UsageServiceGetClusterLevelUsageOK struct {
	Payload *models.MeteringUsageGetClusterUsageResponse
}

// IsSuccess returns true when this usage service get cluster level usage o k response has a 2xx status code
func (o *UsageServiceGetClusterLevelUsageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this usage service get cluster level usage o k response has a 3xx status code
func (o *UsageServiceGetClusterLevelUsageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this usage service get cluster level usage o k response has a 4xx status code
func (o *UsageServiceGetClusterLevelUsageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this usage service get cluster level usage o k response has a 5xx status code
func (o *UsageServiceGetClusterLevelUsageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this usage service get cluster level usage o k response a status code equal to that given
func (o *UsageServiceGetClusterLevelUsageOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsageServiceGetClusterLevelUsageOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/metering/usage/cluster][%d] usageServiceGetClusterLevelUsageOK  %+v", 200, o.Payload)
}

func (o *UsageServiceGetClusterLevelUsageOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/metering/usage/cluster][%d] usageServiceGetClusterLevelUsageOK  %+v", 200, o.Payload)
}

func (o *UsageServiceGetClusterLevelUsageOK) GetPayload() *models.MeteringUsageGetClusterUsageResponse {
	return o.Payload
}

func (o *UsageServiceGetClusterLevelUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MeteringUsageGetClusterUsageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsageServiceGetClusterLevelUsageDefault creates a UsageServiceGetClusterLevelUsageDefault with default headers values
func NewUsageServiceGetClusterLevelUsageDefault(code int) *UsageServiceGetClusterLevelUsageDefault {
	return &UsageServiceGetClusterLevelUsageDefault{
		_statusCode: code,
	}
}

/*
UsageServiceGetClusterLevelUsageDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UsageServiceGetClusterLevelUsageDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the usage service get cluster level usage default response
func (o *UsageServiceGetClusterLevelUsageDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this usage service get cluster level usage default response has a 2xx status code
func (o *UsageServiceGetClusterLevelUsageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this usage service get cluster level usage default response has a 3xx status code
func (o *UsageServiceGetClusterLevelUsageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this usage service get cluster level usage default response has a 4xx status code
func (o *UsageServiceGetClusterLevelUsageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this usage service get cluster level usage default response has a 5xx status code
func (o *UsageServiceGetClusterLevelUsageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this usage service get cluster level usage default response a status code equal to that given
func (o *UsageServiceGetClusterLevelUsageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UsageServiceGetClusterLevelUsageDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/metering/usage/cluster][%d] UsageService_GetClusterLevelUsage default  %+v", o._statusCode, o.Payload)
}

func (o *UsageServiceGetClusterLevelUsageDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/metering/usage/cluster][%d] UsageService_GetClusterLevelUsage default  %+v", o._statusCode, o.Payload)
}

func (o *UsageServiceGetClusterLevelUsageDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *UsageServiceGetClusterLevelUsageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
