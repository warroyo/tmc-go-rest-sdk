// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AntreaIoAntreaPkgApisCrdV1alpha1Source Source describes the source spec of the traceflow.
//
// swagger:model antrea_io.antrea.pkg.apis.crd.v1alpha1.Source
type AntreaIoAntreaPkgApisCrdV1alpha1Source struct {

	// IP is the source IPv4 or IPv6 address. IP as the source is supported
	// only for live-traffic Traceflow.
	IP string `json:"ip,omitempty"`

	// Namespace is the source namespace.
	Namespace string `json:"namespace,omitempty"`

	// Pod is the source pod.
	Pod string `json:"pod,omitempty"`
}

// Validate validates this antrea io antrea pkg apis crd v1alpha1 source
func (m *AntreaIoAntreaPkgApisCrdV1alpha1Source) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this antrea io antrea pkg apis crd v1alpha1 source based on context it is used
func (m *AntreaIoAntreaPkgApisCrdV1alpha1Source) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AntreaIoAntreaPkgApisCrdV1alpha1Source) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntreaIoAntreaPkgApisCrdV1alpha1Source) UnmarshalBinary(b []byte) error {
	var res AntreaIoAntreaPkgApisCrdV1alpha1Source
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
