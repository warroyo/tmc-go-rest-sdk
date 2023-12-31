// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonTanzupackagePackageMetadataSpec Spec of the Package Metadata.
//
// swagger:model common.tanzupackage.PackageMetadataSpec
type CommonTanzupackagePackageMetadataSpec struct {

	// Categories of the Package Metadata.
	Categories []string `json:"categories"`

	// Display name of the Package Metadata.
	DisplayName string `json:"displayName,omitempty"`

	// Icon or logo of the Package Metadata in base64 encoded format. Image should be in .svg format.
	Icon string `json:"icon,omitempty"`

	// Long description of Package Metadata containing key features, common use cases etc.
	LongDescription string `json:"longDescription,omitempty"`

	// Personally identifiable details of Package Metadata maintainers.
	Maintainers []*CommonTanzupackageMaintainer `json:"maintainers"`

	// Name of the Package Metadata author.
	ProviderName string `json:"providerName,omitempty"`

	// Name of package repository to which this package metadata belongs.
	RepositoryName string `json:"repositoryName,omitempty"`

	// Short description of the Package Metadata.
	ShortDescription string `json:"shortDescription,omitempty"`

	// Includes information about how consumer can get support. ex: email, slack etc.
	SupportDescription string `json:"supportDescription,omitempty"`
}

// Validate validates this common tanzupackage package metadata spec
func (m *CommonTanzupackagePackageMetadataSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaintainers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonTanzupackagePackageMetadataSpec) validateMaintainers(formats strfmt.Registry) error {
	if swag.IsZero(m.Maintainers) { // not required
		return nil
	}

	for i := 0; i < len(m.Maintainers); i++ {
		if swag.IsZero(m.Maintainers[i]) { // not required
			continue
		}

		if m.Maintainers[i] != nil {
			if err := m.Maintainers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("maintainers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("maintainers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this common tanzupackage package metadata spec based on the context it is used
func (m *CommonTanzupackagePackageMetadataSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMaintainers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonTanzupackagePackageMetadataSpec) contextValidateMaintainers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Maintainers); i++ {

		if m.Maintainers[i] != nil {
			if err := m.Maintainers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("maintainers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("maintainers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonTanzupackagePackageMetadataSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonTanzupackagePackageMetadataSpec) UnmarshalBinary(b []byte) error {
	var res CommonTanzupackagePackageMetadataSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
