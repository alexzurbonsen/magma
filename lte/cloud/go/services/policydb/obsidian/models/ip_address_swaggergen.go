// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IPAddress IP address
//
// swagger:model ip_address
type IPAddress struct {

	// address
	// Example: 192.168.0.1/24
	Address string `json:"address,omitempty"`

	// version
	// Enum: [IPv4 IPv6]
	Version string `json:"version,omitempty"`
}

// Validate validates this ip address
func (m *IPAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ipAddressTypeVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IPv4","IPv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipAddressTypeVersionPropEnum = append(ipAddressTypeVersionPropEnum, v)
	}
}

const (

	// IPAddressVersionIPV4 captures enum value "IPv4"
	IPAddressVersionIPV4 string = "IPv4"

	// IPAddressVersionIPV6 captures enum value "IPv6"
	IPAddressVersionIPV6 string = "IPv6"
)

// prop value enum
func (m *IPAddress) validateVersionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ipAddressTypeVersionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IPAddress) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.Version) { // not required
		return nil
	}

	// value enum
	if err := m.validateVersionEnum("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ip address based on context it is used
func (m *IPAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IPAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAddress) UnmarshalBinary(b []byte) error {
	var res IPAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
