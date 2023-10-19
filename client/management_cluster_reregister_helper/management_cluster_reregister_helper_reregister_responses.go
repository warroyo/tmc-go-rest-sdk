// Code generated by go-swagger; DO NOT EDIT.

package management_cluster_reregister_helper

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// ManagementClusterReregisterHelperReregisterReader is a Reader for the ManagementClusterReregisterHelperReregister structure.
type ManagementClusterReregisterHelperReregisterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ManagementClusterReregisterHelperReregisterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewManagementClusterReregisterHelperReregisterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewManagementClusterReregisterHelperReregisterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewManagementClusterReregisterHelperReregisterOK creates a ManagementClusterReregisterHelperReregisterOK with default headers values
func NewManagementClusterReregisterHelperReregisterOK() *ManagementClusterReregisterHelperReregisterOK {
	return &ManagementClusterReregisterHelperReregisterOK{}
}

/*
ManagementClusterReregisterHelperReregisterOK describes a response with status code 200, with default header values.

A successful response.
*/
type ManagementClusterReregisterHelperReregisterOK struct {
	Payload *models.ManagementclusterManagementClusterReregisterResponse
}

// IsSuccess returns true when this management cluster reregister helper reregister o k response has a 2xx status code
func (o *ManagementClusterReregisterHelperReregisterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this management cluster reregister helper reregister o k response has a 3xx status code
func (o *ManagementClusterReregisterHelperReregisterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this management cluster reregister helper reregister o k response has a 4xx status code
func (o *ManagementClusterReregisterHelperReregisterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this management cluster reregister helper reregister o k response has a 5xx status code
func (o *ManagementClusterReregisterHelperReregisterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this management cluster reregister helper reregister o k response a status code equal to that given
func (o *ManagementClusterReregisterHelperReregisterOK) IsCode(code int) bool {
	return code == 200
}

func (o *ManagementClusterReregisterHelperReregisterOK) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/managementclusters:reregister/{fullName.name}][%d] managementClusterReregisterHelperReregisterOK  %+v", 200, o.Payload)
}

func (o *ManagementClusterReregisterHelperReregisterOK) String() string {
	return fmt.Sprintf("[POST /v1alpha1/managementclusters:reregister/{fullName.name}][%d] managementClusterReregisterHelperReregisterOK  %+v", 200, o.Payload)
}

func (o *ManagementClusterReregisterHelperReregisterOK) GetPayload() *models.ManagementclusterManagementClusterReregisterResponse {
	return o.Payload
}

func (o *ManagementClusterReregisterHelperReregisterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagementclusterManagementClusterReregisterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManagementClusterReregisterHelperReregisterDefault creates a ManagementClusterReregisterHelperReregisterDefault with default headers values
func NewManagementClusterReregisterHelperReregisterDefault(code int) *ManagementClusterReregisterHelperReregisterDefault {
	return &ManagementClusterReregisterHelperReregisterDefault{
		_statusCode: code,
	}
}

/*
ManagementClusterReregisterHelperReregisterDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ManagementClusterReregisterHelperReregisterDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the management cluster reregister helper reregister default response
func (o *ManagementClusterReregisterHelperReregisterDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this management cluster reregister helper reregister default response has a 2xx status code
func (o *ManagementClusterReregisterHelperReregisterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this management cluster reregister helper reregister default response has a 3xx status code
func (o *ManagementClusterReregisterHelperReregisterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this management cluster reregister helper reregister default response has a 4xx status code
func (o *ManagementClusterReregisterHelperReregisterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this management cluster reregister helper reregister default response has a 5xx status code
func (o *ManagementClusterReregisterHelperReregisterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this management cluster reregister helper reregister default response a status code equal to that given
func (o *ManagementClusterReregisterHelperReregisterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ManagementClusterReregisterHelperReregisterDefault) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/managementclusters:reregister/{fullName.name}][%d] ManagementClusterReregisterHelper_Reregister default  %+v", o._statusCode, o.Payload)
}

func (o *ManagementClusterReregisterHelperReregisterDefault) String() string {
	return fmt.Sprintf("[POST /v1alpha1/managementclusters:reregister/{fullName.name}][%d] ManagementClusterReregisterHelper_Reregister default  %+v", o._statusCode, o.Payload)
}

func (o *ManagementClusterReregisterHelperReregisterDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ManagementClusterReregisterHelperReregisterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
