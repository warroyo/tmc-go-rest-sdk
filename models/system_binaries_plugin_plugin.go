// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SystemBinariesPluginPlugin Plugin denotes a Tanzu cli plugin.
//
// swagger:model system.binaries.plugin.Plugin
type SystemBinariesPluginPlugin struct {

	// Artifacts contains available artifacts for every supported version.
	Artifacts interface{} `json:"artifacts,omitempty"`

	// Description is the plugin's description.
	Description string `json:"description,omitempty"`

	// Name is the name of the plugin.
	Name string `json:"name,omitempty"`

	// Optional specifies whether the plugin is mandatory or optional.
	Optional bool `json:"optional,omitempty"`

	// Recommended version that Tanzu CLI should use if available.
	// The value should be a valid semantic version as defined in
	// https://semver.org/. E.g., 2.0.1
	RecommendedVersion string `json:"recommendedVersion,omitempty"`

	// Target specifies target of the plugin.
	Target string `json:"target,omitempty"`
}

// Validate validates this system binaries plugin plugin
func (m *SystemBinariesPluginPlugin) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this system binaries plugin plugin based on context it is used
func (m *SystemBinariesPluginPlugin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemBinariesPluginPlugin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemBinariesPluginPlugin) UnmarshalBinary(b []byte) error {
	var res SystemBinariesPluginPlugin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}