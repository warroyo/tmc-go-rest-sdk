// Code generated by go-swagger; DO NOT EDIT.

package setting_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new setting resource service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for setting resource service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SettingResourceServiceCreate(params *SettingResourceServiceCreateParams, opts ...ClientOption) (*SettingResourceServiceCreateOK, error)

	SettingResourceServiceDelete(params *SettingResourceServiceDeleteParams, opts ...ClientOption) (*SettingResourceServiceDeleteOK, error)

	SettingResourceServiceGet(params *SettingResourceServiceGetParams, opts ...ClientOption) (*SettingResourceServiceGetOK, error)

	SettingResourceServiceUpdate(params *SettingResourceServiceUpdateParams, opts ...ClientOption) (*SettingResourceServiceUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
SettingResourceServiceCreate creates a setting
*/
func (a *Client) SettingResourceServiceCreate(params *SettingResourceServiceCreateParams, opts ...ClientOption) (*SettingResourceServiceCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSettingResourceServiceCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SettingResourceService_Create",
		Method:             "POST",
		PathPattern:        "/v1alpha1/organization/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SettingResourceServiceCreateReader{formats: a.formats},
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
	success, ok := result.(*SettingResourceServiceCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SettingResourceServiceCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SettingResourceServiceDelete deletes a setting
*/
func (a *Client) SettingResourceServiceDelete(params *SettingResourceServiceDeleteParams, opts ...ClientOption) (*SettingResourceServiceDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSettingResourceServiceDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SettingResourceService_Delete",
		Method:             "DELETE",
		PathPattern:        "/v1alpha1/organization/settings/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SettingResourceServiceDeleteReader{formats: a.formats},
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
	success, ok := result.(*SettingResourceServiceDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SettingResourceServiceDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SettingResourceServiceGet gets a setting
*/
func (a *Client) SettingResourceServiceGet(params *SettingResourceServiceGetParams, opts ...ClientOption) (*SettingResourceServiceGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSettingResourceServiceGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SettingResourceService_Get",
		Method:             "GET",
		PathPattern:        "/v1alpha1/organization/settings/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SettingResourceServiceGetReader{formats: a.formats},
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
	success, ok := result.(*SettingResourceServiceGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SettingResourceServiceGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SettingResourceServiceUpdate updates overwrite a setting
*/
func (a *Client) SettingResourceServiceUpdate(params *SettingResourceServiceUpdateParams, opts ...ClientOption) (*SettingResourceServiceUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSettingResourceServiceUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SettingResourceService_Update",
		Method:             "PUT",
		PathPattern:        "/v1alpha1/organization/settings/{setting.fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SettingResourceServiceUpdateReader{formats: a.formats},
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
	success, ok := result.(*SettingResourceServiceUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SettingResourceServiceUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
