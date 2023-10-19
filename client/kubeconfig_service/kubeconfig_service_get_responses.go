// Code generated by go-swagger; DO NOT EDIT.

package kubeconfig_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// KubeconfigServiceGetReader is a Reader for the KubeconfigServiceGet structure.
type KubeconfigServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubeconfigServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubeconfigServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKubeconfigServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKubeconfigServiceGetOK creates a KubeconfigServiceGetOK with default headers values
func NewKubeconfigServiceGetOK() *KubeconfigServiceGetOK {
	return &KubeconfigServiceGetOK{}
}

/*
KubeconfigServiceGetOK describes a response with status code 200, with default header values.

A successful response.
*/
type KubeconfigServiceGetOK struct {
	Payload *models.ClusterKubeconfigGetKubeconfigResponse
}

// IsSuccess returns true when this kubeconfig service get o k response has a 2xx status code
func (o *KubeconfigServiceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubeconfig service get o k response has a 3xx status code
func (o *KubeconfigServiceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubeconfig service get o k response has a 4xx status code
func (o *KubeconfigServiceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubeconfig service get o k response has a 5xx status code
func (o *KubeconfigServiceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubeconfig service get o k response a status code equal to that given
func (o *KubeconfigServiceGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubeconfigServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.name}/kubeconfig][%d] kubeconfigServiceGetOK  %+v", 200, o.Payload)
}

func (o *KubeconfigServiceGetOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.name}/kubeconfig][%d] kubeconfigServiceGetOK  %+v", 200, o.Payload)
}

func (o *KubeconfigServiceGetOK) GetPayload() *models.ClusterKubeconfigGetKubeconfigResponse {
	return o.Payload
}

func (o *KubeconfigServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterKubeconfigGetKubeconfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeconfigServiceGetDefault creates a KubeconfigServiceGetDefault with default headers values
func NewKubeconfigServiceGetDefault(code int) *KubeconfigServiceGetDefault {
	return &KubeconfigServiceGetDefault{
		_statusCode: code,
	}
}

/*
KubeconfigServiceGetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type KubeconfigServiceGetDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the kubeconfig service get default response
func (o *KubeconfigServiceGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this kubeconfig service get default response has a 2xx status code
func (o *KubeconfigServiceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this kubeconfig service get default response has a 3xx status code
func (o *KubeconfigServiceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this kubeconfig service get default response has a 4xx status code
func (o *KubeconfigServiceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this kubeconfig service get default response has a 5xx status code
func (o *KubeconfigServiceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this kubeconfig service get default response a status code equal to that given
func (o *KubeconfigServiceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *KubeconfigServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.name}/kubeconfig][%d] KubeconfigService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *KubeconfigServiceGetDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/clusters/{fullName.name}/kubeconfig][%d] KubeconfigService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *KubeconfigServiceGetDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *KubeconfigServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
