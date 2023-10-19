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

// K8sIoAPICoreV1NodeStatus NodeStatus is information about the current status of a node.
//
// swagger:model k8s.io.api.core.v1.NodeStatus
type K8sIoAPICoreV1NodeStatus struct {

	// List of addresses reachable to the node.
	// Queried from cloud provider, if available.
	// More info: https://kubernetes.io/docs/concepts/nodes/node/#addresses
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Addresses []*K8sIoAPICoreV1NodeAddress `json:"addresses"`

	// Allocatable represents the resources of a node that are available for scheduling.
	// Defaults to Capacity.
	// +optional
	Allocatable map[string]K8sIoApimachineryPkgAPIResourceQuantity `json:"allocatable,omitempty"`

	// Capacity represents the total resources of a node.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
	// +optional
	Capacity map[string]K8sIoApimachineryPkgAPIResourceQuantity `json:"capacity,omitempty"`

	// Conditions is an array of current observed node conditions.
	// More info: https://kubernetes.io/docs/concepts/nodes/node/#condition
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []*K8sIoAPICoreV1NodeCondition `json:"conditions"`

	// Status of the config assigned to the node via the dynamic Kubelet config feature.
	// +optional
	Config *K8sIoAPICoreV1NodeConfigStatus `json:"config,omitempty"`

	// Endpoints of daemons running on the Node.
	// +optional
	DaemonEndpoints *K8sIoAPICoreV1NodeDaemonEndpoints `json:"daemonEndpoints,omitempty"`

	// List of container images on this node
	// +optional
	Images []*K8sIoAPICoreV1ContainerImage `json:"images"`

	// Set of ids/uuids to uniquely identify the node.
	// More info: https://kubernetes.io/docs/concepts/nodes/node/#info
	// +optional
	NodeInfo *K8sIoAPICoreV1NodeSystemInfo `json:"nodeInfo,omitempty"`

	// NodePhase is the recently observed lifecycle phase of the node.
	// More info: https://kubernetes.io/docs/concepts/nodes/node/#phase
	// The field is never populated, and now is deprecated.
	// +optional
	Phase string `json:"phase,omitempty"`

	// List of volumes that are attached to the node.
	// +optional
	VolumesAttached []*K8sIoAPICoreV1AttachedVolume `json:"volumesAttached"`

	// List of attachable volumes in use (mounted) by the node.
	// +optional
	VolumesInUse []string `json:"volumesInUse"`
}

// Validate validates this k8s io api core v1 node status
func (m *K8sIoAPICoreV1NodeStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllocatable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDaemonEndpoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumesAttached(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1NodeStatus) validateAddresses(formats strfmt.Registry) error {
	if swag.IsZero(m.Addresses) { // not required
		return nil
	}

	for i := 0; i < len(m.Addresses); i++ {
		if swag.IsZero(m.Addresses[i]) { // not required
			continue
		}

		if m.Addresses[i] != nil {
			if err := m.Addresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *K8sIoAPICoreV1NodeStatus) validateAllocatable(formats strfmt.Registry) error {
	if swag.IsZero(m.Allocatable) { // not required
		return nil
	}

	for k := range m.Allocatable {

		if err := validate.Required("allocatable"+"."+k, "body", m.Allocatable[k]); err != nil {
			return err
		}
		if val, ok := m.Allocatable[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allocatable" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allocatable" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *K8sIoAPICoreV1NodeStatus) validateCapacity(formats strfmt.Registry) error {
	if swag.IsZero(m.Capacity) { // not required
		return nil
	}

	for k := range m.Capacity {

		if err := validate.Required("capacity"+"."+k, "body", m.Capacity[k]); err != nil {
			return err
		}
		if val, ok := m.Capacity[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("capacity" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("capacity" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *K8sIoAPICoreV1NodeStatus) validateConditions(formats strfmt.Registry) error {
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

func (m *K8sIoAPICoreV1NodeStatus) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *K8sIoAPICoreV1NodeStatus) validateDaemonEndpoints(formats strfmt.Registry) error {
	if swag.IsZero(m.DaemonEndpoints) { // not required
		return nil
	}

	if m.DaemonEndpoints != nil {
		if err := m.DaemonEndpoints.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("daemonEndpoints")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("daemonEndpoints")
			}
			return err
		}
	}

	return nil
}

func (m *K8sIoAPICoreV1NodeStatus) validateImages(formats strfmt.Registry) error {
	if swag.IsZero(m.Images) { // not required
		return nil
	}

	for i := 0; i < len(m.Images); i++ {
		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {
			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *K8sIoAPICoreV1NodeStatus) validateNodeInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeInfo) { // not required
		return nil
	}

	if m.NodeInfo != nil {
		if err := m.NodeInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeInfo")
			}
			return err
		}
	}

	return nil
}

func (m *K8sIoAPICoreV1NodeStatus) validateVolumesAttached(formats strfmt.Registry) error {
	if swag.IsZero(m.VolumesAttached) { // not required
		return nil
	}

	for i := 0; i < len(m.VolumesAttached); i++ {
		if swag.IsZero(m.VolumesAttached[i]) { // not required
			continue
		}

		if m.VolumesAttached[i] != nil {
			if err := m.VolumesAttached[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumesAttached" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumesAttached" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this k8s io api core v1 node status based on the context it is used
func (m *K8sIoAPICoreV1NodeStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAllocatable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCapacity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDaemonEndpoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumesAttached(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8sIoAPICoreV1NodeStatus) contextValidateAddresses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Addresses); i++ {

		if m.Addresses[i] != nil {
			if err := m.Addresses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *K8sIoAPICoreV1NodeStatus) contextValidateAllocatable(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Allocatable {

		if val, ok := m.Allocatable[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *K8sIoAPICoreV1NodeStatus) contextValidateCapacity(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Capacity {

		if val, ok := m.Capacity[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *K8sIoAPICoreV1NodeStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

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

func (m *K8sIoAPICoreV1NodeStatus) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {
		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *K8sIoAPICoreV1NodeStatus) contextValidateDaemonEndpoints(ctx context.Context, formats strfmt.Registry) error {

	if m.DaemonEndpoints != nil {
		if err := m.DaemonEndpoints.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("daemonEndpoints")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("daemonEndpoints")
			}
			return err
		}
	}

	return nil
}

func (m *K8sIoAPICoreV1NodeStatus) contextValidateImages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Images); i++ {

		if m.Images[i] != nil {
			if err := m.Images[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *K8sIoAPICoreV1NodeStatus) contextValidateNodeInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeInfo != nil {
		if err := m.NodeInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeInfo")
			}
			return err
		}
	}

	return nil
}

func (m *K8sIoAPICoreV1NodeStatus) contextValidateVolumesAttached(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VolumesAttached); i++ {

		if m.VolumesAttached[i] != nil {
			if err := m.VolumesAttached[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumesAttached" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumesAttached" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoAPICoreV1NodeStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoAPICoreV1NodeStatus) UnmarshalBinary(b []byte) error {
	var res K8sIoAPICoreV1NodeStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}