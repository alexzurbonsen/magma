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

// Subscriber subscriber
//
// swagger:model subscriber
type Subscriber struct {

	// active apns
	ActiveAPNS APNList `json:"active_apns,omitempty"`

	// active base names
	ActiveBaseNames BaseNames `json:"active_base_names,omitempty"`

	// active policies
	ActivePolicies PolicyIds `json:"active_policies,omitempty"`

	// active policies by apn
	ActivePoliciesByAPN PolicyIdsByAPN `json:"active_policies_by_apn,omitempty"`

	// config
	// Required: true
	Config *SubscriberConfig `json:"config"`

	// forbidden network types
	ForbiddenNetworkTypes CoreNetworkTypes `json:"forbidden_network_types,omitempty"`

	// id
	// Required: true
	ID SubscriberID `json:"id"`

	// lte
	// Required: true
	LTE *LTESubscription `json:"lte"`

	// monitoring
	Monitoring *SubscriberStatus `json:"monitoring,omitempty"`

	// msisdn
	Msisdn Msisdn `json:"msisdn,omitempty"`

	// Optional name associated with the subscriber
	Name string `json:"name,omitempty"`

	// state
	State *SubscriberState `json:"state,omitempty"`
}

// Validate validates this subscriber
func (m *Subscriber) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveAPNS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActiveBaseNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActivePolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActivePoliciesByAPN(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForbiddenNetworkTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLTE(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonitoring(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMsisdn(formats); err != nil {
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

func (m *Subscriber) validateActiveAPNS(formats strfmt.Registry) error {

	if swag.IsZero(m.ActiveAPNS) { // not required
		return nil
	}

	if err := m.ActiveAPNS.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("active_apns")
		}
		return err
	}

	return nil
}

func (m *Subscriber) validateActiveBaseNames(formats strfmt.Registry) error {

	if swag.IsZero(m.ActiveBaseNames) { // not required
		return nil
	}

	if err := m.ActiveBaseNames.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("active_base_names")
		}
		return err
	}

	return nil
}

func (m *Subscriber) validateActivePolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivePolicies) { // not required
		return nil
	}

	if err := m.ActivePolicies.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("active_policies")
		}
		return err
	}

	return nil
}

func (m *Subscriber) validateActivePoliciesByAPN(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivePoliciesByAPN) { // not required
		return nil
	}

	if err := m.ActivePoliciesByAPN.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("active_policies_by_apn")
		}
		return err
	}

	return nil
}

func (m *Subscriber) validateConfig(formats strfmt.Registry) error {

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

func (m *Subscriber) validateForbiddenNetworkTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.ForbiddenNetworkTypes) { // not required
		return nil
	}

	if err := m.ForbiddenNetworkTypes.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("forbidden_network_types")
		}
		return err
	}

	return nil
}

func (m *Subscriber) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *Subscriber) validateLTE(formats strfmt.Registry) error {

	if err := validate.Required("lte", "body", m.LTE); err != nil {
		return err
	}

	if m.LTE != nil {
		if err := m.LTE.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lte")
			}
			return err
		}
	}

	return nil
}

func (m *Subscriber) validateMonitoring(formats strfmt.Registry) error {

	if swag.IsZero(m.Monitoring) { // not required
		return nil
	}

	if m.Monitoring != nil {
		if err := m.Monitoring.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monitoring")
			}
			return err
		}
	}

	return nil
}

func (m *Subscriber) validateMsisdn(formats strfmt.Registry) error {

	if swag.IsZero(m.Msisdn) { // not required
		return nil
	}

	if err := m.Msisdn.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("msisdn")
		}
		return err
	}

	return nil
}

func (m *Subscriber) validateState(formats strfmt.Registry) error {

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
func (m *Subscriber) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subscriber) UnmarshalBinary(b []byte) error {
	var res Subscriber
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
