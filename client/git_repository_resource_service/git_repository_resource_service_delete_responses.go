// Code generated by go-swagger; DO NOT EDIT.

package git_repository_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// GitRepositoryResourceServiceDeleteReader is a Reader for the GitRepositoryResourceServiceDelete structure.
type GitRepositoryResourceServiceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GitRepositoryResourceServiceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGitRepositoryResourceServiceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGitRepositoryResourceServiceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGitRepositoryResourceServiceDeleteOK creates a GitRepositoryResourceServiceDeleteOK with default headers values
func NewGitRepositoryResourceServiceDeleteOK() *GitRepositoryResourceServiceDeleteOK {
	return &GitRepositoryResourceServiceDeleteOK{}
}

/*
GitRepositoryResourceServiceDeleteOK describes a response with status code 200, with default header values.

A successful response.
*/
type GitRepositoryResourceServiceDeleteOK struct {
	Payload *models.ClusterNamespaceFluxcdGitrepositoryDeleteGitRepositoryResponse
}

// IsSuccess returns true when this git repository resource service delete o k response has a 2xx status code
func (o *GitRepositoryResourceServiceDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this git repository resource service delete o k response has a 3xx status code
func (o *GitRepositoryResourceServiceDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this git repository resource service delete o k response has a 4xx status code
func (o *GitRepositoryResourceServiceDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this git repository resource service delete o k response has a 5xx status code
func (o *GitRepositoryResourceServiceDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this git repository resource service delete o k response a status code equal to that given
func (o *GitRepositoryResourceServiceDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *GitRepositoryResourceServiceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/fluxcd/gitrepositories/{fullName.name}][%d] gitRepositoryResourceServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *GitRepositoryResourceServiceDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/fluxcd/gitrepositories/{fullName.name}][%d] gitRepositoryResourceServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *GitRepositoryResourceServiceDeleteOK) GetPayload() *models.ClusterNamespaceFluxcdGitrepositoryDeleteGitRepositoryResponse {
	return o.Payload
}

func (o *GitRepositoryResourceServiceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterNamespaceFluxcdGitrepositoryDeleteGitRepositoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGitRepositoryResourceServiceDeleteDefault creates a GitRepositoryResourceServiceDeleteDefault with default headers values
func NewGitRepositoryResourceServiceDeleteDefault(code int) *GitRepositoryResourceServiceDeleteDefault {
	return &GitRepositoryResourceServiceDeleteDefault{
		_statusCode: code,
	}
}

/*
GitRepositoryResourceServiceDeleteDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GitRepositoryResourceServiceDeleteDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the git repository resource service delete default response
func (o *GitRepositoryResourceServiceDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this git repository resource service delete default response has a 2xx status code
func (o *GitRepositoryResourceServiceDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this git repository resource service delete default response has a 3xx status code
func (o *GitRepositoryResourceServiceDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this git repository resource service delete default response has a 4xx status code
func (o *GitRepositoryResourceServiceDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this git repository resource service delete default response has a 5xx status code
func (o *GitRepositoryResourceServiceDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this git repository resource service delete default response a status code equal to that given
func (o *GitRepositoryResourceServiceDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GitRepositoryResourceServiceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/fluxcd/gitrepositories/{fullName.name}][%d] GitRepositoryResourceService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *GitRepositoryResourceServiceDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /v1alpha1/clusters/{fullName.clusterName}/namespaces/{fullName.namespaceName}/fluxcd/gitrepositories/{fullName.name}][%d] GitRepositoryResourceService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *GitRepositoryResourceServiceDeleteDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *GitRepositoryResourceServiceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
