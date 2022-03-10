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

// FegNetwork Federation Network spec
//
// swagger:model feg_network
type FegNetwork struct {

	// description
	// Required: true
	Description NetworkDescription `json:"description"`

	// dns
	// Required: true
	DNS *NetworkDNSConfig `json:"dns"`

	// features
	Features *NetworkFeatures `json:"features,omitempty"`

	// federation
	// Required: true
	Federation *NetworkFederationConfigs `json:"federation"`

	// id
	// Required: true
	ID NetworkID `json:"id"`

	// name
	// Required: true
	Name NetworkName `json:"name"`

	// subscriber config
	SubscriberConfig *NetworkSubscriberConfig `json:"subscriber_config,omitempty"`
}

// Validate validates this feg network
func (m *FegNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFederation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FegNetwork) validateDescription(formats strfmt.Registry) error {

	if err := m.Description.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("description")
		}
		return err
	}

	return nil
}

func (m *FegNetwork) validateDNS(formats strfmt.Registry) error {

	if err := validate.Required("dns", "body", m.DNS); err != nil {
		return err
	}

	if m.DNS != nil {
		if err := m.DNS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns")
			}
			return err
		}
	}

	return nil
}

func (m *FegNetwork) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	if m.Features != nil {
		if err := m.Features.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features")
			}
			return err
		}
	}

	return nil
}

func (m *FegNetwork) validateFederation(formats strfmt.Registry) error {

	if err := validate.Required("federation", "body", m.Federation); err != nil {
		return err
	}

	if m.Federation != nil {
		if err := m.Federation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("federation")
			}
			return err
		}
	}

	return nil
}

func (m *FegNetwork) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *FegNetwork) validateName(formats strfmt.Registry) error {

	if err := m.Name.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("name")
		}
		return err
	}

	return nil
}

func (m *FegNetwork) validateSubscriberConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriberConfig) { // not required
		return nil
	}

	if m.SubscriberConfig != nil {
		if err := m.SubscriberConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subscriber_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FegNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FegNetwork) UnmarshalBinary(b []byte) error {
	var res FegNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
