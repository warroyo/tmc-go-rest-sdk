// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new events API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for events API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	EventsList(params *EventsListParams, opts ...ClientOption) (*EventsListOK, error)

	EventsStream(params *EventsStreamParams, opts ...ClientOption) (*EventsStreamOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
EventsList lists events
*/
func (a *Client) EventsList(params *EventsListParams, opts ...ClientOption) (*EventsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEventsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Events_List",
		Method:             "GET",
		PathPattern:        "/v1alpha1/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EventsListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EventsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EventsListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EventsStream streams allows a client to stream events for one or more resource types
*/
func (a *Client) EventsStream(params *EventsStreamParams, opts ...ClientOption) (*EventsStreamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEventsStreamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Events_Stream",
		Method:             "GET",
		PathPattern:        "/v1alpha1/events/stream",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EventsStreamReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EventsStreamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EventsStreamDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
