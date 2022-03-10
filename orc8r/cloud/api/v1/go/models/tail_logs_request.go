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

// TailLogsRequest tail logs request
//
// swagger:model tail_logs_request
type TailLogsRequest struct {

	// service
	// Min Length: 1
	Service string `json:"service,omitempty"`
}

// Validate validates this tail logs request
func (m *TailLogsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TailLogsRequest) validateService(formats strfmt.Registry) error {

	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if err := validate.MinLength("service", "body", string(m.Service), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TailLogsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TailLogsRequest) UnmarshalBinary(b []byte) error {
	var res TailLogsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
