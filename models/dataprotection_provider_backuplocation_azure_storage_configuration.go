// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DataprotectionProviderBackuplocationAzureStorageConfiguration Azure specific storage configuration details.
//
// swagger:model dataprotection.provider.backuplocation.AzureStorageConfiguration
type DataprotectionProviderBackuplocationAzureStorageConfiguration struct {

	// Name of the resource group containing the storage account for this backup storage location.
	ResourceGroup string `json:"resourceGroup,omitempty"`

	// Name of the storage account for this backup storage location.
	StorageAccount string `json:"storageAccount,omitempty"`

	// Subscription ID under which all the resources are being managed in azure. Optional.
	SubscriptionID string `json:"subscriptionId,omitempty"`
}

// Validate validates this dataprotection provider backuplocation azure storage configuration
func (m *DataprotectionProviderBackuplocationAzureStorageConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dataprotection provider backuplocation azure storage configuration based on context it is used
func (m *DataprotectionProviderBackuplocationAzureStorageConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataprotectionProviderBackuplocationAzureStorageConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataprotectionProviderBackuplocationAzureStorageConfiguration) UnmarshalBinary(b []byte) error {
	var res DataprotectionProviderBackuplocationAzureStorageConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}