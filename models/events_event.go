// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EventsEvent Event is any event that has occured within TMC.
// This is closely modeled after CloudEvents, however since there is no official
// protobuf specification; we define our own message.
//
// swagger:model events.Event
type EventsEvent struct {

	// All other context metadata attributes.
	Attributes map[string]interface{} `json:"attributes,omitempty"`

	// Domain specific information.
	Data interface{} `json:"data,omitempty"`

	// A globally unique event identifier (uuid).
	ID string `json:"id,omitempty"`

	// The source, preferably an absolute URI reference.
	// e.g. https://example.tmc.cloud.vmware.com/clusters/create?accountName=accname
	Source string `json:"source,omitempty"`

	// The subject is useful as a qualifier for multiple types of events coming from
	// the same source without resorting to parsing the data payload.
	// e.g. ClusterLifecycleEvent
	Subject string `json:"subject,omitempty"`

	// The timestamp of when the event occured according to the source (UTC).
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`

	// The type is similar to having an event namespace or package
	// it enables more validation on the expected contents of data.
	// e.g. com.vmware.tmc.audit
	Type string `json:"type,omitempty"`
}

// Validate validates this events event
func (m *EventsEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventsEvent) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this events event based on context it is used
func (m *EventsEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventsEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventsEvent) UnmarshalBinary(b []byte) error {
	var res EventsEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}