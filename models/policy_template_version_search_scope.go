// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicyTemplateVersionSearchScope Scope to search by, any fields left empty will be considered all (*).
//
// swagger:model policy.template.version.SearchScope
type PolicyTemplateVersionSearchScope struct {

	// Scope search to the specified name; supports globbing; default (*).
	Name string `json:"name,omitempty"`

	// Scope search to the specified template_name; supports globbing; default (*).
	TemplateName string `json:"templateName,omitempty"`
}

// Validate validates this policy template version search scope
func (m *PolicyTemplateVersionSearchScope) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this policy template version search scope based on context it is used
func (m *PolicyTemplateVersionSearchScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyTemplateVersionSearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyTemplateVersionSearchScope) UnmarshalBinary(b []byte) error {
	var res PolicyTemplateVersionSearchScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
