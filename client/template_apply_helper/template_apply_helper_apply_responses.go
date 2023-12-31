// Code generated by go-swagger; DO NOT EDIT.

package template_apply_helper

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// TemplateApplyHelperApplyReader is a Reader for the TemplateApplyHelperApply structure.
type TemplateApplyHelperApplyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TemplateApplyHelperApplyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTemplateApplyHelperApplyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTemplateApplyHelperApplyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTemplateApplyHelperApplyOK creates a TemplateApplyHelperApplyOK with default headers values
func NewTemplateApplyHelperApplyOK() *TemplateApplyHelperApplyOK {
	return &TemplateApplyHelperApplyOK{}
}

/*
TemplateApplyHelperApplyOK describes a response with status code 200, with default header values.

A successful response.
*/
type TemplateApplyHelperApplyOK struct {
	Payload *models.PolicyTemplateApplyTemplateResponse
}

// IsSuccess returns true when this template apply helper apply o k response has a 2xx status code
func (o *TemplateApplyHelperApplyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this template apply helper apply o k response has a 3xx status code
func (o *TemplateApplyHelperApplyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this template apply helper apply o k response has a 4xx status code
func (o *TemplateApplyHelperApplyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this template apply helper apply o k response has a 5xx status code
func (o *TemplateApplyHelperApplyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this template apply helper apply o k response a status code equal to that given
func (o *TemplateApplyHelperApplyOK) IsCode(code int) bool {
	return code == 200
}

func (o *TemplateApplyHelperApplyOK) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/policy/templates:apply][%d] templateApplyHelperApplyOK  %+v", 200, o.Payload)
}

func (o *TemplateApplyHelperApplyOK) String() string {
	return fmt.Sprintf("[POST /v1alpha1/policy/templates:apply][%d] templateApplyHelperApplyOK  %+v", 200, o.Payload)
}

func (o *TemplateApplyHelperApplyOK) GetPayload() *models.PolicyTemplateApplyTemplateResponse {
	return o.Payload
}

func (o *TemplateApplyHelperApplyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyTemplateApplyTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTemplateApplyHelperApplyDefault creates a TemplateApplyHelperApplyDefault with default headers values
func NewTemplateApplyHelperApplyDefault(code int) *TemplateApplyHelperApplyDefault {
	return &TemplateApplyHelperApplyDefault{
		_statusCode: code,
	}
}

/*
TemplateApplyHelperApplyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type TemplateApplyHelperApplyDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the template apply helper apply default response
func (o *TemplateApplyHelperApplyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this template apply helper apply default response has a 2xx status code
func (o *TemplateApplyHelperApplyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this template apply helper apply default response has a 3xx status code
func (o *TemplateApplyHelperApplyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this template apply helper apply default response has a 4xx status code
func (o *TemplateApplyHelperApplyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this template apply helper apply default response has a 5xx status code
func (o *TemplateApplyHelperApplyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this template apply helper apply default response a status code equal to that given
func (o *TemplateApplyHelperApplyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *TemplateApplyHelperApplyDefault) Error() string {
	return fmt.Sprintf("[POST /v1alpha1/policy/templates:apply][%d] TemplateApplyHelper_Apply default  %+v", o._statusCode, o.Payload)
}

func (o *TemplateApplyHelperApplyDefault) String() string {
	return fmt.Sprintf("[POST /v1alpha1/policy/templates:apply][%d] TemplateApplyHelper_Apply default  %+v", o._statusCode, o.Payload)
}

func (o *TemplateApplyHelperApplyDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *TemplateApplyHelperApplyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
