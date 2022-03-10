// Code generated by go-swagger; DO NOT EDIT.

package gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new gateways API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for gateways API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteNetworksNetworkIDGatewaysGatewayID(params *DeleteNetworksNetworkIDGatewaysGatewayIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNetworksNetworkIDGatewaysGatewayIDNoContent, error)

	GetNetworksNetworkIDGateways(params *GetNetworksNetworkIDGatewaysParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDGatewaysOK, error)

	GetNetworksNetworkIDGatewaysGatewayID(params *GetNetworksNetworkIDGatewaysGatewayIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDGatewaysGatewayIDOK, error)

	GetNetworksNetworkIDGatewaysGatewayIDDescription(params *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDGatewaysGatewayIDDescriptionOK, error)

	GetNetworksNetworkIDGatewaysGatewayIDDevice(params *GetNetworksNetworkIDGatewaysGatewayIDDeviceParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDGatewaysGatewayIDDeviceOK, error)

	GetNetworksNetworkIDGatewaysGatewayIDMagmad(params *GetNetworksNetworkIDGatewaysGatewayIDMagmadParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDGatewaysGatewayIDMagmadOK, error)

	GetNetworksNetworkIDGatewaysGatewayIDName(params *GetNetworksNetworkIDGatewaysGatewayIDNameParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDGatewaysGatewayIDNameOK, error)

	GetNetworksNetworkIDGatewaysGatewayIDStatus(params *GetNetworksNetworkIDGatewaysGatewayIDStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDGatewaysGatewayIDStatusOK, error)

	GetNetworksNetworkIDGatewaysGatewayIDTier(params *GetNetworksNetworkIDGatewaysGatewayIDTierParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDGatewaysGatewayIDTierOK, error)

	PostNetworksNetworkIDGateways(params *PostNetworksNetworkIDGatewaysParams, authInfo runtime.ClientAuthInfoWriter) (*PostNetworksNetworkIDGatewaysCreated, error)

	PutNetworksNetworkIDGatewaysGatewayID(params *PutNetworksNetworkIDGatewaysGatewayIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutNetworksNetworkIDGatewaysGatewayIDNoContent, error)

	PutNetworksNetworkIDGatewaysGatewayIDDescription(params *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams, authInfo runtime.ClientAuthInfoWriter) (*PutNetworksNetworkIDGatewaysGatewayIDDescriptionNoContent, error)

	PutNetworksNetworkIDGatewaysGatewayIDDevice(params *PutNetworksNetworkIDGatewaysGatewayIDDeviceParams, authInfo runtime.ClientAuthInfoWriter) (*PutNetworksNetworkIDGatewaysGatewayIDDeviceNoContent, error)

	PutNetworksNetworkIDGatewaysGatewayIDMagmad(params *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams, authInfo runtime.ClientAuthInfoWriter) (*PutNetworksNetworkIDGatewaysGatewayIDMagmadNoContent, error)

	PutNetworksNetworkIDGatewaysGatewayIDName(params *PutNetworksNetworkIDGatewaysGatewayIDNameParams, authInfo runtime.ClientAuthInfoWriter) (*PutNetworksNetworkIDGatewaysGatewayIDNameNoContent, error)

	PutNetworksNetworkIDGatewaysGatewayIDTier(params *PutNetworksNetworkIDGatewaysGatewayIDTierParams, authInfo runtime.ClientAuthInfoWriter) (*PutNetworksNetworkIDGatewaysGatewayIDTierNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteNetworksNetworkIDGatewaysGatewayID deletes a gateway
*/
func (a *Client) DeleteNetworksNetworkIDGatewaysGatewayID(params *DeleteNetworksNetworkIDGatewaysGatewayIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNetworksNetworkIDGatewaysGatewayIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworksNetworkIDGatewaysGatewayIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteNetworksNetworkIDGatewaysGatewayID",
		Method:             "DELETE",
		PathPattern:        "/networks/{network_id}/gateways/{gateway_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNetworksNetworkIDGatewaysGatewayIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteNetworksNetworkIDGatewaysGatewayIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteNetworksNetworkIDGatewaysGatewayIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetNetworksNetworkIDGateways lists all gateways for a network
*/
func (a *Client) GetNetworksNetworkIDGateways(params *GetNetworksNetworkIDGatewaysParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDGatewaysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksNetworkIDGatewaysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNetworksNetworkIDGateways",
		Method:             "GET",
		PathPattern:        "/networks/{network_id}/gateways",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworksNetworkIDGatewaysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworksNetworkIDGatewaysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNetworksNetworkIDGatewaysDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetNetworksNetworkIDGatewaysGatewayID gets a specific gateway
*/
func (a *Client) GetNetworksNetworkIDGatewaysGatewayID(params *GetNetworksNetworkIDGatewaysGatewayIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDGatewaysGatewayIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksNetworkIDGatewaysGatewayIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNetworksNetworkIDGatewaysGatewayID",
		Method:             "GET",
		PathPattern:        "/networks/{network_id}/gateways/{gateway_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworksNetworkIDGatewaysGatewayIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworksNetworkIDGatewaysGatewayIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNetworksNetworkIDGatewaysGatewayIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetNetworksNetworkIDGatewaysGatewayIDDescription gets the description of a gateway
*/
func (a *Client) GetNetworksNetworkIDGatewaysGatewayIDDescription(params *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDGatewaysGatewayIDDescriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksNetworkIDGatewaysGatewayIDDescriptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNetworksNetworkIDGatewaysGatewayIDDescription",
		Method:             "GET",
		PathPattern:        "/networks/{network_id}/gateways/{gateway_id}/description",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworksNetworkIDGatewaysGatewayIDDescriptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworksNetworkIDGatewaysGatewayIDDescriptionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNetworksNetworkIDGatewaysGatewayIDDescriptionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetNetworksNetworkIDGatewaysGatewayIDDevice gets the physical device for a gateway
*/
func (a *Client) GetNetworksNetworkIDGatewaysGatewayIDDevice(params *GetNetworksNetworkIDGatewaysGatewayIDDeviceParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDGatewaysGatewayIDDeviceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksNetworkIDGatewaysGatewayIDDeviceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNetworksNetworkIDGatewaysGatewayIDDevice",
		Method:             "GET",
		PathPattern:        "/networks/{network_id}/gateways/{gateway_id}/device",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworksNetworkIDGatewaysGatewayIDDeviceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworksNetworkIDGatewaysGatewayIDDeviceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNetworksNetworkIDGatewaysGatewayIDDeviceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetNetworksNetworkIDGatewaysGatewayIDMagmad gets magmad agent configuration
*/
func (a *Client) GetNetworksNetworkIDGatewaysGatewayIDMagmad(params *GetNetworksNetworkIDGatewaysGatewayIDMagmadParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDGatewaysGatewayIDMagmadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksNetworkIDGatewaysGatewayIDMagmadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNetworksNetworkIDGatewaysGatewayIDMagmad",
		Method:             "GET",
		PathPattern:        "/networks/{network_id}/gateways/{gateway_id}/magmad",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworksNetworkIDGatewaysGatewayIDMagmadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworksNetworkIDGatewaysGatewayIDMagmadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNetworksNetworkIDGatewaysGatewayIDMagmadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetNetworksNetworkIDGatewaysGatewayIDName gets the name of a gateway
*/
func (a *Client) GetNetworksNetworkIDGatewaysGatewayIDName(params *GetNetworksNetworkIDGatewaysGatewayIDNameParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDGatewaysGatewayIDNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksNetworkIDGatewaysGatewayIDNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNetworksNetworkIDGatewaysGatewayIDName",
		Method:             "GET",
		PathPattern:        "/networks/{network_id}/gateways/{gateway_id}/name",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworksNetworkIDGatewaysGatewayIDNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworksNetworkIDGatewaysGatewayIDNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNetworksNetworkIDGatewaysGatewayIDNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetNetworksNetworkIDGatewaysGatewayIDStatus gets the status of a gateway
*/
func (a *Client) GetNetworksNetworkIDGatewaysGatewayIDStatus(params *GetNetworksNetworkIDGatewaysGatewayIDStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDGatewaysGatewayIDStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksNetworkIDGatewaysGatewayIDStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNetworksNetworkIDGatewaysGatewayIDStatus",
		Method:             "GET",
		PathPattern:        "/networks/{network_id}/gateways/{gateway_id}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworksNetworkIDGatewaysGatewayIDStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworksNetworkIDGatewaysGatewayIDStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNetworksNetworkIDGatewaysGatewayIDStatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetNetworksNetworkIDGatewaysGatewayIDTier gets the ID of the upgrade tier a gateway belongs to
*/
func (a *Client) GetNetworksNetworkIDGatewaysGatewayIDTier(params *GetNetworksNetworkIDGatewaysGatewayIDTierParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDGatewaysGatewayIDTierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksNetworkIDGatewaysGatewayIDTierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNetworksNetworkIDGatewaysGatewayIDTier",
		Method:             "GET",
		PathPattern:        "/networks/{network_id}/gateways/{gateway_id}/tier",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworksNetworkIDGatewaysGatewayIDTierReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworksNetworkIDGatewaysGatewayIDTierOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNetworksNetworkIDGatewaysGatewayIDTierDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostNetworksNetworkIDGateways registers a new gateway
*/
func (a *Client) PostNetworksNetworkIDGateways(params *PostNetworksNetworkIDGatewaysParams, authInfo runtime.ClientAuthInfoWriter) (*PostNetworksNetworkIDGatewaysCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNetworksNetworkIDGatewaysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostNetworksNetworkIDGateways",
		Method:             "POST",
		PathPattern:        "/networks/{network_id}/gateways",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostNetworksNetworkIDGatewaysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostNetworksNetworkIDGatewaysCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostNetworksNetworkIDGatewaysDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutNetworksNetworkIDGatewaysGatewayID updates an entire gateway record
*/
func (a *Client) PutNetworksNetworkIDGatewaysGatewayID(params *PutNetworksNetworkIDGatewaysGatewayIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutNetworksNetworkIDGatewaysGatewayIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutNetworksNetworkIDGatewaysGatewayIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutNetworksNetworkIDGatewaysGatewayID",
		Method:             "PUT",
		PathPattern:        "/networks/{network_id}/gateways/{gateway_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutNetworksNetworkIDGatewaysGatewayIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutNetworksNetworkIDGatewaysGatewayIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutNetworksNetworkIDGatewaysGatewayIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutNetworksNetworkIDGatewaysGatewayIDDescription updates the description of a gateway
*/
func (a *Client) PutNetworksNetworkIDGatewaysGatewayIDDescription(params *PutNetworksNetworkIDGatewaysGatewayIDDescriptionParams, authInfo runtime.ClientAuthInfoWriter) (*PutNetworksNetworkIDGatewaysGatewayIDDescriptionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutNetworksNetworkIDGatewaysGatewayIDDescriptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutNetworksNetworkIDGatewaysGatewayIDDescription",
		Method:             "PUT",
		PathPattern:        "/networks/{network_id}/gateways/{gateway_id}/description",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutNetworksNetworkIDGatewaysGatewayIDDescriptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutNetworksNetworkIDGatewaysGatewayIDDescriptionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutNetworksNetworkIDGatewaysGatewayIDDescriptionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutNetworksNetworkIDGatewaysGatewayIDDevice updates the physical device for a gateway
*/
func (a *Client) PutNetworksNetworkIDGatewaysGatewayIDDevice(params *PutNetworksNetworkIDGatewaysGatewayIDDeviceParams, authInfo runtime.ClientAuthInfoWriter) (*PutNetworksNetworkIDGatewaysGatewayIDDeviceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutNetworksNetworkIDGatewaysGatewayIDDeviceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutNetworksNetworkIDGatewaysGatewayIDDevice",
		Method:             "PUT",
		PathPattern:        "/networks/{network_id}/gateways/{gateway_id}/device",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutNetworksNetworkIDGatewaysGatewayIDDeviceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutNetworksNetworkIDGatewaysGatewayIDDeviceNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutNetworksNetworkIDGatewaysGatewayIDDeviceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutNetworksNetworkIDGatewaysGatewayIDMagmad reconfigures magmad agent
*/
func (a *Client) PutNetworksNetworkIDGatewaysGatewayIDMagmad(params *PutNetworksNetworkIDGatewaysGatewayIDMagmadParams, authInfo runtime.ClientAuthInfoWriter) (*PutNetworksNetworkIDGatewaysGatewayIDMagmadNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutNetworksNetworkIDGatewaysGatewayIDMagmadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutNetworksNetworkIDGatewaysGatewayIDMagmad",
		Method:             "PUT",
		PathPattern:        "/networks/{network_id}/gateways/{gateway_id}/magmad",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutNetworksNetworkIDGatewaysGatewayIDMagmadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutNetworksNetworkIDGatewaysGatewayIDMagmadNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutNetworksNetworkIDGatewaysGatewayIDMagmadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutNetworksNetworkIDGatewaysGatewayIDName updates the name of a gateway
*/
func (a *Client) PutNetworksNetworkIDGatewaysGatewayIDName(params *PutNetworksNetworkIDGatewaysGatewayIDNameParams, authInfo runtime.ClientAuthInfoWriter) (*PutNetworksNetworkIDGatewaysGatewayIDNameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutNetworksNetworkIDGatewaysGatewayIDNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutNetworksNetworkIDGatewaysGatewayIDName",
		Method:             "PUT",
		PathPattern:        "/networks/{network_id}/gateways/{gateway_id}/name",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutNetworksNetworkIDGatewaysGatewayIDNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutNetworksNetworkIDGatewaysGatewayIDNameNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutNetworksNetworkIDGatewaysGatewayIDNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutNetworksNetworkIDGatewaysGatewayIDTier updates the ID of the upgrade tier a gateway belongs to
*/
func (a *Client) PutNetworksNetworkIDGatewaysGatewayIDTier(params *PutNetworksNetworkIDGatewaysGatewayIDTierParams, authInfo runtime.ClientAuthInfoWriter) (*PutNetworksNetworkIDGatewaysGatewayIDTierNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutNetworksNetworkIDGatewaysGatewayIDTierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutNetworksNetworkIDGatewaysGatewayIDTier",
		Method:             "PUT",
		PathPattern:        "/networks/{network_id}/gateways/{gateway_id}/tier",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutNetworksNetworkIDGatewaysGatewayIDTierReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutNetworksNetworkIDGatewaysGatewayIDTierNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutNetworksNetworkIDGatewaysGatewayIDTierDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
