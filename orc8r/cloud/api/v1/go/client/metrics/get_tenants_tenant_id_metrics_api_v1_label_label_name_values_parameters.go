// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams creates a new GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams object
// with the default values initialized.
func NewGetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams() *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams {
	var ()
	return &GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParamsWithTimeout creates a new GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParamsWithTimeout(timeout time.Duration) *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams {
	var ()
	return &GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams{

		timeout: timeout,
	}
}

// NewGetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParamsWithContext creates a new GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParamsWithContext(ctx context.Context) *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams {
	var ()
	return &GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams{

		Context: ctx,
	}
}

// NewGetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParamsWithHTTPClient creates a new GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParamsWithHTTPClient(client *http.Client) *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams {
	var ()
	return &GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams{
		HTTPClient: client,
	}
}

/*GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams contains all the parameters to send to the API endpoint
for the get tenants tenant ID metrics API v1 label label name values operation typically these are written to a http.Request
*/
type GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams struct {

	/*End
	  end time of the requested range (UnixTime or RFC3339)

	*/
	End *string
	/*LabelName
	  Label name to get values of

	*/
	LabelName string
	/*Start
	  start time of the requested range (UnixTime or RFC3339)

	*/
	Start *string
	/*TenantID
	  Tenant ID

	*/
	TenantID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get tenants tenant ID metrics API v1 label label name values params
func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams) WithTimeout(timeout time.Duration) *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tenants tenant ID metrics API v1 label label name values params
func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tenants tenant ID metrics API v1 label label name values params
func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams) WithContext(ctx context.Context) *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tenants tenant ID metrics API v1 label label name values params
func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tenants tenant ID metrics API v1 label label name values params
func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams) WithHTTPClient(client *http.Client) *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tenants tenant ID metrics API v1 label label name values params
func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnd adds the end to the get tenants tenant ID metrics API v1 label label name values params
func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams) WithEnd(end *string) *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get tenants tenant ID metrics API v1 label label name values params
func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams) SetEnd(end *string) {
	o.End = end
}

// WithLabelName adds the labelName to the get tenants tenant ID metrics API v1 label label name values params
func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams) WithLabelName(labelName string) *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams {
	o.SetLabelName(labelName)
	return o
}

// SetLabelName adds the labelName to the get tenants tenant ID metrics API v1 label label name values params
func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams) SetLabelName(labelName string) {
	o.LabelName = labelName
}

// WithStart adds the start to the get tenants tenant ID metrics API v1 label label name values params
func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams) WithStart(start *string) *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get tenants tenant ID metrics API v1 label label name values params
func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams) SetStart(start *string) {
	o.Start = start
}

// WithTenantID adds the tenantID to the get tenants tenant ID metrics API v1 label label name values params
func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams) WithTenantID(tenantID int64) *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get tenants tenant ID metrics API v1 label label name values params
func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams) SetTenantID(tenantID int64) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.End != nil {

		// query param end
		var qrEnd string
		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := qrEnd
		if qEnd != "" {
			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
		}

	}

	// path param label_name
	if err := r.SetPathParam("label_name", o.LabelName); err != nil {
		return err
	}

	if o.Start != nil {

		// query param start
		var qrStart string
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := qrStart
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	// path param tenant_id
	if err := r.SetPathParam("tenant_id", swag.FormatInt64(o.TenantID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
