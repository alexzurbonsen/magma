// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new alerts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for alerts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteNetworksNetworkIDAlertsSilence(params *DeleteNetworksNetworkIDAlertsSilenceParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNetworksNetworkIDAlertsSilenceOK, error)

	DeleteNetworksNetworkIDPrometheusAlertConfig(params *DeleteNetworksNetworkIDPrometheusAlertConfigParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNetworksNetworkIDPrometheusAlertConfigOK, error)

	DeleteNetworksNetworkIDPrometheusAlertReceiver(params *DeleteNetworksNetworkIDPrometheusAlertReceiverParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNetworksNetworkIDPrometheusAlertReceiverOK, error)

	GetNetworksNetworkIDAlerts(params *GetNetworksNetworkIDAlertsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDAlertsOK, error)

	GetNetworksNetworkIDAlertsSilence(params *GetNetworksNetworkIDAlertsSilenceParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDAlertsSilenceOK, error)

	GetNetworksNetworkIDPrometheusAlertConfig(params *GetNetworksNetworkIDPrometheusAlertConfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDPrometheusAlertConfigOK, error)

	GetNetworksNetworkIDPrometheusAlertReceiver(params *GetNetworksNetworkIDPrometheusAlertReceiverParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDPrometheusAlertReceiverOK, error)

	GetNetworksNetworkIDPrometheusAlertReceiverRoute(params *GetNetworksNetworkIDPrometheusAlertReceiverRouteParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDPrometheusAlertReceiverRouteOK, error)

	PostNetworksNetworkIDAlertsSilence(params *PostNetworksNetworkIDAlertsSilenceParams, authInfo runtime.ClientAuthInfoWriter) (*PostNetworksNetworkIDAlertsSilenceOK, error)

	PostNetworksNetworkIDPrometheusAlertConfig(params *PostNetworksNetworkIDPrometheusAlertConfigParams, authInfo runtime.ClientAuthInfoWriter) (*PostNetworksNetworkIDPrometheusAlertConfigCreated, error)

	PostNetworksNetworkIDPrometheusAlertReceiver(params *PostNetworksNetworkIDPrometheusAlertReceiverParams, authInfo runtime.ClientAuthInfoWriter) (*PostNetworksNetworkIDPrometheusAlertReceiverCreated, error)

	PostNetworksNetworkIDPrometheusAlertReceiverRoute(params *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams, authInfo runtime.ClientAuthInfoWriter) (*PostNetworksNetworkIDPrometheusAlertReceiverRouteOK, error)

	PutNetworksNetworkIDPrometheusAlertConfigAlertName(params *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams, authInfo runtime.ClientAuthInfoWriter) (*PutNetworksNetworkIDPrometheusAlertConfigAlertNameOK, error)

	PutNetworksNetworkIDPrometheusAlertConfigBulk(params *PutNetworksNetworkIDPrometheusAlertConfigBulkParams, authInfo runtime.ClientAuthInfoWriter) (*PutNetworksNetworkIDPrometheusAlertConfigBulkOK, error)

	PutNetworksNetworkIDPrometheusAlertReceiverReceiver(params *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams, authInfo runtime.ClientAuthInfoWriter) (*PutNetworksNetworkIDPrometheusAlertReceiverReceiverOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteNetworksNetworkIDAlertsSilence deletes an alert silencer
*/
func (a *Client) DeleteNetworksNetworkIDAlertsSilence(params *DeleteNetworksNetworkIDAlertsSilenceParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNetworksNetworkIDAlertsSilenceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworksNetworkIDAlertsSilenceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteNetworksNetworkIDAlertsSilence",
		Method:             "DELETE",
		PathPattern:        "/networks/{network_id}/alerts/silence",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNetworksNetworkIDAlertsSilenceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteNetworksNetworkIDAlertsSilenceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteNetworksNetworkIDAlertsSilenceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteNetworksNetworkIDPrometheusAlertConfig deletes an alerting rule
*/
func (a *Client) DeleteNetworksNetworkIDPrometheusAlertConfig(params *DeleteNetworksNetworkIDPrometheusAlertConfigParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNetworksNetworkIDPrometheusAlertConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworksNetworkIDPrometheusAlertConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteNetworksNetworkIDPrometheusAlertConfig",
		Method:             "DELETE",
		PathPattern:        "/networks/{network_id}/prometheus/alert_config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNetworksNetworkIDPrometheusAlertConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteNetworksNetworkIDPrometheusAlertConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteNetworksNetworkIDPrometheusAlertConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteNetworksNetworkIDPrometheusAlertReceiver deletes alert receiver
*/
func (a *Client) DeleteNetworksNetworkIDPrometheusAlertReceiver(params *DeleteNetworksNetworkIDPrometheusAlertReceiverParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNetworksNetworkIDPrometheusAlertReceiverOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworksNetworkIDPrometheusAlertReceiverParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteNetworksNetworkIDPrometheusAlertReceiver",
		Method:             "DELETE",
		PathPattern:        "/networks/{network_id}/prometheus/alert_receiver",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNetworksNetworkIDPrometheusAlertReceiverReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteNetworksNetworkIDPrometheusAlertReceiverOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteNetworksNetworkIDPrometheusAlertReceiverDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetNetworksNetworkIDAlerts views currently firing alerts
*/
func (a *Client) GetNetworksNetworkIDAlerts(params *GetNetworksNetworkIDAlertsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDAlertsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksNetworkIDAlertsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNetworksNetworkIDAlerts",
		Method:             "GET",
		PathPattern:        "/networks/{network_id}/alerts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworksNetworkIDAlertsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworksNetworkIDAlertsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNetworksNetworkIDAlertsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetNetworksNetworkIDAlertsSilence views active alert silencers
*/
func (a *Client) GetNetworksNetworkIDAlertsSilence(params *GetNetworksNetworkIDAlertsSilenceParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDAlertsSilenceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksNetworkIDAlertsSilenceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNetworksNetworkIDAlertsSilence",
		Method:             "GET",
		PathPattern:        "/networks/{network_id}/alerts/silence",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworksNetworkIDAlertsSilenceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworksNetworkIDAlertsSilenceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNetworksNetworkIDAlertsSilenceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetNetworksNetworkIDPrometheusAlertConfig retrives alerting rule configurations

  If no query parameters are included, all alerting rules for the given network are returned.
*/
func (a *Client) GetNetworksNetworkIDPrometheusAlertConfig(params *GetNetworksNetworkIDPrometheusAlertConfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDPrometheusAlertConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksNetworkIDPrometheusAlertConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNetworksNetworkIDPrometheusAlertConfig",
		Method:             "GET",
		PathPattern:        "/networks/{network_id}/prometheus/alert_config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworksNetworkIDPrometheusAlertConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworksNetworkIDPrometheusAlertConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNetworksNetworkIDPrometheusAlertConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetNetworksNetworkIDPrometheusAlertReceiver retrives alert receivers
*/
func (a *Client) GetNetworksNetworkIDPrometheusAlertReceiver(params *GetNetworksNetworkIDPrometheusAlertReceiverParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDPrometheusAlertReceiverOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksNetworkIDPrometheusAlertReceiverParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNetworksNetworkIDPrometheusAlertReceiver",
		Method:             "GET",
		PathPattern:        "/networks/{network_id}/prometheus/alert_receiver",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworksNetworkIDPrometheusAlertReceiverReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworksNetworkIDPrometheusAlertReceiverOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNetworksNetworkIDPrometheusAlertReceiverDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetNetworksNetworkIDPrometheusAlertReceiverRoute retrieves alert routing tree
*/
func (a *Client) GetNetworksNetworkIDPrometheusAlertReceiverRoute(params *GetNetworksNetworkIDPrometheusAlertReceiverRouteParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworksNetworkIDPrometheusAlertReceiverRouteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksNetworkIDPrometheusAlertReceiverRouteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNetworksNetworkIDPrometheusAlertReceiverRoute",
		Method:             "GET",
		PathPattern:        "/networks/{network_id}/prometheus/alert_receiver/route",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNetworksNetworkIDPrometheusAlertReceiverRouteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworksNetworkIDPrometheusAlertReceiverRouteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetNetworksNetworkIDPrometheusAlertReceiverRoute: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostNetworksNetworkIDAlertsSilence creates a new alert silencer
*/
func (a *Client) PostNetworksNetworkIDAlertsSilence(params *PostNetworksNetworkIDAlertsSilenceParams, authInfo runtime.ClientAuthInfoWriter) (*PostNetworksNetworkIDAlertsSilenceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNetworksNetworkIDAlertsSilenceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostNetworksNetworkIDAlertsSilence",
		Method:             "POST",
		PathPattern:        "/networks/{network_id}/alerts/silence",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostNetworksNetworkIDAlertsSilenceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostNetworksNetworkIDAlertsSilenceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostNetworksNetworkIDAlertsSilenceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostNetworksNetworkIDPrometheusAlertConfig creates new alerting rule
*/
func (a *Client) PostNetworksNetworkIDPrometheusAlertConfig(params *PostNetworksNetworkIDPrometheusAlertConfigParams, authInfo runtime.ClientAuthInfoWriter) (*PostNetworksNetworkIDPrometheusAlertConfigCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNetworksNetworkIDPrometheusAlertConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostNetworksNetworkIDPrometheusAlertConfig",
		Method:             "POST",
		PathPattern:        "/networks/{network_id}/prometheus/alert_config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostNetworksNetworkIDPrometheusAlertConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostNetworksNetworkIDPrometheusAlertConfigCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostNetworksNetworkIDPrometheusAlertConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostNetworksNetworkIDPrometheusAlertReceiver creates new alert receiver
*/
func (a *Client) PostNetworksNetworkIDPrometheusAlertReceiver(params *PostNetworksNetworkIDPrometheusAlertReceiverParams, authInfo runtime.ClientAuthInfoWriter) (*PostNetworksNetworkIDPrometheusAlertReceiverCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNetworksNetworkIDPrometheusAlertReceiverParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostNetworksNetworkIDPrometheusAlertReceiver",
		Method:             "POST",
		PathPattern:        "/networks/{network_id}/prometheus/alert_receiver",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostNetworksNetworkIDPrometheusAlertReceiverReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostNetworksNetworkIDPrometheusAlertReceiverCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostNetworksNetworkIDPrometheusAlertReceiverDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostNetworksNetworkIDPrometheusAlertReceiverRoute modifies alert routing tree
*/
func (a *Client) PostNetworksNetworkIDPrometheusAlertReceiverRoute(params *PostNetworksNetworkIDPrometheusAlertReceiverRouteParams, authInfo runtime.ClientAuthInfoWriter) (*PostNetworksNetworkIDPrometheusAlertReceiverRouteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNetworksNetworkIDPrometheusAlertReceiverRouteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostNetworksNetworkIDPrometheusAlertReceiverRoute",
		Method:             "POST",
		PathPattern:        "/networks/{network_id}/prometheus/alert_receiver/route",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostNetworksNetworkIDPrometheusAlertReceiverRouteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostNetworksNetworkIDPrometheusAlertReceiverRouteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostNetworksNetworkIDPrometheusAlertReceiverRouteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutNetworksNetworkIDPrometheusAlertConfigAlertName updates an alerting rule
*/
func (a *Client) PutNetworksNetworkIDPrometheusAlertConfigAlertName(params *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams, authInfo runtime.ClientAuthInfoWriter) (*PutNetworksNetworkIDPrometheusAlertConfigAlertNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutNetworksNetworkIDPrometheusAlertConfigAlertNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutNetworksNetworkIDPrometheusAlertConfigAlertName",
		Method:             "PUT",
		PathPattern:        "/networks/{network_id}/prometheus/alert_config/{alert_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutNetworksNetworkIDPrometheusAlertConfigAlertNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutNetworksNetworkIDPrometheusAlertConfigAlertNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutNetworksNetworkIDPrometheusAlertConfigAlertNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutNetworksNetworkIDPrometheusAlertConfigBulk bulks update create alerting rules
*/
func (a *Client) PutNetworksNetworkIDPrometheusAlertConfigBulk(params *PutNetworksNetworkIDPrometheusAlertConfigBulkParams, authInfo runtime.ClientAuthInfoWriter) (*PutNetworksNetworkIDPrometheusAlertConfigBulkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutNetworksNetworkIDPrometheusAlertConfigBulkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutNetworksNetworkIDPrometheusAlertConfigBulk",
		Method:             "PUT",
		PathPattern:        "/networks/{network_id}/prometheus/alert_config/bulk",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutNetworksNetworkIDPrometheusAlertConfigBulkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutNetworksNetworkIDPrometheusAlertConfigBulkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutNetworksNetworkIDPrometheusAlertConfigBulkDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutNetworksNetworkIDPrometheusAlertReceiverReceiver updates existing alert receiver
*/
func (a *Client) PutNetworksNetworkIDPrometheusAlertReceiverReceiver(params *PutNetworksNetworkIDPrometheusAlertReceiverReceiverParams, authInfo runtime.ClientAuthInfoWriter) (*PutNetworksNetworkIDPrometheusAlertReceiverReceiverOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutNetworksNetworkIDPrometheusAlertReceiverReceiverParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutNetworksNetworkIDPrometheusAlertReceiverReceiver",
		Method:             "PUT",
		PathPattern:        "/networks/{network_id}/prometheus/alert_receiver/{receiver}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutNetworksNetworkIDPrometheusAlertReceiverReceiverReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutNetworksNetworkIDPrometheusAlertReceiverReceiverOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutNetworksNetworkIDPrometheusAlertReceiverReceiverDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
