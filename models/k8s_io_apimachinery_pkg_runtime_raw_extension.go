// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// K8sIoApimachineryPkgRuntimeRawExtension RawExtension is used to hold extensions in external versions.
//
// To use this, make a field which has RawExtension as its type in your external, versioned
// struct, and Object in your internal struct. You also need to register your
// various plugin types.
//
// // Internal package:
//
//	type MyAPIObject struct {
//		runtime.TypeMeta `json:",inline"`
//		MyPlugin runtime.Object `json:"myPlugin"`
//	}
//
//	type PluginA struct {
//		AOption string `json:"aOption"`
//	}
//
// // External package:
//
//	type MyAPIObject struct {
//		runtime.TypeMeta `json:",inline"`
//		MyPlugin runtime.RawExtension `json:"myPlugin"`
//	}
//
//	type PluginA struct {
//		AOption string `json:"aOption"`
//	}
//
// // On the wire, the JSON will look something like this:
//
//	{
//		"kind":"MyAPIObject",
//		"apiVersion":"v1",
//		"myPlugin": {
//			"kind":"PluginA",
//			"aOption":"foo",
//		},
//	}
//
// So what happens? Decode first uses json or yaml to unmarshal the serialized data into
// your external MyAPIObject. That causes the raw JSON to be stored, but not unpacked.
// The next step is to copy (using pkg/conversion) into the internal struct. The runtime
// package's DefaultScheme has conversion functions installed which will unpack the
// JSON stored in RawExtension, turning it into the correct object type, and storing it
// in the Object. (TODO: In the case where the object is of an unknown type, a
// runtime.Unknown object will be created and stored.)
//
// +k8s:deepcopy-gen=true
// +protobuf=true
// +k8s:openapi-gen=true
//
// swagger:model k8s.io.apimachinery.pkg.runtime.RawExtension
type K8sIoApimachineryPkgRuntimeRawExtension struct {

	// Raw is the underlying serialization of this object.
	//
	// TODO: Determine how to detect ContentType and ContentEncoding of 'Raw' data.
	// Format: byte
	Raw strfmt.Base64 `json:"raw,omitempty"`
}

// Validate validates this k8s io apimachinery pkg runtime raw extension
func (m *K8sIoApimachineryPkgRuntimeRawExtension) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this k8s io apimachinery pkg runtime raw extension based on context it is used
func (m *K8sIoApimachineryPkgRuntimeRawExtension) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *K8sIoApimachineryPkgRuntimeRawExtension) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8sIoApimachineryPkgRuntimeRawExtension) UnmarshalBinary(b []byte) error {
	var res K8sIoApimachineryPkgRuntimeRawExtension
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
