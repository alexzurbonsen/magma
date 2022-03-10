// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_gateways

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

// NewGetCwfNetworkIDGatewaysGatewayIDNameParams creates a new GetCwfNetworkIDGatewaysGatewayIDNameParams object
// with the default values initialized.
func NewGetCwfNetworkIDGatewaysGatewayIDNameParams() *GetCwfNetworkIDGatewaysGatewayIDNameParams {
	var ()
	return &GetCwfNetworkIDGatewaysGatewayIDNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCwfNetworkIDGatewaysGatewayIDNameParamsWithTimeout creates a new GetCwfNetworkIDGatewaysGatewayIDNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCwfNetworkIDGatewaysGatewayIDNameParamsWithTimeout(timeout time.Duration) *GetCwfNetworkIDGatewaysGatewayIDNameParams {
	var ()
	return &GetCwfNetworkIDGatewaysGatewayIDNameParams{

		timeout: timeout,
	}
}

// NewGetCwfNetworkIDGatewaysGatewayIDNameParamsWithContext creates a new GetCwfNetworkIDGatewaysGatewayIDNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCwfNetworkIDGatewaysGatewayIDNameParamsWithContext(ctx context.Context) *GetCwfNetworkIDGatewaysGatewayIDNameParams {
	var ()
	return &GetCwfNetworkIDGatewaysGatewayIDNameParams{

		Context: ctx,
	}
}

// NewGetCwfNetworkIDGatewaysGatewayIDNameParamsWithHTTPClient creates a new GetCwfNetworkIDGatewaysGatewayIDNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCwfNetworkIDGatewaysGatewayIDNameParamsWithHTTPClient(client *http.Client) *GetCwfNetworkIDGatewaysGatewayIDNameParams {
	var ()
	return &GetCwfNetworkIDGatewaysGatewayIDNameParams{
		HTTPClient: client,
	}
}

/*GetCwfNetworkIDGatewaysGatewayIDNameParams contains all the parameters to send to the API endpoint
for the get cwf network ID gateways gateway ID name operation typically these are written to a http.Request
*/
type GetCwfNetworkIDGatewaysGatewayIDNameParams struct {

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

// WithTimeout adds the timeout to the get cwf network ID gateways gateway ID name params
func (o *GetCwfNetworkIDGatewaysGatewayIDNameParams) WithTimeout(timeout time.Duration) *GetCwfNetworkIDGatewaysGatewayIDNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cwf network ID gateways gateway ID name params
func (o *GetCwfNetworkIDGatewaysGatewayIDNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cwf network ID gateways gateway ID name params
func (o *GetCwfNetworkIDGatewaysGatewayIDNameParams) WithContext(ctx context.Context) *GetCwfNetworkIDGatewaysGatewayIDNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cwf network ID gateways gateway ID name params
func (o *GetCwfNetworkIDGatewaysGatewayIDNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cwf network ID gateways gateway ID name params
func (o *GetCwfNetworkIDGatewaysGatewayIDNameParams) WithHTTPClient(client *http.Client) *GetCwfNetworkIDGatewaysGatewayIDNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cwf network ID gateways gateway ID name params
func (o *GetCwfNetworkIDGatewaysGatewayIDNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the get cwf network ID gateways gateway ID name params
func (o *GetCwfNetworkIDGatewaysGatewayIDNameParams) WithGatewayID(gatewayID string) *GetCwfNetworkIDGatewaysGatewayIDNameParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the get cwf network ID gateways gateway ID name params
func (o *GetCwfNetworkIDGatewaysGatewayIDNameParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the get cwf network ID gateways gateway ID name params
func (o *GetCwfNetworkIDGatewaysGatewayIDNameParams) WithNetworkID(networkID string) *GetCwfNetworkIDGatewaysGatewayIDNameParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get cwf network ID gateways gateway ID name params
func (o *GetCwfNetworkIDGatewaysGatewayIDNameParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCwfNetworkIDGatewaysGatewayIDNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
