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

// CwfHaPair HA Gateway pair in a Carrier Wifi Network
//
// swagger:model cwf_ha_pair
type CwfHaPair struct {

	// config
	// Required: true
	Config *CwfHaPairConfigs `json:"config"`

	// gateway id 1
	// Required: true
	GatewayID1 string `json:"gateway_id_1"`

	// gateway id 2
	// Required: true
	GatewayID2 string `json:"gateway_id_2"`

	// ha pair id
	// Required: true
	HaPairID string `json:"ha_pair_id"`

	// state
	State *CarrierWifiHaPairState `json:"state,omitempty"`
}

// Validate validates this cwf ha pair
func (m *CwfHaPair) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGatewayID1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGatewayID2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHaPairID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CwfHaPair) validateConfig(formats strfmt.Registry) error {

	if err := validate.Required("config", "body", m.Config); err != nil {
		return err
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *CwfHaPair) validateGatewayID1(formats strfmt.Registry) error {

	if err := validate.RequiredString("gateway_id_1", "body", string(m.GatewayID1)); err != nil {
		return err
	}

	return nil
}

func (m *CwfHaPair) validateGatewayID2(formats strfmt.Registry) error {

	if err := validate.RequiredString("gateway_id_2", "body", string(m.GatewayID2)); err != nil {
		return err
	}

	return nil
}

func (m *CwfHaPair) validateHaPairID(formats strfmt.Registry) error {

	if err := validate.RequiredString("ha_pair_id", "body", string(m.HaPairID)); err != nil {
		return err
	}

	return nil
}

func (m *CwfHaPair) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CwfHaPair) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CwfHaPair) UnmarshalBinary(b []byte) error {
	var res CwfHaPair
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
