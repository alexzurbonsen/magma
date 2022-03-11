// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magma/orc8r/cloud/api/v1/go/models"
)

// PutCwfNetworkIDGatewaysGatewayIDTierReader is a Reader for the PutCwfNetworkIDGatewaysGatewayIDTier structure.
type PutCwfNetworkIDGatewaysGatewayIDTierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCwfNetworkIDGatewaysGatewayIDTierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutCwfNetworkIDGatewaysGatewayIDTierNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCwfNetworkIDGatewaysGatewayIDTierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCwfNetworkIDGatewaysGatewayIDTierNoContent creates a PutCwfNetworkIDGatewaysGatewayIDTierNoContent with default headers values
func NewPutCwfNetworkIDGatewaysGatewayIDTierNoContent() *PutCwfNetworkIDGatewaysGatewayIDTierNoContent {
	return &PutCwfNetworkIDGatewaysGatewayIDTierNoContent{}
}

/* PutCwfNetworkIDGatewaysGatewayIDTierNoContent describes a response with status code 204, with default header values.

Success
*/
type PutCwfNetworkIDGatewaysGatewayIDTierNoContent struct {
}

func (o *PutCwfNetworkIDGatewaysGatewayIDTierNoContent) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}/gateways/{gateway_id}/tier][%d] putCwfNetworkIdGatewaysGatewayIdTierNoContent ", 204)
}

func (o *PutCwfNetworkIDGatewaysGatewayIDTierNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCwfNetworkIDGatewaysGatewayIDTierDefault creates a PutCwfNetworkIDGatewaysGatewayIDTierDefault with default headers values
func NewPutCwfNetworkIDGatewaysGatewayIDTierDefault(code int) *PutCwfNetworkIDGatewaysGatewayIDTierDefault {
	return &PutCwfNetworkIDGatewaysGatewayIDTierDefault{
		_statusCode: code,
	}
}

/* PutCwfNetworkIDGatewaysGatewayIDTierDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type PutCwfNetworkIDGatewaysGatewayIDTierDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put cwf network ID gateways gateway ID tier default response
func (o *PutCwfNetworkIDGatewaysGatewayIDTierDefault) Code() int {
	return o._statusCode
}

func (o *PutCwfNetworkIDGatewaysGatewayIDTierDefault) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}/gateways/{gateway_id}/tier][%d] PutCwfNetworkIDGatewaysGatewayIDTier default  %+v", o._statusCode, o.Payload)
}
func (o *PutCwfNetworkIDGatewaysGatewayIDTierDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutCwfNetworkIDGatewaysGatewayIDTierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
