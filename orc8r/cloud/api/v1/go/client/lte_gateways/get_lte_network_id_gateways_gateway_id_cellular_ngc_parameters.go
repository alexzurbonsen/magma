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

// NewGetLTENetworkIDGatewaysGatewayIDCellularNgcParams creates a new GetLTENetworkIDGatewaysGatewayIDCellularNgcParams object
// with the default values initialized.
func NewGetLTENetworkIDGatewaysGatewayIDCellularNgcParams() *GetLTENetworkIDGatewaysGatewayIDCellularNgcParams {
	var ()
	return &GetLTENetworkIDGatewaysGatewayIDCellularNgcParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDCellularNgcParamsWithTimeout creates a new GetLTENetworkIDGatewaysGatewayIDCellularNgcParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLTENetworkIDGatewaysGatewayIDCellularNgcParamsWithTimeout(timeout time.Duration) *GetLTENetworkIDGatewaysGatewayIDCellularNgcParams {
	var ()
	return &GetLTENetworkIDGatewaysGatewayIDCellularNgcParams{

		timeout: timeout,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDCellularNgcParamsWithContext creates a new GetLTENetworkIDGatewaysGatewayIDCellularNgcParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLTENetworkIDGatewaysGatewayIDCellularNgcParamsWithContext(ctx context.Context) *GetLTENetworkIDGatewaysGatewayIDCellularNgcParams {
	var ()
	return &GetLTENetworkIDGatewaysGatewayIDCellularNgcParams{

		Context: ctx,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDCellularNgcParamsWithHTTPClient creates a new GetLTENetworkIDGatewaysGatewayIDCellularNgcParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLTENetworkIDGatewaysGatewayIDCellularNgcParamsWithHTTPClient(client *http.Client) *GetLTENetworkIDGatewaysGatewayIDCellularNgcParams {
	var ()
	return &GetLTENetworkIDGatewaysGatewayIDCellularNgcParams{
		HTTPClient: client,
	}
}

/*GetLTENetworkIDGatewaysGatewayIDCellularNgcParams contains all the parameters to send to the API endpoint
for the get LTE network ID gateways gateway ID cellular ngc operation typically these are written to a http.Request
*/
type GetLTENetworkIDGatewaysGatewayIDCellularNgcParams struct {

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

// WithTimeout adds the timeout to the get LTE network ID gateways gateway ID cellular ngc params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularNgcParams) WithTimeout(timeout time.Duration) *GetLTENetworkIDGatewaysGatewayIDCellularNgcParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get LTE network ID gateways gateway ID cellular ngc params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularNgcParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get LTE network ID gateways gateway ID cellular ngc params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularNgcParams) WithContext(ctx context.Context) *GetLTENetworkIDGatewaysGatewayIDCellularNgcParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get LTE network ID gateways gateway ID cellular ngc params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularNgcParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get LTE network ID gateways gateway ID cellular ngc params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularNgcParams) WithHTTPClient(client *http.Client) *GetLTENetworkIDGatewaysGatewayIDCellularNgcParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get LTE network ID gateways gateway ID cellular ngc params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularNgcParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the get LTE network ID gateways gateway ID cellular ngc params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularNgcParams) WithGatewayID(gatewayID string) *GetLTENetworkIDGatewaysGatewayIDCellularNgcParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the get LTE network ID gateways gateway ID cellular ngc params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularNgcParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the get LTE network ID gateways gateway ID cellular ngc params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularNgcParams) WithNetworkID(networkID string) *GetLTENetworkIDGatewaysGatewayIDCellularNgcParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get LTE network ID gateways gateway ID cellular ngc params
func (o *GetLTENetworkIDGatewaysGatewayIDCellularNgcParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLTENetworkIDGatewaysGatewayIDCellularNgcParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
