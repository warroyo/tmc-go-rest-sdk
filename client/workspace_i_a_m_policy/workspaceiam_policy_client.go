// Code generated by go-swagger; DO NOT EDIT.

package workspace_i_a_m_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new workspace i a m policy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for workspace i a m policy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	WorkspaceIAMPolicyGet(params *WorkspaceIAMPolicyGetParams, opts ...ClientOption) (*WorkspaceIAMPolicyGetOK, error)

	WorkspaceIAMPolicyPatch(params *WorkspaceIAMPolicyPatchParams, opts ...ClientOption) (*WorkspaceIAMPolicyPatchOK, error)

	WorkspaceIAMPolicyTestPermissions(params *WorkspaceIAMPolicyTestPermissionsParams, opts ...ClientOption) (*WorkspaceIAMPolicyTestPermissionsOK, error)

	WorkspaceIAMPolicyUpdate(params *WorkspaceIAMPolicyUpdateParams, opts ...ClientOption) (*WorkspaceIAMPolicyUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
WorkspaceIAMPolicyGet gets policy for a workspace
*/
func (a *Client) WorkspaceIAMPolicyGet(params *WorkspaceIAMPolicyGetParams, opts ...ClientOption) (*WorkspaceIAMPolicyGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkspaceIAMPolicyGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "WorkspaceIAMPolicy_Get",
		Method:             "GET",
		PathPattern:        "/v1alpha1/workspaces:iam/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &WorkspaceIAMPolicyGetReader{formats: a.formats},
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
	success, ok := result.(*WorkspaceIAMPolicyGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*WorkspaceIAMPolicyGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
WorkspaceIAMPolicyPatch patches a workspace policy
*/
func (a *Client) WorkspaceIAMPolicyPatch(params *WorkspaceIAMPolicyPatchParams, opts ...ClientOption) (*WorkspaceIAMPolicyPatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkspaceIAMPolicyPatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "WorkspaceIAMPolicy_Patch",
		Method:             "PATCH",
		PathPattern:        "/v1alpha1/workspaces:iam/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &WorkspaceIAMPolicyPatchReader{formats: a.formats},
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
	success, ok := result.(*WorkspaceIAMPolicyPatchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*WorkspaceIAMPolicyPatchDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
WorkspaceIAMPolicyTestPermissions tests permissions for a workspace
*/
func (a *Client) WorkspaceIAMPolicyTestPermissions(params *WorkspaceIAMPolicyTestPermissionsParams, opts ...ClientOption) (*WorkspaceIAMPolicyTestPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkspaceIAMPolicyTestPermissionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "WorkspaceIAMPolicy_TestPermissions",
		Method:             "POST",
		PathPattern:        "/v1alpha1/workspaces:iam/{fullName.name}/testPermissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &WorkspaceIAMPolicyTestPermissionsReader{formats: a.formats},
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
	success, ok := result.(*WorkspaceIAMPolicyTestPermissionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*WorkspaceIAMPolicyTestPermissionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
WorkspaceIAMPolicyUpdate updates overwrite policy for a workspace deleted if body is empty
*/
func (a *Client) WorkspaceIAMPolicyUpdate(params *WorkspaceIAMPolicyUpdateParams, opts ...ClientOption) (*WorkspaceIAMPolicyUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkspaceIAMPolicyUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "WorkspaceIAMPolicy_Update",
		Method:             "PUT",
		PathPattern:        "/v1alpha1/workspaces:iam/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &WorkspaceIAMPolicyUpdateReader{formats: a.formats},
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
	success, ok := result.(*WorkspaceIAMPolicyUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*WorkspaceIAMPolicyUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
