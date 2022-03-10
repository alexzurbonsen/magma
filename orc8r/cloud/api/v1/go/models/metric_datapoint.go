// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// MetricDatapoint metric datapoint
//
// swagger:model metric_datapoint
type MetricDatapoint []string

// Validate validates this metric datapoint
func (m MetricDatapoint) Validate(formats strfmt.Registry) error {
	var res []error

	iMetricDatapointSize := int64(len(m))

	if err := validate.MinItems("", "body", iMetricDatapointSize, 2); err != nil {
		return err
	}

	if err := validate.MaxItems("", "body", iMetricDatapointSize, 2); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
