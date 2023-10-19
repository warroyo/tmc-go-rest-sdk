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

// DataprotectionProviderBackuplocationPatchBackupLocationRequest Request to patch (partially update) a BackupLocation.
//
// swagger:model dataprotection.provider.backuplocation.PatchBackupLocationRequest
type DataprotectionProviderBackuplocationPatchBackupLocationRequest struct {

	// BackupLocation to patch.
	FullName *DataprotectionProviderBackuplocationFullName `json:"fullName,omitempty"`

	// List of operations to apply.
	Patch []*VmwareTanzuCoreV1alpha1PatchOperation `json:"patch"`
}

// Validate validates this dataprotection provider backuplocation patch backup location request
func (m *DataprotectionProviderBackuplocationPatchBackupLocationRequest) Validate(formats strfmt.Registry) error {
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

func (m *DataprotectionProviderBackuplocationPatchBackupLocationRequest) validateFullName(formats strfmt.Registry) error {
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

func (m *DataprotectionProviderBackuplocationPatchBackupLocationRequest) validatePatch(formats strfmt.Registry) error {
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

// ContextValidate validate this dataprotection provider backuplocation patch backup location request based on the context it is used
func (m *DataprotectionProviderBackuplocationPatchBackupLocationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *DataprotectionProviderBackuplocationPatchBackupLocationRequest) contextValidateFullName(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DataprotectionProviderBackuplocationPatchBackupLocationRequest) contextValidatePatch(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DataprotectionProviderBackuplocationPatchBackupLocationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataprotectionProviderBackuplocationPatchBackupLocationRequest) UnmarshalBinary(b []byte) error {
	var res DataprotectionProviderBackuplocationPatchBackupLocationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
