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

// PutCwfNetworkIDGatewaysGatewayIDReader is a Reader for the PutCwfNetworkIDGatewaysGatewayID structure.
type PutCwfNetworkIDGatewaysGatewayIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCwfNetworkIDGatewaysGatewayIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutCwfNetworkIDGatewaysGatewayIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCwfNetworkIDGatewaysGatewayIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCwfNetworkIDGatewaysGatewayIDNoContent creates a PutCwfNetworkIDGatewaysGatewayIDNoContent with default headers values
func NewPutCwfNetworkIDGatewaysGatewayIDNoContent() *PutCwfNetworkIDGatewaysGatewayIDNoContent {
	return &PutCwfNetworkIDGatewaysGatewayIDNoContent{}
}

/*PutCwfNetworkIDGatewaysGatewayIDNoContent handles this case with default header values.

Success
*/
type PutCwfNetworkIDGatewaysGatewayIDNoContent struct {
}

func (o *PutCwfNetworkIDGatewaysGatewayIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}/gateways/{gateway_id}][%d] putCwfNetworkIdGatewaysGatewayIdNoContent ", 204)
}

func (o *PutCwfNetworkIDGatewaysGatewayIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCwfNetworkIDGatewaysGatewayIDDefault creates a PutCwfNetworkIDGatewaysGatewayIDDefault with default headers values
func NewPutCwfNetworkIDGatewaysGatewayIDDefault(code int) *PutCwfNetworkIDGatewaysGatewayIDDefault {
	return &PutCwfNetworkIDGatewaysGatewayIDDefault{
		_statusCode: code,
	}
}

/*PutCwfNetworkIDGatewaysGatewayIDDefault handles this case with default header values.

Unexpected Error
*/
type PutCwfNetworkIDGatewaysGatewayIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put cwf network ID gateways gateway ID default response
func (o *PutCwfNetworkIDGatewaysGatewayIDDefault) Code() int {
	return o._statusCode
}

func (o *PutCwfNetworkIDGatewaysGatewayIDDefault) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}/gateways/{gateway_id}][%d] PutCwfNetworkIDGatewaysGatewayID default  %+v", o._statusCode, o.Payload)
}

func (o *PutCwfNetworkIDGatewaysGatewayIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutCwfNetworkIDGatewaysGatewayIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
