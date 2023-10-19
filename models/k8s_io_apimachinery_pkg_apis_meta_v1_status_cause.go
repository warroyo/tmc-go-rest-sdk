// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// K8sIoApimachineryPkgApisMetaV1StatusCause StatusCause provides more information about an api.Status failure, including
// cases when multiple errors are encountered.
//
// swagger:model k8s.io.apimachinery.pkg.apis.meta.v1.StatusCause
type K8sIoApimachineryPkgApisMetaV1StatusCause struct {

	// The field of the resource that has caused this error, as named by its JSON
	// serialization. May include dot and postfix notation for nested attributes.
	// Arrays are zero-indexed.  Fields may appear more than once in an array of
	// causes due to fields having multiple errors.
	// Optional.
	//
	// Examples:
	//   "name" - the field "name" on the current resource
	//   "items[0].name" - the field "name" on the first array entry in "items"
	// +optional
	Field string `json:"field,omitempty"`

	// A human-readable description of the cause of the error.  This field may be
	// presented as-is to a reader.
	// +optional
	Message string `json:"message,omitempty"`

	// A machine-readable description of the cause of the error. If this value is
	// empty there is no information available.
	// +optional
	Reason string `json:"reason,omitempty"`
}

// Validate validates this k8s io apimachinery pkg apis meta v1 status cause
func (m *K8sIoApimachineryPkgApisMetaV1StatusCause) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this k8s io apimachinery pkg apis meta v1 status cause based on context it is used
func (m *K8sIoApimachineryPkgApisMetaV1StatusCause) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoApimachineryPkgApisMetaV1StatusCause) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoApimachineryPkgApisMetaV1StatusCause) UnmarshalBinary(b []byte) error {
	var res K8sIoApimachineryPkgApisMetaV1StatusCause
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
