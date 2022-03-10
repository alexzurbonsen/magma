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

// GatewayDNSConfigs DNS configuration for a gateway
//
// swagger:model gateway_dns_configs
type GatewayDNSConfigs struct {

	// dhcp server enabled
	// Required: true
	DhcpServerEnabled *bool `json:"dhcp_server_enabled"`

	// enable caching
	// Required: true
	EnableCaching *bool `json:"enable_caching"`

	// local ttl
	// Required: true
	LocalTTL *int32 `json:"local_ttl"`

	// records
	Records GatewayDNSRecords `json:"records,omitempty"`
}

// Validate validates this gateway dns configs
func (m *GatewayDNSConfigs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDhcpServerEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnableCaching(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GatewayDNSConfigs) validateDhcpServerEnabled(formats strfmt.Registry) error {

	if err := validate.Required("dhcp_server_enabled", "body", m.DhcpServerEnabled); err != nil {
		return err
	}

	return nil
}

func (m *GatewayDNSConfigs) validateEnableCaching(formats strfmt.Registry) error {

	if err := validate.Required("enable_caching", "body", m.EnableCaching); err != nil {
		return err
	}

	return nil
}

func (m *GatewayDNSConfigs) validateLocalTTL(formats strfmt.Registry) error {

	if err := validate.Required("local_ttl", "body", m.LocalTTL); err != nil {
		return err
	}

	return nil
}

func (m *GatewayDNSConfigs) validateRecords(formats strfmt.Registry) error {

	if swag.IsZero(m.Records) { // not required
		return nil
	}

	if err := m.Records.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("records")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GatewayDNSConfigs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GatewayDNSConfigs) UnmarshalBinary(b []byte) error {
	var res GatewayDNSConfigs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
