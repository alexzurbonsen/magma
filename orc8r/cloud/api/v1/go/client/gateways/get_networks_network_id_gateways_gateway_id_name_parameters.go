// Code generated by go-swagger; DO NOT EDIT.

package gateways

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

// NewGetNetworksNetworkIDGatewaysGatewayIDNameParams creates a new GetNetworksNetworkIDGatewaysGatewayIDNameParams object
// with the default values initialized.
func NewGetNetworksNetworkIDGatewaysGatewayIDNameParams() *GetNetworksNetworkIDGatewaysGatewayIDNameParams {
	var ()
	return &GetNetworksNetworkIDGatewaysGatewayIDNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworksNetworkIDGatewaysGatewayIDNameParamsWithTimeout creates a new GetNetworksNetworkIDGatewaysGatewayIDNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworksNetworkIDGatewaysGatewayIDNameParamsWithTimeout(timeout time.Duration) *GetNetworksNetworkIDGatewaysGatewayIDNameParams {
	var ()
	return &GetNetworksNetworkIDGatewaysGatewayIDNameParams{

		timeout: timeout,
	}
}

// NewGetNetworksNetworkIDGatewaysGatewayIDNameParamsWithContext creates a new GetNetworksNetworkIDGatewaysGatewayIDNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworksNetworkIDGatewaysGatewayIDNameParamsWithContext(ctx context.Context) *GetNetworksNetworkIDGatewaysGatewayIDNameParams {
	var ()
	return &GetNetworksNetworkIDGatewaysGatewayIDNameParams{

		Context: ctx,
	}
}

// NewGetNetworksNetworkIDGatewaysGatewayIDNameParamsWithHTTPClient creates a new GetNetworksNetworkIDGatewaysGatewayIDNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworksNetworkIDGatewaysGatewayIDNameParamsWithHTTPClient(client *http.Client) *GetNetworksNetworkIDGatewaysGatewayIDNameParams {
	var ()
	return &GetNetworksNetworkIDGatewaysGatewayIDNameParams{
		HTTPClient: client,
	}
}

/*GetNetworksNetworkIDGatewaysGatewayIDNameParams contains all the parameters to send to the API endpoint
for the get networks network ID gateways gateway ID name operation typically these are written to a http.Request
*/
type GetNetworksNetworkIDGatewaysGatewayIDNameParams struct {

	/*GatewayID
	  Gateway ID

	*/
	GatewayID string
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get networks network ID gateways gateway ID name params
func (o *GetNetworksNetworkIDGatewaysGatewayIDNameParams) WithTimeout(timeout time.Duration) *GetNetworksNetworkIDGatewaysGatewayIDNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get networks network ID gateways gateway ID name params
func (o *GetNetworksNetworkIDGatewaysGatewayIDNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get networks network ID gateways gateway ID name params
func (o *GetNetworksNetworkIDGatewaysGatewayIDNameParams) WithContext(ctx context.Context) *GetNetworksNetworkIDGatewaysGatewayIDNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get networks network ID gateways gateway ID name params
func (o *GetNetworksNetworkIDGatewaysGatewayIDNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get networks network ID gateways gateway ID name params
func (o *GetNetworksNetworkIDGatewaysGatewayIDNameParams) WithHTTPClient(client *http.Client) *GetNetworksNetworkIDGatewaysGatewayIDNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get networks network ID gateways gateway ID name params
func (o *GetNetworksNetworkIDGatewaysGatewayIDNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the get networks network ID gateways gateway ID name params
func (o *GetNetworksNetworkIDGatewaysGatewayIDNameParams) WithGatewayID(gatewayID string) *GetNetworksNetworkIDGatewaysGatewayIDNameParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the get networks network ID gateways gateway ID name params
func (o *GetNetworksNetworkIDGatewaysGatewayIDNameParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the get networks network ID gateways gateway ID name params
func (o *GetNetworksNetworkIDGatewaysGatewayIDNameParams) WithNetworkID(networkID string) *GetNetworksNetworkIDGatewaysGatewayIDNameParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get networks network ID gateways gateway ID name params
func (o *GetNetworksNetworkIDGatewaysGatewayIDNameParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworksNetworkIDGatewaysGatewayIDNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gateway_id
	if err := r.SetPathParam("gateway_id", o.GatewayID); err != nil {
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
