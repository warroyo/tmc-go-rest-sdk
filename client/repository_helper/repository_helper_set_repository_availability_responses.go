// Code generated by go-swagger; DO NOT EDIT.

package repository_helper

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// RepositoryHelperSetRepositoryAvailabilityReader is a Reader for the RepositoryHelperSetRepositoryAvailability structure.
type RepositoryHelperSetRepositoryAvailabilityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepositoryHelperSetRepositoryAvailabilityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepositoryHelperSetRepositoryAvailabilityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRepositoryHelperSetRepositoryAvailabilityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRepositoryHelperSetRepositoryAvailabilityOK creates a RepositoryHelperSetRepositoryAvailabilityOK with default headers values
func NewRepositoryHelperSetRepositoryAvailabilityOK() *RepositoryHelperSetRepositoryAvailabilityOK {
	return &RepositoryHelperSetRepositoryAvailabilityOK{}
}

/*
RepositoryHelperSetRepositoryAvailabilityOK describes a response with status code 200, with default header values.

A successful response.
*/
type RepositoryHelperSetRepositoryAvailabilityOK struct {
	Payload *models.ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityResponse
}

// IsSuccess returns true when this repository helper set repository availability o k response has a 2xx status code
func (o *RepositoryHelperSetRepositoryAvailabilityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repository helper set repository availability o k response has a 3xx status code
func (o *RepositoryHelperSetRepositoryAvailabilityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repository helper set repository availability o k response has a 4xx status code
func (o *RepositoryHelperSetRepositoryAvailabilityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repository helper set repository availability o k response has a 5xx status code
func (o *RepositoryHelperSetRepositoryAvailabilityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repository helper set repository availability o k response a status code equal to that given
func (o *RepositoryHelperSetRepositoryAvailabilityOK) IsCode(code int) bool {
	return code == 200
}

func (o *RepositoryHelperSetRepositoryAvailabilityOK) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/tanzupackage/repositories:setavailability/{fullName.name}][%d] repositoryHelperSetRepositoryAvailabilityOK  %+v", 200, o.Payload)
}

func (o *RepositoryHelperSetRepositoryAvailabilityOK) String() string {
	return fmt.Sprintf("[POST /v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/tanzupackage/repositories:setavailability/{fullName.name}][%d] repositoryHelperSetRepositoryAvailabilityOK  %+v", 200, o.Payload)
}

func (o *RepositoryHelperSetRepositoryAvailabilityOK) GetPayload() *models.ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityResponse {
	return o.Payload
}

func (o *RepositoryHelperSetRepositoryAvailabilityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterNamespaceTanzupackageRepositorySetRepositoryAvailabilityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepositoryHelperSetRepositoryAvailabilityDefault creates a RepositoryHelperSetRepositoryAvailabilityDefault with default headers values
func NewRepositoryHelperSetRepositoryAvailabilityDefault(code int) *RepositoryHelperSetRepositoryAvailabilityDefault {
	return &RepositoryHelperSetRepositoryAvailabilityDefault{
		_statusCode: code,
	}
}

/*
RepositoryHelperSetRepositoryAvailabilityDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RepositoryHelperSetRepositoryAvailabilityDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the repository helper set repository availability default response
func (o *RepositoryHelperSetRepositoryAvailabilityDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this repository helper set repository availability default response has a 2xx status code
func (o *RepositoryHelperSetRepositoryAvailabilityDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this repository helper set repository availability default response has a 3xx status code
func (o *RepositoryHelperSetRepositoryAvailabilityDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this repository helper set repository availability default response has a 4xx status code
func (o *RepositoryHelperSetRepositoryAvailabilityDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this repository helper set repository availability default response has a 5xx status code
func (o *RepositoryHelperSetRepositoryAvailabilityDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this repository helper set repository availability default response a status code equal to that given
func (o *RepositoryHelperSetRepositoryAvailabilityDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *RepositoryHelperSetRepositoryAvailabilityDefault) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/tanzupackage/repositories:setavailability/{fullName.name}][%d] RepositoryHelper_SetRepositoryAvailability default  %+v", o._statusCode, o.Payload)
}

func (o *RepositoryHelperSetRepositoryAvailabilityDefault) String() string {
	return fmt.Sprintf("[POST /v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/tanzupackage/repositories:setavailability/{fullName.name}][%d] RepositoryHelper_SetRepositoryAvailability default  %+v", o._statusCode, o.Payload)
}

func (o *RepositoryHelperSetRepositoryAvailabilityDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *RepositoryHelperSetRepositoryAvailabilityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
