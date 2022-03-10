// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RegistrationInfo registration info
//
// swagger:model registration_info
type RegistrationInfo struct {

	// domain_name is the domain name where the stated orc8r can be accessed
	// Required: true
	DomainName *string `json:"domain_name"`

	// registration_token is a token for the operator to give to the AGW; it keys to logical and network ID
	// Required: true
	// Min Length: 4
	RegistrationToken *string `json:"registration_token"`

	// root_ca is a certificate that access gateways (AGW) can use to handshake and communicate with the stated orc8r
	// Required: true
	RootCa *string `json:"root_ca"`
}

// Validate validates this registration info
func (m *RegistrationInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomainName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistrationToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootCa(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistrationInfo) validateDomainName(formats strfmt.Registry) error {

	if err := validate.Required("domain_name", "body", m.DomainName); err != nil {
		return err
	}

	return nil
}

func (m *RegistrationInfo) validateRegistrationToken(formats strfmt.Registry) error {

	if err := validate.Required("registration_token", "body", m.RegistrationToken); err != nil {
		return err
	}

	if err := validate.MinLength("registration_token", "body", string(*m.RegistrationToken), 4); err != nil {
		return err
	}

	return nil
}

func (m *RegistrationInfo) validateRootCa(formats strfmt.Registry) error {

	if err := validate.Required("root_ca", "body", m.RootCa); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegistrationInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistrationInfo) UnmarshalBinary(b []byte) error {
	var res RegistrationInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
