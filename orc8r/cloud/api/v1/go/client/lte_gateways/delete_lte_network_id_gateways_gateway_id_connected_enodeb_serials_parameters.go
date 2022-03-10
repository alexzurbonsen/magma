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

// NewDeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams creates a new DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams object
// with the default values initialized.
func NewDeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams() *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams {
	var ()
	return &DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParamsWithTimeout creates a new DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParamsWithTimeout(timeout time.Duration) *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams {
	var ()
	return &DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams{

		timeout: timeout,
	}
}

// NewDeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParamsWithContext creates a new DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParamsWithContext(ctx context.Context) *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams {
	var ()
	return &DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams{

		Context: ctx,
	}
}

// NewDeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParamsWithHTTPClient creates a new DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParamsWithHTTPClient(client *http.Client) *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams {
	var ()
	return &DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams{
		HTTPClient: client,
	}
}

/*DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams contains all the parameters to send to the API endpoint
for the delete LTE network ID gateways gateway ID connected ENODEB serials operation typically these are written to a http.Request
*/
type DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams struct {

	/*GatewayID
	  Gateway ID

	*/
	GatewayID string
	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*Serial*/
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete LTE network ID gateways gateway ID connected ENODEB serials params
func (o *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) WithTimeout(timeout time.Duration) *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete LTE network ID gateways gateway ID connected ENODEB serials params
func (o *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete LTE network ID gateways gateway ID connected ENODEB serials params
func (o *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) WithContext(ctx context.Context) *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete LTE network ID gateways gateway ID connected ENODEB serials params
func (o *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete LTE network ID gateways gateway ID connected ENODEB serials params
func (o *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) WithHTTPClient(client *http.Client) *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete LTE network ID gateways gateway ID connected ENODEB serials params
func (o *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the delete LTE network ID gateways gateway ID connected ENODEB serials params
func (o *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) WithGatewayID(gatewayID string) *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the delete LTE network ID gateways gateway ID connected ENODEB serials params
func (o *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the delete LTE network ID gateways gateway ID connected ENODEB serials params
func (o *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) WithNetworkID(networkID string) *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete LTE network ID gateways gateway ID connected ENODEB serials params
func (o *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithSerial adds the serial to the delete LTE network ID gateways gateway ID connected ENODEB serials params
func (o *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) WithSerial(serial string) *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the delete LTE network ID gateways gateway ID connected ENODEB serials params
func (o *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLTENetworkIDGatewaysGatewayIDConnectedENODEBSerialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if err := r.SetBodyParam(o.Serial); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
