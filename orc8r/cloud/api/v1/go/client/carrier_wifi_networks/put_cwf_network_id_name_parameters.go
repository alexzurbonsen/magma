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

// NewPutCwfNetworkIDNameParams creates a new PutCwfNetworkIDNameParams object
// with the default values initialized.
func NewPutCwfNetworkIDNameParams() *PutCwfNetworkIDNameParams {
	var ()
	return &PutCwfNetworkIDNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCwfNetworkIDNameParamsWithTimeout creates a new PutCwfNetworkIDNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCwfNetworkIDNameParamsWithTimeout(timeout time.Duration) *PutCwfNetworkIDNameParams {
	var ()
	return &PutCwfNetworkIDNameParams{

		timeout: timeout,
	}
}

// NewPutCwfNetworkIDNameParamsWithContext creates a new PutCwfNetworkIDNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCwfNetworkIDNameParamsWithContext(ctx context.Context) *PutCwfNetworkIDNameParams {
	var ()
	return &PutCwfNetworkIDNameParams{

		Context: ctx,
	}
}

// NewPutCwfNetworkIDNameParamsWithHTTPClient creates a new PutCwfNetworkIDNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCwfNetworkIDNameParamsWithHTTPClient(client *http.Client) *PutCwfNetworkIDNameParams {
	var ()
	return &PutCwfNetworkIDNameParams{
		HTTPClient: client,
	}
}

/*PutCwfNetworkIDNameParams contains all the parameters to send to the API endpoint
for the put cwf network ID name operation typically these are written to a http.Request
*/
type PutCwfNetworkIDNameParams struct {

	/*Name
	  New name for the network

	*/
	Name models.NetworkName
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put cwf network ID name params
func (o *PutCwfNetworkIDNameParams) WithTimeout(timeout time.Duration) *PutCwfNetworkIDNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put cwf network ID name params
func (o *PutCwfNetworkIDNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put cwf network ID name params
func (o *PutCwfNetworkIDNameParams) WithContext(ctx context.Context) *PutCwfNetworkIDNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put cwf network ID name params
func (o *PutCwfNetworkIDNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put cwf network ID name params
func (o *PutCwfNetworkIDNameParams) WithHTTPClient(client *http.Client) *PutCwfNetworkIDNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put cwf network ID name params
func (o *PutCwfNetworkIDNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the put cwf network ID name params
func (o *PutCwfNetworkIDNameParams) WithName(name models.NetworkName) *PutCwfNetworkIDNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put cwf network ID name params
func (o *PutCwfNetworkIDNameParams) SetName(name models.NetworkName) {
	o.Name = name
}

// WithNetworkID adds the networkID to the put cwf network ID name params
func (o *PutCwfNetworkIDNameParams) WithNetworkID(networkID string) *PutCwfNetworkIDNameParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put cwf network ID name params
func (o *PutCwfNetworkIDNameParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutCwfNetworkIDNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
