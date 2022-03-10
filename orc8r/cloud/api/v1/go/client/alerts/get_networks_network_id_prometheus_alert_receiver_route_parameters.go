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

// NewGetNetworksNetworkIDPrometheusAlertReceiverRouteParams creates a new GetNetworksNetworkIDPrometheusAlertReceiverRouteParams object
// with the default values initialized.
func NewGetNetworksNetworkIDPrometheusAlertReceiverRouteParams() *GetNetworksNetworkIDPrometheusAlertReceiverRouteParams {
	var ()
	return &GetNetworksNetworkIDPrometheusAlertReceiverRouteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworksNetworkIDPrometheusAlertReceiverRouteParamsWithTimeout creates a new GetNetworksNetworkIDPrometheusAlertReceiverRouteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworksNetworkIDPrometheusAlertReceiverRouteParamsWithTimeout(timeout time.Duration) *GetNetworksNetworkIDPrometheusAlertReceiverRouteParams {
	var ()
	return &GetNetworksNetworkIDPrometheusAlertReceiverRouteParams{

		timeout: timeout,
	}
}

// NewGetNetworksNetworkIDPrometheusAlertReceiverRouteParamsWithContext creates a new GetNetworksNetworkIDPrometheusAlertReceiverRouteParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworksNetworkIDPrometheusAlertReceiverRouteParamsWithContext(ctx context.Context) *GetNetworksNetworkIDPrometheusAlertReceiverRouteParams {
	var ()
	return &GetNetworksNetworkIDPrometheusAlertReceiverRouteParams{

		Context: ctx,
	}
}

// NewGetNetworksNetworkIDPrometheusAlertReceiverRouteParamsWithHTTPClient creates a new GetNetworksNetworkIDPrometheusAlertReceiverRouteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworksNetworkIDPrometheusAlertReceiverRouteParamsWithHTTPClient(client *http.Client) *GetNetworksNetworkIDPrometheusAlertReceiverRouteParams {
	var ()
	return &GetNetworksNetworkIDPrometheusAlertReceiverRouteParams{
		HTTPClient: client,
	}
}

/*GetNetworksNetworkIDPrometheusAlertReceiverRouteParams contains all the parameters to send to the API endpoint
for the get networks network ID prometheus alert receiver route operation typically these are written to a http.Request
*/
type GetNetworksNetworkIDPrometheusAlertReceiverRouteParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get networks network ID prometheus alert receiver route params
func (o *GetNetworksNetworkIDPrometheusAlertReceiverRouteParams) WithTimeout(timeout time.Duration) *GetNetworksNetworkIDPrometheusAlertReceiverRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get networks network ID prometheus alert receiver route params
func (o *GetNetworksNetworkIDPrometheusAlertReceiverRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get networks network ID prometheus alert receiver route params
func (o *GetNetworksNetworkIDPrometheusAlertReceiverRouteParams) WithContext(ctx context.Context) *GetNetworksNetworkIDPrometheusAlertReceiverRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get networks network ID prometheus alert receiver route params
func (o *GetNetworksNetworkIDPrometheusAlertReceiverRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get networks network ID prometheus alert receiver route params
func (o *GetNetworksNetworkIDPrometheusAlertReceiverRouteParams) WithHTTPClient(client *http.Client) *GetNetworksNetworkIDPrometheusAlertReceiverRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get networks network ID prometheus alert receiver route params
func (o *GetNetworksNetworkIDPrometheusAlertReceiverRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get networks network ID prometheus alert receiver route params
func (o *GetNetworksNetworkIDPrometheusAlertReceiverRouteParams) WithNetworkID(networkID string) *GetNetworksNetworkIDPrometheusAlertReceiverRouteParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get networks network ID prometheus alert receiver route params
func (o *GetNetworksNetworkIDPrometheusAlertReceiverRouteParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworksNetworkIDPrometheusAlertReceiverRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
