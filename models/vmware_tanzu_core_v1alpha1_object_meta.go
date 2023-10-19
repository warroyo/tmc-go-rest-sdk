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
	"github.com/go-openapi/validate"
)

// VmwareTanzuCoreV1alpha1ObjectMeta Holds general shared object metadatas.
//
// swagger:model vmware.tanzu.core.v1alpha1.object.Meta
type VmwareTanzuCoreV1alpha1ObjectMeta struct {

	// Annotations for the object. Annotations hold system level information provisioned by controllers.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Creation time of the object.
	// Format: date-time
	CreationTime strfmt.DateTime `json:"creationTime,omitempty"`

	// Description of the resource.
	Description string `json:"description,omitempty"`

	// Generation of the resource as specified by the user, increments on changes.
	Generation string `json:"generation,omitempty"`

	// Labels to apply to the object.
	Labels map[string]string `json:"labels,omitempty"`

	// Hard object references to parents of this resource.
	ParentReferences []*VmwareTanzuCoreV1alpha1ObjectReference `json:"parentReferences"`

	// A string that identifies the internal version of this object that can be used by clients to
	// determine when objects have changed. This value MUST be treated as opaque by clients and
	// passed unmodified back to the server.
	ResourceVersion string `json:"resourceVersion,omitempty"`

	// UID for the object.
	UID string `json:"uid,omitempty"`

	// Update time of the object.
	// Format: date-time
	UpdateTime strfmt.DateTime `json:"updateTime,omitempty"`
}

// Validate validates this vmware tanzu core v1alpha1 object meta
func (m *VmwareTanzuCoreV1alpha1ObjectMeta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentReferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmwareTanzuCoreV1alpha1ObjectMeta) validateCreationTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationTime) { // not required
		return nil
	}

	if err := validate.FormatOf("creationTime", "body", "date-time", m.CreationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VmwareTanzuCoreV1alpha1ObjectMeta) validateParentReferences(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentReferences) { // not required
		return nil
	}

	for i := 0; i < len(m.ParentReferences); i++ {
		if swag.IsZero(m.ParentReferences[i]) { // not required
			continue
		}

		if m.ParentReferences[i] != nil {
			if err := m.ParentReferences[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parentReferences" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parentReferences" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VmwareTanzuCoreV1alpha1ObjectMeta) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("updateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this vmware tanzu core v1alpha1 object meta based on the context it is used
func (m *VmwareTanzuCoreV1alpha1ObjectMeta) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParentReferences(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmwareTanzuCoreV1alpha1ObjectMeta) contextValidateParentReferences(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ParentReferences); i++ {

		if m.ParentReferences[i] != nil {
			if err := m.ParentReferences[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parentReferences" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parentReferences" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VmwareTanzuCoreV1alpha1ObjectMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuCoreV1alpha1ObjectMeta) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuCoreV1alpha1ObjectMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
