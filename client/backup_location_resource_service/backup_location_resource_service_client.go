// Code generated by go-swagger; DO NOT EDIT.

package backup_location_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new backup location resource service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for backup location resource service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	BackupLocationResourceServiceCreate(params *BackupLocationResourceServiceCreateParams, opts ...ClientOption) (*BackupLocationResourceServiceCreateOK, error)

	BackupLocationResourceServiceDelete(params *BackupLocationResourceServiceDeleteParams, opts ...ClientOption) (*BackupLocationResourceServiceDeleteOK, error)

	BackupLocationResourceServiceGet(params *BackupLocationResourceServiceGetParams, opts ...ClientOption) (*BackupLocationResourceServiceGetOK, error)

	BackupLocationResourceServiceList(params *BackupLocationResourceServiceListParams, opts ...ClientOption) (*BackupLocationResourceServiceListOK, error)

	BackupLocationResourceServicePatch(params *BackupLocationResourceServicePatchParams, opts ...ClientOption) (*BackupLocationResourceServicePatchOK, error)

	BackupLocationResourceServiceUpdate(params *BackupLocationResourceServiceUpdateParams, opts ...ClientOption) (*BackupLocationResourceServiceUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
BackupLocationResourceServiceCreate creates a backup location
*/
func (a *Client) BackupLocationResourceServiceCreate(params *BackupLocationResourceServiceCreateParams, opts ...ClientOption) (*BackupLocationResourceServiceCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupLocationResourceServiceCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BackupLocationResourceService_Create",
		Method:             "POST",
		PathPattern:        "/v1alpha1/dataprotection/providers/{backupLocation.fullName.providerName}/backuplocations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BackupLocationResourceServiceCreateReader{formats: a.formats},
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
	success, ok := result.(*BackupLocationResourceServiceCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BackupLocationResourceServiceCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BackupLocationResourceServiceDelete deletes a backup location
*/
func (a *Client) BackupLocationResourceServiceDelete(params *BackupLocationResourceServiceDeleteParams, opts ...ClientOption) (*BackupLocationResourceServiceDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupLocationResourceServiceDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BackupLocationResourceService_Delete",
		Method:             "DELETE",
		PathPattern:        "/v1alpha1/dataprotection/providers/{fullName.providerName}/backuplocations/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BackupLocationResourceServiceDeleteReader{formats: a.formats},
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
	success, ok := result.(*BackupLocationResourceServiceDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BackupLocationResourceServiceDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BackupLocationResourceServiceGet gets a backup location
*/
func (a *Client) BackupLocationResourceServiceGet(params *BackupLocationResourceServiceGetParams, opts ...ClientOption) (*BackupLocationResourceServiceGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupLocationResourceServiceGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BackupLocationResourceService_Get",
		Method:             "GET",
		PathPattern:        "/v1alpha1/dataprotection/providers/{fullName.providerName}/backuplocations/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BackupLocationResourceServiceGetReader{formats: a.formats},
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
	success, ok := result.(*BackupLocationResourceServiceGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BackupLocationResourceServiceGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BackupLocationResourceServiceList lists backup locations
*/
func (a *Client) BackupLocationResourceServiceList(params *BackupLocationResourceServiceListParams, opts ...ClientOption) (*BackupLocationResourceServiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupLocationResourceServiceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BackupLocationResourceService_List",
		Method:             "GET",
		PathPattern:        "/v1alpha1/dataprotection/providers/{searchScope.providerName}/backuplocations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BackupLocationResourceServiceListReader{formats: a.formats},
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
	success, ok := result.(*BackupLocationResourceServiceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BackupLocationResourceServiceListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BackupLocationResourceServicePatch patches partially update a backup location
*/
func (a *Client) BackupLocationResourceServicePatch(params *BackupLocationResourceServicePatchParams, opts ...ClientOption) (*BackupLocationResourceServicePatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupLocationResourceServicePatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BackupLocationResourceService_Patch",
		Method:             "PATCH",
		PathPattern:        "/v1alpha1/dataprotection/providers/{fullName.providerName}/backuplocations/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BackupLocationResourceServicePatchReader{formats: a.formats},
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
	success, ok := result.(*BackupLocationResourceServicePatchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BackupLocationResourceServicePatchDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
BackupLocationResourceServiceUpdate updates overwrite a backup location
*/
func (a *Client) BackupLocationResourceServiceUpdate(params *BackupLocationResourceServiceUpdateParams, opts ...ClientOption) (*BackupLocationResourceServiceUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupLocationResourceServiceUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BackupLocationResourceService_Update",
		Method:             "PUT",
		PathPattern:        "/v1alpha1/dataprotection/providers/{backupLocation.fullName.providerName}/backuplocations/{backupLocation.fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BackupLocationResourceServiceUpdateReader{formats: a.formats},
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
	success, ok := result.(*BackupLocationResourceServiceUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BackupLocationResourceServiceUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
