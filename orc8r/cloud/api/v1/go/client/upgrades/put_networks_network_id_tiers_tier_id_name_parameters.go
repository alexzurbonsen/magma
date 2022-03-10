// Code generated by go-swagger; DO NOT EDIT.

package upgrades

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

// NewPutNetworksNetworkIDTiersTierIDNameParams creates a new PutNetworksNetworkIDTiersTierIDNameParams object
// with the default values initialized.
func NewPutNetworksNetworkIDTiersTierIDNameParams() *PutNetworksNetworkIDTiersTierIDNameParams {
	var ()
	return &PutNetworksNetworkIDTiersTierIDNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutNetworksNetworkIDTiersTierIDNameParamsWithTimeout creates a new PutNetworksNetworkIDTiersTierIDNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutNetworksNetworkIDTiersTierIDNameParamsWithTimeout(timeout time.Duration) *PutNetworksNetworkIDTiersTierIDNameParams {
	var ()
	return &PutNetworksNetworkIDTiersTierIDNameParams{

		timeout: timeout,
	}
}

// NewPutNetworksNetworkIDTiersTierIDNameParamsWithContext creates a new PutNetworksNetworkIDTiersTierIDNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutNetworksNetworkIDTiersTierIDNameParamsWithContext(ctx context.Context) *PutNetworksNetworkIDTiersTierIDNameParams {
	var ()
	return &PutNetworksNetworkIDTiersTierIDNameParams{

		Context: ctx,
	}
}

// NewPutNetworksNetworkIDTiersTierIDNameParamsWithHTTPClient creates a new PutNetworksNetworkIDTiersTierIDNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutNetworksNetworkIDTiersTierIDNameParamsWithHTTPClient(client *http.Client) *PutNetworksNetworkIDTiersTierIDNameParams {
	var ()
	return &PutNetworksNetworkIDTiersTierIDNameParams{
		HTTPClient: client,
	}
}

/*PutNetworksNetworkIDTiersTierIDNameParams contains all the parameters to send to the API endpoint
for the put networks network ID tiers tier ID name operation typically these are written to a http.Request
*/
type PutNetworksNetworkIDTiersTierIDNameParams struct {

	/*Name
	  New name for the tier

	*/
	Name models.TierName
	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*TierID
	  Tier ID

	*/
	TierID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put networks network ID tiers tier ID name params
func (o *PutNetworksNetworkIDTiersTierIDNameParams) WithTimeout(timeout time.Duration) *PutNetworksNetworkIDTiersTierIDNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put networks network ID tiers tier ID name params
func (o *PutNetworksNetworkIDTiersTierIDNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put networks network ID tiers tier ID name params
func (o *PutNetworksNetworkIDTiersTierIDNameParams) WithContext(ctx context.Context) *PutNetworksNetworkIDTiersTierIDNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put networks network ID tiers tier ID name params
func (o *PutNetworksNetworkIDTiersTierIDNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put networks network ID tiers tier ID name params
func (o *PutNetworksNetworkIDTiersTierIDNameParams) WithHTTPClient(client *http.Client) *PutNetworksNetworkIDTiersTierIDNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put networks network ID tiers tier ID name params
func (o *PutNetworksNetworkIDTiersTierIDNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the put networks network ID tiers tier ID name params
func (o *PutNetworksNetworkIDTiersTierIDNameParams) WithName(name models.TierName) *PutNetworksNetworkIDTiersTierIDNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put networks network ID tiers tier ID name params
func (o *PutNetworksNetworkIDTiersTierIDNameParams) SetName(name models.TierName) {
	o.Name = name
}

// WithNetworkID adds the networkID to the put networks network ID tiers tier ID name params
func (o *PutNetworksNetworkIDTiersTierIDNameParams) WithNetworkID(networkID string) *PutNetworksNetworkIDTiersTierIDNameParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put networks network ID tiers tier ID name params
func (o *PutNetworksNetworkIDTiersTierIDNameParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithTierID adds the tierID to the put networks network ID tiers tier ID name params
func (o *PutNetworksNetworkIDTiersTierIDNameParams) WithTierID(tierID string) *PutNetworksNetworkIDTiersTierIDNameParams {
	o.SetTierID(tierID)
	return o
}

// SetTierID adds the tierId to the put networks network ID tiers tier ID name params
func (o *PutNetworksNetworkIDTiersTierIDNameParams) SetTierID(tierID string) {
	o.TierID = tierID
}

// WriteToRequest writes these params to a swagger request
func (o *PutNetworksNetworkIDTiersTierIDNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param tier_id
	if err := r.SetPathParam("tier_id", o.TierID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
