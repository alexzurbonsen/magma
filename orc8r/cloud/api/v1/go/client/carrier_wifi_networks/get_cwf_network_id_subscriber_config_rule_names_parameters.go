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
)

// NewGetCwfNetworkIDSubscriberConfigRuleNamesParams creates a new GetCwfNetworkIDSubscriberConfigRuleNamesParams object
// with the default values initialized.
func NewGetCwfNetworkIDSubscriberConfigRuleNamesParams() *GetCwfNetworkIDSubscriberConfigRuleNamesParams {
	var ()
	return &GetCwfNetworkIDSubscriberConfigRuleNamesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCwfNetworkIDSubscriberConfigRuleNamesParamsWithTimeout creates a new GetCwfNetworkIDSubscriberConfigRuleNamesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCwfNetworkIDSubscriberConfigRuleNamesParamsWithTimeout(timeout time.Duration) *GetCwfNetworkIDSubscriberConfigRuleNamesParams {
	var ()
	return &GetCwfNetworkIDSubscriberConfigRuleNamesParams{

		timeout: timeout,
	}
}

// NewGetCwfNetworkIDSubscriberConfigRuleNamesParamsWithContext creates a new GetCwfNetworkIDSubscriberConfigRuleNamesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCwfNetworkIDSubscriberConfigRuleNamesParamsWithContext(ctx context.Context) *GetCwfNetworkIDSubscriberConfigRuleNamesParams {
	var ()
	return &GetCwfNetworkIDSubscriberConfigRuleNamesParams{

		Context: ctx,
	}
}

// NewGetCwfNetworkIDSubscriberConfigRuleNamesParamsWithHTTPClient creates a new GetCwfNetworkIDSubscriberConfigRuleNamesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCwfNetworkIDSubscriberConfigRuleNamesParamsWithHTTPClient(client *http.Client) *GetCwfNetworkIDSubscriberConfigRuleNamesParams {
	var ()
	return &GetCwfNetworkIDSubscriberConfigRuleNamesParams{
		HTTPClient: client,
	}
}

/*GetCwfNetworkIDSubscriberConfigRuleNamesParams contains all the parameters to send to the API endpoint
for the get cwf network ID subscriber config rule names operation typically these are written to a http.Request
*/
type GetCwfNetworkIDSubscriberConfigRuleNamesParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cwf network ID subscriber config rule names params
func (o *GetCwfNetworkIDSubscriberConfigRuleNamesParams) WithTimeout(timeout time.Duration) *GetCwfNetworkIDSubscriberConfigRuleNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cwf network ID subscriber config rule names params
func (o *GetCwfNetworkIDSubscriberConfigRuleNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cwf network ID subscriber config rule names params
func (o *GetCwfNetworkIDSubscriberConfigRuleNamesParams) WithContext(ctx context.Context) *GetCwfNetworkIDSubscriberConfigRuleNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cwf network ID subscriber config rule names params
func (o *GetCwfNetworkIDSubscriberConfigRuleNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cwf network ID subscriber config rule names params
func (o *GetCwfNetworkIDSubscriberConfigRuleNamesParams) WithHTTPClient(client *http.Client) *GetCwfNetworkIDSubscriberConfigRuleNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cwf network ID subscriber config rule names params
func (o *GetCwfNetworkIDSubscriberConfigRuleNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get cwf network ID subscriber config rule names params
func (o *GetCwfNetworkIDSubscriberConfigRuleNamesParams) WithNetworkID(networkID string) *GetCwfNetworkIDSubscriberConfigRuleNamesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get cwf network ID subscriber config rule names params
func (o *GetCwfNetworkIDSubscriberConfigRuleNamesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCwfNetworkIDSubscriberConfigRuleNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
