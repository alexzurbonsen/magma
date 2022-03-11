// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewGetNetworksNetworkIDFeaturesParams creates a new GetNetworksNetworkIDFeaturesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworksNetworkIDFeaturesParams() *GetNetworksNetworkIDFeaturesParams {
	return &GetNetworksNetworkIDFeaturesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworksNetworkIDFeaturesParamsWithTimeout creates a new GetNetworksNetworkIDFeaturesParams object
// with the ability to set a timeout on a request.
func NewGetNetworksNetworkIDFeaturesParamsWithTimeout(timeout time.Duration) *GetNetworksNetworkIDFeaturesParams {
	return &GetNetworksNetworkIDFeaturesParams{
		timeout: timeout,
	}
}

// NewGetNetworksNetworkIDFeaturesParamsWithContext creates a new GetNetworksNetworkIDFeaturesParams object
// with the ability to set a context for a request.
func NewGetNetworksNetworkIDFeaturesParamsWithContext(ctx context.Context) *GetNetworksNetworkIDFeaturesParams {
	return &GetNetworksNetworkIDFeaturesParams{
		Context: ctx,
	}
}

// NewGetNetworksNetworkIDFeaturesParamsWithHTTPClient creates a new GetNetworksNetworkIDFeaturesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworksNetworkIDFeaturesParamsWithHTTPClient(client *http.Client) *GetNetworksNetworkIDFeaturesParams {
	return &GetNetworksNetworkIDFeaturesParams{
		HTTPClient: client,
	}
}

/* GetNetworksNetworkIDFeaturesParams contains all the parameters to send to the API endpoint
   for the get networks network ID features operation.

   Typically these are written to a http.Request.
*/
type GetNetworksNetworkIDFeaturesParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get networks network ID features params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworksNetworkIDFeaturesParams) WithDefaults() *GetNetworksNetworkIDFeaturesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get networks network ID features params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworksNetworkIDFeaturesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get networks network ID features params
func (o *GetNetworksNetworkIDFeaturesParams) WithTimeout(timeout time.Duration) *GetNetworksNetworkIDFeaturesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get networks network ID features params
func (o *GetNetworksNetworkIDFeaturesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get networks network ID features params
func (o *GetNetworksNetworkIDFeaturesParams) WithContext(ctx context.Context) *GetNetworksNetworkIDFeaturesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get networks network ID features params
func (o *GetNetworksNetworkIDFeaturesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get networks network ID features params
func (o *GetNetworksNetworkIDFeaturesParams) WithHTTPClient(client *http.Client) *GetNetworksNetworkIDFeaturesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get networks network ID features params
func (o *GetNetworksNetworkIDFeaturesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get networks network ID features params
func (o *GetNetworksNetworkIDFeaturesParams) WithNetworkID(networkID string) *GetNetworksNetworkIDFeaturesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get networks network ID features params
func (o *GetNetworksNetworkIDFeaturesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworksNetworkIDFeaturesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
