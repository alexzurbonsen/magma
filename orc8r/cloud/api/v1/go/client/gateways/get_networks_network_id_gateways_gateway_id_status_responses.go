// Code generated by go-swagger; DO NOT EDIT.

package gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetNetworksNetworkIDGatewaysGatewayIDStatusReader is a Reader for the GetNetworksNetworkIDGatewaysGatewayIDStatus structure.
type GetNetworksNetworkIDGatewaysGatewayIDStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDGatewaysGatewayIDStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDGatewaysGatewayIDStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDGatewaysGatewayIDStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDGatewaysGatewayIDStatusOK creates a GetNetworksNetworkIDGatewaysGatewayIDStatusOK with default headers values
func NewGetNetworksNetworkIDGatewaysGatewayIDStatusOK() *GetNetworksNetworkIDGatewaysGatewayIDStatusOK {
	return &GetNetworksNetworkIDGatewaysGatewayIDStatusOK{}
}

/* GetNetworksNetworkIDGatewaysGatewayIDStatusOK describes a response with status code 200, with default header values.

The status of the gateway
*/
type GetNetworksNetworkIDGatewaysGatewayIDStatusOK struct {
	Payload *models.GatewayStatus
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDStatusOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/gateways/{gateway_id}/status][%d] getNetworksNetworkIdGatewaysGatewayIdStatusOK  %+v", 200, o.Payload)
}
func (o *GetNetworksNetworkIDGatewaysGatewayIDStatusOK) GetPayload() *models.GatewayStatus {
	return o.Payload
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDGatewaysGatewayIDStatusDefault creates a GetNetworksNetworkIDGatewaysGatewayIDStatusDefault with default headers values
func NewGetNetworksNetworkIDGatewaysGatewayIDStatusDefault(code int) *GetNetworksNetworkIDGatewaysGatewayIDStatusDefault {
	return &GetNetworksNetworkIDGatewaysGatewayIDStatusDefault{
		_statusCode: code,
	}
}

/* GetNetworksNetworkIDGatewaysGatewayIDStatusDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDGatewaysGatewayIDStatusDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID gateways gateway ID status default response
func (o *GetNetworksNetworkIDGatewaysGatewayIDStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDStatusDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/gateways/{gateway_id}/status][%d] GetNetworksNetworkIDGatewaysGatewayIDStatus default  %+v", o._statusCode, o.Payload)
}
func (o *GetNetworksNetworkIDGatewaysGatewayIDStatusDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
