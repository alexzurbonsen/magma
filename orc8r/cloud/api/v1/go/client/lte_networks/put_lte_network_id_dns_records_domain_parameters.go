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

// NewPutLTENetworkIDDNSRecordsDomainParams creates a new PutLTENetworkIDDNSRecordsDomainParams object
// with the default values initialized.
func NewPutLTENetworkIDDNSRecordsDomainParams() *PutLTENetworkIDDNSRecordsDomainParams {
	var ()
	return &PutLTENetworkIDDNSRecordsDomainParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLTENetworkIDDNSRecordsDomainParamsWithTimeout creates a new PutLTENetworkIDDNSRecordsDomainParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLTENetworkIDDNSRecordsDomainParamsWithTimeout(timeout time.Duration) *PutLTENetworkIDDNSRecordsDomainParams {
	var ()
	return &PutLTENetworkIDDNSRecordsDomainParams{

		timeout: timeout,
	}
}

// NewPutLTENetworkIDDNSRecordsDomainParamsWithContext creates a new PutLTENetworkIDDNSRecordsDomainParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLTENetworkIDDNSRecordsDomainParamsWithContext(ctx context.Context) *PutLTENetworkIDDNSRecordsDomainParams {
	var ()
	return &PutLTENetworkIDDNSRecordsDomainParams{

		Context: ctx,
	}
}

// NewPutLTENetworkIDDNSRecordsDomainParamsWithHTTPClient creates a new PutLTENetworkIDDNSRecordsDomainParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLTENetworkIDDNSRecordsDomainParamsWithHTTPClient(client *http.Client) *PutLTENetworkIDDNSRecordsDomainParams {
	var ()
	return &PutLTENetworkIDDNSRecordsDomainParams{
		HTTPClient: client,
	}
}

/*PutLTENetworkIDDNSRecordsDomainParams contains all the parameters to send to the API endpoint
for the put LTE network ID DNS records domain operation typically these are written to a http.Request
*/
type PutLTENetworkIDDNSRecordsDomainParams struct {

	/*Domain
	  DNS record domain

	*/
	Domain string
	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*Record
	  Custom DNS record for the domain

	*/
	Record *models.DNSConfigRecord

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put LTE network ID DNS records domain params
func (o *PutLTENetworkIDDNSRecordsDomainParams) WithTimeout(timeout time.Duration) *PutLTENetworkIDDNSRecordsDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put LTE network ID DNS records domain params
func (o *PutLTENetworkIDDNSRecordsDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put LTE network ID DNS records domain params
func (o *PutLTENetworkIDDNSRecordsDomainParams) WithContext(ctx context.Context) *PutLTENetworkIDDNSRecordsDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put LTE network ID DNS records domain params
func (o *PutLTENetworkIDDNSRecordsDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put LTE network ID DNS records domain params
func (o *PutLTENetworkIDDNSRecordsDomainParams) WithHTTPClient(client *http.Client) *PutLTENetworkIDDNSRecordsDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put LTE network ID DNS records domain params
func (o *PutLTENetworkIDDNSRecordsDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomain adds the domain to the put LTE network ID DNS records domain params
func (o *PutLTENetworkIDDNSRecordsDomainParams) WithDomain(domain string) *PutLTENetworkIDDNSRecordsDomainParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the put LTE network ID DNS records domain params
func (o *PutLTENetworkIDDNSRecordsDomainParams) SetDomain(domain string) {
	o.Domain = domain
}

// WithNetworkID adds the networkID to the put LTE network ID DNS records domain params
func (o *PutLTENetworkIDDNSRecordsDomainParams) WithNetworkID(networkID string) *PutLTENetworkIDDNSRecordsDomainParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put LTE network ID DNS records domain params
func (o *PutLTENetworkIDDNSRecordsDomainParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRecord adds the record to the put LTE network ID DNS records domain params
func (o *PutLTENetworkIDDNSRecordsDomainParams) WithRecord(record *models.DNSConfigRecord) *PutLTENetworkIDDNSRecordsDomainParams {
	o.SetRecord(record)
	return o
}

// SetRecord adds the record to the put LTE network ID DNS records domain params
func (o *PutLTENetworkIDDNSRecordsDomainParams) SetRecord(record *models.DNSConfigRecord) {
	o.Record = record
}

// WriteToRequest writes these params to a swagger request
func (o *PutLTENetworkIDDNSRecordsDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domain
	if err := r.SetPathParam("domain", o.Domain); err != nil {
		return err
	}

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
