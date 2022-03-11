// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkProbeTaskDetails network probe task details
//
// swagger:model network_probe_task_details
type NetworkProbeTaskDetails struct {

	// correlation id
	// Example: 605394647632969700
	CorrelationID uint64 `json:"correlation_id,omitempty"`

	// delivery type
	// Example: events_only
	// Required: true
	// Enum: [all events_only]
	DeliveryType string `json:"delivery_type"`

	// domain id
	DomainID string `json:"domain_id,omitempty"`

	// the duration in seconds after which the task will expire.
	// Example: 300
	// Minimum: 0
	Duration *int64 `json:"duration,omitempty"`

	// represents the mobile operator identifier
	OperatorID uint32 `json:"operator_id,omitempty"`

	// target id
	// Required: true
	TargetID string `json:"target_id"`

	// target type
	// Example: imsi
	// Required: true
	// Enum: [imsi imei msisdn]
	TargetType string `json:"target_type"`

	// The timestamp in ISO 8601 format
	// Example: 2020-03-11T00:36:59.65Z
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this network probe task details
func (m *NetworkProbeTaskDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeliveryType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var networkProbeTaskDetailsTypeDeliveryTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","events_only"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkProbeTaskDetailsTypeDeliveryTypePropEnum = append(networkProbeTaskDetailsTypeDeliveryTypePropEnum, v)
	}
}

const (

	// NetworkProbeTaskDetailsDeliveryTypeAll captures enum value "all"
	NetworkProbeTaskDetailsDeliveryTypeAll string = "all"

	// NetworkProbeTaskDetailsDeliveryTypeEventsOnly captures enum value "events_only"
	NetworkProbeTaskDetailsDeliveryTypeEventsOnly string = "events_only"
)

// prop value enum
func (m *NetworkProbeTaskDetails) validateDeliveryTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkProbeTaskDetailsTypeDeliveryTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkProbeTaskDetails) validateDeliveryType(formats strfmt.Registry) error {

	if err := validate.RequiredString("delivery_type", "body", m.DeliveryType); err != nil {
		return err
	}

	// value enum
	if err := m.validateDeliveryTypeEnum("delivery_type", "body", m.DeliveryType); err != nil {
		return err
	}

	return nil
}

func (m *NetworkProbeTaskDetails) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	if err := validate.MinimumInt("duration", "body", *m.Duration, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *NetworkProbeTaskDetails) validateTargetID(formats strfmt.Registry) error {

	if err := validate.RequiredString("target_id", "body", m.TargetID); err != nil {
		return err
	}

	return nil
}

var networkProbeTaskDetailsTypeTargetTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["imsi","imei","msisdn"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkProbeTaskDetailsTypeTargetTypePropEnum = append(networkProbeTaskDetailsTypeTargetTypePropEnum, v)
	}
}

const (

	// NetworkProbeTaskDetailsTargetTypeImsi captures enum value "imsi"
	NetworkProbeTaskDetailsTargetTypeImsi string = "imsi"

	// NetworkProbeTaskDetailsTargetTypeImei captures enum value "imei"
	NetworkProbeTaskDetailsTargetTypeImei string = "imei"

	// NetworkProbeTaskDetailsTargetTypeMsisdn captures enum value "msisdn"
	NetworkProbeTaskDetailsTargetTypeMsisdn string = "msisdn"
)

// prop value enum
func (m *NetworkProbeTaskDetails) validateTargetTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkProbeTaskDetailsTypeTargetTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkProbeTaskDetails) validateTargetType(formats strfmt.Registry) error {

	if err := validate.RequiredString("target_type", "body", m.TargetType); err != nil {
		return err
	}

	// value enum
	if err := m.validateTargetTypeEnum("target_type", "body", m.TargetType); err != nil {
		return err
	}

	return nil
}

func (m *NetworkProbeTaskDetails) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this network probe task details based on context it is used
func (m *NetworkProbeTaskDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkProbeTaskDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkProbeTaskDetails) UnmarshalBinary(b []byte) error {
	var res NetworkProbeTaskDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
