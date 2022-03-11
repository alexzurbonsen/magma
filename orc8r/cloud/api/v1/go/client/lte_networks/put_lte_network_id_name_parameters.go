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

// NewPutLTENetworkIDNameParams creates a new PutLTENetworkIDNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutLTENetworkIDNameParams() *PutLTENetworkIDNameParams {
	return &PutLTENetworkIDNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutLTENetworkIDNameParamsWithTimeout creates a new PutLTENetworkIDNameParams object
// with the ability to set a timeout on a request.
func NewPutLTENetworkIDNameParamsWithTimeout(timeout time.Duration) *PutLTENetworkIDNameParams {
	return &PutLTENetworkIDNameParams{
		timeout: timeout,
	}
}

// NewPutLTENetworkIDNameParamsWithContext creates a new PutLTENetworkIDNameParams object
// with the ability to set a context for a request.
func NewPutLTENetworkIDNameParamsWithContext(ctx context.Context) *PutLTENetworkIDNameParams {
	return &PutLTENetworkIDNameParams{
		Context: ctx,
	}
}

// NewPutLTENetworkIDNameParamsWithHTTPClient creates a new PutLTENetworkIDNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutLTENetworkIDNameParamsWithHTTPClient(client *http.Client) *PutLTENetworkIDNameParams {
	return &PutLTENetworkIDNameParams{
		HTTPClient: client,
	}
}

/* PutLTENetworkIDNameParams contains all the parameters to send to the API endpoint
   for the put LTE network ID name operation.

   Typically these are written to a http.Request.
*/
type PutLTENetworkIDNameParams struct {

	/* Name.

	   New name for the network
	*/
	Name models.NetworkName

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put LTE network ID name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutLTENetworkIDNameParams) WithDefaults() *PutLTENetworkIDNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put LTE network ID name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutLTENetworkIDNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put LTE network ID name params
func (o *PutLTENetworkIDNameParams) WithTimeout(timeout time.Duration) *PutLTENetworkIDNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put LTE network ID name params
func (o *PutLTENetworkIDNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put LTE network ID name params
func (o *PutLTENetworkIDNameParams) WithContext(ctx context.Context) *PutLTENetworkIDNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put LTE network ID name params
func (o *PutLTENetworkIDNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put LTE network ID name params
func (o *PutLTENetworkIDNameParams) WithHTTPClient(client *http.Client) *PutLTENetworkIDNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put LTE network ID name params
func (o *PutLTENetworkIDNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the put LTE network ID name params
func (o *PutLTENetworkIDNameParams) WithName(name models.NetworkName) *PutLTENetworkIDNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put LTE network ID name params
func (o *PutLTENetworkIDNameParams) SetName(name models.NetworkName) {
	o.Name = name
}

// WithNetworkID adds the networkID to the put LTE network ID name params
func (o *PutLTENetworkIDNameParams) WithNetworkID(networkID string) *PutLTENetworkIDNameParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put LTE network ID name params
func (o *PutLTENetworkIDNameParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutLTENetworkIDNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Name); err != nil {
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
