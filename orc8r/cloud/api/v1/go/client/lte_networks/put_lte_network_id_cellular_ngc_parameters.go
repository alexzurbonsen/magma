// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

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

// NewPutLTENetworkIDCellularNgcParams creates a new PutLTENetworkIDCellularNgcParams object
// with the default values initialized.
func NewPutLTENetworkIDCellularNgcParams() *PutLTENetworkIDCellularNgcParams {
	var ()
	return &PutLTENetworkIDCellularNgcParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLTENetworkIDCellularNgcParamsWithTimeout creates a new PutLTENetworkIDCellularNgcParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLTENetworkIDCellularNgcParamsWithTimeout(timeout time.Duration) *PutLTENetworkIDCellularNgcParams {
	var ()
	return &PutLTENetworkIDCellularNgcParams{

		timeout: timeout,
	}
}

// NewPutLTENetworkIDCellularNgcParamsWithContext creates a new PutLTENetworkIDCellularNgcParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLTENetworkIDCellularNgcParamsWithContext(ctx context.Context) *PutLTENetworkIDCellularNgcParams {
	var ()
	return &PutLTENetworkIDCellularNgcParams{

		Context: ctx,
	}
}

// NewPutLTENetworkIDCellularNgcParamsWithHTTPClient creates a new PutLTENetworkIDCellularNgcParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLTENetworkIDCellularNgcParamsWithHTTPClient(client *http.Client) *PutLTENetworkIDCellularNgcParams {
	var ()
	return &PutLTENetworkIDCellularNgcParams{
		HTTPClient: client,
	}
}

/*PutLTENetworkIDCellularNgcParams contains all the parameters to send to the API endpoint
for the put LTE network ID cellular ngc operation typically these are written to a http.Request
*/
type PutLTENetworkIDCellularNgcParams struct {

	/*Config
	  New NGC configuration for the network

	*/
	Config *models.NetworkNgcConfigs
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put LTE network ID cellular ngc params
func (o *PutLTENetworkIDCellularNgcParams) WithTimeout(timeout time.Duration) *PutLTENetworkIDCellularNgcParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put LTE network ID cellular ngc params
func (o *PutLTENetworkIDCellularNgcParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put LTE network ID cellular ngc params
func (o *PutLTENetworkIDCellularNgcParams) WithContext(ctx context.Context) *PutLTENetworkIDCellularNgcParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put LTE network ID cellular ngc params
func (o *PutLTENetworkIDCellularNgcParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put LTE network ID cellular ngc params
func (o *PutLTENetworkIDCellularNgcParams) WithHTTPClient(client *http.Client) *PutLTENetworkIDCellularNgcParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put LTE network ID cellular ngc params
func (o *PutLTENetworkIDCellularNgcParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the put LTE network ID cellular ngc params
func (o *PutLTENetworkIDCellularNgcParams) WithConfig(config *models.NetworkNgcConfigs) *PutLTENetworkIDCellularNgcParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the put LTE network ID cellular ngc params
func (o *PutLTENetworkIDCellularNgcParams) SetConfig(config *models.NetworkNgcConfigs) {
	o.Config = config
}

// WithNetworkID adds the networkID to the put LTE network ID cellular ngc params
func (o *PutLTENetworkIDCellularNgcParams) WithNetworkID(networkID string) *PutLTENetworkIDCellularNgcParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put LTE network ID cellular ngc params
func (o *PutLTENetworkIDCellularNgcParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutLTENetworkIDCellularNgcParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Config != nil {
		if err := r.SetBodyParam(o.Config); err != nil {
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
