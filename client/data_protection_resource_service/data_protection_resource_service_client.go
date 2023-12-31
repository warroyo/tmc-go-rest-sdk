// Code generated by go-swagger; DO NOT EDIT.

package data_protection_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new data protection resource service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for data protection resource service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DataProtectionResourceServiceCreate(params *DataProtectionResourceServiceCreateParams, opts ...ClientOption) (*DataProtectionResourceServiceCreateOK, error)

	DataProtectionResourceServiceDelete(params *DataProtectionResourceServiceDeleteParams, opts ...ClientOption) (*DataProtectionResourceServiceDeleteOK, error)

	DataProtectionResourceServiceList(params *DataProtectionResourceServiceListParams, opts ...ClientOption) (*DataProtectionResourceServiceListOK, error)

	DataProtectionResourceServiceUpdate(params *DataProtectionResourceServiceUpdateParams, opts ...ClientOption) (*DataProtectionResourceServiceUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DataProtectionResourceServiceCreate creates a data protection
*/
func (a *Client) DataProtectionResourceServiceCreate(params *DataProtectionResourceServiceCreateParams, opts ...ClientOption) (*DataProtectionResourceServiceCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataProtectionResourceServiceCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DataProtectionResourceService_Create",
		Method:             "POST",
		PathPattern:        "/v1alpha1/clusters/{dataProtection.fullName.clusterName}/dataprotection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DataProtectionResourceServiceCreateReader{formats: a.formats},
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
	success, ok := result.(*DataProtectionResourceServiceCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DataProtectionResourceServiceCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DataProtectionResourceServiceDelete deletes a data protection
*/
func (a *Client) DataProtectionResourceServiceDelete(params *DataProtectionResourceServiceDeleteParams, opts ...ClientOption) (*DataProtectionResourceServiceDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataProtectionResourceServiceDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DataProtectionResourceService_Delete",
		Method:             "DELETE",
		PathPattern:        "/v1alpha1/clusters/{fullName.clusterName}/dataprotection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DataProtectionResourceServiceDeleteReader{formats: a.formats},
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
	success, ok := result.(*DataProtectionResourceServiceDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DataProtectionResourceServiceDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DataProtectionResourceServiceList lists data protections
*/
func (a *Client) DataProtectionResourceServiceList(params *DataProtectionResourceServiceListParams, opts ...ClientOption) (*DataProtectionResourceServiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataProtectionResourceServiceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DataProtectionResourceService_List",
		Method:             "GET",
		PathPattern:        "/v1alpha1/clusters/{searchScope.clusterName}/dataprotection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DataProtectionResourceServiceListReader{formats: a.formats},
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
	success, ok := result.(*DataProtectionResourceServiceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DataProtectionResourceServiceListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DataProtectionResourceServiceUpdate updates overwrite a data protection
*/
func (a *Client) DataProtectionResourceServiceUpdate(params *DataProtectionResourceServiceUpdateParams, opts ...ClientOption) (*DataProtectionResourceServiceUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataProtectionResourceServiceUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DataProtectionResourceService_Update",
		Method:             "PUT",
		PathPattern:        "/v1alpha1/clusters/{dataProtection.fullName.clusterName}/dataprotection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DataProtectionResourceServiceUpdateReader{formats: a.formats},
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
	success, ok := result.(*DataProtectionResourceServiceUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DataProtectionResourceServiceUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
