// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkNgcConfigs NextGeneration Core configuration for network
//
// swagger:model network_ngc_configs
type NetworkNgcConfigs struct {

	// List of SuciProfiles shall be configured only for 5G
	// Min Items: 1
	SuciProfiles []*SuciProfile `json:"suci_profiles"`
}

// Validate validates this network ngc configs
func (m *NetworkNgcConfigs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSuciProfiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkNgcConfigs) validateSuciProfiles(formats strfmt.Registry) error {
	if swag.IsZero(m.SuciProfiles) { // not required
		return nil
	}

	iSuciProfilesSize := int64(len(m.SuciProfiles))

	if err := validate.MinItems("suci_profiles", "body", iSuciProfilesSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.SuciProfiles); i++ {
		if swag.IsZero(m.SuciProfiles[i]) { // not required
			continue
		}

		if m.SuciProfiles[i] != nil {
			if err := m.SuciProfiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("suci_profiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("suci_profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this network ngc configs based on the context it is used
func (m *NetworkNgcConfigs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSuciProfiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkNgcConfigs) contextValidateSuciProfiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SuciProfiles); i++ {

		if m.SuciProfiles[i] != nil {
			if err := m.SuciProfiles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("suci_profiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("suci_profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkNgcConfigs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkNgcConfigs) UnmarshalBinary(b []byte) error {
	var res NetworkNgcConfigs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
