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

// FederationNetworkClusterStatus Status of a Federation HA cluster
//
// swagger:model federation_network_cluster_status
type FederationNetworkClusterStatus struct {

	// active gateway
	// Required: true
	ActiveGateway string `json:"active_gateway"`
}

// Validate validates this federation network cluster status
func (m *FederationNetworkClusterStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveGateway(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FederationNetworkClusterStatus) validateActiveGateway(formats strfmt.Registry) error {

	if err := validate.RequiredString("active_gateway", "body", string(m.ActiveGateway)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FederationNetworkClusterStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FederationNetworkClusterStatus) UnmarshalBinary(b []byte) error {
	var res FederationNetworkClusterStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
