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

// NewGetTenantsTenantIDMetricsSeriesParams creates a new GetTenantsTenantIDMetricsSeriesParams object
// with the default values initialized.
func NewGetTenantsTenantIDMetricsSeriesParams() *GetTenantsTenantIDMetricsSeriesParams {
	var ()
	return &GetTenantsTenantIDMetricsSeriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTenantsTenantIDMetricsSeriesParamsWithTimeout creates a new GetTenantsTenantIDMetricsSeriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTenantsTenantIDMetricsSeriesParamsWithTimeout(timeout time.Duration) *GetTenantsTenantIDMetricsSeriesParams {
	var ()
	return &GetTenantsTenantIDMetricsSeriesParams{

		timeout: timeout,
	}
}

// NewGetTenantsTenantIDMetricsSeriesParamsWithContext creates a new GetTenantsTenantIDMetricsSeriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTenantsTenantIDMetricsSeriesParamsWithContext(ctx context.Context) *GetTenantsTenantIDMetricsSeriesParams {
	var ()
	return &GetTenantsTenantIDMetricsSeriesParams{

		Context: ctx,
	}
}

// NewGetTenantsTenantIDMetricsSeriesParamsWithHTTPClient creates a new GetTenantsTenantIDMetricsSeriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTenantsTenantIDMetricsSeriesParamsWithHTTPClient(client *http.Client) *GetTenantsTenantIDMetricsSeriesParams {
	var ()
	return &GetTenantsTenantIDMetricsSeriesParams{
		HTTPClient: client,
	}
}

/*GetTenantsTenantIDMetricsSeriesParams contains all the parameters to send to the API endpoint
for the get tenants tenant ID metrics series operation typically these are written to a http.Request
*/
type GetTenantsTenantIDMetricsSeriesParams struct {

	/*End
	  end time of the requested range (UnixTime or RFC3339)

	*/
	End *string
	/*Match
	  Matcher for metric series query

	*/
	Match []string
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

// WithTimeout adds the timeout to the get tenants tenant ID metrics series params
func (o *GetTenantsTenantIDMetricsSeriesParams) WithTimeout(timeout time.Duration) *GetTenantsTenantIDMetricsSeriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tenants tenant ID metrics series params
func (o *GetTenantsTenantIDMetricsSeriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tenants tenant ID metrics series params
func (o *GetTenantsTenantIDMetricsSeriesParams) WithContext(ctx context.Context) *GetTenantsTenantIDMetricsSeriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tenants tenant ID metrics series params
func (o *GetTenantsTenantIDMetricsSeriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tenants tenant ID metrics series params
func (o *GetTenantsTenantIDMetricsSeriesParams) WithHTTPClient(client *http.Client) *GetTenantsTenantIDMetricsSeriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tenants tenant ID metrics series params
func (o *GetTenantsTenantIDMetricsSeriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnd adds the end to the get tenants tenant ID metrics series params
func (o *GetTenantsTenantIDMetricsSeriesParams) WithEnd(end *string) *GetTenantsTenantIDMetricsSeriesParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get tenants tenant ID metrics series params
func (o *GetTenantsTenantIDMetricsSeriesParams) SetEnd(end *string) {
	o.End = end
}

// WithMatch adds the match to the get tenants tenant ID metrics series params
func (o *GetTenantsTenantIDMetricsSeriesParams) WithMatch(match []string) *GetTenantsTenantIDMetricsSeriesParams {
	o.SetMatch(match)
	return o
}

// SetMatch adds the match to the get tenants tenant ID metrics series params
func (o *GetTenantsTenantIDMetricsSeriesParams) SetMatch(match []string) {
	o.Match = match
}

// WithStart adds the start to the get tenants tenant ID metrics series params
func (o *GetTenantsTenantIDMetricsSeriesParams) WithStart(start *string) *GetTenantsTenantIDMetricsSeriesParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get tenants tenant ID metrics series params
func (o *GetTenantsTenantIDMetricsSeriesParams) SetStart(start *string) {
	o.Start = start
}

// WithTenantID adds the tenantID to the get tenants tenant ID metrics series params
func (o *GetTenantsTenantIDMetricsSeriesParams) WithTenantID(tenantID int64) *GetTenantsTenantIDMetricsSeriesParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get tenants tenant ID metrics series params
func (o *GetTenantsTenantIDMetricsSeriesParams) SetTenantID(tenantID int64) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTenantsTenantIDMetricsSeriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesMatch := o.Match

	joinedMatch := swag.JoinByFormat(valuesMatch, "ssv")
	// query array param match
	if err := r.SetQueryParam("match", joinedMatch...); err != nil {
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
