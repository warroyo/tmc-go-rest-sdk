// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonScheduleSchedule Holds the schedule options for scheduling a task.
//
// swagger:model common.schedule.Schedule
type CommonScheduleSchedule struct {

	// A Cron expression defining when to run a task.
	Rate string `json:"rate,omitempty"`
}

// Validate validates this common schedule schedule
func (m *CommonScheduleSchedule) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this common schedule schedule based on context it is used
func (m *CommonScheduleSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonScheduleSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonScheduleSchedule) UnmarshalBinary(b []byte) error {
	var res CommonScheduleSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
