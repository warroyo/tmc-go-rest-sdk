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

// ClusterObjectObjectJob Job Object.
//
// swagger:model cluster.object.ObjectJob
type ClusterObjectObjectJob struct {

	// Job object.
	ResourceJob *K8sIoAPIBatchV1Job `json:"resourceJob,omitempty"`
}

// Validate validates this cluster object object job
func (m *ClusterObjectObjectJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceJob(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterObjectObjectJob) validateResourceJob(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceJob) { // not required
		return nil
	}

	if m.ResourceJob != nil {
		if err := m.ResourceJob.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceJob")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceJob")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster object object job based on the context it is used
func (m *ClusterObjectObjectJob) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResourceJob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterObjectObjectJob) contextValidateResourceJob(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceJob != nil {
		if err := m.ResourceJob.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceJob")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceJob")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterObjectObjectJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterObjectObjectJob) UnmarshalBinary(b []byte) error {
	var res ClusterObjectObjectJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}