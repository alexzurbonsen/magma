// Code generated by go-swagger; DO NOT EDIT.

package federation_networks

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

// NewPutFegNetworkIDParams creates a new PutFegNetworkIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutFegNetworkIDParams() *PutFegNetworkIDParams {
	return &PutFegNetworkIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutFegNetworkIDParamsWithTimeout creates a new PutFegNetworkIDParams object
// with the ability to set a timeout on a request.
func NewPutFegNetworkIDParamsWithTimeout(timeout time.Duration) *PutFegNetworkIDParams {
	return &PutFegNetworkIDParams{
		timeout: timeout,
	}
}

// NewPutFegNetworkIDParamsWithContext creates a new PutFegNetworkIDParams object
// with the ability to set a context for a request.
func NewPutFegNetworkIDParamsWithContext(ctx context.Context) *PutFegNetworkIDParams {
	return &PutFegNetworkIDParams{
		Context: ctx,
	}
}

// NewPutFegNetworkIDParamsWithHTTPClient creates a new PutFegNetworkIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutFegNetworkIDParamsWithHTTPClient(client *http.Client) *PutFegNetworkIDParams {
	return &PutFegNetworkIDParams{
		HTTPClient: client,
	}
}

/* PutFegNetworkIDParams contains all the parameters to send to the API endpoint
   for the put feg network ID operation.

   Typically these are written to a http.Request.
*/
type PutFegNetworkIDParams struct {

	/* FegNetwork.

	   Full desired configuration of the network
	*/
	FegNetwork *models.FegNetwork

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put feg network ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutFegNetworkIDParams) WithDefaults() *PutFegNetworkIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put feg network ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutFegNetworkIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put feg network ID params
func (o *PutFegNetworkIDParams) WithTimeout(timeout time.Duration) *PutFegNetworkIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put feg network ID params
func (o *PutFegNetworkIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put feg network ID params
func (o *PutFegNetworkIDParams) WithContext(ctx context.Context) *PutFegNetworkIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put feg network ID params
func (o *PutFegNetworkIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put feg network ID params
func (o *PutFegNetworkIDParams) WithHTTPClient(client *http.Client) *PutFegNetworkIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put feg network ID params
func (o *PutFegNetworkIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFegNetwork adds the fegNetwork to the put feg network ID params
func (o *PutFegNetworkIDParams) WithFegNetwork(fegNetwork *models.FegNetwork) *PutFegNetworkIDParams {
	o.SetFegNetwork(fegNetwork)
	return o
}

// SetFegNetwork adds the fegNetwork to the put feg network ID params
func (o *PutFegNetworkIDParams) SetFegNetwork(fegNetwork *models.FegNetwork) {
	o.FegNetwork = fegNetwork
}

// WithNetworkID adds the networkID to the put feg network ID params
func (o *PutFegNetworkIDParams) WithNetworkID(networkID string) *PutFegNetworkIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put feg network ID params
func (o *PutFegNetworkIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutFegNetworkIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.FegNetwork != nil {
		if err := r.SetBodyParam(o.FegNetwork); err != nil {
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
