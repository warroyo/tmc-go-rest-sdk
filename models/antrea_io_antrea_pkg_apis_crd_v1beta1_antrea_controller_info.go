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

// AntreaIoAntreaPkgApisCrdV1beta1AntreaControllerInfo antrea io antrea pkg apis crd v1beta1 antrea controller info
//
// swagger:model antrea_io.antrea.pkg.apis.crd.v1beta1.AntreaControllerInfo
type AntreaIoAntreaPkgApisCrdV1beta1AntreaControllerInfo struct {

	// Controller condition contains types like ControllerHealthy
	APIPort string `json:"apiPort,omitempty"`

	// Antrea Controller NetworkPolicy information
	ConnectedAgentNum int32 `json:"connectedAgentNum,omitempty"`

	// Number of agents which are connected to this controller
	ControllerConditions []*AntreaIoAntreaPkgApisCrdV1beta1ControllerCondition `json:"controllerConditions"`

	// metadata
	Metadata *K8sIoApimachineryPkgApisMetaV1ObjectMeta `json:"metadata,omitempty"`

	// Antrea Controller Service
	NetworkPolicyControllerInfo *AntreaIoAntreaPkgApisCrdV1beta1NetworkPolicyControllerInfo `json:"networkPolicyControllerInfo,omitempty"`

	// The Pod that Antrea Controller is running in
	NodeRef *K8sIoAPICoreV1ObjectReference `json:"nodeRef,omitempty"`

	// Antrea binary version
	PodRef *K8sIoAPICoreV1ObjectReference `json:"podRef,omitempty"`

	// The Node that Antrea Controller is running in
	ServiceRef *K8sIoAPICoreV1ObjectReference `json:"serviceRef,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this antrea io antrea pkg apis crd v1beta1 antrea controller info
func (m *AntreaIoAntreaPkgApisCrdV1beta1AntreaControllerInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateControllerConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkPolicyControllerInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntreaIoAntreaPkgApisCrdV1beta1AntreaControllerInfo) validateControllerConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.ControllerConditions) { // not required
		return nil
	}

	for i := 0; i < len(m.ControllerConditions); i++ {
		if swag.IsZero(m.ControllerConditions[i]) { // not required
			continue
		}

		if m.ControllerConditions[i] != nil {
			if err := m.ControllerConditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("controllerConditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("controllerConditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AntreaIoAntreaPkgApisCrdV1beta1AntreaControllerInfo) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *AntreaIoAntreaPkgApisCrdV1beta1AntreaControllerInfo) validateNetworkPolicyControllerInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkPolicyControllerInfo) { // not required
		return nil
	}

	if m.NetworkPolicyControllerInfo != nil {
		if err := m.NetworkPolicyControllerInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkPolicyControllerInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkPolicyControllerInfo")
			}
			return err
		}
	}

	return nil
}

func (m *AntreaIoAntreaPkgApisCrdV1beta1AntreaControllerInfo) validateNodeRef(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeRef) { // not required
		return nil
	}

	if m.NodeRef != nil {
		if err := m.NodeRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeRef")
			}
			return err
		}
	}

	return nil
}

func (m *AntreaIoAntreaPkgApisCrdV1beta1AntreaControllerInfo) validatePodRef(formats strfmt.Registry) error {
	if swag.IsZero(m.PodRef) { // not required
		return nil
	}

	if m.PodRef != nil {
		if err := m.PodRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("podRef")
			}
			return err
		}
	}

	return nil
}

func (m *AntreaIoAntreaPkgApisCrdV1beta1AntreaControllerInfo) validateServiceRef(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceRef) { // not required
		return nil
	}

	if m.ServiceRef != nil {
		if err := m.ServiceRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serviceRef")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this antrea io antrea pkg apis crd v1beta1 antrea controller info based on the context it is used
func (m *AntreaIoAntreaPkgApisCrdV1beta1AntreaControllerInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateControllerConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkPolicyControllerInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePodRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntreaIoAntreaPkgApisCrdV1beta1AntreaControllerInfo) contextValidateControllerConditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ControllerConditions); i++ {

		if m.ControllerConditions[i] != nil {
			if err := m.ControllerConditions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("controllerConditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("controllerConditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AntreaIoAntreaPkgApisCrdV1beta1AntreaControllerInfo) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *AntreaIoAntreaPkgApisCrdV1beta1AntreaControllerInfo) contextValidateNetworkPolicyControllerInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkPolicyControllerInfo != nil {
		if err := m.NetworkPolicyControllerInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkPolicyControllerInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkPolicyControllerInfo")
			}
			return err
		}
	}

	return nil
}

func (m *AntreaIoAntreaPkgApisCrdV1beta1AntreaControllerInfo) contextValidateNodeRef(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeRef != nil {
		if err := m.NodeRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeRef")
			}
			return err
		}
	}

	return nil
}

func (m *AntreaIoAntreaPkgApisCrdV1beta1AntreaControllerInfo) contextValidatePodRef(ctx context.Context, formats strfmt.Registry) error {

	if m.PodRef != nil {
		if err := m.PodRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("podRef")
			}
			return err
		}
	}

	return nil
}

func (m *AntreaIoAntreaPkgApisCrdV1beta1AntreaControllerInfo) contextValidateServiceRef(ctx context.Context, formats strfmt.Registry) error {

	if m.ServiceRef != nil {
		if err := m.ServiceRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serviceRef")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AntreaIoAntreaPkgApisCrdV1beta1AntreaControllerInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntreaIoAntreaPkgApisCrdV1beta1AntreaControllerInfo) UnmarshalBinary(b []byte) error {
	var res AntreaIoAntreaPkgApisCrdV1beta1AntreaControllerInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
