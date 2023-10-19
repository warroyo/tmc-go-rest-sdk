// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterNamespaceTanzupackageRepositoryImgPkgBundleSpec Package Repository bundle is an image package bundle that holds Package CRs.
//
// swagger:model cluster.namespace.tanzupackage.repository.ImgPkgBundleSpec
type ClusterNamespaceTanzupackageRepositoryImgPkgBundleSpec struct {

	// Docker image url; unqualified, tagged, or digest references supported.
	Image string `json:"image,omitempty"`
}

// Validate validates this cluster namespace tanzupackage repository img pkg bundle spec
func (m *ClusterNamespaceTanzupackageRepositoryImgPkgBundleSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster namespace tanzupackage repository img pkg bundle spec based on context it is used
func (m *ClusterNamespaceTanzupackageRepositoryImgPkgBundleSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNamespaceTanzupackageRepositoryImgPkgBundleSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNamespaceTanzupackageRepositoryImgPkgBundleSpec) UnmarshalBinary(b []byte) error {
	var res ClusterNamespaceTanzupackageRepositoryImgPkgBundleSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
