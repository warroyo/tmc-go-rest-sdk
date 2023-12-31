// Code generated by go-swagger; DO NOT EDIT.

package eks_cluster_apply_helper

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// EksClusterApplyHelperApplyReader is a Reader for the EksClusterApplyHelperApply structure.
type EksClusterApplyHelperApplyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EksClusterApplyHelperApplyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEksClusterApplyHelperApplyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEksClusterApplyHelperApplyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEksClusterApplyHelperApplyOK creates a EksClusterApplyHelperApplyOK with default headers values
func NewEksClusterApplyHelperApplyOK() *EksClusterApplyHelperApplyOK {
	return &EksClusterApplyHelperApplyOK{}
}

/*
EksClusterApplyHelperApplyOK describes a response with status code 200, with default header values.

A successful response.
*/
type EksClusterApplyHelperApplyOK struct {
	Payload *models.EksclusterApplyEksClusterResponse
}

// IsSuccess returns true when this eks cluster apply helper apply o k response has a 2xx status code
func (o *EksClusterApplyHelperApplyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this eks cluster apply helper apply o k response has a 3xx status code
func (o *EksClusterApplyHelperApplyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this eks cluster apply helper apply o k response has a 4xx status code
func (o *EksClusterApplyHelperApplyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this eks cluster apply helper apply o k response has a 5xx status code
func (o *EksClusterApplyHelperApplyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this eks cluster apply helper apply o k response a status code equal to that given
func (o *EksClusterApplyHelperApplyOK) IsCode(code int) bool {
	return code == 200
}

func (o *EksClusterApplyHelperApplyOK) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/eksclusters:apply][%d] eksClusterApplyHelperApplyOK  %+v", 200, o.Payload)
}

func (o *EksClusterApplyHelperApplyOK) String() string {
	return fmt.Sprintf("[POST /v1alpha1/eksclusters:apply][%d] eksClusterApplyHelperApplyOK  %+v", 200, o.Payload)
}

func (o *EksClusterApplyHelperApplyOK) GetPayload() *models.EksclusterApplyEksClusterResponse {
	return o.Payload
}

func (o *EksClusterApplyHelperApplyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EksclusterApplyEksClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEksClusterApplyHelperApplyDefault creates a EksClusterApplyHelperApplyDefault with default headers values
func NewEksClusterApplyHelperApplyDefault(code int) *EksClusterApplyHelperApplyDefault {
	return &EksClusterApplyHelperApplyDefault{
		_statusCode: code,
	}
}

/*
EksClusterApplyHelperApplyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EksClusterApplyHelperApplyDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the eks cluster apply helper apply default response
func (o *EksClusterApplyHelperApplyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this eks cluster apply helper apply default response has a 2xx status code
func (o *EksClusterApplyHelperApplyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this eks cluster apply helper apply default response has a 3xx status code
func (o *EksClusterApplyHelperApplyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this eks cluster apply helper apply default response has a 4xx status code
func (o *EksClusterApplyHelperApplyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this eks cluster apply helper apply default response has a 5xx status code
func (o *EksClusterApplyHelperApplyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this eks cluster apply helper apply default response a status code equal to that given
func (o *EksClusterApplyHelperApplyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EksClusterApplyHelperApplyDefault) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/eksclusters:apply][%d] EksClusterApplyHelper_Apply default  %+v", o._statusCode, o.Payload)
}

func (o *EksClusterApplyHelperApplyDefault) String() string {
	return fmt.Sprintf("[POST /v1alpha1/eksclusters:apply][%d] EksClusterApplyHelper_Apply default  %+v", o._statusCode, o.Payload)
}

func (o *EksClusterApplyHelperApplyDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *EksClusterApplyHelperApplyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
