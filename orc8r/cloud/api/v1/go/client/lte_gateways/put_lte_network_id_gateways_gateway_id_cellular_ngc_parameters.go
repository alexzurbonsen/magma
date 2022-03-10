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

	"magma/orc8r/cloud/api/v1/go/models"
)

// NewPutLTENetworkIDGatewaysGatewayIDCellularNgcParams creates a new PutLTENetworkIDGatewaysGatewayIDCellularNgcParams object
// with the default values initialized.
func NewPutLTENetworkIDGatewaysGatewayIDCellularNgcParams() *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams {
	var ()
	return &PutLTENetworkIDGatewaysGatewayIDCellularNgcParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDCellularNgcParamsWithTimeout creates a new PutLTENetworkIDGatewaysGatewayIDCellularNgcParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLTENetworkIDGatewaysGatewayIDCellularNgcParamsWithTimeout(timeout time.Duration) *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams {
	var ()
	return &PutLTENetworkIDGatewaysGatewayIDCellularNgcParams{

		timeout: timeout,
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDCellularNgcParamsWithContext creates a new PutLTENetworkIDGatewaysGatewayIDCellularNgcParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLTENetworkIDGatewaysGatewayIDCellularNgcParamsWithContext(ctx context.Context) *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams {
	var ()
	return &PutLTENetworkIDGatewaysGatewayIDCellularNgcParams{

		Context: ctx,
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDCellularNgcParamsWithHTTPClient creates a new PutLTENetworkIDGatewaysGatewayIDCellularNgcParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLTENetworkIDGatewaysGatewayIDCellularNgcParamsWithHTTPClient(client *http.Client) *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams {
	var ()
	return &PutLTENetworkIDGatewaysGatewayIDCellularNgcParams{
		HTTPClient: client,
	}
}

/*PutLTENetworkIDGatewaysGatewayIDCellularNgcParams contains all the parameters to send to the API endpoint
for the put LTE network ID gateways gateway ID cellular ngc operation typically these are written to a http.Request
*/
type PutLTENetworkIDGatewaysGatewayIDCellularNgcParams struct {

	/*Config
	  New NGC configuration

	*/
	Config *models.GatewayNgcConfigs
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

// WithTimeout adds the timeout to the put LTE network ID gateways gateway ID cellular ngc params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams) WithTimeout(timeout time.Duration) *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put LTE network ID gateways gateway ID cellular ngc params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put LTE network ID gateways gateway ID cellular ngc params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams) WithContext(ctx context.Context) *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put LTE network ID gateways gateway ID cellular ngc params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put LTE network ID gateways gateway ID cellular ngc params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams) WithHTTPClient(client *http.Client) *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put LTE network ID gateways gateway ID cellular ngc params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the put LTE network ID gateways gateway ID cellular ngc params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams) WithConfig(config *models.GatewayNgcConfigs) *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the put LTE network ID gateways gateway ID cellular ngc params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams) SetConfig(config *models.GatewayNgcConfigs) {
	o.Config = config
}

// WithGatewayID adds the gatewayID to the put LTE network ID gateways gateway ID cellular ngc params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams) WithGatewayID(gatewayID string) *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the put LTE network ID gateways gateway ID cellular ngc params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the put LTE network ID gateways gateway ID cellular ngc params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams) WithNetworkID(networkID string) *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put LTE network ID gateways gateway ID cellular ngc params
func (o *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutLTENetworkIDGatewaysGatewayIDCellularNgcParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Config != nil {
		if err := r.SetBodyParam(o.Config); err != nil {
			return err
		}
	}

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
