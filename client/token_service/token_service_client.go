// Code generated by go-swagger; DO NOT EDIT.

package token_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new token service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for token service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	TokenServiceReview(params *TokenServiceReviewParams, opts ...ClientOption) (*TokenServiceReviewOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
TokenServiceReview reviews a token to validate it and look up users groups this API is configured as a webhook on the apiserver and expects request response in the specified format https kubernetes io docs reference access authn authz authentication webhook token authentication
*/
func (a *Client) TokenServiceReview(params *TokenServiceReviewParams, opts ...ClientOption) (*TokenServiceReviewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTokenServiceReviewParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TokenService_Review",
		Method:             "POST",
		PathPattern:        "/v1alpha1/clusters/{clusterUid}/auth/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TokenServiceReviewReader{formats: a.formats},
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
	success, ok := result.(*TokenServiceReviewOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*TokenServiceReviewDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}