// Code generated by go-swagger; DO NOT EDIT.

package lte_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// GetLTENetworkIDGatewaysGatewayIDCellularDNSReader is a Reader for the GetLTENetworkIDGatewaysGatewayIDCellularDNS structure.
type GetLTENetworkIDGatewaysGatewayIDCellularDNSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDGatewaysGatewayIDCellularDNSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDGatewaysGatewayIDCellularDNSDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDCellularDNSOK creates a GetLTENetworkIDGatewaysGatewayIDCellularDNSOK with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDCellularDNSOK() *GetLTENetworkIDGatewaysGatewayIDCellularDNSOK {
	return &GetLTENetworkIDGatewaysGatewayIDCellularDNSOK{}
}

/*GetLTENetworkIDGatewaysGatewayIDCellularDNSOK handles this case with default header values.

DNS configuration of the gateway
*/
type GetLTENetworkIDGatewaysGatewayIDCellularDNSOK struct {
	Payload *models.GatewayDNSConfigs
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/cellular/dns][%d] getLteNetworkIdGatewaysGatewayIdCellularDnsOK  %+v", 200, o.Payload)
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSOK) GetPayload() *models.GatewayDNSConfigs {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayDNSConfigs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDGatewaysGatewayIDCellularDNSDefault creates a GetLTENetworkIDGatewaysGatewayIDCellularDNSDefault with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDCellularDNSDefault(code int) *GetLTENetworkIDGatewaysGatewayIDCellularDNSDefault {
	return &GetLTENetworkIDGatewaysGatewayIDCellularDNSDefault{
		_statusCode: code,
	}
}

/*GetLTENetworkIDGatewaysGatewayIDCellularDNSDefault handles this case with default header values.

Unexpected Error
*/
type GetLTENetworkIDGatewaysGatewayIDCellularDNSDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID gateways gateway ID cellular DNS default response
func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/cellular/dns][%d] GetLTENetworkIDGatewaysGatewayIDCellularDNS default  %+v", o._statusCode, o.Payload)
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDCellularDNSDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
