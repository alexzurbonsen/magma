// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	models1 "magma/lte/cloud/go/services/policydb/obsidian/models"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SmsMessage sms message
//
// swagger:model sms_message
type SmsMessage struct {

	// attempt count
	// Required: true
	// Minimum: 0
	AttemptCount int64 `json:"attempt_count"`

	// error status
	ErrorStatus string `json:"error_status,omitempty"`

	// imsi
	// Required: true
	Imsi models1.SubscriberID `json:"imsi"`

	// message
	// Required: true
	// Min Length: 1
	Message string `json:"message"`

	// pk
	// Required: true
	// Min Length: 1
	Pk string `json:"pk"`

	// source msisdn
	// Required: true
	// Min Length: 1
	SourceMsisdn string `json:"source_msisdn"`

	// status
	// Required: true
	// Enum: [Waiting Delivered Failed]
	Status *string `json:"status"`

	// time created
	// Required: true
	// Format: date-time
	TimeCreated *strfmt.DateTime `json:"time_created"`

	// time last attempted
	// Format: date-time
	TimeLastAttempted strfmt.DateTime `json:"time_last_attempted,omitempty"`
}

// Validate validates this sms message
func (m *SmsMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttemptCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImsi(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceMsisdn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeLastAttempted(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmsMessage) validateAttemptCount(formats strfmt.Registry) error {

	if err := validate.Required("attempt_count", "body", int64(m.AttemptCount)); err != nil {
		return err
	}

	if err := validate.MinimumInt("attempt_count", "body", int64(m.AttemptCount), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *SmsMessage) validateImsi(formats strfmt.Registry) error {

	if err := m.Imsi.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("imsi")
		}
		return err
	}

	return nil
}

func (m *SmsMessage) validateMessage(formats strfmt.Registry) error {

	if err := validate.RequiredString("message", "body", string(m.Message)); err != nil {
		return err
	}

	if err := validate.MinLength("message", "body", string(m.Message), 1); err != nil {
		return err
	}

	return nil
}

func (m *SmsMessage) validatePk(formats strfmt.Registry) error {

	if err := validate.RequiredString("pk", "body", string(m.Pk)); err != nil {
		return err
	}

	if err := validate.MinLength("pk", "body", string(m.Pk), 1); err != nil {
		return err
	}

	return nil
}

func (m *SmsMessage) validateSourceMsisdn(formats strfmt.Registry) error {

	if err := validate.RequiredString("source_msisdn", "body", string(m.SourceMsisdn)); err != nil {
		return err
	}

	if err := validate.MinLength("source_msisdn", "body", string(m.SourceMsisdn), 1); err != nil {
		return err
	}

	return nil
}

var smsMessageTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Waiting","Delivered","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smsMessageTypeStatusPropEnum = append(smsMessageTypeStatusPropEnum, v)
	}
}

const (

	// SmsMessageStatusWaiting captures enum value "Waiting"
	SmsMessageStatusWaiting string = "Waiting"

	// SmsMessageStatusDelivered captures enum value "Delivered"
	SmsMessageStatusDelivered string = "Delivered"

	// SmsMessageStatusFailed captures enum value "Failed"
	SmsMessageStatusFailed string = "Failed"
)

// prop value enum
func (m *SmsMessage) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, smsMessageTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SmsMessage) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *SmsMessage) validateTimeCreated(formats strfmt.Registry) error {

	if err := validate.Required("time_created", "body", m.TimeCreated); err != nil {
		return err
	}

	if err := validate.FormatOf("time_created", "body", "date-time", m.TimeCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SmsMessage) validateTimeLastAttempted(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeLastAttempted) { // not required
		return nil
	}

	if err := validate.FormatOf("time_last_attempted", "body", "date-time", m.TimeLastAttempted.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SmsMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmsMessage) UnmarshalBinary(b []byte) error {
	var res SmsMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
