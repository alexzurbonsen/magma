// Code generated by go-swagger; DO NOT EDIT.

package logs

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

// NewGetNetworksNetworkIDLogsSearchParams creates a new GetNetworksNetworkIDLogsSearchParams object
// with the default values initialized.
func NewGetNetworksNetworkIDLogsSearchParams() *GetNetworksNetworkIDLogsSearchParams {
	var ()
	return &GetNetworksNetworkIDLogsSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworksNetworkIDLogsSearchParamsWithTimeout creates a new GetNetworksNetworkIDLogsSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworksNetworkIDLogsSearchParamsWithTimeout(timeout time.Duration) *GetNetworksNetworkIDLogsSearchParams {
	var ()
	return &GetNetworksNetworkIDLogsSearchParams{

		timeout: timeout,
	}
}

// NewGetNetworksNetworkIDLogsSearchParamsWithContext creates a new GetNetworksNetworkIDLogsSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworksNetworkIDLogsSearchParamsWithContext(ctx context.Context) *GetNetworksNetworkIDLogsSearchParams {
	var ()
	return &GetNetworksNetworkIDLogsSearchParams{

		Context: ctx,
	}
}

// NewGetNetworksNetworkIDLogsSearchParamsWithHTTPClient creates a new GetNetworksNetworkIDLogsSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworksNetworkIDLogsSearchParamsWithHTTPClient(client *http.Client) *GetNetworksNetworkIDLogsSearchParams {
	var ()
	return &GetNetworksNetworkIDLogsSearchParams{
		HTTPClient: client,
	}
}

/*GetNetworksNetworkIDLogsSearchParams contains all the parameters to send to the API endpoint
for the get networks network ID logs search operation typically these are written to a http.Request
*/
type GetNetworksNetworkIDLogsSearchParams struct {

	/*End
	  Time to end searching

	*/
	End *string
	/*Fields
	  Comma-separated list of fields to search with the simple query. Defaults to the log field.

	*/
	Fields *string
	/*Filters
	  Comma-separated list of key:value pairs to filter the query with.

	*/
	Filters *string
	/*From
	  The starting offset to fetch from the search hits. This param along with size can be used for pagination.

	*/
	From *string
	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*SimpleQuery
	  Simple Query String to execute

	*/
	SimpleQuery *string
	/*Size
	  Maximum number of hits returned. Defaults to 10.

	*/
	Size *string
	/*Start
	  Time to start searching

	*/
	Start *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) WithTimeout(timeout time.Duration) *GetNetworksNetworkIDLogsSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) WithContext(ctx context.Context) *GetNetworksNetworkIDLogsSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) WithHTTPClient(client *http.Client) *GetNetworksNetworkIDLogsSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnd adds the end to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) WithEnd(end *string) *GetNetworksNetworkIDLogsSearchParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) SetEnd(end *string) {
	o.End = end
}

// WithFields adds the fields to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) WithFields(fields *string) *GetNetworksNetworkIDLogsSearchParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilters adds the filters to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) WithFilters(filters *string) *GetNetworksNetworkIDLogsSearchParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WithFrom adds the from to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) WithFrom(from *string) *GetNetworksNetworkIDLogsSearchParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) SetFrom(from *string) {
	o.From = from
}

// WithNetworkID adds the networkID to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) WithNetworkID(networkID string) *GetNetworksNetworkIDLogsSearchParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithSimpleQuery adds the simpleQuery to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) WithSimpleQuery(simpleQuery *string) *GetNetworksNetworkIDLogsSearchParams {
	o.SetSimpleQuery(simpleQuery)
	return o
}

// SetSimpleQuery adds the simpleQuery to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) SetSimpleQuery(simpleQuery *string) {
	o.SimpleQuery = simpleQuery
}

// WithSize adds the size to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) WithSize(size *string) *GetNetworksNetworkIDLogsSearchParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) SetSize(size *string) {
	o.Size = size
}

// WithStart adds the start to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) WithStart(start *string) *GetNetworksNetworkIDLogsSearchParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get networks network ID logs search params
func (o *GetNetworksNetworkIDLogsSearchParams) SetStart(start *string) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworksNetworkIDLogsSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.End != nil {

		// query param end
		var qrEnd string
		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := qrEnd
		if qEnd != "" {
			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
		}

	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if o.Filters != nil {

		// query param filters
		var qrFilters string
		if o.Filters != nil {
			qrFilters = *o.Filters
		}
		qFilters := qrFilters
		if qFilters != "" {
			if err := r.SetQueryParam("filters", qFilters); err != nil {
				return err
			}
		}

	}

	if o.From != nil {

		// query param from
		var qrFrom string
		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom
		if qFrom != "" {
			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}

	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if o.SimpleQuery != nil {

		// query param simple_query
		var qrSimpleQuery string
		if o.SimpleQuery != nil {
			qrSimpleQuery = *o.SimpleQuery
		}
		qSimpleQuery := qrSimpleQuery
		if qSimpleQuery != "" {
			if err := r.SetQueryParam("simple_query", qSimpleQuery); err != nil {
				return err
			}
		}

	}

	if o.Size != nil {

		// query param size
		var qrSize string
		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := qrSize
		if qSize != "" {
			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}

	}

	if o.Start != nil {

		// query param start
		var qrStart string
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := qrStart
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
