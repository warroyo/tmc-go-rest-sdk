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

// ClusterObjectObjectCronJob CronJob Object.
//
// swagger:model cluster.object.ObjectCronJob
type ClusterObjectObjectCronJob struct {

	// CronJob object.
	ResourceCronJob *K8sIoAPIBatchV1beta1CronJob `json:"resourceCronJob,omitempty"`
}

// Validate validates this cluster object object cron job
func (m *ClusterObjectObjectCronJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceCronJob(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterObjectObjectCronJob) validateResourceCronJob(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceCronJob) { // not required
		return nil
	}

	if m.ResourceCronJob != nil {
		if err := m.ResourceCronJob.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceCronJob")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceCronJob")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster object object cron job based on the context it is used
func (m *ClusterObjectObjectCronJob) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResourceCronJob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterObjectObjectCronJob) contextValidateResourceCronJob(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceCronJob != nil {
		if err := m.ResourceCronJob.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceCronJob")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceCronJob")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterObjectObjectCronJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterObjectObjectCronJob) UnmarshalBinary(b []byte) error {
	var res ClusterObjectObjectCronJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
