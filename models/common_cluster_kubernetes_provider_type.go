// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// CommonClusterKubernetesProviderType KubernetesProviderType definition - indicates the k8s provider Type.
//
//   - KUBERNETES_PROVIDER_UNSPECIFIED: Unspecified k8s provider (default).
//   - VMWARE_TANZU_KUBERNETES_GRID: VMware Tanzu Kubernetes Grid.
//   - VMWARE_TANZU_KUBERNETES_GRID_SERVICE: VMware Tanzu Kubernetes Grid Service (Guest Cluster Management).
//   - VMWARE_TANZU_KUBERNETES_GRID_HOSTED: VMware Tanzu Kubernetes Grid hosted in TMC.
//   - VMWARE_TANZU_COMMUNITY_EDITION: VMware Tanzu Community Edition.
//   - VMWARE_MANAGED_KUBERNETES_SERVICE: VMware Managed Kubernetes Service.
//
// swagger:model common.cluster.KubernetesProviderType
type CommonClusterKubernetesProviderType string

func NewCommonClusterKubernetesProviderType(value CommonClusterKubernetesProviderType) *CommonClusterKubernetesProviderType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated CommonClusterKubernetesProviderType.
func (m CommonClusterKubernetesProviderType) Pointer() *CommonClusterKubernetesProviderType {
	return &m
}

const (

	// CommonClusterKubernetesProviderTypeKUBERNETESPROVIDERUNSPECIFIED captures enum value "KUBERNETES_PROVIDER_UNSPECIFIED"
	CommonClusterKubernetesProviderTypeKUBERNETESPROVIDERUNSPECIFIED CommonClusterKubernetesProviderType = "KUBERNETES_PROVIDER_UNSPECIFIED"

	// CommonClusterKubernetesProviderTypeVMWARETANZUKUBERNETESGRID captures enum value "VMWARE_TANZU_KUBERNETES_GRID"
	CommonClusterKubernetesProviderTypeVMWARETANZUKUBERNETESGRID CommonClusterKubernetesProviderType = "VMWARE_TANZU_KUBERNETES_GRID"

	// CommonClusterKubernetesProviderTypeVMWARETANZUKUBERNETESGRIDSERVICE captures enum value "VMWARE_TANZU_KUBERNETES_GRID_SERVICE"
	CommonClusterKubernetesProviderTypeVMWARETANZUKUBERNETESGRIDSERVICE CommonClusterKubernetesProviderType = "VMWARE_TANZU_KUBERNETES_GRID_SERVICE"

	// CommonClusterKubernetesProviderTypeVMWARETANZUKUBERNETESGRIDHOSTED captures enum value "VMWARE_TANZU_KUBERNETES_GRID_HOSTED"
	CommonClusterKubernetesProviderTypeVMWARETANZUKUBERNETESGRIDHOSTED CommonClusterKubernetesProviderType = "VMWARE_TANZU_KUBERNETES_GRID_HOSTED"

	// CommonClusterKubernetesProviderTypeVMWARETANZUCOMMUNITYEDITION captures enum value "VMWARE_TANZU_COMMUNITY_EDITION"
	CommonClusterKubernetesProviderTypeVMWARETANZUCOMMUNITYEDITION CommonClusterKubernetesProviderType = "VMWARE_TANZU_COMMUNITY_EDITION"

	// CommonClusterKubernetesProviderTypeVMWAREMANAGEDKUBERNETESSERVICE captures enum value "VMWARE_MANAGED_KUBERNETES_SERVICE"
	CommonClusterKubernetesProviderTypeVMWAREMANAGEDKUBERNETESSERVICE CommonClusterKubernetesProviderType = "VMWARE_MANAGED_KUBERNETES_SERVICE"
)

// for schema
var commonClusterKubernetesProviderTypeEnum []interface{}

func init() {
	var res []CommonClusterKubernetesProviderType
	if err := json.Unmarshal([]byte(`["KUBERNETES_PROVIDER_UNSPECIFIED","VMWARE_TANZU_KUBERNETES_GRID","VMWARE_TANZU_KUBERNETES_GRID_SERVICE","VMWARE_TANZU_KUBERNETES_GRID_HOSTED","VMWARE_TANZU_COMMUNITY_EDITION","VMWARE_MANAGED_KUBERNETES_SERVICE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		commonClusterKubernetesProviderTypeEnum = append(commonClusterKubernetesProviderTypeEnum, v)
	}
}

func (m CommonClusterKubernetesProviderType) validateCommonClusterKubernetesProviderTypeEnum(path, location string, value CommonClusterKubernetesProviderType) error {
	if err := validate.EnumCase(path, location, value, commonClusterKubernetesProviderTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this common cluster kubernetes provider type
func (m CommonClusterKubernetesProviderType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCommonClusterKubernetesProviderTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this common cluster kubernetes provider type based on context it is used
func (m CommonClusterKubernetesProviderType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}