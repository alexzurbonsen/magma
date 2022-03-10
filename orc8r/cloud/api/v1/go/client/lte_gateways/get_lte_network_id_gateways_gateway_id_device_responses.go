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

// GetLTENetworkIDGatewaysGatewayIDDeviceReader is a Reader for the GetLTENetworkIDGatewaysGatewayIDDevice structure.
type GetLTENetworkIDGatewaysGatewayIDDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDGatewaysGatewayIDDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDGatewaysGatewayIDDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDGatewaysGatewayIDDeviceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDDeviceOK creates a GetLTENetworkIDGatewaysGatewayIDDeviceOK with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDDeviceOK() *GetLTENetworkIDGatewaysGatewayIDDeviceOK {
	return &GetLTENetworkIDGatewaysGatewayIDDeviceOK{}
}

/*GetLTENetworkIDGatewaysGatewayIDDeviceOK handles this case with default header values.

The physical device for the gateway
*/
type GetLTENetworkIDGatewaysGatewayIDDeviceOK struct {
	Payload *models.GatewayDevice
}

func (o *GetLTENetworkIDGatewaysGatewayIDDeviceOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/device][%d] getLteNetworkIdGatewaysGatewayIdDeviceOK  %+v", 200, o.Payload)
}

func (o *GetLTENetworkIDGatewaysGatewayIDDeviceOK) GetPayload() *models.GatewayDevice {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayDevice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDGatewaysGatewayIDDeviceDefault creates a GetLTENetworkIDGatewaysGatewayIDDeviceDefault with default headers values
func NewGetLTENetworkIDGatewaysGatewayIDDeviceDefault(code int) *GetLTENetworkIDGatewaysGatewayIDDeviceDefault {
	return &GetLTENetworkIDGatewaysGatewayIDDeviceDefault{
		_statusCode: code,
	}
}

/*GetLTENetworkIDGatewaysGatewayIDDeviceDefault handles this case with default header values.

Unexpected Error
*/
type GetLTENetworkIDGatewaysGatewayIDDeviceDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID gateways gateway ID device default response
func (o *GetLTENetworkIDGatewaysGatewayIDDeviceDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDGatewaysGatewayIDDeviceDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/gateways/{gateway_id}/device][%d] GetLTENetworkIDGatewaysGatewayIDDevice default  %+v", o._statusCode, o.Payload)
}

func (o *GetLTENetworkIDGatewaysGatewayIDDeviceDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDGatewaysGatewayIDDeviceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
