// Code generated by go-swagger; DO NOT EDIT.

package network_probes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new network probes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for network probes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteLTENetworkIDNetworkProbeDestinationsDestinationID(params *DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDNoContent, error)

	DeleteLTENetworkIDNetworkProbeTasksTaskID(params *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLTENetworkIDNetworkProbeTasksTaskIDNoContent, error)

	GetLTENetworkIDNetworkProbeDestinations(params *GetLTENetworkIDNetworkProbeDestinationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetLTENetworkIDNetworkProbeDestinationsOK, error)

	GetLTENetworkIDNetworkProbeDestinationsDestinationID(params *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLTENetworkIDNetworkProbeDestinationsDestinationIDOK, error)

	GetLTENetworkIDNetworkProbeTasks(params *GetLTENetworkIDNetworkProbeTasksParams, authInfo runtime.ClientAuthInfoWriter) (*GetLTENetworkIDNetworkProbeTasksOK, error)

	GetLTENetworkIDNetworkProbeTasksTaskID(params *GetLTENetworkIDNetworkProbeTasksTaskIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLTENetworkIDNetworkProbeTasksTaskIDOK, error)

	PostLTENetworkIDNetworkProbeDestinations(params *PostLTENetworkIDNetworkProbeDestinationsParams, authInfo runtime.ClientAuthInfoWriter) (*PostLTENetworkIDNetworkProbeDestinationsCreated, error)

	PostLTENetworkIDNetworkProbeTasks(params *PostLTENetworkIDNetworkProbeTasksParams, authInfo runtime.ClientAuthInfoWriter) (*PostLTENetworkIDNetworkProbeTasksCreated, error)

	PutLTENetworkIDNetworkProbeDestinationsDestinationID(params *PutLTENetworkIDNetworkProbeDestinationsDestinationIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutLTENetworkIDNetworkProbeDestinationsDestinationIDNoContent, error)

	PutLTENetworkIDNetworkProbeTasksTaskID(params *PutLTENetworkIDNetworkProbeTasksTaskIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutLTENetworkIDNetworkProbeTasksTaskIDNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteLTENetworkIDNetworkProbeDestinationsDestinationID removes a network probe destination from the network
*/
func (a *Client) DeleteLTENetworkIDNetworkProbeDestinationsDestinationID(params *DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLTENetworkIDNetworkProbeDestinationsDestinationIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteLTENetworkIDNetworkProbeDestinationsDestinationID",
		Method:             "DELETE",
		PathPattern:        "/lte/{network_id}/network_probe/destinations/{destination_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteLTENetworkIDNetworkProbeDestinationsDestinationIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteLTENetworkIDNetworkProbeTasksTaskID removes an network probe task from the network
*/
func (a *Client) DeleteLTENetworkIDNetworkProbeTasksTaskID(params *DeleteLTENetworkIDNetworkProbeTasksTaskIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLTENetworkIDNetworkProbeTasksTaskIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLTENetworkIDNetworkProbeTasksTaskIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteLTENetworkIDNetworkProbeTasksTaskID",
		Method:             "DELETE",
		PathPattern:        "/lte/{network_id}/network_probe/tasks/{task_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteLTENetworkIDNetworkProbeTasksTaskIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteLTENetworkIDNetworkProbeTasksTaskIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteLTENetworkIDNetworkProbeTasksTaskIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetLTENetworkIDNetworkProbeDestinations lists network probe destinations in the network
*/
func (a *Client) GetLTENetworkIDNetworkProbeDestinations(params *GetLTENetworkIDNetworkProbeDestinationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetLTENetworkIDNetworkProbeDestinationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLTENetworkIDNetworkProbeDestinationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLTENetworkIDNetworkProbeDestinations",
		Method:             "GET",
		PathPattern:        "/lte/{network_id}/network_probe/destinations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLTENetworkIDNetworkProbeDestinationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLTENetworkIDNetworkProbeDestinationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLTENetworkIDNetworkProbeDestinationsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetLTENetworkIDNetworkProbeDestinationsDestinationID retrieves a network probe destination
*/
func (a *Client) GetLTENetworkIDNetworkProbeDestinationsDestinationID(params *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLTENetworkIDNetworkProbeDestinationsDestinationIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLTENetworkIDNetworkProbeDestinationsDestinationIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLTENetworkIDNetworkProbeDestinationsDestinationID",
		Method:             "GET",
		PathPattern:        "/lte/{network_id}/network_probe/destinations/{destination_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLTENetworkIDNetworkProbeDestinationsDestinationIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLTENetworkIDNetworkProbeDestinationsDestinationIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLTENetworkIDNetworkProbeDestinationsDestinationIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetLTENetworkIDNetworkProbeTasks lists network probe task in the network
*/
func (a *Client) GetLTENetworkIDNetworkProbeTasks(params *GetLTENetworkIDNetworkProbeTasksParams, authInfo runtime.ClientAuthInfoWriter) (*GetLTENetworkIDNetworkProbeTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLTENetworkIDNetworkProbeTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLTENetworkIDNetworkProbeTasks",
		Method:             "GET",
		PathPattern:        "/lte/{network_id}/network_probe/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLTENetworkIDNetworkProbeTasksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLTENetworkIDNetworkProbeTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLTENetworkIDNetworkProbeTasksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetLTENetworkIDNetworkProbeTasksTaskID retrieves the network probe task info
*/
func (a *Client) GetLTENetworkIDNetworkProbeTasksTaskID(params *GetLTENetworkIDNetworkProbeTasksTaskIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLTENetworkIDNetworkProbeTasksTaskIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLTENetworkIDNetworkProbeTasksTaskIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLTENetworkIDNetworkProbeTasksTaskID",
		Method:             "GET",
		PathPattern:        "/lte/{network_id}/network_probe/tasks/{task_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLTENetworkIDNetworkProbeTasksTaskIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLTENetworkIDNetworkProbeTasksTaskIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLTENetworkIDNetworkProbeTasksTaskIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostLTENetworkIDNetworkProbeDestinations adds a new network probe destination to the network
*/
func (a *Client) PostLTENetworkIDNetworkProbeDestinations(params *PostLTENetworkIDNetworkProbeDestinationsParams, authInfo runtime.ClientAuthInfoWriter) (*PostLTENetworkIDNetworkProbeDestinationsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLTENetworkIDNetworkProbeDestinationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLTENetworkIDNetworkProbeDestinations",
		Method:             "POST",
		PathPattern:        "/lte/{network_id}/network_probe/destinations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostLTENetworkIDNetworkProbeDestinationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostLTENetworkIDNetworkProbeDestinationsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostLTENetworkIDNetworkProbeDestinationsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostLTENetworkIDNetworkProbeTasks adds a new network probe task to the network
*/
func (a *Client) PostLTENetworkIDNetworkProbeTasks(params *PostLTENetworkIDNetworkProbeTasksParams, authInfo runtime.ClientAuthInfoWriter) (*PostLTENetworkIDNetworkProbeTasksCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLTENetworkIDNetworkProbeTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLTENetworkIDNetworkProbeTasks",
		Method:             "POST",
		PathPattern:        "/lte/{network_id}/network_probe/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostLTENetworkIDNetworkProbeTasksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostLTENetworkIDNetworkProbeTasksCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostLTENetworkIDNetworkProbeTasksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutLTENetworkIDNetworkProbeDestinationsDestinationID updates an existing network probe destination in the network
*/
func (a *Client) PutLTENetworkIDNetworkProbeDestinationsDestinationID(params *PutLTENetworkIDNetworkProbeDestinationsDestinationIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutLTENetworkIDNetworkProbeDestinationsDestinationIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutLTENetworkIDNetworkProbeDestinationsDestinationIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutLTENetworkIDNetworkProbeDestinationsDestinationID",
		Method:             "PUT",
		PathPattern:        "/lte/{network_id}/network_probe/destinations/{destination_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutLTENetworkIDNetworkProbeDestinationsDestinationIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutLTENetworkIDNetworkProbeDestinationsDestinationIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutLTENetworkIDNetworkProbeDestinationsDestinationIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutLTENetworkIDNetworkProbeTasksTaskID updates an existing network probe task in the network
*/
func (a *Client) PutLTENetworkIDNetworkProbeTasksTaskID(params *PutLTENetworkIDNetworkProbeTasksTaskIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutLTENetworkIDNetworkProbeTasksTaskIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutLTENetworkIDNetworkProbeTasksTaskIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutLTENetworkIDNetworkProbeTasksTaskID",
		Method:             "PUT",
		PathPattern:        "/lte/{network_id}/network_probe/tasks/{task_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutLTENetworkIDNetworkProbeTasksTaskIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutLTENetworkIDNetworkProbeTasksTaskIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutLTENetworkIDNetworkProbeTasksTaskIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
