// Code generated by go-swagger; DO NOT EDIT.

package federation_gateways

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

// NewGetFegNetworkIDGatewaysGatewayIDParams creates a new GetFegNetworkIDGatewaysGatewayIDParams object
// with the default values initialized.
func NewGetFegNetworkIDGatewaysGatewayIDParams() *GetFegNetworkIDGatewaysGatewayIDParams {
	var ()
	return &GetFegNetworkIDGatewaysGatewayIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFegNetworkIDGatewaysGatewayIDParamsWithTimeout creates a new GetFegNetworkIDGatewaysGatewayIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFegNetworkIDGatewaysGatewayIDParamsWithTimeout(timeout time.Duration) *GetFegNetworkIDGatewaysGatewayIDParams {
	var ()
	return &GetFegNetworkIDGatewaysGatewayIDParams{

		timeout: timeout,
	}
}

// NewGetFegNetworkIDGatewaysGatewayIDParamsWithContext creates a new GetFegNetworkIDGatewaysGatewayIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFegNetworkIDGatewaysGatewayIDParamsWithContext(ctx context.Context) *GetFegNetworkIDGatewaysGatewayIDParams {
	var ()
	return &GetFegNetworkIDGatewaysGatewayIDParams{

		Context: ctx,
	}
}

// NewGetFegNetworkIDGatewaysGatewayIDParamsWithHTTPClient creates a new GetFegNetworkIDGatewaysGatewayIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFegNetworkIDGatewaysGatewayIDParamsWithHTTPClient(client *http.Client) *GetFegNetworkIDGatewaysGatewayIDParams {
	var ()
	return &GetFegNetworkIDGatewaysGatewayIDParams{
		HTTPClient: client,
	}
}

/*GetFegNetworkIDGatewaysGatewayIDParams contains all the parameters to send to the API endpoint
for the get feg network ID gateways gateway ID operation typically these are written to a http.Request
*/
type GetFegNetworkIDGatewaysGatewayIDParams struct {

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

// WithTimeout adds the timeout to the get feg network ID gateways gateway ID params
func (o *GetFegNetworkIDGatewaysGatewayIDParams) WithTimeout(timeout time.Duration) *GetFegNetworkIDGatewaysGatewayIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get feg network ID gateways gateway ID params
func (o *GetFegNetworkIDGatewaysGatewayIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get feg network ID gateways gateway ID params
func (o *GetFegNetworkIDGatewaysGatewayIDParams) WithContext(ctx context.Context) *GetFegNetworkIDGatewaysGatewayIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get feg network ID gateways gateway ID params
func (o *GetFegNetworkIDGatewaysGatewayIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get feg network ID gateways gateway ID params
func (o *GetFegNetworkIDGatewaysGatewayIDParams) WithHTTPClient(client *http.Client) *GetFegNetworkIDGatewaysGatewayIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get feg network ID gateways gateway ID params
func (o *GetFegNetworkIDGatewaysGatewayIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the get feg network ID gateways gateway ID params
func (o *GetFegNetworkIDGatewaysGatewayIDParams) WithGatewayID(gatewayID string) *GetFegNetworkIDGatewaysGatewayIDParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the get feg network ID gateways gateway ID params
func (o *GetFegNetworkIDGatewaysGatewayIDParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the get feg network ID gateways gateway ID params
func (o *GetFegNetworkIDGatewaysGatewayIDParams) WithNetworkID(networkID string) *GetFegNetworkIDGatewaysGatewayIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get feg network ID gateways gateway ID params
func (o *GetFegNetworkIDGatewaysGatewayIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFegNetworkIDGatewaysGatewayIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
