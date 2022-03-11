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

// DeleteNetworksNetworkIDGatewaysGatewayIDReader is a Reader for the DeleteNetworksNetworkIDGatewaysGatewayID structure.
type DeleteNetworksNetworkIDGatewaysGatewayIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworksNetworkIDGatewaysGatewayIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworksNetworkIDGatewaysGatewayIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteNetworksNetworkIDGatewaysGatewayIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNetworksNetworkIDGatewaysGatewayIDNoContent creates a DeleteNetworksNetworkIDGatewaysGatewayIDNoContent with default headers values
func NewDeleteNetworksNetworkIDGatewaysGatewayIDNoContent() *DeleteNetworksNetworkIDGatewaysGatewayIDNoContent {
	return &DeleteNetworksNetworkIDGatewaysGatewayIDNoContent{}
}

/* DeleteNetworksNetworkIDGatewaysGatewayIDNoContent describes a response with status code 204, with default header values.

Success
*/
type DeleteNetworksNetworkIDGatewaysGatewayIDNoContent struct {
}

func (o *DeleteNetworksNetworkIDGatewaysGatewayIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{network_id}/gateways/{gateway_id}][%d] deleteNetworksNetworkIdGatewaysGatewayIdNoContent ", 204)
}

func (o *DeleteNetworksNetworkIDGatewaysGatewayIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNetworksNetworkIDGatewaysGatewayIDDefault creates a DeleteNetworksNetworkIDGatewaysGatewayIDDefault with default headers values
func NewDeleteNetworksNetworkIDGatewaysGatewayIDDefault(code int) *DeleteNetworksNetworkIDGatewaysGatewayIDDefault {
	return &DeleteNetworksNetworkIDGatewaysGatewayIDDefault{
		_statusCode: code,
	}
}

/* DeleteNetworksNetworkIDGatewaysGatewayIDDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type DeleteNetworksNetworkIDGatewaysGatewayIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete networks network ID gateways gateway ID default response
func (o *DeleteNetworksNetworkIDGatewaysGatewayIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNetworksNetworkIDGatewaysGatewayIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /networks/{network_id}/gateways/{gateway_id}][%d] DeleteNetworksNetworkIDGatewaysGatewayID default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteNetworksNetworkIDGatewaysGatewayIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteNetworksNetworkIDGatewaysGatewayIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
