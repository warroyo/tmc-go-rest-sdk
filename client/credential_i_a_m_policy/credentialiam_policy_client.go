// Code generated by go-swagger; DO NOT EDIT.

package credential_i_a_m_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new credential i a m policy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for credential i a m policy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CredentialIAMPolicyGet(params *CredentialIAMPolicyGetParams, opts ...ClientOption) (*CredentialIAMPolicyGetOK, error)

	CredentialIAMPolicyPatch(params *CredentialIAMPolicyPatchParams, opts ...ClientOption) (*CredentialIAMPolicyPatchOK, error)

	CredentialIAMPolicyTestPermissions(params *CredentialIAMPolicyTestPermissionsParams, opts ...ClientOption) (*CredentialIAMPolicyTestPermissionsOK, error)

	CredentialIAMPolicyUpdate(params *CredentialIAMPolicyUpdateParams, opts ...ClientOption) (*CredentialIAMPolicyUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CredentialIAMPolicyGet gets policy for a credential
*/
func (a *Client) CredentialIAMPolicyGet(params *CredentialIAMPolicyGetParams, opts ...ClientOption) (*CredentialIAMPolicyGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCredentialIAMPolicyGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CredentialIAMPolicy_Get",
		Method:             "GET",
		PathPattern:        "/v1alpha1/account/managementcluster/provisioner/credentials:iam/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CredentialIAMPolicyGetReader{formats: a.formats},
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
	success, ok := result.(*CredentialIAMPolicyGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CredentialIAMPolicyGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CredentialIAMPolicyPatch patches a credential policy
*/
func (a *Client) CredentialIAMPolicyPatch(params *CredentialIAMPolicyPatchParams, opts ...ClientOption) (*CredentialIAMPolicyPatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCredentialIAMPolicyPatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CredentialIAMPolicy_Patch",
		Method:             "PATCH",
		PathPattern:        "/v1alpha1/account/managementcluster/provisioner/credentials:iam/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CredentialIAMPolicyPatchReader{formats: a.formats},
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
	success, ok := result.(*CredentialIAMPolicyPatchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CredentialIAMPolicyPatchDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CredentialIAMPolicyTestPermissions tests permissions for a credential
*/
func (a *Client) CredentialIAMPolicyTestPermissions(params *CredentialIAMPolicyTestPermissionsParams, opts ...ClientOption) (*CredentialIAMPolicyTestPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCredentialIAMPolicyTestPermissionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CredentialIAMPolicy_TestPermissions",
		Method:             "POST",
		PathPattern:        "/v1alpha1/account/managementcluster/provisioner/credentials:iam/{fullName.name}/testPermissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CredentialIAMPolicyTestPermissionsReader{formats: a.formats},
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
	success, ok := result.(*CredentialIAMPolicyTestPermissionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CredentialIAMPolicyTestPermissionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CredentialIAMPolicyUpdate updates overwrite policy for a credential deleted if body is empty
*/
func (a *Client) CredentialIAMPolicyUpdate(params *CredentialIAMPolicyUpdateParams, opts ...ClientOption) (*CredentialIAMPolicyUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCredentialIAMPolicyUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CredentialIAMPolicy_Update",
		Method:             "PUT",
		PathPattern:        "/v1alpha1/account/managementcluster/provisioner/credentials:iam/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CredentialIAMPolicyUpdateReader{formats: a.formats},
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
	success, ok := result.(*CredentialIAMPolicyUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CredentialIAMPolicyUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
