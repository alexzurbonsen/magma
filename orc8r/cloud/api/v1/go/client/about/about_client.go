// Code generated by go-swagger; DO NOT EDIT.

package about

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new about API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for about API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAboutVersion(params *GetAboutVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAboutVersionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAboutVersion gets version
*/
func (a *Client) GetAboutVersion(params *GetAboutVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAboutVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAboutVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAboutVersion",
		Method:             "GET",
		PathPattern:        "/about/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAboutVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAboutVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAboutVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
