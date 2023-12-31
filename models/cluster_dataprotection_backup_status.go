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

// ClusterDataprotectionBackupStatus Status of the backup resource.
//
// swagger:model cluster.dataprotection.backup.Status
type ClusterDataprotectionBackupStatus struct {

	// A list of available phases for backup object.
	AvailablePhases []*ClusterDataprotectionBackupStatusPhase `json:"availablePhases"`

	// The URL to download the backup logs.
	BackupLogsURL string `json:"backupLogsUrl,omitempty"`

	// The timestamp when a backup was completed.
	// Format: date-time
	CompletionTimestamp strfmt.DateTime `json:"completionTimestamp,omitempty"`

	// The conditions attached to this backup object.
	// The description of the conditions is as follows:
	// - "Scheduled" with status 'Unknown' indicates the backup request has not been applied to the cluster yet
	// - "Scheduled" with status 'False' indicates the request could not be forwarded to the cluster (e.g. intent generation failure)
	// - "Scheduled" with status 'True' and "Ready" with status 'Unknown' indicates the backup create / delete intent has been applied / deleted but not yet acted upon
	// - "Ready" with status 'True' indicates the the creation of data protection is complete
	// - "Ready" with status 'False' indicates the the creation of data protection is in error state.
	Conditions map[string]VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// The total number of attempted CSI volume snapshots for this backup.
	CsiVolumeSnapshotsAttempted int32 `json:"csiVolumeSnapshotsAttempted,omitempty"`

	// The total number of successfully completed CSI volume snapshots for this backup.
	CsiVolumeSnapshotsCompleted int32 `json:"csiVolumeSnapshotsCompleted,omitempty"`

	// The expiration time associated with this Backup object if it is eligible for garbage-collection.
	// Format: date-time
	Expiration strfmt.DateTime `json:"expiration,omitempty"`

	// The error that caused the entire backup to fail.
	FailureReason string `json:"failureReason,omitempty"`

	// The backup's format version, including major, minor, and patch version.
	FormatVersion string `json:"formatVersion,omitempty"`

	// This holds the status of ListBackupResources action which is triggered on the cluster once the backup is completed.
	// This will specify the state of that process such that the user is indicated, if the action to get the resources is processed or failed.
	GatherBackupResourcesStatus *ClusterDataprotectionBackupGatherBackupResourcesStatus `json:"gatherBackupResourcesStatus,omitempty"`

	// The resource generation the current status applies to.
	ObservedGeneration string `json:"observedGeneration,omitempty"`

	// The current state of the Backup.
	Phase *ClusterDataprotectionBackupStatusPhase `json:"phase,omitempty"`

	// Additional info about the phase.
	PhaseInfo string `json:"phaseInfo,omitempty"`

	// The backup progress so far.
	Progress *ClusterDataprotectionBackupProgress `json:"progress,omitempty"`

	// The URL to download the list of resources that were backed up.
	ResourceListURL string `json:"resourceListUrl,omitempty"`

	// Map of important resources kind with their names. Map<k8s_Kind, list of names>.
	// Currently the expected kinds in a backup are Namespace, PersistentVolumeClaim and PersistentVolume.
	// Other resources which are part of the backup will be saved in an appropriate format in TMC S3 bucket and will be accessible
	// to the user on demand.
	Resources map[string]ClusterDataprotectionBackupResources `json:"resources,omitempty"`

	// A list of all the volume backups attempted by restic.
	ResticBackupsAttempted []string `json:"resticBackupsAttempted"`

	// A list of all the volume backups completed by restic.
	ResticBackupsCompleted []string `json:"resticBackupsCompleted"`

	// The timestamp when a backup was started.
	// Format: date-time
	StartTimestamp strfmt.DateTime `json:"startTimestamp,omitempty"`

	// A list of all validation errors (if applicable).
	ValidationErrors []string `json:"validationErrors"`

	// Information about volumes backed up.
	VolumeBackups []*ClusterDataprotectionBackupVolumeBackup `json:"volumeBackups"`

	// The total number of attempted volume snapshots for this backup.
	VolumeSnapshotsAttempted int32 `json:"volumeSnapshotsAttempted,omitempty"`

	// The total number of successfully completed volume snapshots for this backup.
	VolumeSnapshotsCompleted int32 `json:"volumeSnapshotsCompleted,omitempty"`
}

