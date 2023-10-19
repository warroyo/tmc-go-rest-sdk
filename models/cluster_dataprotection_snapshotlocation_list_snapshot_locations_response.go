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

// ClusterDataprotectionSnapshotlocationListSnapshotLocationsResponse Response from listing SnapshotLocations.
//
// swagger:model cluster.dataprotection.snapshotlocation.ListSnapshotLocationsResponse
type ClusterDataprotectionSnapshotlocationListSnapshotLocationsResponse struct {

	// List of snapshotlocations.
	SnapshotLocations []*ClusterDataprotectionSnapshotlocationSnapshotLocation `json:"snapshotLocations"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// Validate validates this cluster dataprotection snapshotlocation list snapshot locations response
func (m *ClusterDataprotectionSnapshotlocationListSnapshotLocationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSnapshotLocations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDataprotectionSnapshotlocationListSnapshotLocationsResponse) validateSnapshotLocations(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotLocations) { // not required
		return nil
	}

	for i := 0; i < len(m.SnapshotLocations); i++ {
		if swag.IsZero(m.SnapshotLocations[i]) { // not required
			continue
		}

		if m.SnapshotLocations[i] != nil {
			if err := m.SnapshotLocations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshotLocations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("snapshotLocations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cluster dataprotection snapshotlocation list snapshot locations response based on the context it is used
func (m *ClusterDataprotectionSnapshotlocationListSnapshotLocationsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSnapshotLocations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDataprotectionSnapshotlocationListSnapshotLocationsResponse) contextValidateSnapshotLocations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SnapshotLocations); i++ {

		if m.SnapshotLocations[i] != nil {
			if err := m.SnapshotLocations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshotLocations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("snapshotLocations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterDataprotectionSnapshotlocationListSnapshotLocationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterDataprotectionSnapshotlocationListSnapshotLocationsResponse) UnmarshalBinary(b []byte) error {
	var res ClusterDataprotectionSnapshotlocationListSnapshotLocationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}