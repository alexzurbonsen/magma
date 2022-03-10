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

// NewGetNetworksParams creates a new GetNetworksParams object
// with the default values initialized.
func NewGetNetworksParams() *GetNetworksParams {

	return &GetNetworksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworksParamsWithTimeout creates a new GetNetworksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworksParamsWithTimeout(timeout time.Duration) *GetNetworksParams {

	return &GetNetworksParams{

		timeout: timeout,
	}
}

// NewGetNetworksParamsWithContext creates a new GetNetworksParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworksParamsWithContext(ctx context.Context) *GetNetworksParams {

	return &GetNetworksParams{

		Context: ctx,
	}
}

// NewGetNetworksParamsWithHTTPClient creates a new GetNetworksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworksParamsWithHTTPClient(client *http.Client) *GetNetworksParams {

	return &GetNetworksParams{
		HTTPClient: client,
	}
}

/*GetNetworksParams contains all the parameters to send to the API endpoint
for the get networks operation typically these are written to a http.Request
*/
type GetNetworksParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get networks params
func (o *GetNetworksParams) WithTimeout(timeout time.Duration) *GetNetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get networks params
func (o *GetNetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get networks params
func (o *GetNetworksParams) WithContext(ctx context.Context) *GetNetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get networks params
func (o *GetNetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get networks params
func (o *GetNetworksParams) WithHTTPClient(client *http.Client) *GetNetworksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get networks params
func (o *GetNetworksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
