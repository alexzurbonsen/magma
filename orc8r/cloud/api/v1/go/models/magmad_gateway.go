// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MagmadGateway Full representation of a generic gateway
//
// swagger:model magmad_gateway
type MagmadGateway struct {

	// description
	// Required: true
	Description GatewayDescription `json:"description"`

	// device
	Device *GatewayDevice `json:"device,omitempty"`

	// id
	// Required: true
	ID GatewayID `json:"id"`

	// magmad
	// Required: true
	Magmad *MagmadGatewayConfigs `json:"magmad"`

	// name
	// Required: true
	Name GatewayName `json:"name"`

	// registration info
	RegistrationInfo *RegistrationInfo `json:"registration_info,omitempty"`

	// status
	Status *GatewayStatus `json:"status,omitempty"`

	// tier
	// Required: true
	Tier TierID `json:"tier"`
}

// Validate validates this magmad gateway
func (m *MagmadGateway) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMagmad(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistrationInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MagmadGateway) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", GatewayDescription(m.Description)); err != nil {
		return err
	}

	if err := m.Description.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("description")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("description")
		}
		return err
	}

	return nil
}

func (m *MagmadGateway) validateDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.Device) { // not required
		return nil
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *MagmadGateway) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", GatewayID(m.ID)); err != nil {
		return err
	}

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *MagmadGateway) validateMagmad(formats strfmt.Registry) error {

	if err := validate.Required("magmad", "body", m.Magmad); err != nil {
		return err
	}

	if m.Magmad != nil {
		if err := m.Magmad.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("magmad")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("magmad")
			}
			return err
		}
	}

	return nil
}

func (m *MagmadGateway) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", GatewayName(m.Name)); err != nil {
		return err
	}

	if err := m.Name.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("name")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("name")
		}
		return err
	}

	return nil
}

func (m *MagmadGateway) validateRegistrationInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.RegistrationInfo) { // not required
		return nil
	}

	if m.RegistrationInfo != nil {
		if err := m.RegistrationInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registration_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registration_info")
			}
			return err
		}
	}

	return nil
}

func (m *MagmadGateway) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *MagmadGateway) validateTier(formats strfmt.Registry) error {

	if err := validate.Required("tier", "body", TierID(m.Tier)); err != nil {
		return err
	}

	if err := m.Tier.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tier")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("tier")
		}
		return err
	}

	return nil
}

// ContextValidate validate this magmad gateway based on the context it is used
func (m *MagmadGateway) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMagmad(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegistrationInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MagmadGateway) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Description.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("description")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("description")
		}
		return err
	}

	return nil
}

func (m *MagmadGateway) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.Device != nil {
		if err := m.Device.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *MagmadGateway) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *MagmadGateway) contextValidateMagmad(ctx context.Context, formats strfmt.Registry) error {

	if m.Magmad != nil {
		if err := m.Magmad.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("magmad")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("magmad")
			}
			return err
		}
	}

	return nil
}

func (m *MagmadGateway) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Name.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("name")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("name")
		}
		return err
	}

	return nil
}

func (m *MagmadGateway) contextValidateRegistrationInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.RegistrationInfo != nil {
		if err := m.RegistrationInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registration_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registration_info")
			}
			return err
		}
	}

	return nil
}

func (m *MagmadGateway) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *MagmadGateway) contextValidateTier(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Tier.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tier")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("tier")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MagmadGateway) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MagmadGateway) UnmarshalBinary(b []byte) error {
	var res MagmadGateway
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
