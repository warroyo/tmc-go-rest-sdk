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

// ClusterObjectDataKind Enum for resource kind. We are not using all caps to match k8s kinds
//
// - KIND_UNSPECIFIED: KIND_UNSPECIFIED is the default kind.
//   - Node: Node Kind.
//   - Pod: Pod Kind.
//   - Deployment: Deployment Kind.
//   - Service: Service kind.
//   - ReplicaSet: ReplicaSet kind.
//   - ReplicationController: ReplicationController kind.
//   - DaemonSet: DaemonSet kind.
//   - StatefulSet: StatefulSet kind.
//   - CronJob: CronJob kind.
//   - Job: Job kind.
//   - Namespace: Namespace kind.
//   - CustomResource: CustomResource kind.
//
// swagger:model cluster.object.Data.Kind
type ClusterObjectDataKind string

func NewClusterObjectDataKind(value ClusterObjectDataKind) *ClusterObjectDataKind {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ClusterObjectDataKind.
func (m ClusterObjectDataKind) Pointer() *ClusterObjectDataKind {
	return &m
}

const (

	// ClusterObjectDataKindKINDUNSPECIFIED captures enum value "KIND_UNSPECIFIED"
	ClusterObjectDataKindKINDUNSPECIFIED ClusterObjectDataKind = "KIND_UNSPECIFIED"

	// ClusterObjectDataKindNode captures enum value "Node"
	ClusterObjectDataKindNode ClusterObjectDataKind = "Node"

	// ClusterObjectDataKindPod captures enum value "Pod"
	ClusterObjectDataKindPod ClusterObjectDataKind = "Pod"

	// ClusterObjectDataKindDeployment captures enum value "Deployment"
	ClusterObjectDataKindDeployment ClusterObjectDataKind = "Deployment"

	// ClusterObjectDataKindService captures enum value "Service"
	ClusterObjectDataKindService ClusterObjectDataKind = "Service"

	// ClusterObjectDataKindReplicaSet captures enum value "ReplicaSet"
	ClusterObjectDataKindReplicaSet ClusterObjectDataKind = "ReplicaSet"

	// ClusterObjectDataKindReplicationController captures enum value "ReplicationController"
	ClusterObjectDataKindReplicationController ClusterObjectDataKind = "ReplicationController"

	// ClusterObjectDataKindDaemonSet captures enum value "DaemonSet"
	ClusterObjectDataKindDaemonSet ClusterObjectDataKind = "DaemonSet"

	// ClusterObjectDataKindStatefulSet captures enum value "StatefulSet"
	ClusterObjectDataKindStatefulSet ClusterObjectDataKind = "StatefulSet"

	// ClusterObjectDataKindCronJob captures enum value "CronJob"
	ClusterObjectDataKindCronJob ClusterObjectDataKind = "CronJob"

	// ClusterObjectDataKindJob captures enum value "Job"
	ClusterObjectDataKindJob ClusterObjectDataKind = "Job"

	// ClusterObjectDataKindNamespace captures enum value "Namespace"
	ClusterObjectDataKindNamespace ClusterObjectDataKind = "Namespace"

	// ClusterObjectDataKindCustomResource captures enum value "CustomResource"
	ClusterObjectDataKindCustomResource ClusterObjectDataKind = "CustomResource"
)

// for schema
var clusterObjectDataKindEnum []interface{}

func init() {
	var res []ClusterObjectDataKind
	if err := json.Unmarshal([]byte(`["KIND_UNSPECIFIED","Node","Pod","Deployment","Service","ReplicaSet","ReplicationController","DaemonSet","StatefulSet","CronJob","Job","Namespace","CustomResource"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterObjectDataKindEnum = append(clusterObjectDataKindEnum, v)
	}
}

func (m ClusterObjectDataKind) validateClusterObjectDataKindEnum(path, location string, value ClusterObjectDataKind) error {
	if err := validate.EnumCase(path, location, value, clusterObjectDataKindEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this cluster object data kind
func (m ClusterObjectDataKind) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateClusterObjectDataKindEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this cluster object data kind based on context it is used
func (m ClusterObjectDataKind) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
