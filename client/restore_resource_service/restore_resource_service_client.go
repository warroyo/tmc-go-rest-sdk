// Code generated by go-swagger; DO NOT EDIT.

package restore_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new restore resource service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for restore resource service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	RestoreResourceServiceCreate(params *RestoreResourceServiceCreateParams, opts ...ClientOption) (*RestoreResourceServiceCreateOK, error)

	RestoreResourceServiceDelete(params *RestoreResourceServiceDeleteParams, opts ...ClientOption) (*RestoreResourceServiceDeleteOK, error)

	RestoreResourceServiceGet(params *RestoreResourceServiceGetParams, opts ...ClientOption) (*RestoreResourceServiceGetOK, error)

	RestoreResourceServiceList(params *RestoreResourceServiceListParams, opts ...ClientOption) (*RestoreResourceServiceListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
RestoreResourceServiceCreate creates a restore
*/
func (a *Client) RestoreResourceServiceCreate(params *RestoreResourceServiceCreateParams, opts ...ClientOption) (*RestoreResourceServiceCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestoreResourceServiceCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RestoreResourceService_Create",
		Method:             "POST",
		PathPattern:        "/v1alpha1/clusters/{restore.fullName.clusterName}/dataprotection/restores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RestoreResourceServiceCreateReader{formats: a.formats},
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
	success, ok := result.(*RestoreResourceServiceCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RestoreResourceServiceCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RestoreResourceServiceDelete deletes a restore
*/
func (a *Client) RestoreResourceServiceDelete(params *RestoreResourceServiceDeleteParams, opts ...ClientOption) (*RestoreResourceServiceDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestoreResourceServiceDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RestoreResourceService_Delete",
		Method:             "DELETE",
		PathPattern:        "/v1alpha1/clusters/{fullName.clusterName}/dataprotection/restores/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RestoreResourceServiceDeleteReader{formats: a.formats},
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
	success, ok := result.(*RestoreResourceServiceDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RestoreResourceServiceDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RestoreResourceServiceGet gets a restore
*/
func (a *Client) RestoreResourceServiceGet(params *RestoreResourceServiceGetParams, opts ...ClientOption) (*RestoreResourceServiceGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestoreResourceServiceGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RestoreResourceService_Get",
		Method:             "GET",
		PathPattern:        "/v1alpha1/clusters/{fullName.clusterName}/dataprotection/restores/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RestoreResourceServiceGetReader{formats: a.formats},
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
	success, ok := result.(*RestoreResourceServiceGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RestoreResourceServiceGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RestoreResourceServiceList lists restores
*/
func (a *Client) RestoreResourceServiceList(params *RestoreResourceServiceListParams, opts ...ClientOption) (*RestoreResourceServiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestoreResourceServiceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RestoreResourceService_List",
		Method:             "GET",
		PathPattern:        "/v1alpha1/clusters/{searchScope.clusterName}/dataprotection/restores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RestoreResourceServiceListReader{formats: a.formats},
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
	success, ok := result.(*RestoreResourceServiceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RestoreResourceServiceListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}