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

// DeleteCwfNetworkIDGatewaysGatewayIDReader is a Reader for the DeleteCwfNetworkIDGatewaysGatewayID structure.
type DeleteCwfNetworkIDGatewaysGatewayIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCwfNetworkIDGatewaysGatewayIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCwfNetworkIDGatewaysGatewayIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteCwfNetworkIDGatewaysGatewayIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCwfNetworkIDGatewaysGatewayIDNoContent creates a DeleteCwfNetworkIDGatewaysGatewayIDNoContent with default headers values
func NewDeleteCwfNetworkIDGatewaysGatewayIDNoContent() *DeleteCwfNetworkIDGatewaysGatewayIDNoContent {
	return &DeleteCwfNetworkIDGatewaysGatewayIDNoContent{}
}

/* DeleteCwfNetworkIDGatewaysGatewayIDNoContent describes a response with status code 204, with default header values.

Success
*/
type DeleteCwfNetworkIDGatewaysGatewayIDNoContent struct {
}

func (o *DeleteCwfNetworkIDGatewaysGatewayIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /cwf/{network_id}/gateways/{gateway_id}][%d] deleteCwfNetworkIdGatewaysGatewayIdNoContent ", 204)
}

func (o *DeleteCwfNetworkIDGatewaysGatewayIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCwfNetworkIDGatewaysGatewayIDDefault creates a DeleteCwfNetworkIDGatewaysGatewayIDDefault with default headers values
func NewDeleteCwfNetworkIDGatewaysGatewayIDDefault(code int) *DeleteCwfNetworkIDGatewaysGatewayIDDefault {
	return &DeleteCwfNetworkIDGatewaysGatewayIDDefault{
		_statusCode: code,
	}
}

/* DeleteCwfNetworkIDGatewaysGatewayIDDefault describes a response with status code -1, with default header values.

Unexpected Error
*/
type DeleteCwfNetworkIDGatewaysGatewayIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete cwf network ID gateways gateway ID default response
func (o *DeleteCwfNetworkIDGatewaysGatewayIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCwfNetworkIDGatewaysGatewayIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /cwf/{network_id}/gateways/{gateway_id}][%d] DeleteCwfNetworkIDGatewaysGatewayID default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteCwfNetworkIDGatewaysGatewayIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteCwfNetworkIDGatewaysGatewayIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
