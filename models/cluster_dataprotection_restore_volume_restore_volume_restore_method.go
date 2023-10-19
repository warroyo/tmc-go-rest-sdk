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

// ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethod The possible methods used to perform a volume restore.
//
//   - METHOD_UNSPECIFIED: Unspecified method.
//   - RESTIC: A file system restore using velero's restic integration.
//   - CSI_VOLUME_SNAPSHOT: A volume snapshot performed by the CSI driver.
//
// swagger:model cluster.dataprotection.restore.VolumeRestore.VolumeRestoreMethod
type ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethod string

func NewClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethod(value ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethod) *ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethod {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethod.
func (m ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethod) Pointer() *ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethod {
	return &m
}

const (

	// ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethodMETHODUNSPECIFIED captures enum value "METHOD_UNSPECIFIED"
	ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethodMETHODUNSPECIFIED ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethod = "METHOD_UNSPECIFIED"

	// ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethodRESTIC captures enum value "RESTIC"
	ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethodRESTIC ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethod = "RESTIC"

	// ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethodCSIVOLUMESNAPSHOT captures enum value "CSI_VOLUME_SNAPSHOT"
	ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethodCSIVOLUMESNAPSHOT ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethod = "CSI_VOLUME_SNAPSHOT"
)

// for schema
var clusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethodEnum []interface{}

func init() {
	var res []ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethod
	if err := json.Unmarshal([]byte(`["METHOD_UNSPECIFIED","RESTIC","CSI_VOLUME_SNAPSHOT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethodEnum = append(clusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethodEnum, v)
	}
}

func (m ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethod) validateClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethodEnum(path, location string, value ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethod) error {
	if err := validate.EnumCase(path, location, value, clusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethodEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this cluster dataprotection restore volume restore volume restore method
func (m ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethodEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this cluster dataprotection restore volume restore volume restore method based on context it is used
func (m ClusterDataprotectionRestoreVolumeRestoreVolumeRestoreMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}