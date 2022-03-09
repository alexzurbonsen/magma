// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModeMapItem Item containing {mode, [plmnA, plmB], [imsi1, imsi2], [apnY, apnZ]}
//
// swagger:model modeMapItem
type ModeMapItem struct {

	// apn
	Apn string `json:"apn,omitempty"`

	// imsi range
	ImsiRange string `json:"imsi_range,omitempty"`

	// mode
	// Enum: [local_subscriber s8_subscriber]
	Mode string `json:"mode,omitempty"`

	// plmn
	// Max Length: 6
	// Min Length: 5
	// Pattern: ^(\d{5,6})$
	Plmn string `json:"plmn,omitempty"`
}

// Validate validates this mode map item
func (m *ModeMapItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlmn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var modeMapItemTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["local_subscriber","s8_subscriber"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modeMapItemTypeModePropEnum = append(modeMapItemTypeModePropEnum, v)
	}
}

const (

	// ModeMapItemModeLocalSubscriber captures enum value "local_subscriber"
	ModeMapItemModeLocalSubscriber string = "local_subscriber"

	// ModeMapItemModeS8Subscriber captures enum value "s8_subscriber"
	ModeMapItemModeS8Subscriber string = "s8_subscriber"
)

// prop value enum
func (m *ModeMapItem) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, modeMapItemTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ModeMapItem) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *ModeMapItem) validatePlmn(formats strfmt.Registry) error {

	if swag.IsZero(m.Plmn) { // not required
		return nil
	}

	if err := validate.MinLength("plmn", "body", string(m.Plmn), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("plmn", "body", string(m.Plmn), 6); err != nil {
		return err
	}

	if err := validate.Pattern("plmn", "body", string(m.Plmn), `^(\d{5,6})$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModeMapItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModeMapItem) UnmarshalBinary(b []byte) error {
	var res ModeMapItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
