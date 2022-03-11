// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

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
)

// NewGetCwfParams creates a new GetCwfParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCwfParams() *GetCwfParams {
	return &GetCwfParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCwfParamsWithTimeout creates a new GetCwfParams object
// with the ability to set a timeout on a request.
func NewGetCwfParamsWithTimeout(timeout time.Duration) *GetCwfParams {
	return &GetCwfParams{
		timeout: timeout,
	}
}

// NewGetCwfParamsWithContext creates a new GetCwfParams object
// with the ability to set a context for a request.
func NewGetCwfParamsWithContext(ctx context.Context) *GetCwfParams {
	return &GetCwfParams{
		Context: ctx,
	}
}

// NewGetCwfParamsWithHTTPClient creates a new GetCwfParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCwfParamsWithHTTPClient(client *http.Client) *GetCwfParams {
	return &GetCwfParams{
		HTTPClient: client,
	}
}

/* GetCwfParams contains all the parameters to send to the API endpoint
   for the get cwf operation.

   Typically these are written to a http.Request.
*/
type GetCwfParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cwf params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCwfParams) WithDefaults() *GetCwfParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cwf params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCwfParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cwf params
func (o *GetCwfParams) WithTimeout(timeout time.Duration) *GetCwfParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cwf params
func (o *GetCwfParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cwf params
func (o *GetCwfParams) WithContext(ctx context.Context) *GetCwfParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cwf params
func (o *GetCwfParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cwf params
func (o *GetCwfParams) WithHTTPClient(client *http.Client) *GetCwfParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cwf params
func (o *GetCwfParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCwfParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
