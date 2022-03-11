// Code generated by go-swagger; DO NOT EDIT.

package lte_gateways

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

// NewGetLTENetworkIDGatewaysParams creates a new GetLTENetworkIDGatewaysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLTENetworkIDGatewaysParams() *GetLTENetworkIDGatewaysParams {
	return &GetLTENetworkIDGatewaysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLTENetworkIDGatewaysParamsWithTimeout creates a new GetLTENetworkIDGatewaysParams object
// with the ability to set a timeout on a request.
func NewGetLTENetworkIDGatewaysParamsWithTimeout(timeout time.Duration) *GetLTENetworkIDGatewaysParams {
	return &GetLTENetworkIDGatewaysParams{
		timeout: timeout,
	}
}

// NewGetLTENetworkIDGatewaysParamsWithContext creates a new GetLTENetworkIDGatewaysParams object
// with the ability to set a context for a request.
func NewGetLTENetworkIDGatewaysParamsWithContext(ctx context.Context) *GetLTENetworkIDGatewaysParams {
	return &GetLTENetworkIDGatewaysParams{
		Context: ctx,
	}
}

// NewGetLTENetworkIDGatewaysParamsWithHTTPClient creates a new GetLTENetworkIDGatewaysParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLTENetworkIDGatewaysParamsWithHTTPClient(client *http.Client) *GetLTENetworkIDGatewaysParams {
	return &GetLTENetworkIDGatewaysParams{
		HTTPClient: client,
	}
}

/* GetLTENetworkIDGatewaysParams contains all the parameters to send to the API endpoint
   for the get LTE network ID gateways operation.

   Typically these are written to a http.Request.
*/
type GetLTENetworkIDGatewaysParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get LTE network ID gateways params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLTENetworkIDGatewaysParams) WithDefaults() *GetLTENetworkIDGatewaysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get LTE network ID gateways params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLTENetworkIDGatewaysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get LTE network ID gateways params
func (o *GetLTENetworkIDGatewaysParams) WithTimeout(timeout time.Duration) *GetLTENetworkIDGatewaysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get LTE network ID gateways params
func (o *GetLTENetworkIDGatewaysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get LTE network ID gateways params
func (o *GetLTENetworkIDGatewaysParams) WithContext(ctx context.Context) *GetLTENetworkIDGatewaysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get LTE network ID gateways params
func (o *GetLTENetworkIDGatewaysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get LTE network ID gateways params
func (o *GetLTENetworkIDGatewaysParams) WithHTTPClient(client *http.Client) *GetLTENetworkIDGatewaysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get LTE network ID gateways params
func (o *GetLTENetworkIDGatewaysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get LTE network ID gateways params
func (o *GetLTENetworkIDGatewaysParams) WithNetworkID(networkID string) *GetLTENetworkIDGatewaysParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get LTE network ID gateways params
func (o *GetLTENetworkIDGatewaysParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLTENetworkIDGatewaysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
