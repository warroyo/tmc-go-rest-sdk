// Code generated by go-swagger; DO NOT EDIT.

package package_helper_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// PackageHelperServiceValidatePrerequisitesReader is a Reader for the PackageHelperServiceValidatePrerequisites structure.
type PackageHelperServiceValidatePrerequisitesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackageHelperServiceValidatePrerequisitesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackageHelperServiceValidatePrerequisitesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackageHelperServiceValidatePrerequisitesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackageHelperServiceValidatePrerequisitesOK creates a PackageHelperServiceValidatePrerequisitesOK with default headers values
func NewPackageHelperServiceValidatePrerequisitesOK() *PackageHelperServiceValidatePrerequisitesOK {
	return &PackageHelperServiceValidatePrerequisitesOK{}
}

/*
PackageHelperServiceValidatePrerequisitesOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackageHelperServiceValidatePrerequisitesOK struct {
	Payload *models.TanzupackageTapValidatePrerequisitesResponse
}

// IsSuccess returns true when this package helper service validate prerequisites o k response has a 2xx status code
func (o *PackageHelperServiceValidatePrerequisitesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this package helper service validate prerequisites o k response has a 3xx status code
func (o *PackageHelperServiceValidatePrerequisitesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this package helper service validate prerequisites o k response has a 4xx status code
func (o *PackageHelperServiceValidatePrerequisitesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this package helper service validate prerequisites o k response has a 5xx status code
func (o *PackageHelperServiceValidatePrerequisitesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this package helper service validate prerequisites o k response a status code equal to that given
func (o *PackageHelperServiceValidatePrerequisitesOK) IsCode(code int) bool {
	return code == 200
}

func (o *PackageHelperServiceValidatePrerequisitesOK) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/tanzupackage/tap:validateprerequisites][%d] packageHelperServiceValidatePrerequisitesOK  %+v", 200, o.Payload)
}

func (o *PackageHelperServiceValidatePrerequisitesOK) String() string {
	return fmt.Sprintf("[POST /v1alpha1/tanzupackage/tap:validateprerequisites][%d] packageHelperServiceValidatePrerequisitesOK  %+v", 200, o.Payload)
}

func (o *PackageHelperServiceValidatePrerequisitesOK) GetPayload() *models.TanzupackageTapValidatePrerequisitesResponse {
	return o.Payload
}

func (o *PackageHelperServiceValidatePrerequisitesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TanzupackageTapValidatePrerequisitesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackageHelperServiceValidatePrerequisitesDefault creates a PackageHelperServiceValidatePrerequisitesDefault with default headers values
func NewPackageHelperServiceValidatePrerequisitesDefault(code int) *PackageHelperServiceValidatePrerequisitesDefault {
	return &PackageHelperServiceValidatePrerequisitesDefault{
		_statusCode: code,
	}
}

/*
PackageHelperServiceValidatePrerequisitesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackageHelperServiceValidatePrerequisitesDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the package helper service validate prerequisites default response
func (o *PackageHelperServiceValidatePrerequisitesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this package helper service validate prerequisites default response has a 2xx status code
func (o *PackageHelperServiceValidatePrerequisitesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this package helper service validate prerequisites default response has a 3xx status code
func (o *PackageHelperServiceValidatePrerequisitesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this package helper service validate prerequisites default response has a 4xx status code
func (o *PackageHelperServiceValidatePrerequisitesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this package helper service validate prerequisites default response has a 5xx status code
func (o *PackageHelperServiceValidatePrerequisitesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this package helper service validate prerequisites default response a status code equal to that given
func (o *PackageHelperServiceValidatePrerequisitesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PackageHelperServiceValidatePrerequisitesDefault) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/tanzupackage/tap:validateprerequisites][%d] PackageHelperService_ValidatePrerequisites default  %+v", o._statusCode, o.Payload)
}

func (o *PackageHelperServiceValidatePrerequisitesDefault) String() string {
	return fmt.Sprintf("[POST /v1alpha1/tanzupackage/tap:validateprerequisites][%d] PackageHelperService_ValidatePrerequisites default  %+v", o._statusCode, o.Payload)
}

func (o *PackageHelperServiceValidatePrerequisitesDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *PackageHelperServiceValidatePrerequisitesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}