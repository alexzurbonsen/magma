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
)

// NewGetFegParams creates a new GetFegParams object
// with the default values initialized.
func NewGetFegParams() *GetFegParams {

	return &GetFegParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFegParamsWithTimeout creates a new GetFegParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFegParamsWithTimeout(timeout time.Duration) *GetFegParams {

	return &GetFegParams{

		timeout: timeout,
	}
}

// NewGetFegParamsWithContext creates a new GetFegParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFegParamsWithContext(ctx context.Context) *GetFegParams {

	return &GetFegParams{

		Context: ctx,
	}
}

// NewGetFegParamsWithHTTPClient creates a new GetFegParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFegParamsWithHTTPClient(client *http.Client) *GetFegParams {

	return &GetFegParams{
		HTTPClient: client,
	}
}

/*GetFegParams contains all the parameters to send to the API endpoint
for the get feg operation typically these are written to a http.Request
*/
type GetFegParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get feg params
func (o *GetFegParams) WithTimeout(timeout time.Duration) *GetFegParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get feg params
func (o *GetFegParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get feg params
func (o *GetFegParams) WithContext(ctx context.Context) *GetFegParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get feg params
func (o *GetFegParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get feg params
func (o *GetFegParams) WithHTTPClient(client *http.Client) *GetFegParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get feg params
func (o *GetFegParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetFegParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
