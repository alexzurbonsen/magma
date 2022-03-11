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

// NewPutFegNetworkIDSubscriberConfigRuleNamesParams creates a new PutFegNetworkIDSubscriberConfigRuleNamesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutFegNetworkIDSubscriberConfigRuleNamesParams() *PutFegNetworkIDSubscriberConfigRuleNamesParams {
	return &PutFegNetworkIDSubscriberConfigRuleNamesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutFegNetworkIDSubscriberConfigRuleNamesParamsWithTimeout creates a new PutFegNetworkIDSubscriberConfigRuleNamesParams object
// with the ability to set a timeout on a request.
func NewPutFegNetworkIDSubscriberConfigRuleNamesParamsWithTimeout(timeout time.Duration) *PutFegNetworkIDSubscriberConfigRuleNamesParams {
	return &PutFegNetworkIDSubscriberConfigRuleNamesParams{
		timeout: timeout,
	}
}

// NewPutFegNetworkIDSubscriberConfigRuleNamesParamsWithContext creates a new PutFegNetworkIDSubscriberConfigRuleNamesParams object
// with the ability to set a context for a request.
func NewPutFegNetworkIDSubscriberConfigRuleNamesParamsWithContext(ctx context.Context) *PutFegNetworkIDSubscriberConfigRuleNamesParams {
	return &PutFegNetworkIDSubscriberConfigRuleNamesParams{
		Context: ctx,
	}
}

// NewPutFegNetworkIDSubscriberConfigRuleNamesParamsWithHTTPClient creates a new PutFegNetworkIDSubscriberConfigRuleNamesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutFegNetworkIDSubscriberConfigRuleNamesParamsWithHTTPClient(client *http.Client) *PutFegNetworkIDSubscriberConfigRuleNamesParams {
	return &PutFegNetworkIDSubscriberConfigRuleNamesParams{
		HTTPClient: client,
	}
}

/* PutFegNetworkIDSubscriberConfigRuleNamesParams contains all the parameters to send to the API endpoint
   for the put feg network ID subscriber config rule names operation.

   Typically these are written to a http.Request.
*/
type PutFegNetworkIDSubscriberConfigRuleNamesParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* Record.

	   Subscriber Config for the Network
	*/
	Record models.RuleNames

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put feg network ID subscriber config rule names params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutFegNetworkIDSubscriberConfigRuleNamesParams) WithDefaults() *PutFegNetworkIDSubscriberConfigRuleNamesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put feg network ID subscriber config rule names params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutFegNetworkIDSubscriberConfigRuleNamesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put feg network ID subscriber config rule names params
func (o *PutFegNetworkIDSubscriberConfigRuleNamesParams) WithTimeout(timeout time.Duration) *PutFegNetworkIDSubscriberConfigRuleNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put feg network ID subscriber config rule names params
func (o *PutFegNetworkIDSubscriberConfigRuleNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put feg network ID subscriber config rule names params
func (o *PutFegNetworkIDSubscriberConfigRuleNamesParams) WithContext(ctx context.Context) *PutFegNetworkIDSubscriberConfigRuleNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put feg network ID subscriber config rule names params
func (o *PutFegNetworkIDSubscriberConfigRuleNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put feg network ID subscriber config rule names params
func (o *PutFegNetworkIDSubscriberConfigRuleNamesParams) WithHTTPClient(client *http.Client) *PutFegNetworkIDSubscriberConfigRuleNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put feg network ID subscriber config rule names params
func (o *PutFegNetworkIDSubscriberConfigRuleNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the put feg network ID subscriber config rule names params
func (o *PutFegNetworkIDSubscriberConfigRuleNamesParams) WithNetworkID(networkID string) *PutFegNetworkIDSubscriberConfigRuleNamesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put feg network ID subscriber config rule names params
func (o *PutFegNetworkIDSubscriberConfigRuleNamesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRecord adds the record to the put feg network ID subscriber config rule names params
func (o *PutFegNetworkIDSubscriberConfigRuleNamesParams) WithRecord(record models.RuleNames) *PutFegNetworkIDSubscriberConfigRuleNamesParams {
	o.SetRecord(record)
	return o
}

// SetRecord adds the record to the put feg network ID subscriber config rule names params
func (o *PutFegNetworkIDSubscriberConfigRuleNamesParams) SetRecord(record models.RuleNames) {
	o.Record = record
}

// WriteToRequest writes these params to a swagger request
func (o *PutFegNetworkIDSubscriberConfigRuleNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}
	if o.Record != nil {
		if err := r.SetBodyParam(o.Record); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
