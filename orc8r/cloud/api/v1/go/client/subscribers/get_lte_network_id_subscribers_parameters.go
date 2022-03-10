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
	"github.com/go-openapi/swag"
)

// NewGetLTENetworkIDSubscribersParams creates a new GetLTENetworkIDSubscribersParams object
// with the default values initialized.
func NewGetLTENetworkIDSubscribersParams() *GetLTENetworkIDSubscribersParams {
	var ()
	return &GetLTENetworkIDSubscribersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLTENetworkIDSubscribersParamsWithTimeout creates a new GetLTENetworkIDSubscribersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLTENetworkIDSubscribersParamsWithTimeout(timeout time.Duration) *GetLTENetworkIDSubscribersParams {
	var ()
	return &GetLTENetworkIDSubscribersParams{

		timeout: timeout,
	}
}

// NewGetLTENetworkIDSubscribersParamsWithContext creates a new GetLTENetworkIDSubscribersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLTENetworkIDSubscribersParamsWithContext(ctx context.Context) *GetLTENetworkIDSubscribersParams {
	var ()
	return &GetLTENetworkIDSubscribersParams{

		Context: ctx,
	}
}

// NewGetLTENetworkIDSubscribersParamsWithHTTPClient creates a new GetLTENetworkIDSubscribersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLTENetworkIDSubscribersParamsWithHTTPClient(client *http.Client) *GetLTENetworkIDSubscribersParams {
	var ()
	return &GetLTENetworkIDSubscribersParams{
		HTTPClient: client,
	}
}

/*GetLTENetworkIDSubscribersParams contains all the parameters to send to the API endpoint
for the get LTE network ID subscribers operation typically these are written to a http.Request
*/
type GetLTENetworkIDSubscribersParams struct {

	/*IP
	  Filter to subscribers assigned the passed IP address

	*/
	IP *string
	/*Msisdn
	  Filter to subscribers with the passed MSISDN

	*/
	Msisdn *string
	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*PageSize
	  Maximum number of entities to return

	*/
	PageSize *uint32
	/*PageToken
	  Opaque page token for paginated requests

	*/
	PageToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get LTE network ID subscribers params
func (o *GetLTENetworkIDSubscribersParams) WithTimeout(timeout time.Duration) *GetLTENetworkIDSubscribersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get LTE network ID subscribers params
func (o *GetLTENetworkIDSubscribersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get LTE network ID subscribers params
func (o *GetLTENetworkIDSubscribersParams) WithContext(ctx context.Context) *GetLTENetworkIDSubscribersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get LTE network ID subscribers params
func (o *GetLTENetworkIDSubscribersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get LTE network ID subscribers params
func (o *GetLTENetworkIDSubscribersParams) WithHTTPClient(client *http.Client) *GetLTENetworkIDSubscribersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get LTE network ID subscribers params
func (o *GetLTENetworkIDSubscribersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIP adds the ip to the get LTE network ID subscribers params
func (o *GetLTENetworkIDSubscribersParams) WithIP(ip *string) *GetLTENetworkIDSubscribersParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the get LTE network ID subscribers params
func (o *GetLTENetworkIDSubscribersParams) SetIP(ip *string) {
	o.IP = ip
}

// WithMsisdn adds the msisdn to the get LTE network ID subscribers params
func (o *GetLTENetworkIDSubscribersParams) WithMsisdn(msisdn *string) *GetLTENetworkIDSubscribersParams {
	o.SetMsisdn(msisdn)
	return o
}

// SetMsisdn adds the msisdn to the get LTE network ID subscribers params
func (o *GetLTENetworkIDSubscribersParams) SetMsisdn(msisdn *string) {
	o.Msisdn = msisdn
}

// WithNetworkID adds the networkID to the get LTE network ID subscribers params
func (o *GetLTENetworkIDSubscribersParams) WithNetworkID(networkID string) *GetLTENetworkIDSubscribersParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get LTE network ID subscribers params
func (o *GetLTENetworkIDSubscribersParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithPageSize adds the pageSize to the get LTE network ID subscribers params
func (o *GetLTENetworkIDSubscribersParams) WithPageSize(pageSize *uint32) *GetLTENetworkIDSubscribersParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get LTE network ID subscribers params
func (o *GetLTENetworkIDSubscribersParams) SetPageSize(pageSize *uint32) {
	o.PageSize = pageSize
}

// WithPageToken adds the pageToken to the get LTE network ID subscribers params
func (o *GetLTENetworkIDSubscribersParams) WithPageToken(pageToken *string) *GetLTENetworkIDSubscribersParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the get LTE network ID subscribers params
func (o *GetLTENetworkIDSubscribersParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetLTENetworkIDSubscribersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IP != nil {

		// query param ip
		var qrIP string
		if o.IP != nil {
			qrIP = *o.IP
		}
		qIP := qrIP
		if qIP != "" {
			if err := r.SetQueryParam("ip", qIP); err != nil {
				return err
			}
		}

	}

	if o.Msisdn != nil {

		// query param msisdn
		var qrMsisdn string
		if o.Msisdn != nil {
			qrMsisdn = *o.Msisdn
		}
		qMsisdn := qrMsisdn
		if qMsisdn != "" {
			if err := r.SetQueryParam("msisdn", qMsisdn); err != nil {
				return err
			}
		}

	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize uint32
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatUint32(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.PageToken != nil {

		// query param page_token
		var qrPageToken string
		if o.PageToken != nil {
			qrPageToken = *o.PageToken
		}
		qPageToken := qrPageToken
		if qPageToken != "" {
			if err := r.SetQueryParam("page_token", qPageToken); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
