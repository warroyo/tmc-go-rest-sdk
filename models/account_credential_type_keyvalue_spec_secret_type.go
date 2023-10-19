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

// AccountCredentialTypeKeyvalueSpecSecretType Type of Secret.
//
//   - SECRET_TYPE_UNSPECIFIED: SECRET_TYPE_UNSPECIFIED is default.
//   - OPAQUE_SECRET_TYPE: SECRET_TYPE_OPAQUE maps to the k8s secret type OPAQUE.
//
// It is arbitrary user-defined data.
//   - DOCKERCONFIGJSON_SECRET_TYPE: DOCKERCONFIGJSON_SECRET_TYPE maps to  Kubernetes secrets type kubernetes.io/dockerconfigjson.
//
// swagger:model account.credential.type.keyvalue.Spec.SecretType
type AccountCredentialTypeKeyvalueSpecSecretType string

func NewAccountCredentialTypeKeyvalueSpecSecretType(value AccountCredentialTypeKeyvalueSpecSecretType) *AccountCredentialTypeKeyvalueSpecSecretType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AccountCredentialTypeKeyvalueSpecSecretType.
func (m AccountCredentialTypeKeyvalueSpecSecretType) Pointer() *AccountCredentialTypeKeyvalueSpecSecretType {
	return &m
}

const (

	// AccountCredentialTypeKeyvalueSpecSecretTypeSECRETTYPEUNSPECIFIED captures enum value "SECRET_TYPE_UNSPECIFIED"
	AccountCredentialTypeKeyvalueSpecSecretTypeSECRETTYPEUNSPECIFIED AccountCredentialTypeKeyvalueSpecSecretType = "SECRET_TYPE_UNSPECIFIED"

	// AccountCredentialTypeKeyvalueSpecSecretTypeOPAQUESECRETTYPE captures enum value "OPAQUE_SECRET_TYPE"
	AccountCredentialTypeKeyvalueSpecSecretTypeOPAQUESECRETTYPE AccountCredentialTypeKeyvalueSpecSecretType = "OPAQUE_SECRET_TYPE"

	// AccountCredentialTypeKeyvalueSpecSecretTypeDOCKERCONFIGJSONSECRETTYPE captures enum value "DOCKERCONFIGJSON_SECRET_TYPE"
	AccountCredentialTypeKeyvalueSpecSecretTypeDOCKERCONFIGJSONSECRETTYPE AccountCredentialTypeKeyvalueSpecSecretType = "DOCKERCONFIGJSON_SECRET_TYPE"
)

// for schema
var accountCredentialTypeKeyvalueSpecSecretTypeEnum []interface{}

func init() {
	var res []AccountCredentialTypeKeyvalueSpecSecretType
	if err := json.Unmarshal([]byte(`["SECRET_TYPE_UNSPECIFIED","OPAQUE_SECRET_TYPE","DOCKERCONFIGJSON_SECRET_TYPE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountCredentialTypeKeyvalueSpecSecretTypeEnum = append(accountCredentialTypeKeyvalueSpecSecretTypeEnum, v)
	}
}

func (m AccountCredentialTypeKeyvalueSpecSecretType) validateAccountCredentialTypeKeyvalueSpecSecretTypeEnum(path, location string, value AccountCredentialTypeKeyvalueSpecSecretType) error {
	if err := validate.EnumCase(path, location, value, accountCredentialTypeKeyvalueSpecSecretTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this account credential type keyvalue spec secret type
func (m AccountCredentialTypeKeyvalueSpecSecretType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAccountCredentialTypeKeyvalueSpecSecretTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this account credential type keyvalue spec secret type based on context it is used
func (m AccountCredentialTypeKeyvalueSpecSecretType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}