// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HTTPConfig http config
//
// swagger:model http_config
type HTTPConfig struct {

	// basic auth
	BasicAuth *HTTPConfigBasicAuth `json:"basic_auth,omitempty"`

	// bearer token
	BearerToken string `json:"bearer_token,omitempty"`

	// proxy url
	ProxyURL string `json:"proxy_url,omitempty"`
}

// Validate validates this http config
func (m *HTTPConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBasicAuth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPConfig) validateBasicAuth(formats strfmt.Registry) error {
	if swag.IsZero(m.BasicAuth) { // not required
		return nil
	}

	if m.BasicAuth != nil {
		if err := m.BasicAuth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basic_auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("basic_auth")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this http config based on the context it is used
func (m *HTTPConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBasicAuth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPConfig) contextValidateBasicAuth(ctx context.Context, formats strfmt.Registry) error {

	if m.BasicAuth != nil {
		if err := m.BasicAuth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basic_auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("basic_auth")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HTTPConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPConfig) UnmarshalBinary(b []byte) error {
	var res HTTPConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
