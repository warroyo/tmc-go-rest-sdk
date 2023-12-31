// Code generated by go-swagger; DO NOT EDIT.

package admin_kubeconfig_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// AdminKubeconfigServiceGetAdminKubeconfigReader is a Reader for the AdminKubeconfigServiceGetAdminKubeconfig structure.
type AdminKubeconfigServiceGetAdminKubeconfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminKubeconfigServiceGetAdminKubeconfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminKubeconfigServiceGetAdminKubeconfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAdminKubeconfigServiceGetAdminKubeconfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAdminKubeconfigServiceGetAdminKubeconfigOK creates a AdminKubeconfigServiceGetAdminKubeconfigOK with default headers values
func NewAdminKubeconfigServiceGetAdminKubeconfigOK() *AdminKubeconfigServiceGetAdminKubeconfigOK {
	return &AdminKubeconfigServiceGetAdminKubeconfigOK{}
}

/*
AdminKubeconfigServiceGetAdminKubeconfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type AdminKubeconfigServiceGetAdminKubeconfigOK struct {
	Payload *models.ClusterAdminkubeconfigGetAdminKubeconfigResponse
}

// IsSuccess returns true when this admin kubeconfig service get admin kubeconfig o k response has a 2xx status code
func (o *AdminKubeconfigServiceGetAdminKubeconfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin kubeconfig service get admin kubeconfig o k response has a 3xx status code
func (o *AdminKubeconfigServiceGetAdminKubeconfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin kubeconfig service get admin kubeconfig o k response has a 4xx status code
func (o *AdminKubeconfigServiceGetAdminKubeconfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin kubeconfig service get admin kubeconfig o k response has a 5xx status code
func (o *AdminKubeconfigServiceGetAdminKubeconfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin kubeconfig service get admin kubeconfig o k response a status code equal to that given
func (o *AdminKubeconfigServiceGetAdminKubeconfigOK) IsCode(code int) bool {
	return code == 200
}

func (o *AdminKubeconfigServiceGetAdminKubeconfigOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.name}/adminkubeconfig][%d] adminKubeconfigServiceGetAdminKubeconfigOK  %+v", 200, o.Payload)
}

func (o *AdminKubeconfigServiceGetAdminKubeconfigOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.name}/adminkubeconfig][%d] adminKubeconfigServiceGetAdminKubeconfigOK  %+v", 200, o.Payload)
}

func (o *AdminKubeconfigServiceGetAdminKubeconfigOK) GetPayload() *models.ClusterAdminkubeconfigGetAdminKubeconfigResponse {
	return o.Payload
}

func (o *AdminKubeconfigServiceGetAdminKubeconfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterAdminkubeconfigGetAdminKubeconfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminKubeconfigServiceGetAdminKubeconfigDefault creates a AdminKubeconfigServiceGetAdminKubeconfigDefault with default headers values
func NewAdminKubeconfigServiceGetAdminKubeconfigDefault(code int) *AdminKubeconfigServiceGetAdminKubeconfigDefault {
	return &AdminKubeconfigServiceGetAdminKubeconfigDefault{
		_statusCode: code,
	}
}

/*
AdminKubeconfigServiceGetAdminKubeconfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AdminKubeconfigServiceGetAdminKubeconfigDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the admin kubeconfig service get admin kubeconfig default response
func (o *AdminKubeconfigServiceGetAdminKubeconfigDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this admin kubeconfig service get admin kubeconfig default response has a 2xx status code
func (o *AdminKubeconfigServiceGetAdminKubeconfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this admin kubeconfig service get admin kubeconfig default response has a 3xx status code
func (o *AdminKubeconfigServiceGetAdminKubeconfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this admin kubeconfig service get admin kubeconfig default response has a 4xx status code
func (o *AdminKubeconfigServiceGetAdminKubeconfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this admin kubeconfig service get admin kubeconfig default response has a 5xx status code
func (o *AdminKubeconfigServiceGetAdminKubeconfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this admin kubeconfig service get admin kubeconfig default response a status code equal to that given
func (o *AdminKubeconfigServiceGetAdminKubeconfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AdminKubeconfigServiceGetAdminKubeconfigDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.name}/adminkubeconfig][%d] AdminKubeconfigService_GetAdminKubeconfig default  %+v", o._statusCode, o.Payload)
}

func (o *AdminKubeconfigServiceGetAdminKubeconfigDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.name}/adminkubeconfig][%d] AdminKubeconfigService_GetAdminKubeconfig default  %+v", o._statusCode, o.Payload)
}

func (o *AdminKubeconfigServiceGetAdminKubeconfigDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *AdminKubeconfigServiceGetAdminKubeconfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
