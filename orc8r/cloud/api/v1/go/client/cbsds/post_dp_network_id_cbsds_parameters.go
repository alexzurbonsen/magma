// Code generated by go-swagger; DO NOT EDIT.

package cbsds

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

// NewPostDpNetworkIDCbsdsParams creates a new PostDpNetworkIDCbsdsParams object
// with the default values initialized.
func NewPostDpNetworkIDCbsdsParams() *PostDpNetworkIDCbsdsParams {
	var ()
	return &PostDpNetworkIDCbsdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDpNetworkIDCbsdsParamsWithTimeout creates a new PostDpNetworkIDCbsdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDpNetworkIDCbsdsParamsWithTimeout(timeout time.Duration) *PostDpNetworkIDCbsdsParams {
	var ()
	return &PostDpNetworkIDCbsdsParams{

		timeout: timeout,
	}
}

// NewPostDpNetworkIDCbsdsParamsWithContext creates a new PostDpNetworkIDCbsdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDpNetworkIDCbsdsParamsWithContext(ctx context.Context) *PostDpNetworkIDCbsdsParams {
	var ()
	return &PostDpNetworkIDCbsdsParams{

		Context: ctx,
	}
}

// NewPostDpNetworkIDCbsdsParamsWithHTTPClient creates a new PostDpNetworkIDCbsdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDpNetworkIDCbsdsParamsWithHTTPClient(client *http.Client) *PostDpNetworkIDCbsdsParams {
	var ()
	return &PostDpNetworkIDCbsdsParams{
		HTTPClient: client,
	}
}

/*PostDpNetworkIDCbsdsParams contains all the parameters to send to the API endpoint
for the post dp network ID cbsds operation typically these are written to a http.Request
*/
type PostDpNetworkIDCbsdsParams struct {

	/*Cbsd
	  CBSD

	*/
	Cbsd *models.MutableCbsd
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post dp network ID cbsds params
func (o *PostDpNetworkIDCbsdsParams) WithTimeout(timeout time.Duration) *PostDpNetworkIDCbsdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post dp network ID cbsds params
func (o *PostDpNetworkIDCbsdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post dp network ID cbsds params
func (o *PostDpNetworkIDCbsdsParams) WithContext(ctx context.Context) *PostDpNetworkIDCbsdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post dp network ID cbsds params
func (o *PostDpNetworkIDCbsdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post dp network ID cbsds params
func (o *PostDpNetworkIDCbsdsParams) WithHTTPClient(client *http.Client) *PostDpNetworkIDCbsdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post dp network ID cbsds params
func (o *PostDpNetworkIDCbsdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCbsd adds the cbsd to the post dp network ID cbsds params
func (o *PostDpNetworkIDCbsdsParams) WithCbsd(cbsd *models.MutableCbsd) *PostDpNetworkIDCbsdsParams {
	o.SetCbsd(cbsd)
	return o
}

// SetCbsd adds the cbsd to the post dp network ID cbsds params
func (o *PostDpNetworkIDCbsdsParams) SetCbsd(cbsd *models.MutableCbsd) {
	o.Cbsd = cbsd
}

// WithNetworkID adds the networkID to the post dp network ID cbsds params
func (o *PostDpNetworkIDCbsdsParams) WithNetworkID(networkID string) *PostDpNetworkIDCbsdsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post dp network ID cbsds params
func (o *PostDpNetworkIDCbsdsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PostDpNetworkIDCbsdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cbsd != nil {
		if err := r.SetBodyParam(o.Cbsd); err != nil {
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
