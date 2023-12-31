// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// K8sIoApimachineryPkgApisMetaV1ManagedFieldsEntry ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource
// that the fieldset applies to.
//
// swagger:model k8s.io.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry
type K8sIoApimachineryPkgApisMetaV1ManagedFieldsEntry struct {

	// APIVersion defines the version of this resource that this field set
	// applies to. The format is "group/version" just like the top-level
	// APIVersion field. It is necessary to track the version of a field
	// set because it cannot be automatically converted.
	APIVersion string `json:"apiVersion,omitempty"`

	// Fields identifies a set of fields.
	// +optional
	Fields *K8sIoApimachineryPkgApisMetaV1Fields `json:"fields,omitempty"`

	// Manager is an identifier of the workflow managing these fields.
	Manager string `json:"manager,omitempty"`

	// Operation is the type of operation which lead to this ManagedFieldsEntry being created.
	// The only valid values for this field are 'Apply' and 'Update'.
	Operation string `json:"operation,omitempty"`

	// Time is timestamp of when these fields were set. It should always be empty if Operation is 'Apply'
	// +optional
	Time *K8sIoApimachineryPkgApisMetaV1Time `json:"time,omitempty"`
}

// Validate validates this k8s io apimachinery pkg apis meta v1 managed fields entry
func (m *K8sIoApimachineryPkgApisMetaV1ManagedFieldsEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoApimachineryPkgApisMetaV1ManagedFieldsEntry) validateFields(formats strfmt.Registry) error {
	if swag.IsZero(m.Fields) { // not required
		return nil
	}

	if m.Fields != nil {
		if err := m.Fields.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fields")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fields")
			}
			return err
		}
	}

	return nil
}

func (m *K8sIoApimachineryPkgApisMetaV1ManagedFieldsEntry) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if m.Time != nil {
		if err := m.Time.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("time")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("time")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this k8s io apimachinery pkg apis meta v1 managed fields entry based on the context it is used
func (m *K8sIoApimachineryPkgApisMetaV1ManagedFieldsEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoApimachineryPkgApisMetaV1ManagedFieldsEntry) contextValidateFields(ctx context.Context, formats strfmt.Registry) error {

	if m.Fields != nil {
		if err := m.Fields.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fields")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fields")
			}
			return err
		}
	}

	return nil
}

func (m *K8sIoApimachineryPkgApisMetaV1ManagedFieldsEntry) contextValidateTime(ctx context.Context, formats strfmt.Registry) error {

	if m.Time != nil {
		if err := m.Time.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("time")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("time")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoApimachineryPkgApisMetaV1ManagedFieldsEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoApimachineryPkgApisMetaV1ManagedFieldsEntry) UnmarshalBinary(b []byte) error {
	var res K8sIoApimachineryPkgApisMetaV1ManagedFieldsEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
