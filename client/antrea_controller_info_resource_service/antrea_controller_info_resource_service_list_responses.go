// Code generated by go-swagger; DO NOT EDIT.

package antrea_controller_info_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// AntreaControllerInfoResourceServiceListReader is a Reader for the AntreaControllerInfoResourceServiceList structure.
type AntreaControllerInfoResourceServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AntreaControllerInfoResourceServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAntreaControllerInfoResourceServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAntreaControllerInfoResourceServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAntreaControllerInfoResourceServiceListOK creates a AntreaControllerInfoResourceServiceListOK with default headers values
func NewAntreaControllerInfoResourceServiceListOK() *AntreaControllerInfoResourceServiceListOK {
	return &AntreaControllerInfoResourceServiceListOK{}
}

/*
AntreaControllerInfoResourceServiceListOK describes a response with status code 200, with default header values.

A successful response.
*/
type AntreaControllerInfoResourceServiceListOK struct {
	Payload *models.ClusterNetworkAntreaAntreacontrollerinfoListAntreaControllerInfosResponse
}

// IsSuccess returns true when this antrea controller info resource service list o k response has a 2xx status code
func (o *AntreaControllerInfoResourceServiceListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this antrea controller info resource service list o k response has a 3xx status code
func (o *AntreaControllerInfoResourceServiceListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this antrea controller info resource service list o k response has a 4xx status code
func (o *AntreaControllerInfoResourceServiceListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this antrea controller info resource service list o k response has a 5xx status code
func (o *AntreaControllerInfoResourceServiceListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this antrea controller info resource service list o k response a status code equal to that given
func (o *AntreaControllerInfoResourceServiceListOK) IsCode(code int) bool {
	return code == 200
}

func (o *AntreaControllerInfoResourceServiceListOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{searchScope.clusterName}/network/antrea/antreacontrollerinfos][%d] antreaControllerInfoResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *AntreaControllerInfoResourceServiceListOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{searchScope.clusterName}/network/antrea/antreacontrollerinfos][%d] antreaControllerInfoResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *AntreaControllerInfoResourceServiceListOK) GetPayload() *models.ClusterNetworkAntreaAntreacontrollerinfoListAntreaControllerInfosResponse {
	return o.Payload
}

func (o *AntreaControllerInfoResourceServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterNetworkAntreaAntreacontrollerinfoListAntreaControllerInfosResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAntreaControllerInfoResourceServiceListDefault creates a AntreaControllerInfoResourceServiceListDefault with default headers values
func NewAntreaControllerInfoResourceServiceListDefault(code int) *AntreaControllerInfoResourceServiceListDefault {
	return &AntreaControllerInfoResourceServiceListDefault{
		_statusCode: code,
	}
}

/*
AntreaControllerInfoResourceServiceListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AntreaControllerInfoResourceServiceListDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the antrea controller info resource service list default response
func (o *AntreaControllerInfoResourceServiceListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this antrea controller info resource service list default response has a 2xx status code
func (o *AntreaControllerInfoResourceServiceListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this antrea controller info resource service list default response has a 3xx status code
func (o *AntreaControllerInfoResourceServiceListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this antrea controller info resource service list default response has a 4xx status code
func (o *AntreaControllerInfoResourceServiceListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this antrea controller info resource service list default response has a 5xx status code
func (o *AntreaControllerInfoResourceServiceListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this antrea controller info resource service list default response a status code equal to that given
func (o *AntreaControllerInfoResourceServiceListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AntreaControllerInfoResourceServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{searchScope.clusterName}/network/antrea/antreacontrollerinfos][%d] AntreaControllerInfoResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *AntreaControllerInfoResourceServiceListDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{searchScope.clusterName}/network/antrea/antreacontrollerinfos][%d] AntreaControllerInfoResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *AntreaControllerInfoResourceServiceListDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *AntreaControllerInfoResourceServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}