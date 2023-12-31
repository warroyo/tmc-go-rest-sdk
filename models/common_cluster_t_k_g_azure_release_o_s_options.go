// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonClusterTKGAzureReleaseOSOptions Azure specific os options per release.
//
// swagger:model common.cluster.TKGAzureReleaseOSOptions
type CommonClusterTKGAzureReleaseOSOptions struct {

	// OS arch supporting the release.
	Arch string `json:"arch,omitempty"`

	// OS name supporting the release.
	Name string `json:"name,omitempty"`

	// Terms of the release have been accepted or not.
	TermsAccepted bool `json:"termsAccepted,omitempty"`

	// OS version supporting the release.
	Version string `json:"version,omitempty"`
}

// Validate validates this common cluster t k g azure release o s options
func (m *CommonClusterTKGAzureReleaseOSOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this common cluster t k g azure release o s options based on context it is used
func (m *CommonClusterTKGAzureReleaseOSOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonClusterTKGAzureReleaseOSOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonClusterTKGAzureReleaseOSOptions) UnmarshalBinary(b []byte) error {
	var res CommonClusterTKGAzureReleaseOSOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
