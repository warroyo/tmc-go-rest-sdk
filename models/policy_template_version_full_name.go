// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicyTemplateVersionFullName Full name of the policy template version. This includes the object name along
// with any parents or further identifiers.
//
// swagger:model policy.template.version.FullName
type PolicyTemplateVersionFullName struct {

	// Name of policy template version.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Name of policy template.
	TemplateName string `json:"templateName,omitempty"`
}

// Validate validates this policy template version full name
func (m *PolicyTemplateVersionFullName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this policy template version full name based on context it is used
func (m *PolicyTemplateVersionFullName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyTemplateVersionFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyTemplateVersionFullName) UnmarshalBinary(b []byte) error {
	var res PolicyTemplateVersionFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
