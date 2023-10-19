// Code generated by go-swagger; DO NOT EDIT.

package backup_location_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// BackupLocationResourceServiceListReader is a Reader for the BackupLocationResourceServiceList structure.
type BackupLocationResourceServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupLocationResourceServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupLocationResourceServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBackupLocationResourceServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBackupLocationResourceServiceListOK creates a BackupLocationResourceServiceListOK with default headers values
func NewBackupLocationResourceServiceListOK() *BackupLocationResourceServiceListOK {
	return &BackupLocationResourceServiceListOK{}
}

/*
BackupLocationResourceServiceListOK describes a response with status code 200, with default header values.

A successful response.
*/
type BackupLocationResourceServiceListOK struct {
	Payload *models.DataprotectionProviderBackuplocationListBackupLocationsResponse
}

// IsSuccess returns true when this backup location resource service list o k response has a 2xx status code
func (o *BackupLocationResourceServiceListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backup location resource service list o k response has a 3xx status code
func (o *BackupLocationResourceServiceListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup location resource service list o k response has a 4xx status code
func (o *BackupLocationResourceServiceListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup location resource service list o k response has a 5xx status code
func (o *BackupLocationResourceServiceListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this backup location resource service list o k response a status code equal to that given
func (o *BackupLocationResourceServiceListOK) IsCode(code int) bool {
	return code == 200
}

func (o *BackupLocationResourceServiceListOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/dataprotection/providers/{searchScope.providerName}/backuplocations][%d] backupLocationResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *BackupLocationResourceServiceListOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/dataprotection/providers/{searchScope.providerName}/backuplocations][%d] backupLocationResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *BackupLocationResourceServiceListOK) GetPayload() *models.DataprotectionProviderBackuplocationListBackupLocationsResponse {
	return o.Payload
}

func (o *BackupLocationResourceServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataprotectionProviderBackuplocationListBackupLocationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupLocationResourceServiceListDefault creates a BackupLocationResourceServiceListDefault with default headers values
func NewBackupLocationResourceServiceListDefault(code int) *BackupLocationResourceServiceListDefault {
	return &BackupLocationResourceServiceListDefault{
		_statusCode: code,
	}
}

/*
BackupLocationResourceServiceListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BackupLocationResourceServiceListDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the backup location resource service list default response
func (o *BackupLocationResourceServiceListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this backup location resource service list default response has a 2xx status code
func (o *BackupLocationResourceServiceListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this backup location resource service list default response has a 3xx status code
func (o *BackupLocationResourceServiceListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this backup location resource service list default response has a 4xx status code
func (o *BackupLocationResourceServiceListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this backup location resource service list default response has a 5xx status code
func (o *BackupLocationResourceServiceListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this backup location resource service list default response a status code equal to that given
func (o *BackupLocationResourceServiceListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *BackupLocationResourceServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/dataprotection/providers/{searchScope.providerName}/backuplocations][%d] BackupLocationResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *BackupLocationResourceServiceListDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/dataprotection/providers/{searchScope.providerName}/backuplocations][%d] BackupLocationResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *BackupLocationResourceServiceListDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *BackupLocationResourceServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}