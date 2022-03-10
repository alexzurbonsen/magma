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

// NewPostLTENetworkIDGatewayPoolsParams creates a new PostLTENetworkIDGatewayPoolsParams object
// with the default values initialized.
func NewPostLTENetworkIDGatewayPoolsParams() *PostLTENetworkIDGatewayPoolsParams {
	var ()
	return &PostLTENetworkIDGatewayPoolsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLTENetworkIDGatewayPoolsParamsWithTimeout creates a new PostLTENetworkIDGatewayPoolsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLTENetworkIDGatewayPoolsParamsWithTimeout(timeout time.Duration) *PostLTENetworkIDGatewayPoolsParams {
	var ()
	return &PostLTENetworkIDGatewayPoolsParams{

		timeout: timeout,
	}
}

// NewPostLTENetworkIDGatewayPoolsParamsWithContext creates a new PostLTENetworkIDGatewayPoolsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLTENetworkIDGatewayPoolsParamsWithContext(ctx context.Context) *PostLTENetworkIDGatewayPoolsParams {
	var ()
	return &PostLTENetworkIDGatewayPoolsParams{

		Context: ctx,
	}
}

// NewPostLTENetworkIDGatewayPoolsParamsWithHTTPClient creates a new PostLTENetworkIDGatewayPoolsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLTENetworkIDGatewayPoolsParamsWithHTTPClient(client *http.Client) *PostLTENetworkIDGatewayPoolsParams {
	var ()
	return &PostLTENetworkIDGatewayPoolsParams{
		HTTPClient: client,
	}
}

/*PostLTENetworkIDGatewayPoolsParams contains all the parameters to send to the API endpoint
for the post LTE network ID gateway pools operation typically these are written to a http.Request
*/
type PostLTENetworkIDGatewayPoolsParams struct {

	/*HAGatewayPool
	  LTE HA gateway pool

	*/
	HAGatewayPool *models.MutableCellularGatewayPool
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post LTE network ID gateway pools params
func (o *PostLTENetworkIDGatewayPoolsParams) WithTimeout(timeout time.Duration) *PostLTENetworkIDGatewayPoolsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post LTE network ID gateway pools params
func (o *PostLTENetworkIDGatewayPoolsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post LTE network ID gateway pools params
func (o *PostLTENetworkIDGatewayPoolsParams) WithContext(ctx context.Context) *PostLTENetworkIDGatewayPoolsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post LTE network ID gateway pools params
func (o *PostLTENetworkIDGatewayPoolsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post LTE network ID gateway pools params
func (o *PostLTENetworkIDGatewayPoolsParams) WithHTTPClient(client *http.Client) *PostLTENetworkIDGatewayPoolsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post LTE network ID gateway pools params
func (o *PostLTENetworkIDGatewayPoolsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHAGatewayPool adds the hAGatewayPool to the post LTE network ID gateway pools params
func (o *PostLTENetworkIDGatewayPoolsParams) WithHAGatewayPool(hAGatewayPool *models.MutableCellularGatewayPool) *PostLTENetworkIDGatewayPoolsParams {
	o.SetHAGatewayPool(hAGatewayPool)
	return o
}

// SetHAGatewayPool adds the hAGatewayPool to the post LTE network ID gateway pools params
func (o *PostLTENetworkIDGatewayPoolsParams) SetHAGatewayPool(hAGatewayPool *models.MutableCellularGatewayPool) {
	o.HAGatewayPool = hAGatewayPool
}

// WithNetworkID adds the networkID to the post LTE network ID gateway pools params
func (o *PostLTENetworkIDGatewayPoolsParams) WithNetworkID(networkID string) *PostLTENetworkIDGatewayPoolsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post LTE network ID gateway pools params
func (o *PostLTENetworkIDGatewayPoolsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLTENetworkIDGatewayPoolsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HAGatewayPool != nil {
		if err := r.SetBodyParam(o.HAGatewayPool); err != nil {
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
