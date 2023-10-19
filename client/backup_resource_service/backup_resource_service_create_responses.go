// Code generated by go-swagger; DO NOT EDIT.

package backup_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// BackupResourceServiceCreateReader is a Reader for the BackupResourceServiceCreate structure.
type BackupResourceServiceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupResourceServiceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupResourceServiceCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBackupResourceServiceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBackupResourceServiceCreateOK creates a BackupResourceServiceCreateOK with default headers values
func NewBackupResourceServiceCreateOK() *BackupResourceServiceCreateOK {
	return &BackupResourceServiceCreateOK{}
}

/*
BackupResourceServiceCreateOK describes a response with status code 200, with default header values.

A successful response.
*/
type BackupResourceServiceCreateOK struct {
	Payload *models.ClusterDataprotectionBackupCreateBackupResponse
}

// IsSuccess returns true when this backup resource service create o k response has a 2xx status code
func (o *BackupResourceServiceCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backup resource service create o k response has a 3xx status code
func (o *BackupResourceServiceCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup resource service create o k response has a 4xx status code
func (o *BackupResourceServiceCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup resource service create o k response has a 5xx status code
func (o *BackupResourceServiceCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this backup resource service create o k response a status code equal to that given
func (o *BackupResourceServiceCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *BackupResourceServiceCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/clusters/{backup.fullName.clusterName}/dataprotection/backups][%d] backupResourceServiceCreateOK  %+v", 200, o.Payload)
}

func (o *BackupResourceServiceCreateOK) String() string {
	return fmt.Sprintf("[POST /v1alpha1/clusters/{backup.fullName.clusterName}/dataprotection/backups][%d] backupResourceServiceCreateOK  %+v", 200, o.Payload)
}

func (o *BackupResourceServiceCreateOK) GetPayload() *models.ClusterDataprotectionBackupCreateBackupResponse {
	return o.Payload
}

func (o *BackupResourceServiceCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterDataprotectionBackupCreateBackupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupResourceServiceCreateDefault creates a BackupResourceServiceCreateDefault with default headers values
func NewBackupResourceServiceCreateDefault(code int) *BackupResourceServiceCreateDefault {
	return &BackupResourceServiceCreateDefault{
		_statusCode: code,
	}
}

/*
BackupResourceServiceCreateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BackupResourceServiceCreateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the backup resource service create default response
func (o *BackupResourceServiceCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this backup resource service create default response has a 2xx status code
func (o *BackupResourceServiceCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this backup resource service create default response has a 3xx status code
func (o *BackupResourceServiceCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this backup resource service create default response has a 4xx status code
func (o *BackupResourceServiceCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this backup resource service create default response has a 5xx status code
func (o *BackupResourceServiceCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this backup resource service create default response a status code equal to that given
func (o *BackupResourceServiceCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *BackupResourceServiceCreateDefault) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/clusters/{backup.fullName.clusterName}/dataprotection/backups][%d] BackupResourceService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *BackupResourceServiceCreateDefault) String() string {
	return fmt.Sprintf("[POST /v1alpha1/clusters/{backup.fullName.clusterName}/dataprotection/backups][%d] BackupResourceService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *BackupResourceServiceCreateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *BackupResourceServiceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}