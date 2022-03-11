// Code generated by go-swagger; DO NOT EDIT.

package cbsds

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

// NewGetDpNetworkIDCbsdsCbsdIDParams creates a new GetDpNetworkIDCbsdsCbsdIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDpNetworkIDCbsdsCbsdIDParams() *GetDpNetworkIDCbsdsCbsdIDParams {
	return &GetDpNetworkIDCbsdsCbsdIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDpNetworkIDCbsdsCbsdIDParamsWithTimeout creates a new GetDpNetworkIDCbsdsCbsdIDParams object
// with the ability to set a timeout on a request.
func NewGetDpNetworkIDCbsdsCbsdIDParamsWithTimeout(timeout time.Duration) *GetDpNetworkIDCbsdsCbsdIDParams {
	return &GetDpNetworkIDCbsdsCbsdIDParams{
		timeout: timeout,
	}
}

// NewGetDpNetworkIDCbsdsCbsdIDParamsWithContext creates a new GetDpNetworkIDCbsdsCbsdIDParams object
// with the ability to set a context for a request.
func NewGetDpNetworkIDCbsdsCbsdIDParamsWithContext(ctx context.Context) *GetDpNetworkIDCbsdsCbsdIDParams {
	return &GetDpNetworkIDCbsdsCbsdIDParams{
		Context: ctx,
	}
}

// NewGetDpNetworkIDCbsdsCbsdIDParamsWithHTTPClient creates a new GetDpNetworkIDCbsdsCbsdIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDpNetworkIDCbsdsCbsdIDParamsWithHTTPClient(client *http.Client) *GetDpNetworkIDCbsdsCbsdIDParams {
	return &GetDpNetworkIDCbsdsCbsdIDParams{
		HTTPClient: client,
	}
}

/* GetDpNetworkIDCbsdsCbsdIDParams contains all the parameters to send to the API endpoint
   for the get dp network ID cbsds cbsd ID operation.

   Typically these are written to a http.Request.
*/
type GetDpNetworkIDCbsdsCbsdIDParams struct {

	/* CbsdID.

	   CBSD ID
	*/
	CbsdID int64

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dp network ID cbsds cbsd ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDpNetworkIDCbsdsCbsdIDParams) WithDefaults() *GetDpNetworkIDCbsdsCbsdIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dp network ID cbsds cbsd ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDpNetworkIDCbsdsCbsdIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get dp network ID cbsds cbsd ID params
func (o *GetDpNetworkIDCbsdsCbsdIDParams) WithTimeout(timeout time.Duration) *GetDpNetworkIDCbsdsCbsdIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dp network ID cbsds cbsd ID params
func (o *GetDpNetworkIDCbsdsCbsdIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dp network ID cbsds cbsd ID params
func (o *GetDpNetworkIDCbsdsCbsdIDParams) WithContext(ctx context.Context) *GetDpNetworkIDCbsdsCbsdIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dp network ID cbsds cbsd ID params
func (o *GetDpNetworkIDCbsdsCbsdIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dp network ID cbsds cbsd ID params
func (o *GetDpNetworkIDCbsdsCbsdIDParams) WithHTTPClient(client *http.Client) *GetDpNetworkIDCbsdsCbsdIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dp network ID cbsds cbsd ID params
func (o *GetDpNetworkIDCbsdsCbsdIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCbsdID adds the cbsdID to the get dp network ID cbsds cbsd ID params
func (o *GetDpNetworkIDCbsdsCbsdIDParams) WithCbsdID(cbsdID int64) *GetDpNetworkIDCbsdsCbsdIDParams {
	o.SetCbsdID(cbsdID)
	return o
}

// SetCbsdID adds the cbsdId to the get dp network ID cbsds cbsd ID params
func (o *GetDpNetworkIDCbsdsCbsdIDParams) SetCbsdID(cbsdID int64) {
	o.CbsdID = cbsdID
}

// WithNetworkID adds the networkID to the get dp network ID cbsds cbsd ID params
func (o *GetDpNetworkIDCbsdsCbsdIDParams) WithNetworkID(networkID string) *GetDpNetworkIDCbsdsCbsdIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get dp network ID cbsds cbsd ID params
func (o *GetDpNetworkIDCbsdsCbsdIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDpNetworkIDCbsdsCbsdIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cbsd_id
	if err := r.SetPathParam("cbsd_id", swag.FormatInt64(o.CbsdID)); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
