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

// K8sIoAPIBatchV1JobStatus JobStatus represents the current state of a Job.
//
// swagger:model k8s.io.api.batch.v1.JobStatus
type K8sIoAPIBatchV1JobStatus struct {

	// The number of actively running pods.
	// +optional
	Active int32 `json:"active,omitempty"`

	// Represents time when the job was completed. It is not guaranteed to
	// be set in happens-before order across separate operations.
	// It is represented in RFC3339 form and is in UTC.
	// +optional
	CompletionTime *K8sIoApimachineryPkgApisMetaV1Time `json:"completionTime,omitempty"`

	// The latest available observations of an object's current state.
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []*K8sIoAPIBatchV1JobCondition `json:"conditions"`

	// The number of pods which reached phase Failed.
	// +optional
	Failed int32 `json:"failed,omitempty"`

	// Represents time when the job was acknowledged by the job controller.
	// It is not guaranteed to be set in happens-before order across separate operations.
	// It is represented in RFC3339 form and is in UTC.
	// +optional
	StartTime *K8sIoApimachineryPkgApisMetaV1Time `json:"startTime,omitempty"`

	// The number of pods which reached phase Succeeded.
	// +optional
	Succeeded int32 `json:"succeeded,omitempty"`
}

// Validate validates this k8s io api batch v1 job status
func (m *K8sIoAPIBatchV1JobStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPIBatchV1JobStatus) validateCompletionTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CompletionTime) { // not required
		return nil
	}

	if m.CompletionTime != nil {
		if err := m.CompletionTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("completionTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("completionTime")
			}
			return err
		}
	}

	return nil
}

func (m *K8sIoAPIBatchV1JobStatus) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *K8sIoAPIBatchV1JobStatus) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if m.StartTime != nil {
		if err := m.StartTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("startTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("startTime")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this k8s io api batch v1 job status based on the context it is used
func (m *K8sIoAPIBatchV1JobStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCompletionTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPIBatchV1JobStatus) contextValidateCompletionTime(ctx context.Context, formats strfmt.Registry) error {

	if m.CompletionTime != nil {
		if err := m.CompletionTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("completionTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("completionTime")
			}
			return err
		}
	}

	return nil
}

func (m *K8sIoAPIBatchV1JobStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Conditions); i++ {

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *K8sIoAPIBatchV1JobStatus) contextValidateStartTime(ctx context.Context, formats strfmt.Registry) error {

	if m.StartTime != nil {
		if err := m.StartTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("startTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("startTime")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoAPIBatchV1JobStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPIBatchV1JobStatus) UnmarshalBinary(b []byte) error {
	var res K8sIoAPIBatchV1JobStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
