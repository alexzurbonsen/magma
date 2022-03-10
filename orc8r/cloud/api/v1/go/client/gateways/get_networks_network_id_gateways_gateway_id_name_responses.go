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

// GetNetworksNetworkIDGatewaysGatewayIDNameReader is a Reader for the GetNetworksNetworkIDGatewaysGatewayIDName structure.
type GetNetworksNetworkIDGatewaysGatewayIDNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDGatewaysGatewayIDNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDGatewaysGatewayIDNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDGatewaysGatewayIDNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDGatewaysGatewayIDNameOK creates a GetNetworksNetworkIDGatewaysGatewayIDNameOK with default headers values
func NewGetNetworksNetworkIDGatewaysGatewayIDNameOK() *GetNetworksNetworkIDGatewaysGatewayIDNameOK {
	return &GetNetworksNetworkIDGatewaysGatewayIDNameOK{}
}

/*GetNetworksNetworkIDGatewaysGatewayIDNameOK handles this case with default header values.

The name of the gateway
*/
type GetNetworksNetworkIDGatewaysGatewayIDNameOK struct {
	Payload models.GatewayName
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDNameOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/gateways/{gateway_id}/name][%d] getNetworksNetworkIdGatewaysGatewayIdNameOK  %+v", 200, o.Payload)
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDNameOK) GetPayload() models.GatewayName {
	return o.Payload
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDGatewaysGatewayIDNameDefault creates a GetNetworksNetworkIDGatewaysGatewayIDNameDefault with default headers values
func NewGetNetworksNetworkIDGatewaysGatewayIDNameDefault(code int) *GetNetworksNetworkIDGatewaysGatewayIDNameDefault {
	return &GetNetworksNetworkIDGatewaysGatewayIDNameDefault{
		_statusCode: code,
	}
}

/*GetNetworksNetworkIDGatewaysGatewayIDNameDefault handles this case with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDGatewaysGatewayIDNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID gateways gateway ID name default response
func (o *GetNetworksNetworkIDGatewaysGatewayIDNameDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDNameDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/gateways/{gateway_id}/name][%d] GetNetworksNetworkIDGatewaysGatewayIDName default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
