// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DataprotectionProviderBackuplocationS3Configuration AWS S3 or other S3-compatible storage configuration details.
//
// swagger:model dataprotection.provider.backuplocation.S3Configuration
type DataprotectionProviderBackuplocationS3Configuration struct {

	// The service endpoint used for generating download URLs. This field is primarily for local storage services like Minio.
	PublicURL string `json:"publicUrl,omitempty"`

	// A flag for whether to force path style URLs for S3 objects. It is default to false and set it to true when
	// using local storage service like Minio.
	S3ForcePathStyle bool `json:"s3ForcePathStyle,omitempty"`

	// The service endpoint for non-AWS S3 storage solution.
	S3URL string `json:"s3Url,omitempty"`
}

// Validate validates this dataprotection provider backuplocation s3 configuration
func (m *DataprotectionProviderBackuplocationS3Configuration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dataprotection provider backuplocation s3 configuration based on context it is used
func (m *DataprotectionProviderBackuplocationS3Configuration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataprotectionProviderBackuplocationS3Configuration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataprotectionProviderBackuplocationS3Configuration) UnmarshalBinary(b []byte) error {
	var res DataprotectionProviderBackuplocationS3Configuration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
