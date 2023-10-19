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

// K8sIoAPIAppsV1DeploymentSpec DeploymentSpec is the specification of the desired behavior of the Deployment.
//
// swagger:model k8s.io.api.apps.v1.DeploymentSpec
type K8sIoAPIAppsV1DeploymentSpec struct {

	// Minimum number of seconds for which a newly created pod should be ready
	// without any of its container crashing, for it to be considered available.
	// Defaults to 0 (pod will be considered available as soon as it is ready)
	// +optional
	MinReadySeconds int32 `json:"minReadySeconds,omitempty"`

	// Indicates that the deployment is paused.
	// +optional
	Paused bool `json:"paused,omitempty"`

	// The maximum time in seconds for a deployment to make progress before it
	// is considered to be failed. The deployment controller will continue to
	// process failed deployments and a condition with a ProgressDeadlineExceeded
	// reason will be surfaced in the deployment status. Note that progress will
	// not be estimated during the time a deployment is paused. Defaults to 600s.
	ProgressDeadlineSeconds int32 `json:"progressDeadlineSeconds,omitempty"`

	// Number of desired pods. This is a pointer to distinguish between explicit
	// zero and not specified. Defaults to 1.
	// +optional
	Replicas int32 `json:"replicas,omitempty"`

	// The number of old ReplicaSets to retain to allow rollback.
	// This is a pointer to distinguish between explicit zero and not specified.
	// Defaults to 10.
	// +optional
	RevisionHistoryLimit int32 `json:"revisionHistoryLimit,omitempty"`

	// Label selector for pods. Existing ReplicaSets whose pods are
	// selected by this will be the ones affected by this deployment.
	// It must match the pod template's labels.
	Selector *K8sIoApimachineryPkgApisMetaV1LabelSelector `json:"selector,omitempty"`

	// The deployment strategy to use to replace existing pods with new ones.
	// +optional
	// +patchStrategy=retainKeys
	Strategy *K8sIoAPIAppsV1DeploymentStrategy `json:"strategy,omitempty"`

	// Template describes the pods that will be created.
	Template *K8sIoAPICoreV1PodTemplateSpec `json:"template,omitempty"`
}

// Validate validates this k8s io api apps v1 deployment spec
func (m *K8sIoAPIAppsV1DeploymentSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStrategy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPIAppsV1DeploymentSpec) validateSelector(formats strfmt.Registry) error {
	if swag.IsZero(m.Selector) { // not required
		return nil
	}

	if m.Selector != nil {
		if err := m.Selector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selector")
			}
			return err
		}
	}

	return nil
}

func (m *K8sIoAPIAppsV1DeploymentSpec) validateStrategy(formats strfmt.Registry) error {
	if swag.IsZero(m.Strategy) { // not required
		return nil
	}

	if m.Strategy != nil {
		if err := m.Strategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("strategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("strategy")
			}
			return err
		}
	}

	return nil
}

func (m *K8sIoAPIAppsV1DeploymentSpec) validateTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.Template) { // not required
		return nil
	}

	if m.Template != nil {
		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this k8s io api apps v1 deployment spec based on the context it is used
func (m *K8sIoAPIAppsV1DeploymentSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStrategy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPIAppsV1DeploymentSpec) contextValidateSelector(ctx context.Context, formats strfmt.Registry) error {

	if m.Selector != nil {
		if err := m.Selector.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selector")
			}
			return err
		}
	}

	return nil
}

func (m *K8sIoAPIAppsV1DeploymentSpec) contextValidateStrategy(ctx context.Context, formats strfmt.Registry) error {

	if m.Strategy != nil {
		if err := m.Strategy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("strategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("strategy")
			}
			return err
		}
	}

	return nil
}

func (m *K8sIoAPIAppsV1DeploymentSpec) contextValidateTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.Template != nil {
		if err := m.Template.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoAPIAppsV1DeploymentSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPIAppsV1DeploymentSpec) UnmarshalBinary(b []byte) error {
	var res K8sIoAPIAppsV1DeploymentSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}