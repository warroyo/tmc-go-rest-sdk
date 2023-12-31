// Code generated by go-swagger; DO NOT EDIT.

package cluster_i_a_m_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cluster i a m policy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cluster i a m policy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ClusterIAMPolicyGet(params *ClusterIAMPolicyGetParams, opts ...ClientOption) (*ClusterIAMPolicyGetOK, error)

	ClusterIAMPolicyPatch(params *ClusterIAMPolicyPatchParams, opts ...ClientOption) (*ClusterIAMPolicyPatchOK, error)

	ClusterIAMPolicyTestPermissions(params *ClusterIAMPolicyTestPermissionsParams, opts ...ClientOption) (*ClusterIAMPolicyTestPermissionsOK, error)

	ClusterIAMPolicyUpdate(params *ClusterIAMPolicyUpdateParams, opts ...ClientOption) (*ClusterIAMPolicyUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ClusterIAMPolicyGet gets policy for a cluster
*/
func (a *Client) ClusterIAMPolicyGet(params *ClusterIAMPolicyGetParams, opts ...ClientOption) (*ClusterIAMPolicyGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClusterIAMPolicyGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ClusterIAMPolicy_Get",
		Method:             "GET",
		PathPattern:        "/v1alpha1/clusters:iam/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ClusterIAMPolicyGetReader{formats: a.formats},
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
	success, ok := result.(*ClusterIAMPolicyGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ClusterIAMPolicyGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ClusterIAMPolicyPatch patches a cluster policy
*/
func (a *Client) ClusterIAMPolicyPatch(params *ClusterIAMPolicyPatchParams, opts ...ClientOption) (*ClusterIAMPolicyPatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClusterIAMPolicyPatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ClusterIAMPolicy_Patch",
		Method:             "PATCH",
		PathPattern:        "/v1alpha1/clusters:iam/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ClusterIAMPolicyPatchReader{formats: a.formats},
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
	success, ok := result.(*ClusterIAMPolicyPatchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ClusterIAMPolicyPatchDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ClusterIAMPolicyTestPermissions tests permissions for a cluster
*/
func (a *Client) ClusterIAMPolicyTestPermissions(params *ClusterIAMPolicyTestPermissionsParams, opts ...ClientOption) (*ClusterIAMPolicyTestPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClusterIAMPolicyTestPermissionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ClusterIAMPolicy_TestPermissions",
		Method:             "POST",
		PathPattern:        "/v1alpha1/clusters:iam/{fullName.name}/testPermissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ClusterIAMPolicyTestPermissionsReader{formats: a.formats},
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
	success, ok := result.(*ClusterIAMPolicyTestPermissionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ClusterIAMPolicyTestPermissionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ClusterIAMPolicyUpdate updates overwrite policy for a cluster deleted if body is empty
*/
func (a *Client) ClusterIAMPolicyUpdate(params *ClusterIAMPolicyUpdateParams, opts ...ClientOption) (*ClusterIAMPolicyUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClusterIAMPolicyUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ClusterIAMPolicy_Update",
		Method:             "PUT",
		PathPattern:        "/v1alpha1/clusters:iam/{fullName.name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ClusterIAMPolicyUpdateReader{formats: a.formats},
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
	success, ok := result.(*ClusterIAMPolicyUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ClusterIAMPolicyUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
