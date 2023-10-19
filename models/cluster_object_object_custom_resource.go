// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterObjectObjectCustomResource CustomResource Object.
//
// swagger:model cluster.object.ObjectCustomResource
type ClusterObjectObjectCustomResource struct {

	// CustomResource object.
	ResourceCustomResource interface{} `json:"resourceCustomResource,omitempty"`
}

// Validate validates this cluster object object custom resource
func (m *ClusterObjectObjectCustomResource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster object object custom resource based on context it is used
func (m *ClusterObjectObjectCustomResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterObjectObjectCustomResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterObjectObjectCustomResource) UnmarshalBinary(b []byte) error {
	var res ClusterObjectObjectCustomResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
