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

	"magma/orc8r/cloud/api/v1/go/models"
)

// NewPutNetworksNetworkIDGatewaysGatewayIDDescriptionParams creates a new PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams object
// with the default values initialized.
func NewPutNetworksNetworkIDGatewaysGatewayIDDescriptionParams() *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams {
	var ()
	return &PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutNetworksNetworkIDGatewaysGatewayIDDescriptionParamsWithTimeout creates a new PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutNetworksNetworkIDGatewaysGatewayIDDescriptionParamsWithTimeout(timeout time.Duration) *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams {
	var ()
	return &PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams{

		timeout: timeout,
	}
}

// NewPutNetworksNetworkIDGatewaysGatewayIDDescriptionParamsWithContext creates a new PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutNetworksNetworkIDGatewaysGatewayIDDescriptionParamsWithContext(ctx context.Context) *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams {
	var ()
	return &PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams{

		Context: ctx,
	}
}

// NewPutNetworksNetworkIDGatewaysGatewayIDDescriptionParamsWithHTTPClient creates a new PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutNetworksNetworkIDGatewaysGatewayIDDescriptionParamsWithHTTPClient(client *http.Client) *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams {
	var ()
	return &PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams{
		HTTPClient: client,
	}
}

/*PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams contains all the parameters to send to the API endpoint
for the put networks network ID gateways gateway ID description operation typically these are written to a http.Request
*/
type PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams struct {

	/*Description*/
	Description models.GatewayDescription
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

// WithTimeout adds the timeout to the put networks network ID gateways gateway ID description params
func (o *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams) WithTimeout(timeout time.Duration) *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put networks network ID gateways gateway ID description params
func (o *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put networks network ID gateways gateway ID description params
func (o *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams) WithContext(ctx context.Context) *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put networks network ID gateways gateway ID description params
func (o *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put networks network ID gateways gateway ID description params
func (o *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams) WithHTTPClient(client *http.Client) *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put networks network ID gateways gateway ID description params
func (o *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the put networks network ID gateways gateway ID description params
func (o *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams) WithDescription(description models.GatewayDescription) *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the put networks network ID gateways gateway ID description params
func (o *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams) SetDescription(description models.GatewayDescription) {
	o.Description = description
}

// WithGatewayID adds the gatewayID to the put networks network ID gateways gateway ID description params
func (o *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams) WithGatewayID(gatewayID string) *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the put networks network ID gateways gateway ID description params
func (o *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the put networks network ID gateways gateway ID description params
func (o *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams) WithNetworkID(networkID string) *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put networks network ID gateways gateway ID description params
func (o *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Description); err != nil {
		return err
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
