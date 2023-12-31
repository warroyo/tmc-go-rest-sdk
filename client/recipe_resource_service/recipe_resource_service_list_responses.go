// Code generated by go-swagger; DO NOT EDIT.

package recipe_resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/warroyo/tmc-go-rest-sdk/models"
)

// RecipeResourceServiceListReader is a Reader for the RecipeResourceServiceList structure.
type RecipeResourceServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RecipeResourceServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRecipeResourceServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRecipeResourceServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRecipeResourceServiceListOK creates a RecipeResourceServiceListOK with default headers values
func NewRecipeResourceServiceListOK() *RecipeResourceServiceListOK {
	return &RecipeResourceServiceListOK{}
}

/*
RecipeResourceServiceListOK describes a response with status code 200, with default header values.

A successful response.
*/
type RecipeResourceServiceListOK struct {
	Payload *models.PolicyTypeRecipeListRecipesResponse
}

// IsSuccess returns true when this recipe resource service list o k response has a 2xx status code
func (o *RecipeResourceServiceListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this recipe resource service list o k response has a 3xx status code
func (o *RecipeResourceServiceListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this recipe resource service list o k response has a 4xx status code
func (o *RecipeResourceServiceListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this recipe resource service list o k response has a 5xx status code
func (o *RecipeResourceServiceListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this recipe resource service list o k response a status code equal to that given
func (o *RecipeResourceServiceListOK) IsCode(code int) bool {
	return code == 200
}

func (o *RecipeResourceServiceListOK) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/policy/types/{searchScope.typeName}/recipes][%d] recipeResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *RecipeResourceServiceListOK) String() string {
	return fmt.Sprintf("[GET /v1alpha1/policy/types/{searchScope.typeName}/recipes][%d] recipeResourceServiceListOK  %+v", 200, o.Payload)
}

func (o *RecipeResourceServiceListOK) GetPayload() *models.PolicyTypeRecipeListRecipesResponse {
	return o.Payload
}

func (o *RecipeResourceServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyTypeRecipeListRecipesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRecipeResourceServiceListDefault creates a RecipeResourceServiceListDefault with default headers values
func NewRecipeResourceServiceListDefault(code int) *RecipeResourceServiceListDefault {
	return &RecipeResourceServiceListDefault{
		_statusCode: code,
	}
}

/*
RecipeResourceServiceListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RecipeResourceServiceListDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the recipe resource service list default response
func (o *RecipeResourceServiceListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this recipe resource service list default response has a 2xx status code
func (o *RecipeResourceServiceListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this recipe resource service list default response has a 3xx status code
func (o *RecipeResourceServiceListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this recipe resource service list default response has a 4xx status code
func (o *RecipeResourceServiceListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this recipe resource service list default response has a 5xx status code
func (o *RecipeResourceServiceListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this recipe resource service list default response a status code equal to that given
func (o *RecipeResourceServiceListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *RecipeResourceServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /v1alpha1/policy/types/{searchScope.typeName}/recipes][%d] RecipeResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *RecipeResourceServiceListDefault) String() string {
	return fmt.Sprintf("[GET /v1alpha1/policy/types/{searchScope.typeName}/recipes][%d] RecipeResourceService_List default  %+v", o._statusCode, o.Payload)
}

func (o *RecipeResourceServiceListDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *RecipeResourceServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
