// Code generated by go-swagger; DO NOT EDIT.

package call_tracing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new call tracing API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for call tracing API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteNetworksNetworkIDTracingTraceID(params *DeleteNetworksNetworkIDTracingTraceIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNetworksNetworkIDTracingTraceIDNoContent, error)

	GetNetworksNetworkIDTracing(params *GetNetworksNetworkIDTracingParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDTracingOK, error)

	GetNetworksNetworkIDTracingTraceID(params *GetNetworksNetworkIDTracingTraceIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDTracingTraceIDOK, error)

	GetNetworksNetworkIDTracingTraceIDDownload(params *GetNetworksNetworkIDTracingTraceIDDownloadParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer) (*GetNetworksNetworkIDTracingTraceIDDownloadOK, error)

	PostNetworksNetworkIDTracing(params *PostNetworksNetworkIDTracingParams, authInfo runtime.ClientAuthInfoWriter) (*PostNetworksNetworkIDTracingCreated, error)

	PutNetworksNetworkIDTracingTraceID(params *PutNetworksNetworkIDTracingTraceIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutNetworksNetworkIDTracingTraceIDNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteNetworksNetworkIDTracingTraceID deletes a call trace record
*/
func (a *Client) DeleteNetworksNetworkIDTracingTraceID(params *DeleteNetworksNetworkIDTracingTraceIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNetworksNetworkIDTracingTraceIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworksNetworkIDTracingTraceIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteNetworksNetworkIDTracingTraceID",
		Method:             "DELETE",
		PathPattern:        "/networks/{network_id}/tracing/{trace_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNetworksNetworkIDTracingTraceIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteNetworksNetworkIDTracingTraceIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteNetworksNetworkIDTracingTraceIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetNetworksNetworkIDTracing lists all call traces for a network
*/
func (a *Client) GetNetworksNetworkIDTracing(params *GetNetworksNetworkIDTracingParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDTracingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksNetworkIDTracingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNetworksNetworkIDTracing",
		Method:             "GET",
		PathPattern:        "/networks/{network_id}/tracing",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworksNetworkIDTracingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworksNetworkIDTracingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNetworksNetworkIDTracingDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetNetworksNetworkIDTracingTraceID gets tracing status
*/
func (a *Client) GetNetworksNetworkIDTracingTraceID(params *GetNetworksNetworkIDTracingTraceIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDTracingTraceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksNetworkIDTracingTraceIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNetworksNetworkIDTracingTraceID",
		Method:             "GET",
		PathPattern:        "/networks/{network_id}/tracing/{trace_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworksNetworkIDTracingTraceIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworksNetworkIDTracingTraceIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNetworksNetworkIDTracingTraceIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetNetworksNetworkIDTracingTraceIDDownload gets the call trace in p c a p format
*/
func (a *Client) GetNetworksNetworkIDTracingTraceIDDownload(params *GetNetworksNetworkIDTracingTraceIDDownloadParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer) (*GetNetworksNetworkIDTracingTraceIDDownloadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksNetworkIDTracingTraceIDDownloadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNetworksNetworkIDTracingTraceIDDownload",
		Method:             "GET",
		PathPattern:        "/networks/{network_id}/tracing/{trace_id}/download",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworksNetworkIDTracingTraceIDDownloadReader{formats: a.formats, writer: writer},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworksNetworkIDTracingTraceIDDownloadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNetworksNetworkIDTracingTraceIDDownloadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostNetworksNetworkIDTracing starts a new call trace
*/
func (a *Client) PostNetworksNetworkIDTracing(params *PostNetworksNetworkIDTracingParams, authInfo runtime.ClientAuthInfoWriter) (*PostNetworksNetworkIDTracingCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNetworksNetworkIDTracingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostNetworksNetworkIDTracing",
		Method:             "POST",
		PathPattern:        "/networks/{network_id}/tracing",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostNetworksNetworkIDTracingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostNetworksNetworkIDTracingCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostNetworksNetworkIDTracingDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutNetworksNetworkIDTracingTraceID updates a call trace
*/
func (a *Client) PutNetworksNetworkIDTracingTraceID(params *PutNetworksNetworkIDTracingTraceIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutNetworksNetworkIDTracingTraceIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutNetworksNetworkIDTracingTraceIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutNetworksNetworkIDTracingTraceID",
		Method:             "PUT",
		PathPattern:        "/networks/{network_id}/tracing/{trace_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutNetworksNetworkIDTracingTraceIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutNetworksNetworkIDTracingTraceIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutNetworksNetworkIDTracingTraceIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
