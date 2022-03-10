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

// NewGetTenantsTenantIDMetricsQueryParams creates a new GetTenantsTenantIDMetricsQueryParams object
// with the default values initialized.
func NewGetTenantsTenantIDMetricsQueryParams() *GetTenantsTenantIDMetricsQueryParams {
	var ()
	return &GetTenantsTenantIDMetricsQueryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTenantsTenantIDMetricsQueryParamsWithTimeout creates a new GetTenantsTenantIDMetricsQueryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTenantsTenantIDMetricsQueryParamsWithTimeout(timeout time.Duration) *GetTenantsTenantIDMetricsQueryParams {
	var ()
	return &GetTenantsTenantIDMetricsQueryParams{

		timeout: timeout,
	}
}

// NewGetTenantsTenantIDMetricsQueryParamsWithContext creates a new GetTenantsTenantIDMetricsQueryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTenantsTenantIDMetricsQueryParamsWithContext(ctx context.Context) *GetTenantsTenantIDMetricsQueryParams {
	var ()
	return &GetTenantsTenantIDMetricsQueryParams{

		Context: ctx,
	}
}

// NewGetTenantsTenantIDMetricsQueryParamsWithHTTPClient creates a new GetTenantsTenantIDMetricsQueryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTenantsTenantIDMetricsQueryParamsWithHTTPClient(client *http.Client) *GetTenantsTenantIDMetricsQueryParams {
	var ()
	return &GetTenantsTenantIDMetricsQueryParams{
		HTTPClient: client,
	}
}

/*GetTenantsTenantIDMetricsQueryParams contains all the parameters to send to the API endpoint
for the get tenants tenant ID metrics query operation typically these are written to a http.Request
*/
type GetTenantsTenantIDMetricsQueryParams struct {

	/*Query
	  PromQL query to proxy to prometheus

	*/
	Query string
	/*TenantID
	  Tenant ID

	*/
	TenantID int64
	/*Time
	  time for query (UnixTime or RFC3339)

	*/
	Time *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get tenants tenant ID metrics query params
func (o *GetTenantsTenantIDMetricsQueryParams) WithTimeout(timeout time.Duration) *GetTenantsTenantIDMetricsQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tenants tenant ID metrics query params
func (o *GetTenantsTenantIDMetricsQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tenants tenant ID metrics query params
func (o *GetTenantsTenantIDMetricsQueryParams) WithContext(ctx context.Context) *GetTenantsTenantIDMetricsQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tenants tenant ID metrics query params
func (o *GetTenantsTenantIDMetricsQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tenants tenant ID metrics query params
func (o *GetTenantsTenantIDMetricsQueryParams) WithHTTPClient(client *http.Client) *GetTenantsTenantIDMetricsQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tenants tenant ID metrics query params
func (o *GetTenantsTenantIDMetricsQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the get tenants tenant ID metrics query params
func (o *GetTenantsTenantIDMetricsQueryParams) WithQuery(query string) *GetTenantsTenantIDMetricsQueryParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get tenants tenant ID metrics query params
func (o *GetTenantsTenantIDMetricsQueryParams) SetQuery(query string) {
	o.Query = query
}

// WithTenantID adds the tenantID to the get tenants tenant ID metrics query params
func (o *GetTenantsTenantIDMetricsQueryParams) WithTenantID(tenantID int64) *GetTenantsTenantIDMetricsQueryParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get tenants tenant ID metrics query params
func (o *GetTenantsTenantIDMetricsQueryParams) SetTenantID(tenantID int64) {
	o.TenantID = tenantID
}

// WithTime adds the time to the get tenants tenant ID metrics query params
func (o *GetTenantsTenantIDMetricsQueryParams) WithTime(time *string) *GetTenantsTenantIDMetricsQueryParams {
	o.SetTime(time)
	return o
}

// SetTime adds the time to the get tenants tenant ID metrics query params
func (o *GetTenantsTenantIDMetricsQueryParams) SetTime(time *string) {
	o.Time = time
}

// WriteToRequest writes these params to a swagger request
func (o *GetTenantsTenantIDMetricsQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param query
	qrQuery := o.Query
	qQuery := qrQuery
	if qQuery != "" {
		if err := r.SetQueryParam("query", qQuery); err != nil {
			return err
		}
	}

	// path param tenant_id
	if err := r.SetPathParam("tenant_id", swag.FormatInt64(o.TenantID)); err != nil {
		return err
	}

	if o.Time != nil {

		// query param time
		var qrTime string
		if o.Time != nil {
			qrTime = *o.Time
		}
		qTime := qrTime
		if qTime != "" {
			if err := r.SetQueryParam("time", qTime); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