// Validate validates this cluster dataprotection backup status
func (m *ClusterDataprotectionBackupStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailablePhases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompletionTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGatherBackupResourcesStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeBackups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDataprotectionBackupStatus) validateAvailablePhases(formats strfmt.Registry) error {
	if swag.IsZero(m.AvailablePhases) { // not required
		return nil
	}

	for i := 0; i < len(m.AvailablePhases); i++ {
		if swag.IsZero(m.AvailablePhases[i]) { // not required
			continue
		}

		if m.AvailablePhases[i] != nil {
			if err := m.AvailablePhases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("availablePhases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("availablePhases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterDataprotectionBackupStatus) validateCompletionTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.CompletionTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("completionTimestamp", "body", "date-time", m.CompletionTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClusterDataprotectionBackupStatus) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for k := range m.Conditions {

		if err := validate.Required("conditions"+"."+k, "body", m.Conditions[k]); err != nil {
			return err
		}
		if val, ok := m.Conditions[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterDataprotectionBackupStatus) validateExpiration(formats strfmt.Registry) error {
	if swag.IsZero(m.Expiration) { // not required
		return nil
	}

	if err := validate.FormatOf("expiration", "body", "date-time", m.Expiration.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClusterDataprotectionBackupStatus) validateGatherBackupResourcesStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.GatherBackupResourcesStatus) { // not required
		return nil
	}

	if m.GatherBackupResourcesStatus != nil {
		if err := m.GatherBackupResourcesStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gatherBackupResourcesStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gatherBackupResourcesStatus")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterDataprotectionBackupStatus) validatePhase(formats strfmt.Registry) error {
	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	if m.Phase != nil {
		if err := m.Phase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterDataprotectionBackupStatus) validateProgress(formats strfmt.Registry) error {
	if swag.IsZero(m.Progress) { // not required
		return nil
	}

	if m.Progress != nil {
		if err := m.Progress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("progress")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterDataprotectionBackupStatus) validateResources(formats strfmt.Registry) error {
	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	for k := range m.Resources {

		if err := validate.Required("resources"+"."+k, "body", m.Resources[k]); err != nil {
			return err
		}
		if val, ok := m.Resources[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resources" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterDataprotectionBackupStatus) validateStartTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("startTimestamp", "body", "date-time", m.StartTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClusterDataprotectionBackupStatus) validateVolumeBackups(formats strfmt.Registry) error {
	if swag.IsZero(m.VolumeBackups) { // not required
		return nil
	}

	for i := 0; i < len(m.VolumeBackups); i++ {
		if swag.IsZero(m.VolumeBackups[i]) { // not required
			continue
		}

		if m.VolumeBackups[i] != nil {
			if err := m.VolumeBackups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeBackups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumeBackups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cluster dataprotection backup status based on the context it is used
func (m *ClusterDataprotectionBackupStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAvailablePhases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGatherBackupResourcesStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProgress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumeBackups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDataprotectionBackupStatus) contextValidateAvailablePhases(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AvailablePhases); i++ {

		if m.AvailablePhases[i] != nil {
			if err := m.AvailablePhases[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("availablePhases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("availablePhases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterDataprotectionBackupStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Conditions {

		if val, ok := m.Conditions[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ClusterDataprotectionBackupStatus) contextValidateGatherBackupResourcesStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.GatherBackupResourcesStatus != nil {
		if err := m.GatherBackupResourcesStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gatherBackupResourcesStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gatherBackupResourcesStatus")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterDataprotectionBackupStatus) contextValidatePhase(ctx context.Context, formats strfmt.Registry) error {

	if m.Phase != nil {
		if err := m.Phase.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterDataprotectionBackupStatus) contextValidateProgress(ctx context.Context, formats strfmt.Registry) error {

	if m.Progress != nil {
		if err := m.Progress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("progress")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterDataprotectionBackupStatus) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Resources {

		if val, ok := m.Resources[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ClusterDataprotectionBackupStatus) contextValidateVolumeBackups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VolumeBackups); i++ {

		if m.VolumeBackups[i] != nil {
			if err := m.VolumeBackups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeBackups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumeBackups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterDataprotectionBackupStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterDataprotectionBackupStatus) UnmarshalBinary(b []byte) error {
	var res ClusterDataprotectionBackupStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
