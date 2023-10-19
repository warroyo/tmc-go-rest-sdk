// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// EventsListReader is a Reader for the EventsList structure.
type EventsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EventsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEventsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEventsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEventsListOK creates a EventsListOK with default headers values
func NewEventsListOK() *EventsListOK {
	return &EventsListOK{}
}

/*
EventsListOK describes a response with status code 200, with default header values.

A successful response.
*/
type EventsListOK struct {
	Payload *models.EventsListEventsResponse
}

// IsSuccess returns true when this events list o k response has a 2xx status code
func (o *EventsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this events list o k response has a 3xx status code
func (o *EventsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this events list o k response has a 4xx status code
func (o *EventsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this events list o k response has a 5xx status code
func (o *EventsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this events list o k response a status code equal to that given
func (o *EventsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *EventsListOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/events][%d] eventsListOK  %+v", 200, o.Payload)
}

func (o *EventsListOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/events][%d] eventsListOK  %+v", 200, o.Payload)
}

func (o *EventsListOK) GetPayload() *models.EventsListEventsResponse {
	return o.Payload
}

func (o *EventsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventsListEventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEventsListDefault creates a EventsListDefault with default headers values
func NewEventsListDefault(code int) *EventsListDefault {
	return &EventsListDefault{
		_statusCode: code,
	}
}

/*
EventsListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EventsListDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the events list default response
func (o *EventsListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this events list default response has a 2xx status code
func (o *EventsListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this events list default response has a 3xx status code
func (o *EventsListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this events list default response has a 4xx status code
func (o *EventsListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this events list default response has a 5xx status code
func (o *EventsListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this events list default response a status code equal to that given
func (o *EventsListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EventsListDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/events][%d] Events_List default  %+v", o._statusCode, o.Payload)
}

func (o *EventsListDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/events][%d] Events_List default  %+v", o._statusCode, o.Payload)
}

func (o *EventsListDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *EventsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}