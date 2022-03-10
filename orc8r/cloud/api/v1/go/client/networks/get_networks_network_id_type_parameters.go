// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewGetNetworksNetworkIDTypeParams creates a new GetNetworksNetworkIDTypeParams object
// with the default values initialized.
func NewGetNetworksNetworkIDTypeParams() *GetNetworksNetworkIDTypeParams {
	var ()
	return &GetNetworksNetworkIDTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworksNetworkIDTypeParamsWithTimeout creates a new GetNetworksNetworkIDTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworksNetworkIDTypeParamsWithTimeout(timeout time.Duration) *GetNetworksNetworkIDTypeParams {
	var ()
	return &GetNetworksNetworkIDTypeParams{

		timeout: timeout,
	}
}

// NewGetNetworksNetworkIDTypeParamsWithContext creates a new GetNetworksNetworkIDTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworksNetworkIDTypeParamsWithContext(ctx context.Context) *GetNetworksNetworkIDTypeParams {
	var ()
	return &GetNetworksNetworkIDTypeParams{

		Context: ctx,
	}
}

// NewGetNetworksNetworkIDTypeParamsWithHTTPClient creates a new GetNetworksNetworkIDTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworksNetworkIDTypeParamsWithHTTPClient(client *http.Client) *GetNetworksNetworkIDTypeParams {
	var ()
	return &GetNetworksNetworkIDTypeParams{
		HTTPClient: client,
	}
}

/*GetNetworksNetworkIDTypeParams contains all the parameters to send to the API endpoint
for the get networks network ID type operation typically these are written to a http.Request
*/
type GetNetworksNetworkIDTypeParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get networks network ID type params
func (o *GetNetworksNetworkIDTypeParams) WithTimeout(timeout time.Duration) *GetNetworksNetworkIDTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get networks network ID type params
func (o *GetNetworksNetworkIDTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get networks network ID type params
func (o *GetNetworksNetworkIDTypeParams) WithContext(ctx context.Context) *GetNetworksNetworkIDTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get networks network ID type params
func (o *GetNetworksNetworkIDTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get networks network ID type params
func (o *GetNetworksNetworkIDTypeParams) WithHTTPClient(client *http.Client) *GetNetworksNetworkIDTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get networks network ID type params
func (o *GetNetworksNetworkIDTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get networks network ID type params
func (o *GetNetworksNetworkIDTypeParams) WithNetworkID(networkID string) *GetNetworksNetworkIDTypeParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get networks network ID type params
func (o *GetNetworksNetworkIDTypeParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworksNetworkIDTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
