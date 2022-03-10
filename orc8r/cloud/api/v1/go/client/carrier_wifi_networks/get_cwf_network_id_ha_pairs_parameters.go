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

// NewGetCwfNetworkIDHaPairsParams creates a new GetCwfNetworkIDHaPairsParams object
// with the default values initialized.
func NewGetCwfNetworkIDHaPairsParams() *GetCwfNetworkIDHaPairsParams {
	var ()
	return &GetCwfNetworkIDHaPairsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCwfNetworkIDHaPairsParamsWithTimeout creates a new GetCwfNetworkIDHaPairsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCwfNetworkIDHaPairsParamsWithTimeout(timeout time.Duration) *GetCwfNetworkIDHaPairsParams {
	var ()
	return &GetCwfNetworkIDHaPairsParams{

		timeout: timeout,
	}
}

// NewGetCwfNetworkIDHaPairsParamsWithContext creates a new GetCwfNetworkIDHaPairsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCwfNetworkIDHaPairsParamsWithContext(ctx context.Context) *GetCwfNetworkIDHaPairsParams {
	var ()
	return &GetCwfNetworkIDHaPairsParams{

		Context: ctx,
	}
}

// NewGetCwfNetworkIDHaPairsParamsWithHTTPClient creates a new GetCwfNetworkIDHaPairsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCwfNetworkIDHaPairsParamsWithHTTPClient(client *http.Client) *GetCwfNetworkIDHaPairsParams {
	var ()
	return &GetCwfNetworkIDHaPairsParams{
		HTTPClient: client,
	}
}

/*GetCwfNetworkIDHaPairsParams contains all the parameters to send to the API endpoint
for the get cwf network ID ha pairs operation typically these are written to a http.Request
*/
type GetCwfNetworkIDHaPairsParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cwf network ID ha pairs params
func (o *GetCwfNetworkIDHaPairsParams) WithTimeout(timeout time.Duration) *GetCwfNetworkIDHaPairsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cwf network ID ha pairs params
func (o *GetCwfNetworkIDHaPairsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cwf network ID ha pairs params
func (o *GetCwfNetworkIDHaPairsParams) WithContext(ctx context.Context) *GetCwfNetworkIDHaPairsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cwf network ID ha pairs params
func (o *GetCwfNetworkIDHaPairsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cwf network ID ha pairs params
func (o *GetCwfNetworkIDHaPairsParams) WithHTTPClient(client *http.Client) *GetCwfNetworkIDHaPairsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cwf network ID ha pairs params
func (o *GetCwfNetworkIDHaPairsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get cwf network ID ha pairs params
func (o *GetCwfNetworkIDHaPairsParams) WithNetworkID(networkID string) *GetCwfNetworkIDHaPairsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get cwf network ID ha pairs params
func (o *GetCwfNetworkIDHaPairsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCwfNetworkIDHaPairsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
