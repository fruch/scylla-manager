// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ColumnFamilyInfo column_family_info
//
// Information about column family
// swagger:model ColumnFamilyInfo
type ColumnFamilyInfo struct {

	// The column family
	Cf string `json:"cf,omitempty"`

	// The Keyspace
	Ks string `json:"ks,omitempty"`

	// The column family type
	Type string `json:"type,omitempty"`
}

// Validate validates this column family info
func (m *ColumnFamilyInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ColumnFamilyInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ColumnFamilyInfo) UnmarshalBinary(b []byte) error {
	var res ColumnFamilyInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
