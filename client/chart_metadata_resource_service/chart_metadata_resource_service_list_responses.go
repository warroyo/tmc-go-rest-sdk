// Code generated by go-swagger; DO NOT EDIT.

package chart_metadata_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// ChartMetadataResourceServiceListReader is a Reader for the ChartMetadataResourceServiceList structure.
type ChartMetadataResourceServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChartMetadataResourceServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChartMetadataResourceServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChartMetadataResourceServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChartMetadataResourceServiceListOK creates a ChartMetadataResourceServiceListOK with default headers values
func NewChartMetadataResourceServiceListOK() *ChartMetadataResourceServiceListOK {
	return &ChartMetadataResourceServiceListOK{}
}

/*
ChartMetadataResourceServiceListOK describes a response with status code 200, with default header values.

A successful response.
*/
type ChartMetadataResourceServiceListOK struct {
	Payload *models.OrganizationFluxcdHelmRepositoryChartmetadataListChartMetadatasResponse
}

// IsSuccess returns true when this chart metadata resource service list o k response has a 2xx status code
func (o *ChartMetadataResourceServiceListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this chart metadata resource service list o k response has a 3xx status code
func (o *ChartMetadataResourceServiceListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this chart metadata resource service list o k response has a 4xx status code
func (o *ChartMetadataResourceServiceListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this chart metadata resource service list o k response has a 5xx status code
func (o *ChartMetadataResourceServiceListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this chart metadata resource service list o k response a status code equal to that given
func (o *ChartMetadataResourceServiceListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ChartMetadataResourceServiceListOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/organization/fluxcd/helm/repositories/{searchScope.repositoryName}/chartmetadatas][%d] chartMetadataResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *ChartMetadataResourceServiceListOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/organization/fluxcd/helm/repositories/{searchScope.repositoryName}/chartmetadatas][%d] chartMetadataResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *ChartMetadataResourceServiceListOK) GetPayload() *models.OrganizationFluxcdHelmRepositoryChartmetadataListChartMetadatasResponse {
	return o.Payload
}

func (o *ChartMetadataResourceServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationFluxcdHelmRepositoryChartmetadataListChartMetadatasResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChartMetadataResourceServiceListDefault creates a ChartMetadataResourceServiceListDefault with default headers values
func NewChartMetadataResourceServiceListDefault(code int) *ChartMetadataResourceServiceListDefault {
	return &ChartMetadataResourceServiceListDefault{
		_statusCode: code,
	}
}

/*
ChartMetadataResourceServiceListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ChartMetadataResourceServiceListDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the chart metadata resource service list default response
func (o *ChartMetadataResourceServiceListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this chart metadata resource service list default response has a 2xx status code
func (o *ChartMetadataResourceServiceListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this chart metadata resource service list default response has a 3xx status code
func (o *ChartMetadataResourceServiceListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this chart metadata resource service list default response has a 4xx status code
func (o *ChartMetadataResourceServiceListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this chart metadata resource service list default response has a 5xx status code
func (o *ChartMetadataResourceServiceListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this chart metadata resource service list default response a status code equal to that given
func (o *ChartMetadataResourceServiceListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ChartMetadataResourceServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/organization/fluxcd/helm/repositories/{searchScope.repositoryName}/chartmetadatas][%d] ChartMetadataResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *ChartMetadataResourceServiceListDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/organization/fluxcd/helm/repositories/{searchScope.repositoryName}/chartmetadatas][%d] ChartMetadataResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *ChartMetadataResourceServiceListDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ChartMetadataResourceServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
