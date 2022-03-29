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

// PlmnConfig PLMN Configuration
//
// swagger:model plmnConfig
type PlmnConfig struct {

	// mcc
	// Example: 001
	// Required: true
	// Pattern: ^(\d{3})$
	Mcc string `json:"mcc"`

	// mnc
	// Example: 01
	// Required: true
	// Pattern: ^(\d{2,3})$
	Mnc string `json:"mnc"`
}

// Validate validates this plmn config
func (m *PlmnConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMcc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMnc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlmnConfig) validateMcc(formats strfmt.Registry) error {

	if err := validate.RequiredString("mcc", "body", m.Mcc); err != nil {
		return err
	}

	if err := validate.Pattern("mcc", "body", m.Mcc, `^(\d{3})$`); err != nil {
		return err
	}

	return nil
}

func (m *PlmnConfig) validateMnc(formats strfmt.Registry) error {

	if err := validate.RequiredString("mnc", "body", m.Mnc); err != nil {
		return err
	}

	if err := validate.Pattern("mnc", "body", m.Mnc, `^(\d{2,3})$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this plmn config based on context it is used
func (m *PlmnConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlmnConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlmnConfig) UnmarshalBinary(b []byte) error {
	var res PlmnConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
