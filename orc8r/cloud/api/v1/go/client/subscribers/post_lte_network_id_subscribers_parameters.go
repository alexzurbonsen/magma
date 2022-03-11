// Code generated by go-swagger; DO NOT EDIT.

package subscribers

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

// NewPostLTENetworkIDSubscribersParams creates a new PostLTENetworkIDSubscribersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostLTENetworkIDSubscribersParams() *PostLTENetworkIDSubscribersParams {
	return &PostLTENetworkIDSubscribersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostLTENetworkIDSubscribersParamsWithTimeout creates a new PostLTENetworkIDSubscribersParams object
// with the ability to set a timeout on a request.
func NewPostLTENetworkIDSubscribersParamsWithTimeout(timeout time.Duration) *PostLTENetworkIDSubscribersParams {
	return &PostLTENetworkIDSubscribersParams{
		timeout: timeout,
	}
}

// NewPostLTENetworkIDSubscribersParamsWithContext creates a new PostLTENetworkIDSubscribersParams object
// with the ability to set a context for a request.
func NewPostLTENetworkIDSubscribersParamsWithContext(ctx context.Context) *PostLTENetworkIDSubscribersParams {
	return &PostLTENetworkIDSubscribersParams{
		Context: ctx,
	}
}

// NewPostLTENetworkIDSubscribersParamsWithHTTPClient creates a new PostLTENetworkIDSubscribersParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostLTENetworkIDSubscribersParamsWithHTTPClient(client *http.Client) *PostLTENetworkIDSubscribersParams {
	return &PostLTENetworkIDSubscribersParams{
		HTTPClient: client,
	}
}

/* PostLTENetworkIDSubscribersParams contains all the parameters to send to the API endpoint
   for the post LTE network ID subscribers operation.

   Typically these are written to a http.Request.
*/
type PostLTENetworkIDSubscribersParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* Subscribers.

	   Subscribers to add
	*/
	Subscribers models.MutableSubscribers

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post LTE network ID subscribers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLTENetworkIDSubscribersParams) WithDefaults() *PostLTENetworkIDSubscribersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post LTE network ID subscribers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLTENetworkIDSubscribersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post LTE network ID subscribers params
func (o *PostLTENetworkIDSubscribersParams) WithTimeout(timeout time.Duration) *PostLTENetworkIDSubscribersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post LTE network ID subscribers params
func (o *PostLTENetworkIDSubscribersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post LTE network ID subscribers params
func (o *PostLTENetworkIDSubscribersParams) WithContext(ctx context.Context) *PostLTENetworkIDSubscribersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post LTE network ID subscribers params
func (o *PostLTENetworkIDSubscribersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post LTE network ID subscribers params
func (o *PostLTENetworkIDSubscribersParams) WithHTTPClient(client *http.Client) *PostLTENetworkIDSubscribersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post LTE network ID subscribers params
func (o *PostLTENetworkIDSubscribersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the post LTE network ID subscribers params
func (o *PostLTENetworkIDSubscribersParams) WithNetworkID(networkID string) *PostLTENetworkIDSubscribersParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post LTE network ID subscribers params
func (o *PostLTENetworkIDSubscribersParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithSubscribers adds the subscribers to the post LTE network ID subscribers params
func (o *PostLTENetworkIDSubscribersParams) WithSubscribers(subscribers models.MutableSubscribers) *PostLTENetworkIDSubscribersParams {
	o.SetSubscribers(subscribers)
	return o
}

// SetSubscribers adds the subscribers to the post LTE network ID subscribers params
func (o *PostLTENetworkIDSubscribersParams) SetSubscribers(subscribers models.MutableSubscribers) {
	o.Subscribers = subscribers
}

// WriteToRequest writes these params to a swagger request
func (o *PostLTENetworkIDSubscribersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}
	if o.Subscribers != nil {
		if err := r.SetBodyParam(o.Subscribers); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
