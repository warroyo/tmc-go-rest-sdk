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

// CommonClusterTKGVsphereOptions TKG vSphere provider specific options.
//
// swagger:model common.cluster.TKGVsphereOptions
type CommonClusterTKGVsphereOptions struct {

	// The advanced configuration options for vSphere cloud provider.
	AdvancedConfigOptions []*CommonClusterAdvancedConfigOption `json:"advancedConfigOptions"`

	// AVI related config parameters.
	AviOptions *CommonClusterTKGVsphereAVIOptions `json:"aviOptions,omitempty"`

	// List of TKG vSphere options broken by datacenter.
	DatacenterOptions []*CommonClusterTKGVsphereDatacenterOptions `json:"datacenterOptions"`

	// TKG Instance Types
	// An ordered ascending list of sizes e.g. small, medium, large
	InstanceTypes []*CommonClusterTKGVsphereInstanceType `json:"instanceTypes"`

	// The minimum resources required for TKG vms.
	VMMinimums *CommonClusterTKGVsphereVMConfig `json:"vmMinimums,omitempty"`
}

// Validate validates this common cluster t k g vsphere options
func (m *CommonClusterTKGVsphereOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdvancedConfigOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAviOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatacenterOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMMinimums(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonClusterTKGVsphereOptions) validateAdvancedConfigOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.AdvancedConfigOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.AdvancedConfigOptions); i++ {
		if swag.IsZero(m.AdvancedConfigOptions[i]) { // not required
			continue
		}

		if m.AdvancedConfigOptions[i] != nil {
			if err := m.AdvancedConfigOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("advancedConfigOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("advancedConfigOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonClusterTKGVsphereOptions) validateAviOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.AviOptions) { // not required
		return nil
	}

	if m.AviOptions != nil {
		if err := m.AviOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aviOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aviOptions")
			}
			return err
		}
	}

	return nil
}

func (m *CommonClusterTKGVsphereOptions) validateDatacenterOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.DatacenterOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.DatacenterOptions); i++ {
		if swag.IsZero(m.DatacenterOptions[i]) { // not required
			continue
		}

		if m.DatacenterOptions[i] != nil {
			if err := m.DatacenterOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("datacenterOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("datacenterOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonClusterTKGVsphereOptions) validateInstanceTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.InstanceTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.InstanceTypes); i++ {
		if swag.IsZero(m.InstanceTypes[i]) { // not required
			continue
		}

		if m.InstanceTypes[i] != nil {
			if err := m.InstanceTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("instanceTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonClusterTKGVsphereOptions) validateVMMinimums(formats strfmt.Registry) error {
	if swag.IsZero(m.VMMinimums) { // not required
		return nil
	}

	if m.VMMinimums != nil {
		if err := m.VMMinimums.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmMinimums")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmMinimums")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this common cluster t k g vsphere options based on the context it is used
func (m *CommonClusterTKGVsphereOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdvancedConfigOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAviOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatacenterOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstanceTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMMinimums(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonClusterTKGVsphereOptions) contextValidateAdvancedConfigOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdvancedConfigOptions); i++ {

		if m.AdvancedConfigOptions[i] != nil {
			if err := m.AdvancedConfigOptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("advancedConfigOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("advancedConfigOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonClusterTKGVsphereOptions) contextValidateAviOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.AviOptions != nil {
		if err := m.AviOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aviOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aviOptions")
			}
			return err
		}
	}

	return nil
}

func (m *CommonClusterTKGVsphereOptions) contextValidateDatacenterOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DatacenterOptions); i++ {

		if m.DatacenterOptions[i] != nil {
			if err := m.DatacenterOptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("datacenterOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("datacenterOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonClusterTKGVsphereOptions) contextValidateInstanceTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InstanceTypes); i++ {

		if m.InstanceTypes[i] != nil {
			if err := m.InstanceTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("instanceTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonClusterTKGVsphereOptions) contextValidateVMMinimums(ctx context.Context, formats strfmt.Registry) error {

	if m.VMMinimums != nil {
		if err := m.VMMinimums.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmMinimums")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmMinimums")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonClusterTKGVsphereOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonClusterTKGVsphereOptions) UnmarshalBinary(b []byte) error {
	var res CommonClusterTKGVsphereOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
