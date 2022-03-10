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

// NewPostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams creates a new PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams object
// with the default values initialized.
func NewPostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams() *PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams {
	var ()
	return &PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostFegNetworkIDSubscriberConfigBaseNamesBaseNameParamsWithTimeout creates a new PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostFegNetworkIDSubscriberConfigBaseNamesBaseNameParamsWithTimeout(timeout time.Duration) *PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams {
	var ()
	return &PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams{

		timeout: timeout,
	}
}

// NewPostFegNetworkIDSubscriberConfigBaseNamesBaseNameParamsWithContext creates a new PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostFegNetworkIDSubscriberConfigBaseNamesBaseNameParamsWithContext(ctx context.Context) *PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams {
	var ()
	return &PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams{

		Context: ctx,
	}
}

// NewPostFegNetworkIDSubscriberConfigBaseNamesBaseNameParamsWithHTTPClient creates a new PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostFegNetworkIDSubscriberConfigBaseNamesBaseNameParamsWithHTTPClient(client *http.Client) *PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams {
	var ()
	return &PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams{
		HTTPClient: client,
	}
}

/*PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams contains all the parameters to send to the API endpoint
for the post feg network ID subscriber config base names base name operation typically these are written to a http.Request
*/
type PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams struct {

	/*BaseName
	  Charging Rule Base Name

	*/
	BaseName string
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post feg network ID subscriber config base names base name params
func (o *PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams) WithTimeout(timeout time.Duration) *PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post feg network ID subscriber config base names base name params
func (o *PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post feg network ID subscriber config base names base name params
func (o *PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams) WithContext(ctx context.Context) *PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post feg network ID subscriber config base names base name params
func (o *PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post feg network ID subscriber config base names base name params
func (o *PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams) WithHTTPClient(client *http.Client) *PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post feg network ID subscriber config base names base name params
func (o *PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseName adds the baseName to the post feg network ID subscriber config base names base name params
func (o *PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams) WithBaseName(baseName string) *PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams {
	o.SetBaseName(baseName)
	return o
}

// SetBaseName adds the baseName to the post feg network ID subscriber config base names base name params
func (o *PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams) SetBaseName(baseName string) {
	o.BaseName = baseName
}

// WithNetworkID adds the networkID to the post feg network ID subscriber config base names base name params
func (o *PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams) WithNetworkID(networkID string) *PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post feg network ID subscriber config base names base name params
func (o *PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PostFegNetworkIDSubscriberConfigBaseNamesBaseNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param base_name
	if err := r.SetPathParam("base_name", o.BaseName); err != nil {
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
