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

// GitRepositoryResourceServiceUpdateReader is a Reader for the GitRepositoryResourceServiceUpdate structure.
type GitRepositoryResourceServiceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GitRepositoryResourceServiceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGitRepositoryResourceServiceUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGitRepositoryResourceServiceUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGitRepositoryResourceServiceUpdateOK creates a GitRepositoryResourceServiceUpdateOK with default headers values
func NewGitRepositoryResourceServiceUpdateOK() *GitRepositoryResourceServiceUpdateOK {
	return &GitRepositoryResourceServiceUpdateOK{}
}

/*
GitRepositoryResourceServiceUpdateOK describes a response with status code 200, with default header values.

A successful response.
*/
type GitRepositoryResourceServiceUpdateOK struct {
	Payload *models.ClusterNamespaceFluxcdGitrepositoryUpdateGitRepositoryResponse
}

// IsSuccess returns true when this git repository resource service update o k response has a 2xx status code
func (o *GitRepositoryResourceServiceUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this git repository resource service update o k response has a 3xx status code
func (o *GitRepositoryResourceServiceUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this git repository resource service update o k response has a 4xx status code
func (o *GitRepositoryResourceServiceUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this git repository resource service update o k response has a 5xx status code
func (o *GitRepositoryResourceServiceUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this git repository resource service update o k response a status code equal to that given
func (o *GitRepositoryResourceServiceUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *GitRepositoryResourceServiceUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/clusters/{gitRepository.fullName.clusterName}/namespaces/{gitRepository.fullName.namespaceName}/fluxcd/gitrepositories/{gitRepository.fullName.name}][%d] gitRepositoryResourceServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *GitRepositoryResourceServiceUpdateOK) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/clusters/{gitRepository.fullName.clusterName}/namespaces/{gitRepository.fullName.namespaceName}/fluxcd/gitrepositories/{gitRepository.fullName.name}][%d] gitRepositoryResourceServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *GitRepositoryResourceServiceUpdateOK) GetPayload() *models.ClusterNamespaceFluxcdGitrepositoryUpdateGitRepositoryResponse {
	return o.Payload
}

func (o *GitRepositoryResourceServiceUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterNamespaceFluxcdGitrepositoryUpdateGitRepositoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGitRepositoryResourceServiceUpdateDefault creates a GitRepositoryResourceServiceUpdateDefault with default headers values
func NewGitRepositoryResourceServiceUpdateDefault(code int) *GitRepositoryResourceServiceUpdateDefault {
	return &GitRepositoryResourceServiceUpdateDefault{
		_statusCode: code,
	}
}

/*
GitRepositoryResourceServiceUpdateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GitRepositoryResourceServiceUpdateDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the git repository resource service update default response
func (o *GitRepositoryResourceServiceUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this git repository resource service update default response has a 2xx status code
func (o *GitRepositoryResourceServiceUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this git repository resource service update default response has a 3xx status code
func (o *GitRepositoryResourceServiceUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this git repository resource service update default response has a 4xx status code
func (o *GitRepositoryResourceServiceUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this git repository resource service update default response has a 5xx status code
func (o *GitRepositoryResourceServiceUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this git repository resource service update default response a status code equal to that given
func (o *GitRepositoryResourceServiceUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GitRepositoryResourceServiceUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1alpha1/clusters/{gitRepository.fullName.clusterName}/namespaces/{gitRepository.fullName.namespaceName}/fluxcd/gitrepositories/{gitRepository.fullName.name}][%d] GitRepositoryResourceService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *GitRepositoryResourceServiceUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /v1alpha1/clusters/{gitRepository.fullName.clusterName}/namespaces/{gitRepository.fullName.namespaceName}/fluxcd/gitrepositories/{gitRepository.fullName.name}][%d] GitRepositoryResourceService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *GitRepositoryResourceServiceUpdateDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *GitRepositoryResourceServiceUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
