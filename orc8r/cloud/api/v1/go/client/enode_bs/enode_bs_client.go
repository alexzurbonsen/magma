// Code generated by go-swagger; DO NOT EDIT.

package enode_bs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new enode bs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for enode bs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteLTENetworkIDEnodebsENODEBSerial(params *DeleteLTENetworkIDEnodebsENODEBSerialParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLTENetworkIDEnodebsENODEBSerialNoContent, error)

	GetLTENetworkIDEnodebs(params *GetLTENetworkIDEnodebsParams, authInfo runtime.ClientAuthInfoWriter) (*GetLTENetworkIDEnodebsOK, error)

	GetLTENetworkIDEnodebsENODEBSerial(params *GetLTENetworkIDEnodebsENODEBSerialParams, authInfo runtime.ClientAuthInfoWriter) (*GetLTENetworkIDEnodebsENODEBSerialOK, error)

	GetLTENetworkIDEnodebsENODEBSerialState(params *GetLTENetworkIDEnodebsENODEBSerialStateParams, authInfo runtime.ClientAuthInfoWriter) (*GetLTENetworkIDEnodebsENODEBSerialStateOK, error)

	PostLTENetworkIDEnodebs(params *PostLTENetworkIDEnodebsParams, authInfo runtime.ClientAuthInfoWriter) (*PostLTENetworkIDEnodebsCreated, error)

	PutLTENetworkIDEnodebsENODEBSerial(params *PutLTENetworkIDEnodebsENODEBSerialParams, authInfo runtime.ClientAuthInfoWriter) (*PutLTENetworkIDEnodebsENODEBSerialNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteLTENetworkIDEnodebsENODEBSerial unregisters an enode b
*/
func (a *Client) DeleteLTENetworkIDEnodebsENODEBSerial(params *DeleteLTENetworkIDEnodebsENODEBSerialParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLTENetworkIDEnodebsENODEBSerialNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLTENetworkIDEnodebsENODEBSerialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteLTENetworkIDEnodebsENODEBSerial",
		Method:             "DELETE",
		PathPattern:        "/lte/{network_id}/enodebs/{enodeb_serial}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteLTENetworkIDEnodebsENODEBSerialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteLTENetworkIDEnodebsENODEBSerialNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteLTENetworkIDEnodebsENODEBSerialDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetLTENetworkIDEnodebs lists all enode bs in the network
*/
func (a *Client) GetLTENetworkIDEnodebs(params *GetLTENetworkIDEnodebsParams, authInfo runtime.ClientAuthInfoWriter) (*GetLTENetworkIDEnodebsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLTENetworkIDEnodebsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLTENetworkIDEnodebs",
		Method:             "GET",
		PathPattern:        "/lte/{network_id}/enodebs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLTENetworkIDEnodebsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLTENetworkIDEnodebsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLTENetworkIDEnodebsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetLTENetworkIDEnodebsENODEBSerial retrieves a specific enode b configuration
*/
func (a *Client) GetLTENetworkIDEnodebsENODEBSerial(params *GetLTENetworkIDEnodebsENODEBSerialParams, authInfo runtime.ClientAuthInfoWriter) (*GetLTENetworkIDEnodebsENODEBSerialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLTENetworkIDEnodebsENODEBSerialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLTENetworkIDEnodebsENODEBSerial",
		Method:             "GET",
		PathPattern:        "/lte/{network_id}/enodebs/{enodeb_serial}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLTENetworkIDEnodebsENODEBSerialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLTENetworkIDEnodebsENODEBSerialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLTENetworkIDEnodebsENODEBSerialDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetLTENetworkIDEnodebsENODEBSerialState retrieves reported state from enodeb device
*/
func (a *Client) GetLTENetworkIDEnodebsENODEBSerialState(params *GetLTENetworkIDEnodebsENODEBSerialStateParams, authInfo runtime.ClientAuthInfoWriter) (*GetLTENetworkIDEnodebsENODEBSerialStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLTENetworkIDEnodebsENODEBSerialStateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLTENetworkIDEnodebsENODEBSerialState",
		Method:             "GET",
		PathPattern:        "/lte/{network_id}/enodebs/{enodeb_serial}/state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLTENetworkIDEnodebsENODEBSerialStateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLTENetworkIDEnodebsENODEBSerialStateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLTENetworkIDEnodebsENODEBSerialStateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostLTENetworkIDEnodebs registers a new enode b
*/
func (a *Client) PostLTENetworkIDEnodebs(params *PostLTENetworkIDEnodebsParams, authInfo runtime.ClientAuthInfoWriter) (*PostLTENetworkIDEnodebsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLTENetworkIDEnodebsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLTENetworkIDEnodebs",
		Method:             "POST",
		PathPattern:        "/lte/{network_id}/enodebs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostLTENetworkIDEnodebsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostLTENetworkIDEnodebsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostLTENetworkIDEnodebsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutLTENetworkIDEnodebsENODEBSerial updates an enode b s configuration
*/
func (a *Client) PutLTENetworkIDEnodebsENODEBSerial(params *PutLTENetworkIDEnodebsENODEBSerialParams, authInfo runtime.ClientAuthInfoWriter) (*PutLTENetworkIDEnodebsENODEBSerialNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutLTENetworkIDEnodebsENODEBSerialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutLTENetworkIDEnodebsENODEBSerial",
		Method:             "PUT",
		PathPattern:        "/lte/{network_id}/enodebs/{enodeb_serial}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutLTENetworkIDEnodebsENODEBSerialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutLTENetworkIDEnodebsENODEBSerialNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutLTENetworkIDEnodebsENODEBSerialDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
