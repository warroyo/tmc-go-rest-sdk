// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AntreaIoAntreaPkgApisCrdV1alpha1TCPHeader TCPHeader describes spec of a TCP header.
//
// swagger:model antrea_io.antrea.pkg.apis.crd.v1alpha1.TCPHeader
type AntreaIoAntreaPkgApisCrdV1alpha1TCPHeader struct {

	// DstPort is the destination port.
	DstPort int32 `json:"dstPort,omitempty"`

	// Flags are flags in the header.
	Flags int32 `json:"flags,omitempty"`

	// SrcPort is the source port.
	SrcPort int32 `json:"srcPort,omitempty"`
}

// Validate validates this antrea io antrea pkg apis crd v1alpha1 TCP header
func (m *AntreaIoAntreaPkgApisCrdV1alpha1TCPHeader) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this antrea io antrea pkg apis crd v1alpha1 TCP header based on context it is used
func (m *AntreaIoAntreaPkgApisCrdV1alpha1TCPHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AntreaIoAntreaPkgApisCrdV1alpha1TCPHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntreaIoAntreaPkgApisCrdV1alpha1TCPHeader) UnmarshalBinary(b []byte) error {
	var res AntreaIoAntreaPkgApisCrdV1alpha1TCPHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
