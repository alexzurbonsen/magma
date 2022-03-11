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

// NewPutLTENetworkIDGatewaysGatewayIDDescriptionParams creates a new PutLTENetworkIDGatewaysGatewayIDDescriptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutLTENetworkIDGatewaysGatewayIDDescriptionParams() *PutLTENetworkIDGatewaysGatewayIDDescriptionParams {
	return &PutLTENetworkIDGatewaysGatewayIDDescriptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDDescriptionParamsWithTimeout creates a new PutLTENetworkIDGatewaysGatewayIDDescriptionParams object
// with the ability to set a timeout on a request.
func NewPutLTENetworkIDGatewaysGatewayIDDescriptionParamsWithTimeout(timeout time.Duration) *PutLTENetworkIDGatewaysGatewayIDDescriptionParams {
	return &PutLTENetworkIDGatewaysGatewayIDDescriptionParams{
		timeout: timeout,
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDDescriptionParamsWithContext creates a new PutLTENetworkIDGatewaysGatewayIDDescriptionParams object
// with the ability to set a context for a request.
func NewPutLTENetworkIDGatewaysGatewayIDDescriptionParamsWithContext(ctx context.Context) *PutLTENetworkIDGatewaysGatewayIDDescriptionParams {
	return &PutLTENetworkIDGatewaysGatewayIDDescriptionParams{
		Context: ctx,
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDDescriptionParamsWithHTTPClient creates a new PutLTENetworkIDGatewaysGatewayIDDescriptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutLTENetworkIDGatewaysGatewayIDDescriptionParamsWithHTTPClient(client *http.Client) *PutLTENetworkIDGatewaysGatewayIDDescriptionParams {
	return &PutLTENetworkIDGatewaysGatewayIDDescriptionParams{
		HTTPClient: client,
	}
}

/* PutLTENetworkIDGatewaysGatewayIDDescriptionParams contains all the parameters to send to the API endpoint
   for the put LTE network ID gateways gateway ID description operation.

   Typically these are written to a http.Request.
*/
type PutLTENetworkIDGatewaysGatewayIDDescriptionParams struct {

	// Description.
	Description models.GatewayDescription

	/* GatewayID.

	   Gateway ID
	*/
	GatewayID string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put LTE network ID gateways gateway ID description params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionParams) WithDefaults() *PutLTENetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put LTE network ID gateways gateway ID description params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put LTE network ID gateways gateway ID description params
func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionParams) WithTimeout(timeout time.Duration) *PutLTENetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put LTE network ID gateways gateway ID description params
func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put LTE network ID gateways gateway ID description params
func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionParams) WithContext(ctx context.Context) *PutLTENetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put LTE network ID gateways gateway ID description params
func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put LTE network ID gateways gateway ID description params
func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionParams) WithHTTPClient(client *http.Client) *PutLTENetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put LTE network ID gateways gateway ID description params
func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the put LTE network ID gateways gateway ID description params
func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionParams) WithDescription(description models.GatewayDescription) *PutLTENetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the put LTE network ID gateways gateway ID description params
func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionParams) SetDescription(description models.GatewayDescription) {
	o.Description = description
}

// WithGatewayID adds the gatewayID to the put LTE network ID gateways gateway ID description params
func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionParams) WithGatewayID(gatewayID string) *PutLTENetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the put LTE network ID gateways gateway ID description params
func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the put LTE network ID gateways gateway ID description params
func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionParams) WithNetworkID(networkID string) *PutLTENetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put LTE network ID gateways gateway ID description params
func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutLTENetworkIDGatewaysGatewayIDDescriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
