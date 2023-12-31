// Code generated by go-swagger; DO NOT EDIT.

package subscriptions_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new subscriptions service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for subscriptions service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SubscriptionsServiceList(params *SubscriptionsServiceListParams, opts ...ClientOption) (*SubscriptionsServiceListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
SubscriptionsServiceList lists returns all subscriptions information related to the organization
*/
func (a *Client) SubscriptionsServiceList(params *SubscriptionsServiceListParams, opts ...ClientOption) (*SubscriptionsServiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscriptionsServiceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SubscriptionsService_List",
		Method:             "GET",
		PathPattern:        "/v1alpha1/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SubscriptionsServiceListReader{formats: a.formats},
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
	success, ok := result.(*SubscriptionsServiceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SubscriptionsServiceListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
