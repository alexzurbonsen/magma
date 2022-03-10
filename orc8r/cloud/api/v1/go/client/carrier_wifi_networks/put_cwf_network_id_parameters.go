// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

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

// NewPutCwfNetworkIDParams creates a new PutCwfNetworkIDParams object
// with the default values initialized.
func NewPutCwfNetworkIDParams() *PutCwfNetworkIDParams {
	var ()
	return &PutCwfNetworkIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCwfNetworkIDParamsWithTimeout creates a new PutCwfNetworkIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCwfNetworkIDParamsWithTimeout(timeout time.Duration) *PutCwfNetworkIDParams {
	var ()
	return &PutCwfNetworkIDParams{

		timeout: timeout,
	}
}

// NewPutCwfNetworkIDParamsWithContext creates a new PutCwfNetworkIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCwfNetworkIDParamsWithContext(ctx context.Context) *PutCwfNetworkIDParams {
	var ()
	return &PutCwfNetworkIDParams{

		Context: ctx,
	}
}

// NewPutCwfNetworkIDParamsWithHTTPClient creates a new PutCwfNetworkIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCwfNetworkIDParamsWithHTTPClient(client *http.Client) *PutCwfNetworkIDParams {
	var ()
	return &PutCwfNetworkIDParams{
		HTTPClient: client,
	}
}

/*PutCwfNetworkIDParams contains all the parameters to send to the API endpoint
for the put cwf network ID operation typically these are written to a http.Request
*/
type PutCwfNetworkIDParams struct {

	/*CwfNetwork
	  Full desired configuration of the network

	*/
	CwfNetwork *models.CwfNetwork
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put cwf network ID params
func (o *PutCwfNetworkIDParams) WithTimeout(timeout time.Duration) *PutCwfNetworkIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put cwf network ID params
func (o *PutCwfNetworkIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put cwf network ID params
func (o *PutCwfNetworkIDParams) WithContext(ctx context.Context) *PutCwfNetworkIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put cwf network ID params
func (o *PutCwfNetworkIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put cwf network ID params
func (o *PutCwfNetworkIDParams) WithHTTPClient(client *http.Client) *PutCwfNetworkIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put cwf network ID params
func (o *PutCwfNetworkIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCwfNetwork adds the cwfNetwork to the put cwf network ID params
func (o *PutCwfNetworkIDParams) WithCwfNetwork(cwfNetwork *models.CwfNetwork) *PutCwfNetworkIDParams {
	o.SetCwfNetwork(cwfNetwork)
	return o
}

// SetCwfNetwork adds the cwfNetwork to the put cwf network ID params
func (o *PutCwfNetworkIDParams) SetCwfNetwork(cwfNetwork *models.CwfNetwork) {
	o.CwfNetwork = cwfNetwork
}

// WithNetworkID adds the networkID to the put cwf network ID params
func (o *PutCwfNetworkIDParams) WithNetworkID(networkID string) *PutCwfNetworkIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put cwf network ID params
func (o *PutCwfNetworkIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutCwfNetworkIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CwfNetwork != nil {
		if err := r.SetBodyParam(o.CwfNetwork); err != nil {
			return err
		}
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
