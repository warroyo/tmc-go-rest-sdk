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

// AntreaIoAntreaPkgApisCrdV1alpha1TraceflowStatus TraceflowStatus describes current status of the traceflow.
//
// swagger:model antrea_io.antrea.pkg.apis.crd.v1alpha1.TraceflowStatus
type AntreaIoAntreaPkgApisCrdV1alpha1TraceflowStatus struct {

	// CapturedPacket is the captured packet in live-traffic Traceflow.
	CapturedPacket *AntreaIoAntreaPkgApisCrdV1alpha1Packet `json:"capturedPacket,omitempty"`

	// DataplaneTag is a tag to identify a traceflow session across Nodes.
	DataplaneTag int64 `json:"dataplaneTag,omitempty"`

	// Phase is the Traceflow phase.
	Phase string `json:"phase,omitempty"`

	// Reason is a message indicating the reason of the traceflow's current phase.
	Reason string `json:"reason,omitempty"`

	// Results is the collection of all observations on different nodes.
	Results []*AntreaIoAntreaPkgApisCrdV1alpha1NodeResult `json:"results"`

	// StartTime is the time at which the Traceflow as started by the Antrea Controller.
	// Before K8s v1.20, null values (field not set) are not pruned, and a CR where a
	// metav1.Time field is not set would fail OpenAPI validation (type string). The
	// recommendation seems to be to use a pointer instead, and the field will be omitted when
	// serializing.
	// See https://github.com/kubernetes/kubernetes/issues/86811
	// Format: date-time
	StartTime strfmt.DateTime `json:"startTime,omitempty"`
}

// Validate validates this antrea io antrea pkg apis crd v1alpha1 traceflow status
func (m *AntreaIoAntreaPkgApisCrdV1alpha1TraceflowStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapturedPacket(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntreaIoAntreaPkgApisCrdV1alpha1TraceflowStatus) validateCapturedPacket(formats strfmt.Registry) error {
	if swag.IsZero(m.CapturedPacket) { // not required
		return nil
	}

	if m.CapturedPacket != nil {
		if err := m.CapturedPacket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capturedPacket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capturedPacket")
			}
			return err
		}
	}

	return nil
}

func (m *AntreaIoAntreaPkgApisCrdV1alpha1TraceflowStatus) validateResults(formats strfmt.Registry) error {
	if swag.IsZero(m.Results) { // not required
		return nil
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AntreaIoAntreaPkgApisCrdV1alpha1TraceflowStatus) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this antrea io antrea pkg apis crd v1alpha1 traceflow status based on the context it is used
func (m *AntreaIoAntreaPkgApisCrdV1alpha1TraceflowStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapturedPacket(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntreaIoAntreaPkgApisCrdV1alpha1TraceflowStatus) contextValidateCapturedPacket(ctx context.Context, formats strfmt.Registry) error {

	if m.CapturedPacket != nil {
		if err := m.CapturedPacket.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capturedPacket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capturedPacket")
			}
			return err
		}
	}

	return nil
}

func (m *AntreaIoAntreaPkgApisCrdV1alpha1TraceflowStatus) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Results); i++ {

		if m.Results[i] != nil {
			if err := m.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AntreaIoAntreaPkgApisCrdV1alpha1TraceflowStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntreaIoAntreaPkgApisCrdV1alpha1TraceflowStatus) UnmarshalBinary(b []byte) error {
	var res AntreaIoAntreaPkgApisCrdV1alpha1TraceflowStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
