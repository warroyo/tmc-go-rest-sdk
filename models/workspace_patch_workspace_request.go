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

// WorkspacePatchWorkspaceRequest Request to patch (partially update) a Workspace.
//
// swagger:model workspace.PatchWorkspaceRequest
type WorkspacePatchWorkspaceRequest struct {

	// Workspace to patch.
	FullName *WorkspaceFullName `json:"fullName,omitempty"`

	// List of operations to apply.
	Patch []*VmwareTanzuCoreV1alpha1PatchOperation `json:"patch"`
}

// Validate validates this workspace patch workspace request
func (m *WorkspacePatchWorkspaceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFullName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePatch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkspacePatchWorkspaceRequest) validateFullName(formats strfmt.Registry) error {
	if swag.IsZero(m.FullName) { // not required
		return nil
	}

	if m.FullName != nil {
		if err := m.FullName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fullName")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fullName")
			}
			return err
		}
	}

	return nil
}

func (m *WorkspacePatchWorkspaceRequest) validatePatch(formats strfmt.Registry) error {
	if swag.IsZero(m.Patch) { // not required
		return nil
	}

	for i := 0; i < len(m.Patch); i++ {
		if swag.IsZero(m.Patch[i]) { // not required
			continue
		}

		if m.Patch[i] != nil {
			if err := m.Patch[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("patch" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("patch" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this workspace patch workspace request based on the context it is used
func (m *WorkspacePatchWorkspaceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFullName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePatch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkspacePatchWorkspaceRequest) contextValidateFullName(ctx context.Context, formats strfmt.Registry) error {

	if m.FullName != nil {
		if err := m.FullName.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fullName")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fullName")
			}
			return err
		}
	}

	return nil
}

func (m *WorkspacePatchWorkspaceRequest) contextValidatePatch(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Patch); i++ {

		if m.Patch[i] != nil {
			if err := m.Patch[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("patch" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("patch" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkspacePatchWorkspaceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkspacePatchWorkspaceRequest) UnmarshalBinary(b []byte) error {
	var res WorkspacePatchWorkspaceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}