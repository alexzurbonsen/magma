// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// NewGetNetworksNetworkIDAlertsParams creates a new GetNetworksNetworkIDAlertsParams object
// with the default values initialized.
func NewGetNetworksNetworkIDAlertsParams() *GetNetworksNetworkIDAlertsParams {
	var ()
	return &GetNetworksNetworkIDAlertsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworksNetworkIDAlertsParamsWithTimeout creates a new GetNetworksNetworkIDAlertsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworksNetworkIDAlertsParamsWithTimeout(timeout time.Duration) *GetNetworksNetworkIDAlertsParams {
	var ()
	return &GetNetworksNetworkIDAlertsParams{

		timeout: timeout,
	}
}

// NewGetNetworksNetworkIDAlertsParamsWithContext creates a new GetNetworksNetworkIDAlertsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworksNetworkIDAlertsParamsWithContext(ctx context.Context) *GetNetworksNetworkIDAlertsParams {
	var ()
	return &GetNetworksNetworkIDAlertsParams{

		Context: ctx,
	}
}

// NewGetNetworksNetworkIDAlertsParamsWithHTTPClient creates a new GetNetworksNetworkIDAlertsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworksNetworkIDAlertsParamsWithHTTPClient(client *http.Client) *GetNetworksNetworkIDAlertsParams {
	var ()
	return &GetNetworksNetworkIDAlertsParams{
		HTTPClient: client,
	}
}

/*GetNetworksNetworkIDAlertsParams contains all the parameters to send to the API endpoint
for the get networks network ID alerts operation typically these are written to a http.Request
*/
type GetNetworksNetworkIDAlertsParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get networks network ID alerts params
func (o *GetNetworksNetworkIDAlertsParams) WithTimeout(timeout time.Duration) *GetNetworksNetworkIDAlertsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get networks network ID alerts params
func (o *GetNetworksNetworkIDAlertsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get networks network ID alerts params
func (o *GetNetworksNetworkIDAlertsParams) WithContext(ctx context.Context) *GetNetworksNetworkIDAlertsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get networks network ID alerts params
func (o *GetNetworksNetworkIDAlertsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get networks network ID alerts params
func (o *GetNetworksNetworkIDAlertsParams) WithHTTPClient(client *http.Client) *GetNetworksNetworkIDAlertsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get networks network ID alerts params
func (o *GetNetworksNetworkIDAlertsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get networks network ID alerts params
func (o *GetNetworksNetworkIDAlertsParams) WithNetworkID(networkID string) *GetNetworksNetworkIDAlertsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get networks network ID alerts params
func (o *GetNetworksNetworkIDAlertsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworksNetworkIDAlertsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
