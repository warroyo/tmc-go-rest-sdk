// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// K8sIoApimachineryPkgApisMetaV1Time Time is a wrapper around time.Time which supports correct
// marshaling to YAML and JSON.  Wrappers are provided for many
// of the factory methods that the time package offers.
//
// +protobuf.options.marshal=false
// +protobuf.as=Timestamp
// +protobuf.options.(gogoproto.goproto_stringer)=false
//
// swagger:model k8s.io.apimachinery.pkg.apis.meta.v1.Time
type K8sIoApimachineryPkgApisMetaV1Time struct {

	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999,999,999
	// inclusive. This field may be limited in precision depending on context.
	Nanos int32 `json:"nanos,omitempty"`

	// Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds string `json:"seconds,omitempty"`
}

// Validate validates this k8s io apimachinery pkg apis meta v1 time
func (m *K8sIoApimachineryPkgApisMetaV1Time) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this k8s io apimachinery pkg apis meta v1 time based on context it is used
func (m *K8sIoApimachineryPkgApisMetaV1Time) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoApimachineryPkgApisMetaV1Time) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoApimachineryPkgApisMetaV1Time) UnmarshalBinary(b []byte) error {
	var res K8sIoApimachineryPkgApisMetaV1Time
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
